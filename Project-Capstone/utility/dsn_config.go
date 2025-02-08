package utility

func getDsn() string {

	host, err := resourceManager.GetProperty("host")
	if err != nil {
		Log.Error(err.Error())
		errorMessage, _ := resourceManager.GetProperty("database.connection.error")
		Log.Error(errorMessage)
		panic(errorMessage)
	}

	user, err := resourceManager.GetProperty("user")
	if err != nil {
		Log.Error(err.Error())
		errorMessage, _ := resourceManager.GetProperty("database.connection.error")
		Log.Error(errorMessage)
		panic(errorMessage)
	}

	password, err := resourceManager.GetProperty("password")
	if err != nil {
		Log.Error(err.Error())
		errorMessage, _ := resourceManager.GetProperty("database.connection.error")
		Log.Error(errorMessage)
		panic(errorMessage)
	}

	dbname, err := resourceManager.GetProperty("dbname")
	if err != nil {
		Log.Error(err.Error())
		errorMessage, _ := resourceManager.GetProperty("database.connection.error")
		Log.Error(errorMessage)
		panic(errorMessage)
	}

	port, err := resourceManager.GetProperty("port")
	if err != nil {
		Log.Error(err.Error())
		errorMessage, _ := resourceManager.GetProperty("database.connection.error")
		Log.Error(errorMessage)
		panic(errorMessage)
	}

	sslmode, err := resourceManager.GetProperty("sslmode")
	if err != nil {
		Log.Error(err.Error())
		errorMessage, _ := resourceManager.GetProperty("database.connection.error")
		Log.Error(errorMessage)
		panic(errorMessage)
	}

	return "host=" + host + " " + "user=" + user + " " + "password=" + password + " " + "dbname=" + dbname + " " + "port=" + port + " " + "sslmode=" + sslmode
}
