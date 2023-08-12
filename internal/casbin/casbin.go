package casbin

import (
	"Identity/cmd/config"
	"fmt"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	xormadapter "github.com/casbin/xorm-adapter/v2"
)

func NewEnforcer(conf config.Config) (*casbin.Enforcer, error) {
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s",
		conf.Database.Username,
		conf.Database.Password,
		conf.Database.Host,
		conf.Database.Port,
		conf.Database.Name,
	)
	a, _ := xormadapter.NewAdapter("mssql", dsn, true)

	m, _ := model.NewModelFromString(`
		[request_definition]
		r = sub, obj, act

		[policy_definition]
		p = sub, obj, act

		[role_definition]
		g = _, _

		[policy_effect]
		e = some(where (p.eft == allow))

		[matchers]
		g(r.sub, p.sub) && r.obj == p.obj && r.sub == p.sub
	`)

	return casbin.NewEnforcer(m, a)

	// // Load the policy from DB.
	// e.LoadPolicy()

	// // Check the permission.
	// e.Enforce("alice", "data1", "read")

	// // Modify the policy.
	// // e.AddPolicy(...)
	// // e.RemovePolicy(...)

	// // Save the policy back to DB.
	// e.SavePolicy()
}
