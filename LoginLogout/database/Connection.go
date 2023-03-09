package database

func Connection() {
	// Initialize Database
	Connect("root:123456789@tcp(localhost:3306)/jwt_demo?parseTime=true")
	Migrate()
}
