package main

import (
	"github.com/project-sistem-voucher/api/seeders"
	"github.com/project-sistem-voucher/config"
	_ "github.com/project-sistem-voucher/docs"
	routes "github.com/project-sistem-voucher/router"
)

// @title shipping API
// @version 1.0
// @description this is a shipping api
// @termofService
// @contact.name API support
// @contanct.url https://shipping.com/
// @License.name Lumoshive academy
// @host Localhost:8088
// @schemes http
// @basePath/v1/api/
func init() {
	config.InitiliazeConfig()
	config.InitDB()
	config.SyncDB()
	seeders.SeedVouchers(config.DB)
	seeders.SeedRedeem(config.DB)
}

func main() {
	// this is collaboration app
	routes.Server().Run()
}
