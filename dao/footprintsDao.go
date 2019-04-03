package dao

import (
	"database/sql"
	"fmt"

	"github.com/brettman/nyc_footprints/config"
	// not sure why this doesn't belong here.
	_ "github.com/lib/pq"
)

const connectionFormat = "host=%s port=%d user=%s password=%s dbname=%s sslmode=disable"

// FootprintsDao - our db
type FootprintsDao struct {
	connstr string
}

// NewDao - initialze a configured dao
func NewDao() *FootprintsDao {
	dbcfg := config.Config{}
	dbcfg.Read()
	constr := fmt.Sprintf(connectionFormat,
		dbcfg.Host,
		dbcfg.Port,
		dbcfg.User,
		dbcfg.Password,
		dbcfg.Dbname,
	)

	thedao := &FootprintsDao{
		connstr: constr,
	}

	fmt.Println(thedao.connstr)
	return thedao
}

// Connect - try to connect to the db
func (d *FootprintsDao) Connect() *sql.DB {
	db, err := sql.Open("postgres", d.connstr)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}
	return db
}
