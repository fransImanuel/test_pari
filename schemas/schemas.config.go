package schemas

type SchemaEnvironment struct {
	REST_PORT string
	GO_ENV    string

	DB_USER    string
	DB_PASS    string
	DB_HOST    string
	DB_PORT    string
	DB_NAME    string
	DB_SSLMODE string

	MONGO_USER string
	MONGO_PASS string
	MONGO_HOST string
	MONGO_PORT string
	MONGO_DB   string
	MONGO_SSL  string
	MONGO_AUTH string

	TIMEZONE     string
	SWAGGER_HOST string

	Minio_Host      string
	Minio_Location  string
	Minio_AccessKey string
	Minio_SecretKey string
	Minio_SSL       string
	Minio_Domain    string
}

type SMTPConfig struct {
	Host     string
	Port     int
	Email    string
	Password string
	Name     string
}
