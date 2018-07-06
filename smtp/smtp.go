package smtp

import (
	"net/smtp"
	"regexp"
	"strings"

	"github.com/gomarkdown/markdown"
	cfg "gitlab.telemed.help/devops/ci/config"
)

var (
	smtpHost        string
	fromAddress     string
	jiraLinkRegexp0 *regexp.Regexp
	jiraLinkRegexp1 *regexp.Regexp
	jiraLinkRegexp2 *regexp.Regexp
	jiraURL         string
)

func init() {
	cfg.AddReloadHook(reloadConfig)
}

func reloadConfig() {
	c := cfg.Get()
	smtpHost = c.Smtp.Host
	fromAddress = c.Smtp.FromAddress

	jiraURL = c.Jira.URL

	jiraLinkRegexpPart := c.Jira.IssuesRegexp
	jiraLinkRegexp0 = regexp.MustCompile(`^(` + jiraLinkRegexpPart + `)([ \r\n\t,.\-]+)`)
	jiraLinkRegexp1 = regexp.MustCompile(`([ \r\n\t,.\-]+)(` + jiraLinkRegexpPart + `)$`)
	jiraLinkRegexp2 = regexp.MustCompile(`([ \r\n\t,.\-]+)(` + jiraLinkRegexpPart + `)([ \r\n\t,.\-]+)`)
}

func toHTML(md string) string {
	md = jiraLinkRegexp0.ReplaceAllString(md, "[$1]("+jiraURL+"/browse/$1)$2")
	md = jiraLinkRegexp1.ReplaceAllString(md, "$1[$2]("+jiraURL+"/browse/$2)")
	md = jiraLinkRegexp2.ReplaceAllString(md, "$1[$2]("+jiraURL+"/browse/$2)$4")

	md = string(markdown.ToHTML([]byte(md), nil, nil))
	md = strings.Replace(md, "\n", "<br>\n", -1)

	return md
}

func Send(recipientAddress, subject, message string) error {
	return smtp.SendMail(smtpHost+":25", nil, fromAddress, []string{recipientAddress}, []byte("To: "+recipientAddress+"\r\n"+
		"MIME-version: 1.0;\r\nContent-Type: text/html; charset=\"UTF-8\";\r\nSubject: "+subject+"\r\n"+
		"\r\n"+toHTML(message)))
}
