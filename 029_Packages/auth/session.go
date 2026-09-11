package auth

// now see this extractSession() is private but GetSession() is not private.


func GetSession() string {
	return "loggedin"
}
