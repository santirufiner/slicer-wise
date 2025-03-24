package pg

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/santirufiner/slicerwise/internal/infrastructure/config"
	"github.com/sirupsen/logrus"
)

func Connect(l *logrus.Logger, conf config.Pg) *pgxpool.Pool {
	timeout, _ := context.WithTimeout(context.Background(), conf.Timeout)
	db, err := pgxpool.New(timeout, conf.Url)
	if err != nil {
		l.Fatalf("unable to connect to pg database %v", err)
	}

	go func() {
		for {
			timeout, _ := context.WithTimeout(context.Background(), conf.Timeout)
			if err := db.Ping(timeout); err != nil {
				l.Fatalf("pg connection lost %v", err)
			}
			time.Sleep(conf.Heartbeat)
		}
	}()

	l.Info("connected to pg")
	return db
}
