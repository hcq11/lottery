package main

import (
	"lottery/api"
	"lottery/model"
)

// @title Lottery API
// @version 1.0
// @description TLottery API.

// @contact.name Eden
// @contact.email Eden@ubtrobot.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @BasePath /v1
func main() {
	model.Database()
	defer model.DB.Close()
	httpSrv := api.HTTPSrv{}
	httpSrv.Run()

}
