package repository

import (
	"context"
	"database/sql"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"strings"
	"time"
	"weather/config"

	_ "github.com/lib/pq"
)

const DBPATH = "./database/schema.sql"

func LoadMongo(db config.MongoDB) (*mongo.Collection, error) {
	uri := fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority", db.User, db.Password, db.Host)

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)

	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	collection := client.Database(db.Name).Collection("weather")

	return collection, nil
}

func LoadDB(dbConfig config.DB) (*sql.DB, error) {
	db, err := sql.Open(dbConfig.DriverName, dbConfig.DataSourceName)
	if err != nil {
		return db, err
	}

	if err := db.Ping(); err != nil {
		return db, err
	}

	if err := createTable(db); err != nil {
		return db, err
	}

	return db, err
}

func createTable(db *sql.DB) error {
	fileSql, err := os.ReadFile(DBPATH)
	if err != nil {
		return err
	}

	requests := strings.Split(string(fileSql), ";")
	for _, request := range requests {
		_, err = db.Exec(request)
		if err != nil {
			return err
		}
	}

	return nil
}
