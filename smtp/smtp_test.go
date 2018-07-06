package smtp

import (
	"fmt"
	"testing"

	cfg "gitlab.telemed.help/devops/ci/config"
)

func init() {
	cfg.Reload("../config.yaml-sample")
}

func TestToHTML(t *testing.T) {
	src := `[hey](https://myorg.hey)

* DEVOPS-1 - The first item
* DEVOPS-321... The second item

DEVOPS-322 is also important.
And don't forget about DEVOPS-333

Best regards...`

	fmt.Println(toHTML(src))
}
