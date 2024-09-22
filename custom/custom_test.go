package custom

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type DirSpecImplForTest struct{}

func (d *DirSpecImplForTest) ConfigHome() string {
	return "/config"
}

func (d *DirSpecImplForTest) DataHome() string {
	return "/data"
}

func (d *DirSpecImplForTest) CacheHome() string {
	return "/cache"
}

func (d *DirSpecImplForTest) RuntimeDir() string {
	return "/runtime"
}

func (d *DirSpecImplForTest) ConfigDirs() []string {
	return []string{"/config1", "/config2"}
}

func (d *DirSpecImplForTest) DataDirs() []string {
	return []string{"/data1", "/data2"}
}

func TestNewCustomBaseDirSpecFrom(t *testing.T) {
	custom := NewCustomBaseDirSpecFrom(&DirSpecImplForTest{})
	expectedSpec := &CustomBaseDirSpec{
		ConfigHomeValue: "/config",
		DataHomeValue:   "/data",
		CacheHomeValue:  "/cache",
		RuntimeDirValue: "/runtime",
		ConfigDirsValue: []string{"/config1", "/config2"},
		DataDirsValue:   []string{"/data1", "/data2"},
	}
	assert.Equal(t, expectedSpec, custom)
}

func TestCustomBaseDirSpec_ConfigHome(t *testing.T) {
	custom := NewCustomBaseDirSpecFrom(&DirSpecImplForTest{})
	assert.Equal(t, "/config", custom.ConfigHome())
}

func TestCustomBaseDirSpec_DataHome(t *testing.T) {
	custom := NewCustomBaseDirSpecFrom(&DirSpecImplForTest{})
	assert.Equal(t, "/data", custom.DataHome())
}

func TestCustomBaseDirSpec_CacheHome(t *testing.T) {
	custom := NewCustomBaseDirSpecFrom(&DirSpecImplForTest{})
	assert.Equal(t, "/cache", custom.CacheHome())
}

func TestCustomBaseDirSpec_RuntimeDir(t *testing.T) {
	custom := NewCustomBaseDirSpecFrom(&DirSpecImplForTest{})
	assert.Equal(t, "/runtime", custom.RuntimeDir())
}

func TestCustomBaseDirSpec_ConfigDirs(t *testing.T) {
	custom := NewCustomBaseDirSpecFrom(&DirSpecImplForTest{})
	assert.Equal(t, []string{"/config1", "/config2"}, custom.ConfigDirs())
}

func TestCustomBaseDirSpec_DataDirs(t *testing.T) {
	custom := NewCustomBaseDirSpecFrom(&DirSpecImplForTest{})
	assert.Equal(t, []string{"/data1", "/data2"}, custom.DataDirs())
}
