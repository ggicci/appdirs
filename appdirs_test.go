package appdirs

import (
	"testing"

	"github.com/ggicci/appdirs/internal"
	"github.com/ggicci/appdirs/xdg"
	"github.com/stretchr/testify/assert"
)

const TestAppName = "testapp"

func TestNewWithSpec(t *testing.T) {
	app := createTestApp(t)
	assert.Equal(t, TestAppName, app.App)
	assert.Equal(t, "/home/ggicci/.config/testapp", app.ConfigHome())
	assert.Equal(t, "/home/ggicci/.local/share/testapp", app.DataHome())
	assert.Equal(t, "/home/ggicci/.cache/testapp", app.CacheHome())
	assert.Equal(t, "/run/user/1000/testapp", app.RuntimeDir())
	assert.Equal(t, []string{"/etc/xdg/testapp"}, app.ConfigDirs())
	assert.Equal(t, []string{"/usr/local/share/testapp", "/usr/share/testapp"}, app.DataDirs())
}

func TestConfigFile(t *testing.T) {
	app := createTestApp(t)
	assert.Equal(t, "/home/ggicci/.config/testapp/myconfig.conf", app.ConfigFile("myconfig.conf"))
}

func TestSystemConfigFiles(t *testing.T) {
	app := createTestApp(t)
	assert.Equal(t, []string{
		"/etc/xdg/testapp/myconfig.conf",
	}, app.SystemConfigFiles("myconfig.conf"))
}

func TestConfigFiles(t *testing.T) {
	app := createTestApp(t)
	assert.Equal(t, []string{
		"/home/ggicci/.config/testapp/myconfig.conf",
		"/etc/xdg/testapp/myconfig.conf",
	}, app.ConfigFiles("myconfig.conf"))
}

func TestDataFile(t *testing.T) {
	app := createTestApp(t)
	assert.Equal(t, "/home/ggicci/.local/share/testapp/mydata.dat", app.DataFile("mydata.dat"))
}

func TestSystemDataFiles(t *testing.T) {
	app := createTestApp(t)
	assert.Equal(t, []string{
		"/usr/local/share/testapp/mydata.dat",
		"/usr/share/testapp/mydata.dat",
	}, app.SystemDataFiles("mydata.dat"))
}

func TestDataFiles(t *testing.T) {
	app := createTestApp(t)
	assert.Equal(t, []string{
		"/home/ggicci/.local/share/testapp/mydata.dat",
		"/usr/local/share/testapp/mydata.dat",
		"/usr/share/testapp/mydata.dat",
	}, app.DataFiles("mydata.dat"))
}

func TestCacheFile(t *testing.T) {
	app := createTestApp(t)
	assert.Equal(t, "/home/ggicci/.cache/testapp/mycache.dat", app.CacheFile("mycache.dat"))
}

func TestRuntimeFile(t *testing.T) {
	app := createTestApp(t)
	assert.Equal(t, "/run/user/1000/testapp/myruntime.dat", app.RuntimeFile("myruntime.dat"))
}

func createTestApp(t *testing.T) *AppDirs {
	user := internal.FakeUser("ggicci")
	spec, err := xdg.NewXDGWithUser(user)
	assert.NoError(t, err)
	assert.NotNil(t, spec)

	app, err := NewWithSpec(TestAppName, spec)
	assert.NoError(t, err)
	assert.NotNil(t, app)

	return app
}
