package pg

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/santirufiner/slicerwise/internal/infrastructure/config"
	"github.com/sirupsen/logrus"
)

func Connect(l *logrus.Logger, conf config.Pg) *pgxpool.Pool {
	ctx, cancel := context.WithTimeout(context.Background(), conf.Timeout)
	defer cancel()
	db, err := pgxpool.New(ctx, conf.Url)
	if err != nil {
		l.Fatalf("unable to connect to pg database %v", err)
	}

	go func() {
		for {
			pingCtx, pingCancel := context.WithTimeout(context.Background(), conf.Timeout)
			if err := db.Ping(pingCtx); err != nil {
				l.Fatalf("pg connection lost %v", err)
			}
			pingCancel()
			time.Sleep(conf.Heartbeat)
		}
	}()

	l.Info("connected to pg")
	return db
}
