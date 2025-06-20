package migrate

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
	"github.com/sirupsen/logrus"
)

func Run(connectionString string, sourceMigrationUrl string, l *logrus.Logger) error {
	db, err := sql.Open("pgx", connectionString)
	if err != nil {
		return fmt.Errorf("couldn't open db: %w", err)
	}
	defer func() {
		if err = db.Close(); err != nil {
			l.Errorf("failed to close db connection: %v", err)
		}
	}()
	goose.SetLogger(l)
	err = goose.SetDialect("postgres")
	if err != nil {
		return err
	}
	if err = goose.Up(db, sourceMigrationUrl); err != nil {
		return fmt.Errorf("couldn't run migrations: %w", err)
	}
	err = goose.Status(db, sourceMigrationUrl)
	if err != nil {
		return err
	}
	l.Info("migrations up ran successfully")
	return nil
}
