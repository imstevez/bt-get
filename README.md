# bt-get
Download file from bittorrent.

### All commands
```
NAME:
   bt-get - download file with bittorrent.

USAGE:
   bt-get [global options] command [command options] [arguments...]

VERSION:
   v1.0.0

COMMANDS:
   download  Download all files within a bittorrent by a torrents file or a magnet uri
   help, h   Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

- Download
```
NAME:
   bt-get download - Download all files within a bittorrent by a torrents file or a magnet uri

USAGE:
   bt-get download [command options] torrents_file_path_or_magnet_uri

OPTIONS:
   --typ value  source type: 'file' for torrents file or 'uri' for magnet uri (default: "uri")
   --dir value  download dir (default: "./")
   --help, -h   show help

```
