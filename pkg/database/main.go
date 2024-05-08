package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jpillora/backoff"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBOptions struct {
	Driver        string
	User          string
	Password      string
	Name          string
	Host          string
	Port          string
	SlowThreshold string
	Colorful      string
}

func New(dbOptions *DBOptions) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		dbOptions.Host, dbOptions.User, dbOptions.Password, dbOptions.Name, dbOptions.Port,
	)
	dialector := postgres.New(postgres.Config{
		DSN: dsn,
	})
	b := &backoff.Backoff{
		Factor: 1.5,
		Min:    1 * time.Second,
		Max:    32 * time.Second,
	}
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Duration(1000) * time.Millisecond,
			Colorful:      false,
		},
	)
	for {
		db, err := gorm.Open(dialector, &gorm.Config{
			CreateBatchSize: 100,
			NowFunc: func() time.Time {
				return time.Now().UTC()
			},
			Logger: newLogger,
		})
		if err != nil {
			d := b.Duration()
			fmt.Printf("%s, reconnecting in %s", err, d)
			if d == b.Max {
				panic(err)
			}
			time.Sleep(d)

			continue
		}

		ddb, err := db.DB()
		if err != nil {
			panic(err)
			// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
		}
		ddb.SetMaxIdleConns(20)
		// SetMaxOpenConns sets the maximum number of open connections to the database.
		ddb.SetMaxOpenConns(150)
		// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
		ddb.SetConnMaxLifetime(time.Minute * 5)

		return db.Debug()
	}
}
