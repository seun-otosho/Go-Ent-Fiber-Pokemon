package main

import (
	"GoEntFiberPokeman/ent"
	"GoEntFiberPokeman/ent/ogent"
	"context"
	"database/sql"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"flag"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

type handler struct {
	*ogent.OgentHandler
	db *sql.DB
}

func (h handler) DBHealth(_ context.Context) (ogent.DBHealthRes, error) {
	if err := h.db.Ping(); err != nil {
		return &ogent.DBHealthServiceUnavailable{}, nil
	}
	return &ogent.DBHealthNoContent{}, nil
}

func main() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	ctx := context.Background()

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s",
		viper.Get("DB_HOST"), viper.Get("DB_PORT"), viper.Get("DB_USER"), viper.Get("DB_NAME"), viper.Get("DB_PASSWORD"),
	)

	var args struct {
		Addr string
		DSN  string
	}

	flag.StringVar(&args.Addr, "addr", ":8080", "http address to listen")
	flag.StringVar(&args.DSN, "dsn", "file:ent?mode=memory&cache=shared&_fk=1", "dsn of database")
	flag.Parse()

	// Create ent client.
	drv, err := entsql.Open(dialect.Postgres, dsn)
	if err != nil {
		log.Fatal(err)
	}

	client, err := ent.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// Start listen to incoming requests.
	fmt.Println("Server starting")
	defer fmt.Println("Server stopped")

	srv, _ := ogent.NewServer(handler{
		OgentHandler: ogent.NewOgentHandler(client),
		db:           drv.DB(),
	})

	if err := http.ListenAndServe(args.Addr, srv); err != nil {
		log.Fatal(err)
	}
}
