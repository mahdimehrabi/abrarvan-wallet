package main

import (
	"challange/app"
	"go.uber.org/fx"
)

func main() {
	fx.New(app.BootstrapModule).Run()

}
