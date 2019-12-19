package openapi

import (
	"database/sql"
	"errors"
	"fmt"
)

const (
	queryTemplate4PG  = "SELECT %s FROM %s WHERE %s = $1 LIMIT 1"
	insertTemplate4PG = "INSERT INTO %s (%s, %s)VALUES($1, $2)"
)

// default provided sql
type PgSqlSecretKeeper struct {
	Db        *sql.DB // the client to access database
	TableName string  // the table where the secret stores
	KeyCol    string  // the column name of the key
	SecretCol string  // the column name of the secret
}

// get secret from a sql data source
func (s PgSqlSecretKeeper) GetSecret(key string) (string, error) {
	if s.Db == nil {
		return EmptyString, errors.New("db client should not be nil")
	}
	row := s.Db.QueryRow(s.constructQuery(), key)
	var secret string
	err := row.Scan(&secret)
	if err != nil {
		return EmptyString, err
	}
	return secret, nil
}

// construct query for getting secret
func (s PgSqlSecretKeeper) constructQuery() string {
	return fmt.Sprintf(queryTemplate4PG, s.SecretCol, s.TableName, s.KeyCol)
}

func (s PgSqlSecretKeeper) GeneratePair() *KvPair {
	p := KvPair{
		Key:   string(randomStr(keyLen, kindAll)),
		Value: string(randomStr(secretLen, kindAll)),
	}
	// do the insert work
	insertSql := fmt.Sprintf(insertTemplate4PG, s.TableName, s.KeyCol, s.SecretCol)
	stmt, err := s.Db.Prepare(insertSql)
	if err != nil {
		return nil
	}
	r, err := stmt.Exec(p.Key, p.Value)
	if err != nil || r == nil {
		return nil
	}
	defer stmt.Close()
	a, err := r.RowsAffected()
	// check result
	if err != nil || a < 1 {
		return nil
	}
	return &p
}
