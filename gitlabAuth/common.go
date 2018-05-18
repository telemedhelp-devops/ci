package gitlabAuth

import (
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/gitlab"
	cfg "gitlab.telemed.help/devops/ci/config"
)

func Init() {
	config := cfg.Get()

	goth.UseProviders(
		gitlab.NewCustomisedURL(
			config.GitLab.Key,
			config.GitLab.Secret,
			config.BaseURL+"/auth/gitlab/callback",
			config.GitLab.URL+"/oauth/authorize",
			config.GitLab.URL+"/oauth/token",
			config.GitLab.URL+"/api/v3/user",
			"read_user",
		),
	)
}
