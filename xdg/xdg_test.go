package xdg

import (
	"os"
	"os/user"
	"testing"

	"github.com/ggicci/appdirs/internal"
	"github.com/stretchr/testify/assert"
)

func TestNewXDG(t *testing.T) {
	currentUser, err := user.Current()

	if err != nil {
		t.Logf("user.Current() failed: %v", err)
		spec, err := NewXDG()
		assert.Error(t, err)
		assert.Nil(t, spec)
	} else {
		spec, err := NewXDG()
		assert.NoError(t, err)
		assert.Equal(t, currentUser.Username, spec.user.Username)
	}
}

func TestNewXDGWithUsername(t *testing.T) {
	username := "nonexistentuser"
	spec, err := NewXDGWithUsername(username)
	assert.ErrorIs(t, err, user.UnknownUserError(username))
	assert.Nil(t, spec)

	// current os user
	currentUser, err := user.Current()
	if err == nil {
		spec, err := NewXDGWithUsername(currentUser.Username)
		assert.NoError(t, err)
		assert.NotNil(t, spec)
	}
}

func TestNewXDGWithInvalidUser(t *testing.T) {
	// nil user
	spec, err := NewXDGWithUser(nil)
	assert.Nil(t, spec)
	assert.ErrorIs(t, err, ErrInvalidUser)
	assert.ErrorContains(t, err, "user is nil")

	// empty username
	user := internal.FakeUser("")
	spec, err = NewXDGWithUser(user)
	assert.Nil(t, spec)
	assert.ErrorIs(t, err, ErrInvalidUser)
	assert.ErrorContains(t, err, "username is empty")
}

func TestUserHomeDir(t *testing.T) {
	user := internal.FakeUser("ggicci")
	expectedHomeDir := "/home/ggicci"

	spec, err := NewXDGWithUser(user)
	assert.NoError(t, err)
	assert.Equal(t, expectedHomeDir, spec.UserHomeDir())

	user.HomeDir = "" // reset to empty
	assert.Equal(t, expectedHomeDir, spec.UserHomeDir())
}

func TestGettingDefaultHomeDirs(t *testing.T) {
	spec := createTestSpecWithFakeUser(t, "ggicci")
	assert.Equal(t, "/home/ggicci/.config", spec.DefaultConfigHome())
	assert.Equal(t, "/home/ggicci/.local/share", spec.DefaultDataHome())
	assert.Equal(t, "/home/ggicci/.cache", spec.DefaultCacheHome())
	assert.Equal(t, "/run/user/1000", spec.DefaultRuntimeDir())
	assert.Equal(t, []string{"/etc/xdg"}, spec.DefaultConfigDirs())
	assert.Equal(t, []string{"/usr/local/share", "/usr/share"}, spec.DefaultDataDirs())
}

func TestGettingHomeDirs_noEnvOverrides(t *testing.T) {
	spec := createTestSpecWithFakeUser(t, "ggicci")
	assert.Equal(t, "/home/ggicci/.config", spec.ConfigHome())
	assert.Equal(t, "/home/ggicci/.local/share", spec.DataHome())
	assert.Equal(t, "/home/ggicci/.cache", spec.CacheHome())
	assert.Equal(t, "/run/user/1000", spec.RuntimeDir())
	assert.Equal(t, []string{"/etc/xdg"}, spec.ConfigDirs())
	assert.Equal(t, []string{"/usr/local/share", "/usr/share"}, spec.DataDirs())
}

func TestGettingHomeDirs_withEnvOverrides(t *testing.T) {
	spec := createTestSpecWithFakeUser(t, "ggicci")
	os.Setenv(EnvXDGConfigHome, "/tmp/app/config")
	assert.Equal(t, "/tmp/app/config", spec.ConfigHome())

	os.Setenv(EnvXDGDataHome, "/tmp/app/data")
	assert.Equal(t, "/tmp/app/data", spec.DataHome())

	os.Setenv(EnvXDGCacheHome, "/tmp/app/cache")
	assert.Equal(t, "/tmp/app/cache", spec.CacheHome())

	os.Setenv(EnvXDGRuntimeDir, "/tmp/app/runtime")
	assert.Equal(t, "/tmp/app/runtime", spec.RuntimeDir())

	os.Setenv(EnvXDGConfigDirs, "/etc/legacyApp:/tmp/config/app")
	assert.Equal(t, []string{"/etc/legacyApp", "/tmp/config/app"}, spec.ConfigDirs())

	os.Setenv(EnvXDGDataDirs, "/var/share/app:/tmp/data/app:")
	assert.Equal(t, []string{"/var/share/app", "/tmp/data/app"}, spec.DataDirs())
}

func createTestSpecWithFakeUser(t *testing.T, username string) *XDGBaseDirSpec {
	user := internal.FakeUser(username)
	spec, err := NewXDGWithUser(user)
	assert.NoError(t, err)
	return spec
}
