package mysql

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

type Storage struct {
	db  *sql.DB
	log *log.Logger
}

// MustSetup opens mysql database or panics in negative case
func MustSetup(logger *log.Logger) (s *Storage) {
	const op = "storage.mysql.MustSetup"
	var err error
	s = &Storage{
		log: logger,
	}

	dbPath := fmt.Sprintf(
		"%v:%v@/%v",
		viper.GetString("storage.mysql.user"),
		viper.GetString("storage.mysql.password"),
		viper.GetString("storage.mysql.dbname"),
	)

	s.db, err = sql.Open("mysql", dbPath)
	if err != nil {
		s.log.Panicf("%v: Mysql opening error %v", op, err)
	}

	if err = s.Deploy(); err != nil {
		s.log.Printf("%v: Mysql deploying error %v\n", op, err)
		return
	}
	s.log.Println("Mysql successfully deployed")
	return
}

func (s *Storage) Deploy() (err error) {
	const op = "storage.mysql.Deploy"

	for _, command := range viper.GetStringSlice("postgres.deployment") {
		if _, err = s.db.Exec(command); err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}
	}
	return
}
