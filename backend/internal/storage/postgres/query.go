package postgres

const (
	// AUTH
	getPassword  = "SELECT password FROM auth WHERE username = $1"
	registerUser = "INSERT INTO auth(username, password, gender, age) VALUES($1, $2, $3, $4)"
	isAdmin      = "SELECT isadmin FROM auth WHERE username = $1"
	deleteUser   = "DELETE FROM auth WHERE username = $1"

	//Event
	getEvent    = "SELECT * FROM events WHERE id = $1"
	createEvent = `INSERT INTO events(
							price,
							restrictions,
							date,
							location,
							name,
							img_path,
							description,
							disability,
							deaf,
							blind,
							neural
                   			) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)	
	`
	getId       = "SELECT * FROM events WHERE name = $1 AND date = $2"
	deleteEvent = "DELETE FROM events WHERE id = $1"
)
