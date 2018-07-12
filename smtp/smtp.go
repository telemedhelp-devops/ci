package smtp

import (
	"net/smtp"
	"strings"

	"github.com/gomarkdown/markdown"
	cfg "gitlab.telemed.help/devops/ci/config"
	"gitlab.telemed.help/devops/ci/jira"
)

var (
	smtpHost    string
	fromAddress string
)

func init() {
	cfg.AddReloadHook(reloadConfig)
}

func reloadConfig() {
	c := cfg.Get()
	smtpHost = c.Smtp.Host
	fromAddress = c.Smtp.FromAddress
}

func toHTML(md string) string {
	md = jira.FilterMarkdown(md)

	md = string(markdown.ToHTML([]byte(md), nil, nil))
	md = strings.Replace(md, "\n", "<br>\n", -1)

	return md
}

func Send(recipientAddress, subject, message string) error {
	return smtp.SendMail(smtpHost+":25", nil, fromAddress, []string{recipientAddress}, []byte("To: "+recipientAddress+"\r\n"+
		"MIME-version: 1.0;\r\nContent-Type: text/html; charset=\"UTF-8\";\r\nSubject: "+subject+"\r\n"+
		"\r\n"+toHTML(message)))
}
