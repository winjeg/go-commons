/// @Author winjeg,  winjeg@qq.com
/// All rights reserved to winjeg

package settings

import (
	"database/sql"
	"sync"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	a := []string{"1", "2", "3"}
	b := []string{"1"}
	c := []string{}
	assert.True(t, contains(a, "1"))
	assert.True(t, contains(a, c...))
	assert.True(t, contains(a, b...))
	assert.True(t, !contains(a, "4"))
	assert.True(t, !contains(a, "3", "4"))
}

func TestSettings(t *testing.T) {
	dbConn := getDb()
	err := Init(dbConn)
	assert.Nil(t, err)
	SetVar("a", "b")
	assert.Equal(t, GetVar("a"), "b")
	SetVar("a", "bc")
	assert.Equal(t, GetVar("a"), "bc")
	DelVar("a")
}

var (
	once sync.Once
)

const (
	testAddr = "testuser:123456@tcp(localhost:3306)/cloudb?charset=utf8&parseTime=True&loc=Local"
)

func getDb() *sql.DB {
	var myDb *sql.DB
	once.Do(func() {
		db, err := sql.Open("mysql", testAddr)
		checkErr(err)
		db.SetMaxIdleConns(int(3))
		db.SetMaxOpenConns(int(10))
		pingErr := db.Ping()
		if pingErr != nil {
			panic(pingErr)
		}
		myDb = db
	})
	return myDb
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
