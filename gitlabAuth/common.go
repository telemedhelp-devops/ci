package gitlabAuth

import (
	"github.com/markbates/goth"
	cfg "gitlab.telemed.help/devops/ci/config"
	"github.com/markbates/goth/providers/gitlab"
)

func init() {
	config := cfg.Get()

	goth.UseProviders(
		gitlab.New(config.GitLab.Key, config.GitLab.Secret, config.BaseURL+"/gitlab/callback"),
	)
}
