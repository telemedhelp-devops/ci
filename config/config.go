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
	Token  string
}

type SmtpCfg struct {
	Host        string `yaml:"host"`
	FromAddress string `yaml:"from_address"`
}

type JiraCfg struct {
	URL          string `yaml:"url"`
	IssuesRegexp string `yaml:"issues_regexp"`
}

type SlackCfg struct {
	HookURL   string `yaml:"hook_url"`
	Username  string `yaml:"username"`
	Channel   string `yaml:"channel"`
	IconEmoji string `yaml:"icon_emoji"`
}

type Developer struct {
	Email string `yaml:"email"`
}

type Config struct {
	IsDev             bool      `yaml:"development_mode"`
	BaseURL           string    `yaml:"base_url"`
	Secret            string    `yaml:"secret"`
	LogReaderIPAddr   string    `yaml:"log_reader_ip_address"`
	LogWriterIPAddr   string    `yaml:"log_writer_ip_address"`
	SMSGWLogin        string    `yaml:"sms_gw_login"`
	SMSGWPassword     string    `yaml:"sms_gw_password"`
	GitLab            GitLabCfg `yaml:"gitlab"`
	Db                DbCfg     `yaml:"db"`
	Smtp              SmtpCfg   `yaml:"smtp"`
	Jira              JiraCfg   `yaml:"jira"`
	Slack             SlackCfg  `yaml:"slack"`
	NotificationEmail string    `yaml:"notification_email"`
	DefaultDeveloper  Developer `yaml:"default_developer"`
}

var cfg Config

var reloadHooks []func()

func checkErr(err error) {
	if err == nil {
		return
	}

	panic(err)
}

func Reload(path string) {
	configData, err := ioutil.ReadFile(path)
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
