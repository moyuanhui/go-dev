package config

type Config struct {
	LogLevel string
	LogPath  string

	CollectConf []CollectConf
}

type CollectConf struct {
	LogPath string
	Topic   string
}
