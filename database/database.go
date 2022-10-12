package database

import (
	"GoEntFiberPokeman/ent"
	"context"
	"entgo.io/ent/dialect"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
)

var DbClient *ent.Client

func init() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	//fmt.Println("Connected to database successfully")
	DSN := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s",
		viper.Get("DB_HOST"), viper.Get("DB_PORT"), viper.Get("DB_USER"), viper.Get("DB_NAME"),
		viper.Get("DB_PASSWORD"),
	)
	Client, err := ent.Open(dialect.Postgres, DSN)
	if err != nil {
		log.Fatalf("failed opening connection to DB: %v", err)
	}
	fmt.Println("Connected to database successfully")
	//defer Client.Close()
	//return Client
	if err := Client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	DbClient = Client
}
