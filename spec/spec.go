package spec

type DirSpec interface {
	ConfigHome() string
	DataHome() string
	CacheHome() string
	RuntimeDir() string
	ConfigDirs() []string
	DataDirs() []string
}
