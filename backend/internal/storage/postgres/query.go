package postgres

const (
	getPassword  = `SELECT * FROM Auth WHERE username = $1`
	registerUser = `INSERT INTO Auth(username, password) VALUES($1, $2)`
	isAdmin      = `SELECT isAdmin FROM Auth WHERE username = $1`
	deleteUser   = `DELETE FROM Auth WHERE username = $1`
)
