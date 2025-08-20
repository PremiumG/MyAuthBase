package main

import (
	"AuthBase/cmd/webTest"
	"AuthBase/internal/db"
	"AuthBase/internal/routes"
	"AuthBase/internal/utils"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	//just to run the testing file
	if utils.AppConfig.TestingRun == true {
		utils.L.Info("Runnig the other main script")
		webTest.Main2()
		os.Exit(0)
	}

	if err := db.Initialize(); err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()

	// Run migrations
	if err := db.Migrate(); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Optional: Seed admin user
	if err := db.SeedAdmin("admin@example.com"); err != nil {
		log.Fatal("Failed to seed admin:", err)
	}

	utils.L.Info("Runnig the MAIN script")

	r := gin.Default()

	if utils.AppConfig.Debug == false {
		gin.SetMode(gin.ReleaseMode)
	}

	r.Static("/assets", "./assets")
	r.LoadHTMLGlob("templates/*")
	routes.SetupRoutes(r)

	r.Run(fmt.Sprintf("%s:%v", utils.AppConfig.Host, utils.AppConfig.Port))
}
