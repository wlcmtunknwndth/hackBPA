package postgres

const (
	getPassword  = "SELECT password FROM auth WHERE username = $1"
	registerUser = "INSERT INTO auth(username, password) VALUES($1, $2)"
	isAdmin      = "SELECT isadmin FROM auth WHERE username = $1"
	deleteUser   = "DELETE FROM auth WHERE username = $1"
)
