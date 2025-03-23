package main

import (
	"github.com/santirufiner/slicer-wise/internal/infrastructure/api"
	"github.com/santirufiner/slicer-wise/internal/infrastructure/config"
	"github.com/santirufiner/slicer-wise/internal/infrastructure/pg"
	"github.com/santirufiner/slicer-wise/pkg/sql/migrate"
	"github.com/sirupsen/logrus"
)

func main() {
	conf := config.Load()

	l := logrus.StandardLogger()
	level, _ := logrus.ParseLevel(conf.LogLevel)
	l.SetLevel(level)

	if conf.Pg.RunMigration {
		if err := migrate.Run(conf.Pg.Url, conf.Pg.SourceMigration, l); err != nil {
			l.Panicf("could not run migration database: %v", err)
		}
	}
	db := pg.Connect(l, conf.Pg)
	defer db.Close()

	l.Fatal(api.Run(conf.Port, l))
}
