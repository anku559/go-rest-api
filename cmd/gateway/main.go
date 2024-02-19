package main

import (
	gatewayroutes "github.com/impleopus/go-rest-api/cmd/gateway/gateway-routes"
	"github.com/impleopus/go-rest-api/internal/config/db"
	"github.com/impleopus/go-rest-api/internal/config/env"
)

func init() {
	env.InitializeEnv()
	db.Connect()
	db.SyncDB()

}

func main() {
	app := gatewayroutes.SetupRouter()
	app.Run()
}
