package controller

import "go.uber.org/fx"

var Module = fx.Provide(
	NewAccountController,
)
