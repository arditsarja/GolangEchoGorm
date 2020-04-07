package configuration

type DatabaseConfiguration struct {
	DbHost     string `json:"db_host"`
	DbPort     string `json:"db_port"`
	DbUser     string `json:"db_user"`
	DbName     string `json:"db_name"`
	DbPassword string `json:"db_password"`
}

var DbConfig = DatabaseConfiguration{
	DbHost:     "localhost",
	DbPort:     "3306",
	DbUser:     "root",
	DbName:     "golang_project",
	DbPassword: "root",
}
