package model

func migration() {
	// Migrate the schema
	DB.AutoMigrate(&Prize{}).
		AutoMigrate(&User{})
}