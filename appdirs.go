package appdirs

import (
	"errors"
	"os"

	"github.com/ggicci/appdirs/internal"
	"github.com/ggicci/appdirs/spec"
	"github.com/ggicci/appdirs/xdg"
)

type DirSpec = spec.DirSpec

// AppDirs is a struct that provides the app-specific directories.
type AppDirs struct {
	// App is the name of the application. It's used to create the app-specific
	// directories. Used as the last part of the path. For example, if the root
	// config directory is ~/.config, and the app name is "myapp", then the
	// config directory of the app is ~/.config/myapp.
	App string

	// Root is the DirSpec object that provides the root directories. It's the
	// definition of the parent directories of the app. For example, if the root
	// directory of the config is ~/.config, then the config directory of the
	// app is ~/.config/APP_NAME, where APP_NAME is defined by the App field.
	Root DirSpec
}

func New(app string) (*AppDirs, error) {
	spec, err := xdg.NewXDG()
	if err != nil {
		return nil, nil
	}
	return NewWithSpec(app, spec)
}

func NewWithSpec(app string, spec DirSpec) (*AppDirs, error) {
	if err := validateApp(app); err != nil {
		return nil, err
	}
	return &AppDirs{App: app, Root: spec}, nil
}

func (b *AppDirs) ConfigHome() string {
	return internal.PathItemAppend(b.Root.ConfigHome(), b.App)
}

func (b *AppDirs) DataHome() string {
	return internal.PathItemAppend(b.Root.DataHome(), b.App)
}

func (b *AppDirs) CacheHome() string {
	return internal.PathItemAppend(b.Root.CacheHome(), b.App)
}

func (b *AppDirs) RuntimeDir() string {
	return internal.PathItemAppend(b.Root.RuntimeDir(), b.App)
}

func (b *AppDirs) ConfigDirs() []string {
	return internal.PathItemAppendList(b.Root.ConfigDirs(), b.App)
}

func (b *AppDirs) DataDirs() []string {
	return internal.PathItemAppendList(b.Root.DataDirs(), b.App)
}

// ConfigFile returns the full path of a user-specific configuration file.
// e.g. ~/.config/APP_NAME/myconfig.conf
func (b *AppDirs) ConfigFile(filename string) string {
	return internal.PathItemAppend(b.ConfigHome(), filename)
}

// SystemConfigFiles returns the full paths of system-wide configuration files.
// Remember that paths listed earlier have higher priority. e.g.
// [/etc/xdg/APP_NAME/myconfig.conf, /usr/local/etc/xdg/APP_NAME/myconfig.conf]
func (b *AppDirs) SystemConfigFiles(filename string) []string {
	return internal.PathItemAppendList(b.ConfigDirs(), filename)
}

// ConfigFiles returns the full paths of configuration files. It combines the
// results of UserConfigFile and SystemConfigFiles. Remember that paths listed
// earlier have higher priority. e.g. [~/.config/APP_NAME/myconfig.conf,
// /etc/xdg/APP_NAME/myconfig.conf, /usr/local/etc/xdg/APP_NAME/myconfig.conf]
func (b *AppDirs) ConfigFiles(filename string) []string {
	files := []string{b.ConfigFile(filename)}
	files = append(files, b.SystemConfigFiles(filename)...)
	return files
}

// DataFile returns the full path of a user-specific data file. e.g.
// ~/.local/share/APP_NAME/mydata.dat
func (b *AppDirs) DataFile(filename string) string {
	return internal.PathItemAppend(b.DataHome(), filename)
}

// SystemDataFiles returns the full paths of system-wide data files. Remember
// that paths listed earlier have higher priority. e.g.
// [/usr/local/share/APP_NAME/mydata.dat]
func (b *AppDirs) SystemDataFiles(filename string) []string {
	return internal.PathItemAppendList(b.DataDirs(), filename)
}

// DataFiles returns the full paths of data files. It combines the results of
// UserDataFile and SystemDataFiles. Remember that paths listed earlier have
// higher priority. e.g. [~/.local/share/APP_NAME/mydata.dat,
// /usr/local/share/APP_NAME/mydata.dat]
func (b *AppDirs) DataFiles(filename string) []string {
	files := []string{b.DataFile(filename)}
	files = append(files, b.SystemDataFiles(filename)...)
	return files
}

// CacheFile returns the full path of a cache file. e.g. ~/.cache/APP_NAME/mycache.dat
func (b *AppDirs) CacheFile(filename string) string {
	return internal.PathItemAppend(b.CacheHome(), filename)
}

// RuntimeFile returns the full path of a runtime file. e.g.
// /run/user/1000/APP_NAME/myruntime.dat
func (b *AppDirs) RuntimeFile(filename string) string {
	return internal.PathItemAppend(b.RuntimeDir(), filename)
}

// CreateDirectories creates all directories for the app. It doesn't create the
// runtime directory as it is not commonly used. If you need it, you can create
// it manually. It also doesn't create the system-wide directories as the
// current user may not have the permission to do so.
func (b *AppDirs) CreateDirectories() error {
	dirs := []string{
		b.ConfigHome(),
		b.DataHome(),
		b.CacheHome(),
	}
	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}
	return nil
}

func validateApp(app string) error {
	if app == "" {
		return errors.New("app name cannot be empty")
	}
	return nil
}
