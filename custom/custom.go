package custom

import "github.com/ggicci/appdirs/spec"

type CustomBaseDirSpec struct {
	ConfigHomeValue string
	DataHomeValue   string
	CacheHomeValue  string
	RuntimeDirValue string
	ConfigDirsValue []string
	DataDirsValue   []string
}

func NewCustomBaseDirSpecFrom(spec spec.DirSpec) *CustomBaseDirSpec {
	return &CustomBaseDirSpec{
		ConfigHomeValue: spec.ConfigHome(),
		DataHomeValue:   spec.DataHome(),
		CacheHomeValue:  spec.CacheHome(),
		RuntimeDirValue: spec.RuntimeDir(),
		ConfigDirsValue: spec.ConfigDirs(),
		DataDirsValue:   spec.DataDirs(),
	}
}

func (s *CustomBaseDirSpec) ConfigHome() string {
	return s.ConfigHomeValue
}

func (s *CustomBaseDirSpec) DataHome() string {
	return s.DataHomeValue
}

func (s *CustomBaseDirSpec) CacheHome() string {
	return s.CacheHomeValue
}

func (s *CustomBaseDirSpec) RuntimeDir() string {
	return s.RuntimeDirValue
}

func (s *CustomBaseDirSpec) ConfigDirs() []string {
	return s.ConfigDirsValue
}

func (s *CustomBaseDirSpec) DataDirs() []string {
	return s.DataDirsValue
}
