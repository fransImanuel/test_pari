package main

import (
	"fmt"
	"log"
	"strconv"
	"test-pari/docs"
	"test-pari/drivers"
	"test-pari/utils"

	_ "test-pari/docs"

	// swagger embed files
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	itemHandler "test-pari/modules/items/handler"
	itemRepo "test-pari/modules/items/repository"
	itemServ "test-pari/modules/items/service"
)

func main() {

	ConfigEnv := utils.Environment()
	RESTPort, err := strconv.Atoi(ConfigEnv.REST_PORT)
	if err != nil {
		/**
		* ? usually log package include date & time
		* ? rather than print the error message only ( fmt.Print )
		 */

		log.Println("REST_PORT is not valid ", err.Error())
	}

	app := utils.SetupRouter(ConfigEnv)

	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Title = "Project API"
	docs.SwaggerInfo.Description = "Project API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = ConfigEnv.SWAGGER_HOST
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	// @title Project API
	// @version 1.0
	// @description This is a sample server celler server.
	// @termsOfService http://swagger.io/terms/

	// @contact.name API Support
	// @contact.url http://www.swagger.io/support
	// @contact.email yaour.personal.email@gmail.com

	// @license.name Apache 2.0
	// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

	// @query.collection.format multi

	// @securityDefinitions.apikey ApiKeyAuth
	// @in header
	// @name Authorization
	// @x-extension-openapi {"example": "value on a json format"}

	DBPostgres, err := drivers.SetupDBSQL(ConfigEnv)
	if err != nil {
		log.Fatal(err.Error())
	}

	itemRepository := itemRepo.InitItemRepository(DBPostgres)
	itemService := itemServ.InitItemsService(itemRepository)
	itemHandler.InitItemHandler(app, itemService)

	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	app.Run(fmt.Sprintf(":%v", RESTPort))
}
