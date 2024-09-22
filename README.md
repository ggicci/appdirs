# appdirs

A utility to locate the base directories, as well as the config, data, cache files an app needs.

[![Go](https://github.com/ggicci/appdirs/actions/workflows/go.yaml/badge.svg)](https://github.com/ggicci/appdirs/actions/workflows/go.yaml)
[![codecov](https://codecov.io/gh/ggicci/appdirs/graph/badge.svg?token=YU7FGGOY60)](https://codecov.io/gh/ggicci/appdirs)
[![Go Report Card](https://goreportcard.com/badge/github.com/ggicci/appdirs)](https://goreportcard.com/report/github.com/ggicci/appdirs)
[![Go Reference](https://pkg.go.dev/badge/github.com/ggicci/appdirs.svg)](https://pkg.go.dev/github.com/ggicci/appdirs)

```go
import "github.com/ggicci/appdirs"

dirs := appdirs.New("myapp")
```

## The Directories API

By default, `dirs` will use the following directories (following the [XDG Base Directory Specification](https://specifications.freedesktop.org/basedir-spec/latest/index.html)):

| API                 | Value                                                         | Description                       |
| ------------------- | ------------------------------------------------------------- | --------------------------------- |
| `dirs.ConfigHome()` | `$XDG_CONFIG_HOME/APP` or `$HOME/.config/APP`                 | User-specific configuration files |
| `dirs.DataHome()`   | `$XDG_DATA_HOME/APP` or `$HOME/.local/share/APP`              | User-specific data files          |
| `dirs.CacheHome()`  | `$XDG_CACHE_HOME/APP` or `$HOME/.cache/APP`                   | User-specific non-essential data  |
| `dirs.RuntimeDir()` | `$XDG_RUNTIME_DIR/APP` or `/run/user/$UID/APP`                | User-specific runtime files       |
| `dirs.ConfigDirs()` | `$XDG_CONFIG_DIRS/APP` or `/etc/xdg/APP`                      | System-wide configuration files   |
| `dirs.DataDirs()`   | `$XDG_DATA_DIRS/APP` or `/usr/local/share/APP:/usr/share/APP` | System-wide data files            |

`APP` is the app name passed to `appdirs.New`.

You can override these directories by setting the environment variables `XDG_CONFIG_HOME`, `XDG_DATA_HOME`, `XDG_CACHE_HOME`, `XDG_RUNTIME_DIR`, `XDG_CONFIG_DIRS`, `XDG_DATA_DIRS`.

## The Paths API

The `appdirs` package provides a set of functions to locate the files an app needs:

| API                          | Value                        |
| ---------------------------- | ---------------------------- |
| `dirs.ConfigFile(FILENAME)`  | `dirs.ConfigHome()/FILENAME` |
| `dirs.DataFile(FILENAME)`    | `dirs.DataHome()/FILENAME`   |
| `dirs.CacheFile(FILENAME)`   | `dirs.CacheHome()/FILENAME`  |
| `dirs.RuntimeFile(FILENAME)` | `dirs.RuntimeDir()/FILENAME` |

More functions are available in the [godoc](https://pkg.go.dev/github.com/ggicci/appdirs).
