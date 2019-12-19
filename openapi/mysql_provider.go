package openapi

import (
	"database/sql"
	"errors"
	"fmt"
)

/**
# the sql to create the table app
CREATE TABLE `app` (
  `app_key` varchar(32) NOT NULL,
  `app_secret` varchar(128) NOT NULL,
  `created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`app_key`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8
*/

const (
	queryTemplate  = "SELECT `%s` FROM `%s` WHERE `%s` = ? LIMIT 1"
	insertTemplate = "INSERT INTO `%s`(`%s`, `%s`)VALUE(?, ?)"
)

// default provided sql
type SqlSecretKeeper struct {
	Db        *sql.DB // the client to access database
	TableName string  // the table where the secret stores
	KeyCol    string  // the column name of the key
	SecretCol string  // the column name of the secret
}

// get secret from a sql data source
func (s SqlSecretKeeper) GetSecret(key string) (string, error) {
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
func (s SqlSecretKeeper) constructQuery() string {
	return fmt.Sprintf(queryTemplate, s.SecretCol, s.TableName, s.KeyCol)
}

func (s SqlSecretKeeper) GeneratePair() *KvPair {
	p := KvPair{
		Key:   string(randomStr(keyLen, kindAll)),
		Value: string(randomStr(secretLen, kindAll)),
	}
	// do the insert work
	insertSql := fmt.Sprintf(insertTemplate, s.TableName, s.KeyCol, s.SecretCol)
	r, err := s.Db.Exec(insertSql, p.Key, p.Value)
	if err != nil || r == nil {
		return nil
	}
	a, err := r.RowsAffected()
	// check result
	if err != nil || a < 1 {
		return nil
	}
	return &p
}
