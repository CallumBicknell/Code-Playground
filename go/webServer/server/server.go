package server

import (
	"github.com/CallumBicknell/go-webServer/config"
	"github.com/CallumBicknell/go-webServer/routes"
)

func Init() {
	config := config.GetConfig()
	r := routes.NewRouter()
	r.Run(config.GetString("server.port"))
}
