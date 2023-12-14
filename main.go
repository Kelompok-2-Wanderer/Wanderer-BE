package main

import (
	"wanderer/config"
	"wanderer/routes"
	"wanderer/utils/database"

	lh "wanderer/features/locations/handler"
	lr "wanderer/features/locations/repository"
	ls "wanderer/features/locations/service"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	var dbConfig = new(config.DatabaseMysql)
	if err := dbConfig.LoadFromEnv(); err != nil {
		panic(err)
	}

	dbConnection, err := database.MysqlInit(*dbConfig)
	if err != nil {
		panic(err)
	}

	if err := database.MysqlMigrate(dbConnection); err != nil {
		panic(err)
	}

	var jwtConfig = new(config.JWT)
	if err := jwtConfig.LoadFromEnv(); err != nil {
		panic(err)
	}

	var cldConfig = new(config.Cloudinary)
	if err := cldConfig.LoadFromEnv(); err != nil {
		panic(err)
	}

	cld, err := cloudinary.NewFromParams(cldConfig.CloudName, cldConfig.ApiKey, cldConfig.ApiSecret)
	if err != nil {
		panic(err)
	}

	locationRepository := lr.NewLocationRepository(dbConnection, cld)
	locationService := ls.NewLocationService(locationRepository)
	locationHandler := lh.NewLocationHandler(locationService)

	app := echo.New()
	app.Use(middleware.Recover())
	app.Use(middleware.CORS())

	route := routes.Routes{
		JWTKey:          jwtConfig.Secret,
		Server:          app,
		LocationHandler: locationHandler,
	}

	route.InitRouter()

	app.Logger.Fatal(app.Start(":8000"))
}
