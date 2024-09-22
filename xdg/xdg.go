// XDG Base Directory Specification
package xdg

import (
	"errors"
	"fmt"
	"os"
	"os/user"
	"strings"
)

const (
	EnvXDGConfigHome = "XDG_CONFIG_HOME"
	EnvXDGDataHome   = "XDG_DATA_HOME"
	EnvXDGCacheHome  = "XDG_CACHE_HOME"
	EnvXDGRuntimeDir = "XDG_RUNTIME_DIR"
	EnvXDGConfigDirs = "XDG_CONFIG_DIRS"
	EnvXDGDataDirs   = "XDG_DATA_DIRS"
)

var ErrInvalidUser = errors.New("invalid user")

type XDGBaseDirSpec struct {
	user *user.User
}

// NewXDG returns a new XDGBaseDirSpec instance for the current user.
func NewXDG() (*XDGBaseDirSpec, error) {
	user, err := user.Current()
	if err != nil {
		return nil, err
	}

	return NewXDGWithUser(user)
}

// NewXDGWithUsername returns a new XDGBaseDirSpec instance for the given username.
// When the user does not exist, an error is returned.
func NewXDGWithUsername(username string) (*XDGBaseDirSpec, error) {
	user, err := user.Lookup(username)
	if err != nil {
		return nil, err
	}

	return NewXDGWithUser(user)
}

// NewXDGWithUser returns a new XDGBaseDirSpec instance for the given user.
func NewXDGWithUser(user *user.User) (*XDGBaseDirSpec, error) {
	if err := validateUser(user); err != nil {
		return nil, err
	}
	return &XDGBaseDirSpec{user: user}, nil
}

// UserHomeDir returns the home directory of the user.
func (x *XDGBaseDirSpec) UserHomeDir() string {
	if x.user.HomeDir != "" {
		return x.user.HomeDir
	}
	return "/home/" + x.user.Username
}

// DefaultConfigHome returns the default configuration directory for the user,
// i.e. ~/.config.
func (x *XDGBaseDirSpec) DefaultConfigHome() string {
	return x.UserHomeDir() + "/.config"
}

// DefaultDataHome returns the default data directory for the user, i.e. ~/.local/share.
func (x *XDGBaseDirSpec) DefaultDataHome() string {
	return x.UserHomeDir() + "/.local/share"
}

// DefaultCacheHome returns the default cache directory for the user, i.e. ~/.cache.
func (x *XDGBaseDirSpec) DefaultCacheHome() string {
	return x.UserHomeDir() + "/.cache"
}

// DefaultConfigDirs returns the default configuration directories for the user.
func (x *XDGBaseDirSpec) DefaultConfigDirs() []string {
	return []string{"/etc/xdg"}
}

// DefaultDataDirs returns the default data directories for the user.
func (x *XDGBaseDirSpec) DefaultDataDirs() []string {
	return []string{"/usr/local/share", "/usr/share"}
}

// DefaultRuntimeDir returns the default runtime directory for the user.
// This is the directory where user-specific non-essential runtime files and
// other file objects (such as sockets, named pipes, ...) should be stored.
// In systemd-based systems, systemd-logind(8) will create and manage this
// directory for the user.
func (x *XDGBaseDirSpec) DefaultRuntimeDir() string {
	return "/run/user/" + x.user.Uid
}

func (x *XDGBaseDirSpec) ConfigHome() string {
	return getXDGDir(EnvXDGConfigHome, x.DefaultConfigHome())
}

func (x *XDGBaseDirSpec) DataHome() string {
	return getXDGDir(EnvXDGDataHome, x.DefaultDataHome())
}

func (x *XDGBaseDirSpec) CacheHome() string {
	return getXDGDir(EnvXDGCacheHome, x.DefaultCacheHome())
}

func (x *XDGBaseDirSpec) RuntimeDir() string {
	return getXDGDir(EnvXDGRuntimeDir, x.DefaultRuntimeDir())
}

func (x *XDGBaseDirSpec) ConfigDirs() []string {
	return getXDGDirs(EnvXDGConfigDirs, x.DefaultConfigDirs())
}

func (x *XDGBaseDirSpec) DataDirs() []string {
	return getXDGDirs(EnvXDGDataDirs, x.DefaultDataDirs())
}

func validateUser(user *user.User) error {
	if user == nil {
		return fmt.Errorf("%w: user is nil", ErrInvalidUser)
	}
	if user.Username == "" {
		return fmt.Errorf("%w: username is empty", ErrInvalidUser)
	}
	return nil
}

func getXDGDir(envVar, fallbackValue string) string {
	dir := os.Getenv(envVar)
	if dir == "" {
		dir = fallbackValue
	}
	return dir
}

func getXDGDirs(envVar string, fallbackValue []string) []string {
	directoryPaths := os.Getenv(envVar)
	if directoryPaths == "" {
		return fallbackValue
	}

	// Split the colon-separated string into a slice of strings.
	return removeEmptyStrings(strings.Split(directoryPaths, ":"))
}

func removeEmptyStrings(dirs []string) []string {
	var newDirs []string
	for _, dir := range dirs {
		if dir != "" {
			newDirs = append(newDirs, dir)
		}
	}
	return newDirs
}
