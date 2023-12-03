package main

import (
	"github.com/sikehish/Advertising-Agency-DBMS/configs"
	"github.com/sikehish/Advertising-Agency-DBMS/internal/routes"
)

func main() {
	configs.Initialize() // Initialization of config
	routes.SetupRoutes() //  API routes
}
