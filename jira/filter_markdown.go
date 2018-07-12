package jira

import (
	"regexp"

	cfg "gitlab.telemed.help/devops/ci/config"
)

var (
	jiraLinkRegexp0 *regexp.Regexp
	jiraLinkRegexp1 *regexp.Regexp
	jiraLinkRegexp2 *regexp.Regexp
	jiraURL         string
)

func init() {
	cfg.AddReloadHook(reloadConfig)
}

func reloadConfig() {
	c := cfg.Get().Jira

	jiraURL = c.URL

	jiraLinkRegexpPart := c.IssuesRegexp
	jiraLinkRegexp0 = regexp.MustCompile(`^(` + jiraLinkRegexpPart + `)([ \r\n\t,.\-]+)`)
	jiraLinkRegexp1 = regexp.MustCompile(`([ \r\n\t,.\-]+)(` + jiraLinkRegexpPart + `)$`)
	jiraLinkRegexp2 = regexp.MustCompile(`([ \r\n\t,.\-]+)(` + jiraLinkRegexpPart + `)([ \r\n\t,.\-]+)`)
}

func FilterMarkdown(md string) string {
	md = jiraLinkRegexp0.ReplaceAllString(md, "[$1]("+jiraURL+"/browse/$1)$2")
	md = jiraLinkRegexp1.ReplaceAllString(md, "$1[$2]("+jiraURL+"/browse/$2)")
	md = jiraLinkRegexp2.ReplaceAllString(md, "$1[$2]("+jiraURL+"/browse/$2)$4")
	return md
}
