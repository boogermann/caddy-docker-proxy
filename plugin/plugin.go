package plugin

import (
	"github.com/mholt/caddy"
)

func init() {
	caddy.RegisterCaddyfileLoader("docker-input", CreateDockerLoader())
}