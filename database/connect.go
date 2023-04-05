package database

import (
	"CRM-test/structures"
	"CRM-test/utils"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var Link *sql.DB

func Connection(config *structures.Setting) {
	var e error
	Link, e = sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DbHost,
		config.DbPort,
		config.DbUser,
		config.DbPass,
		config.DbName))
	if e != nil {
		utils.Logger.Println(e)
		return
	}

	e = Link.Ping()
	if e != nil {
		utils.Logger.Println(e)
		return
	}

	errors := make([]string, 0)

	errors = append(errors, prepareUser()...)
	//errors = append(errors, prepareCredit()...)

	if len(errors) > 0 {
		for _, i := range errors {
			utils.Logger.Println(i)
		}
	}

	LoadSession(sessionMap)
}
