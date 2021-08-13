package main

import (
	"github.com/joho/godotenv"
	"github.com/rdurelli/todolist/src/api/bootstrap"
	"go.uber.org/fx"
)

func main() {
	godotenv.Load()
	fx.New(bootstrap.Module).Run()
}
