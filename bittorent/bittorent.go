package bittorent

import (
	"bt-get/clean"
	"context"
	"errors"
	"fmt"
	"github.com/anacrolix/log"
	"github.com/anacrolix/torrent"
	"github.com/dustin/go-humanize"
	"path/filepath"
	"time"
)

var (
	client  *torrent.Client
	dataDir string
)

type Config struct {
	DataDir string
}

func Init(cfg *Config) (err error) {
	clientConfig := torrent.NewDefaultClientConfig()
	logger := log.NewLogger()
	logger.SetHandlers(log.DiscardHandler)
	clientConfig.Logger = logger
	dataDir = filepath.Clean(cfg.DataDir)
	clientConfig.DataDir = dataDir
	clientConfig.ListenPort = 0
	clientConfig.NoUpload = true
	client, err = torrent.NewClient(clientConfig)
	if err != nil {
		return
	}
	clean.PushClose(func() (_ error) {
		client.Close()
		return
	})
	return
}

type BtType string

const (
	BtTypeFile BtType = "file"
	BtTypeUri  BtType = "uri"
)

func Download(ctx context.Context, bt string, btType BtType) (filename string, err error) {
	var t *torrent.Torrent
	switch btType {
	case BtTypeFile:
		t, err = client.AddTorrentFromFile(bt)
	case BtTypeUri:
		t, err = client.AddMagnet(bt)
	default:
		err = errors.New("unknown bt type")
	}
	if err != nil {
		return
	}

	select {
	case <-t.GotInfo():
		t.DownloadAll()
		filename = t.Name()
	case <-ctx.Done():
		err = ctx.Err()
		return
	}

	fmt.Printf("\033[s")
	defer fmt.Printf("\033[0m")

	for {
		select {
		case <-time.After(500 * time.Millisecond):
			fmt.Printf("\033[u\033[2J ")
			compl := t.BytesCompleted()
			total := t.Length()
			fmt.Printf(
				"[ %q ]-[ %s/%s ]-[ %d%% ]\033[K\n",
				filename,
				humanize.Bytes(uint64(compl)),
				humanize.Bytes(uint64(total)),
				compl*100/total,
			)
			fmt.Println("---")
			for _, f := range t.Files() {
				compl := f.BytesCompleted()
				total := f.Length()
				fmt.Printf("[ %q ]-[ %s/%s ]-[ %d%% ]\033[K\n",
					f.Path(),
					humanize.Bytes(uint64(compl)),
					humanize.Bytes(uint64(total)),
					compl*100/total,
				)
			}
			if t.Complete.Bool() {
				return
			}
		case <-ctx.Done():
			err = ctx.Err()
			return
		}
	}
}
