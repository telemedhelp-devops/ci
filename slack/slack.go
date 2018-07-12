package slack

import (
	"github.com/ashwanthkumar/slack-go-webhook"
	cfg "gitlab.telemed.help/devops/ci/config"
	"gitlab.telemed.help/devops/ci/jira"
)

func Send(markdown string) error {
	c := cfg.Get().Slack

	markdown = jira.FilterMarkdown(markdown)

	payload := slack.Payload{
		Text:      markdown,
		Username:  c.Username,
		Channel:   c.Channel,
		IconEmoji: c.IconEmoji,
	}
	errs := slack.Send(c.HookURL, "", payload)
	if len(errs) == 0 {
		return nil
	}
	return errs[0]
}
