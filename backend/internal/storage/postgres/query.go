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
                   			feature,
							city,
                   			address,
							name,
							img_path,
							description
                   			) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)	
	`
	getId              = "SELECT id FROM events WHERE name = $1 AND date = $2"
	changeImgPath      = "UPDATE events SET img_path=$1"
	deleteEvent        = "DELETE FROM events WHERE id = $1"
	getEventsByFeature = "SELECT * FROM events WHERE data BETWEEN $1 AND $2 AND feature = $3"
)
