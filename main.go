package main

import (
	"lottery/api"
	"lottery/model"
)

func main() {
	model.Database()
	defer model.DB.Close()
	httpSrv := api.HTTPSrv{}
	httpSrv.Run()

}
