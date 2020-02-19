package db

import (
	"database/sql"
	"github.com/spf13/viper"

	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
)

func RunMigration() {
	fmt.Println("running migrations")
	db, err := sql.Open("mysql", viper.GetString("mysql.user")+":"+viper.GetString("mysql.password")+"@tcp("+viper.GetString("mysql.host")+":"+viper.GetString("mysql.port")+")/"+viper.GetString("mysql.database"))
	if err != nil {
		//panic(err)
	}
	driver, _ := mysql.WithInstance(db, &mysql.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations",
		"mysql",
		driver,
	)
	if err != nil {
		fmt.Println(err)
	}

	err = m.Steps(1)
	if err != nil {
		//panic(err)
	}
}
