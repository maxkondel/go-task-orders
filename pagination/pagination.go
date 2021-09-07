package pagination

import "time"

// Database settings
const (
	host     = "localhost"
	port     = 5432 // Default port
	user     = "postgres"
	password = "password"
	dbname   = "fiber_task"
)

// Order struct
type Order struct {
	ID       int64
	Items    []string
	CreateAt time.Time
}

// Employees struct
type Orders struct {
	Orders []Order `json:"orders"`
}

// Connect function
func Connect() error {
	// var err error
	// db, err = sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname))
	// if err != nil {
	// 	return err
	// }
	// if err = db.Ping(); err != nil {
	// 	return err
	// }
	return nil
}
