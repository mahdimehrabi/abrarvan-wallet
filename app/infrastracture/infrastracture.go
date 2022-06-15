package infrastracture

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewLogger),
	fx.Provide(NewPgxDB),
	fx.Provide(NewRedis),
	fx.Provide(NewRandom),
)
