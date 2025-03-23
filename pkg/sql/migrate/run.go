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
	defer db.Close()

	goose.SetLogger(l)
	goose.SetDialect("postgres")

	if err := goose.Up(db, sourceMigrationUrl); err != nil {
		return fmt.Errorf("couldn't run migrations: %w", err)
	}
	goose.Status(db, sourceMigrationUrl)
	l.Info("migrations up ran successfully")
	return nil
}
