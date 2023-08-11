package casbin

import (
	"Identity/cmd/config"

	"github.com/casbin/casbin/v2"
)

func NewEnforcer(conf config.Config) (*casbin.Enforcer, error) {
	return casbin.NewEnforcer(conf.Casbin.ConfPath, conf.Casbin.CSVPath)
}
