package main

import (
	lightms "github.com/RajulSahu/lightms/v3"
	"github.com/RajulSahu/lightms/v3/example/infra/config"
	_ "github.com/RajulSahu/lightms/v3/example/infra/config"
	_ "github.com/RajulSahu/lightms/v3/example/infra/secondary"
)

func main() {
	// Adding properties
	lightms.AddConf[config.PropsConfig]()

	//Set properties file path. Default is "./resources/properties.yml"
	//lightms.SetPropFilePath("example/resources/properties.yml")
	lightms.SetPropFilePath("example/resources/properties.json")

	// Run lightms
	lightms.Run()
}
