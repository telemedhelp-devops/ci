package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type DbCfg struct {
	Driver   string
	Protocol string
	Host     string
	Port     int
	Db       string
	User     string
	Password string
	Path     string
}

type GitLabCfg struct {
	URL    string `yaml:"url"`
	Key    string
	Secret string
}

type Config struct {
	IsDev           bool      `yaml:"development_mode"`
	BaseURL         string    `yaml:"base_url"`
	Secret          string    `yaml:"secret"`
	LogReaderIPAddr string    `yaml:"log_reader_ip_address"`
	LogWriterIPAddr string    `yaml:"log_writer_ip_address"`
	GitLab          GitLabCfg `yaml:"gitlab"`
	Db              DbCfg     `yaml:"db"`
}

var cfg Config

var reloadHooks []func()

func checkErr(err error) {
	if err == nil {
		return
	}

	panic(err)
}

func Reload() {
	configData, err := ioutil.ReadFile("config.yaml")
	checkErr(err)

	err = yaml.Unmarshal([]byte(configData), &cfg)
	checkErr(err)

	for _, reloadHook := range reloadHooks {
		reloadHook()
	}
}

func Get() Config {
	return cfg
}

func AddReloadHook(hook func()) {
	reloadHooks = append(reloadHooks, hook)
}
