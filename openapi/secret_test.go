package openapi

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

const (
	dbAddr = "testuser:123456@tcp(localhost:3306)/cloudb?charset=utf8&parseTime=True&loc=Local"
)

var (
	sqlKeeper = SqlSecretKeeper{
		getdb(),
		"app",
		"app_key",
		"app_secret",
	}
)

func TestSqlSecretKeeper_GetSecret(t *testing.T) {
	keeper := SqlSecretKeeper{
		nil,
		"apps",
		"app_key",
		"app_secret",
	}
	_, err := keeper.GetSecret("thekey")
	assert.NotNil(t, err)
	keeper.Db = getdb()
	_, err = keeper.GetSecret("thekey")
	assert.NotNil(t, err)
	val, _ := sqlKeeper.GetSecret("thekey")
	assert.True(t, len(val) > 0)
}

func TestSqlSecretKeeper_GeneratePair(t *testing.T) {
	sqlKeeper.TableName = "app"
	r := sqlKeeper.GeneratePair()
	assert.NotNil(t, r)
}

func getdb() *sql.DB {
	db, err := sql.Open("mysql", dbAddr)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}
