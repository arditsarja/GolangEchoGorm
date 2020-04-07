package configuration

type AdminConfiguration struct {
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Address   string `json:"address"`
	Admin     bool   `json:"admin"`
}

var AdminConfig = AdminConfiguration{
	Username:  "admin",
	FirstName: "adminFirstName",
	LastName:  "adminLastName",
	Password:  "1234.com",
	Email:     "admin@emaildomain.com",
	Address:   "admin address",
	Admin:     true,
}
