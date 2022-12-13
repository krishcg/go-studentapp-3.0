package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoDBConnection(clientOptions *options.ClientOptions) *mongo.Client {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, _ = mongo.Connect(ctx, clientOptions)
	return client
}

// To Get the database
func GetDatabase() *gorm.DB {
	// databasename := "studentpostgresql"
	database := "postgres"
	// databasepassword := "Nitya12$$"
	// databaseuser := "citus"
	// databaseurl := "postgres://" + databaseuser + ":" + databasepassword + "@c." + databasename + ".postgres.database.azure.com:5432/" + databaseuser + "?sslmode=require"
	// databaseurl := "postgres://" + databaseuser + ":" + databasepassword + "@c." + databasename + ".postgres.database.azure.com:5432/" + databaseuser + "?sslmode=require"
	databaseurl := "postgres://citus:Azureadmin123@c.authdb2.postgres.database.azure.com:5432/citus?sslmode=require"
	connection, err := gorm.Open(database, databaseurl)
	if err != nil {
		log.Fatalln("Invalid database url")
	}
	sqldb := connection.DB()
	err = sqldb.Ping()
	if err != nil {
		log.Fatal("Database connected")
	}
	fmt.Println("Database connection successful.")
	return connection
}

//create user table in userdb
func InitialMigration() {
	connection := GetDatabase()
	defer CloseDatabase(connection)
	connection.AutoMigrate(User{})
}

//closes database connection
func CloseDatabase(connection *gorm.DB) {
	sqldb := connection.DB()
	sqldb.Close()
}
