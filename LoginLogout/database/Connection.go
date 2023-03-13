package database

func Connection() {
	// Initialize Database
	// Connect("root:123456789@tcp(localhost:3306)/jwt_demo?parseTime=true")
	Connect("postgres://postgres:123456789@localhost:5432/jwt_demo?sslmode=disable")
	Migrate()
}
