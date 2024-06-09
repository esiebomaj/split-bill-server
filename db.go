package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetUpDB(Envs AppEnvs) *gorm.DB {
	fmt.Println("Setting up DB with", Envs)
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=allow TimeZone=Asia/Shanghai",
		Envs.DB_HOST, Envs.DB_USER,
		Envs.DB_PASSWORD, Envs.DB_NAME,
		Envs.DB_PORT,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Group{}, &ListItem{}, &User{})

	return db
}
