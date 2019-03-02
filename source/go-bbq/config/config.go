package config

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" //wtse-1
	_ "github.com/golang-migrate/migrate/v4/source/file"       //wtse-1
	_ "github.com/lib/pq"                                      //wtse-1

	"github.com/go-redis/cache"
	"github.com/go-redis/redis"
	"github.com/vmihailenco/msgpack"
)

// Config WTSE-1
type Config struct {
	Cache    *cache.Codec
	Database *sql.DB
	UseCache bool
	Port     string
}

// Initialize creates the object and  invokes world::peace()
func (c *Config) Initialize(user, password, dbname, host, redis1, redispw string) {
	connectionString :=
		fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", user, password, host, dbname)

	fmt.Println("Connecting to ", connectionString)

	m, err := migrate.New(
		"file://./migrations",
		connectionString)

	if err != nil {
		fmt.Println(err)
	} else {
		err := m.Up()
		if err != nil {
			fmt.Println(err)
		}
	}
	// nonsense
	c.Database, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	ring := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			redis1: redis1,
		},
		Password:    redispw,
		DialTimeout: time.Second * 30,
	})

	c.Cache = &cache.Codec{
		Redis: ring,
		Marshal: func(v interface{}) ([]byte, error) {
			return msgpack.Marshal(v)
		},
		Unmarshal: func(b []byte, v interface{}) error {
			return msgpack.Unmarshal(b, v)
		},
	}

	c.Port = "21337"
	c.UseCache = false
}
