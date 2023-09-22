package server

import (
	"github.com/joho/godotenv"
	"main/router"
	"os"
)

func Start() {
	_ = godotenv.Load()
	e := router.BaseRouter()

	e.Static(os.Getenv("STATIC_SERVE_FOLDER"), os.Getenv("SERVE_FILE"))

	if err := e.Start(":1234"); err != nil {
		panic(err)
	}
}
