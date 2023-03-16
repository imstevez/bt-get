package main

import (
	"bt-get/bittorent"
	"errors"
	"fmt"
	"github.com/urfave/cli/v2"
	"path/filepath"
)

var cmds = &cli.App{
	Name:    "bt-get",
	Version: "v1.0.0",
	Usage:   "download file with bittorrent.",
	Commands: []*cli.Command{
		{
			Name:  "download",
			Usage: "Download all files within a bittorrent by a torrents file or a magnet uri",
			Flags: []cli.Flag{
				&cli.StringFlag{Name: "typ", Value: string(bittorent.BtTypeUri), Usage: "source type: 'file' for torrents file or 'uri' for magnet uri"},
				&cli.StringFlag{Name: "dir", Value: "./", Usage: "download dir"},
			},
			ArgsUsage: "torrents_file_path_or_magnet_uri",
			Before: func(ctx *cli.Context) (err error) {
				err = bittorent.Init(&bittorent.Config{
					DataDir: filepath.Clean(ctx.String("dir")),
				})
				return
			},
			Action: func(ctx *cli.Context) (err error) {
				bt := ctx.Args().Get(0)
				if bt == "" {
					err = errors.New("empty torrents file path or magnet uri")
					return
				}
				tp := bittorent.BtType(ctx.String("typ"))
				if tp != bittorent.BtTypeUri && tp != bittorent.BtTypeFile {
					err = errors.New("invalid typ")
					return
				}
				filename, err := bittorent.Download(ctx.Context, bt, tp)
				if err != nil {
					err = fmt.Errorf("download: %w", err)
					return
				}
				fmt.Printf("complete download: %s", filename)
				return
			},
		},
	},
}
