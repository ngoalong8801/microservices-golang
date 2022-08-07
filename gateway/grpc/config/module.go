package gconfig

import "go.uber.org/fx"

var Module = fx.Provide(
	NewTransportOption,
	NewAccountCC,
)
