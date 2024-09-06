package utils

import (
	"fmt"
	"os"
	"test-pari/schemas"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Environment() (config schemas.SchemaEnvironment) {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
		panic("Error loading .env file")
	}

	// Read environment variables from docker-compose.yml (if in production) / env from local dev
	config.DB_HOST = os.Getenv("DB_HOST")
	config.DB_PORT = os.Getenv("DB_PORT")
	config.DB_USER = os.Getenv("DB_USER")
	config.DB_NAME = os.Getenv("DB_NAME")
	config.DB_PASS = os.Getenv("DB_PASS")
	config.DB_SSLMODE = os.Getenv("DB_SSLMODE")

	config.MONGO_HOST = os.Getenv("MONGO_HOST")
	config.MONGO_PORT = os.Getenv("MONGO_PORT")
	config.MONGO_USER = os.Getenv("MONGO_USER")
	config.MONGO_PASS = os.Getenv("MONGO_PASS")
	config.MONGO_DB = os.Getenv("MONGO_DB")
	config.MONGO_SSL = os.Getenv("MONGO_SSL")
	config.MONGO_AUTH = os.Getenv("MONGO_AUTH")

	config.TIMEZONE = os.Getenv("TIMEZONE")
	config.REST_PORT = os.Getenv("REST_PORT")
	config.GO_ENV = os.Getenv("GO_ENV")
	config.SWAGGER_HOST = os.Getenv("SWAGGER_HOST")

	config.Minio_Host = os.Getenv("MINIO_HOST")
	config.Minio_SSL = os.Getenv("MINIO_SSL")
	config.Minio_SecretKey = os.Getenv("MINIO_SECRET_KEY")
	config.Minio_AccessKey = os.Getenv("MINIO_ACCESS_KEY")
	config.Minio_Domain = os.Getenv("MINIO_DOMAIN")

	return config
}

/*
? I use gin as a framework in this project example
*/
func SetupRouter(config schemas.SchemaEnvironment) *gin.Engine {

	app := gin.Default()

	/*
	* ? gin.SetMode nentuin seberapa detail logging yang muncul
	 */
	if config.GO_ENV == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else if config.GO_ENV == "test" {
		gin.SetMode(gin.TestMode)
	} else {
		gin.SetMode(gin.DebugMode)

	}

	app.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"*"},
		AllowHeaders:  []string{"*"},
		AllowWildcard: true,
	}))

	return app
}
