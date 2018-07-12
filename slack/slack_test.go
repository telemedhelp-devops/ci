package slack

import (
	"testing"

	cfg "gitlab.telemed.help/devops/ci/config"
)

func init() {
	cfg.Reload("../config.yaml")
}

func TestSend(t *testing.T) {
	text := `[hey](https://myorg.hey)

* DEVOPS-1 - The first item
* DEVOPS-321... The second item

DEVOPS-322 is also important.
And don't forget about DEVOPS-333

Best regards...`

	err := Send(text)
	if err != nil {
		t.Errorf("Cannot send a message to Slack/Mattermost: %v", err.Error())
	}
}
