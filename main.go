package GoEntFiberPokeman

import (
	"GoEntFiberPokeman/ent"
	"context"
	"entgo.io/ent/dialect"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
)

func main() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	ctx := context.Background()
	url := fmt.Sprintf(
		"%s:%s@(%s:%s)/%s?parseTime=True", viper.Get("DB_USER"), viper.Get("DB_PASSWORD"),
		viper.Get("DB_HOST"), viper.Get("DB_PORT"), viper.Get("DB_NAME"),
	)
	client, err := ent.Open(dialect.Postgres, url) // connect to MySQL
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
