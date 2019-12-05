/// @Author winjeg,  winjeg@qq.com
/// All rights reserved to winjeg

package settings

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/winjeg/go-commons/cache"
	"strings"
	"sync"

	"github.com/winjeg/go-commons/log"
	"github.com/winjeg/go-commons/uid"
)

var getSql, getSqlPg, updateSql, updateSqlPg, existSql, existSqlPg, addSql, addSqlPg, addSqlWithId, addSqlWithIdPg, deleteVarSql, deleteVarSqlPg, settingsSql, settingVar, createSettingsTableSql, descSettingsSql, nameCol, valCol string

var (
	logger           = log.GetLogger(nil)
	db       *sql.DB = nil
	withId           = false
	postgres         = false
)

// generate primary key, with this function
// if using databases that won't automatically generated primary key
// this function may suit you, but you must make the primary key at lease 8 byte long

func intSql(tableName string) {
	getSql = fmt.Sprintf("SELECT value FROM %s WHERE name = ?", tableName)
	getSqlPg = fmt.Sprintf("SELECT value FROM %s WHERE name = $1", tableName)
	updateSql = fmt.Sprintf("UPDATE %s SET value = ? WHERE name = ?", tableName)
	updateSqlPg = fmt.Sprintf("UPDATE %s SET value = $1 WHERE name = $2", tableName)
	existSql = fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE name= ?", tableName)
	existSqlPg = fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE name= $1", tableName)
	addSql = fmt.Sprintf("INSERT IGNORE INTO %s(guid,name, value) VALUE(?, ?, ?)", tableName)
	addSqlPg = fmt.Sprintf("INSERT INTO %s(guid, name, value) VALUES($1, $2, $3)", tableName)
	addSqlWithId = fmt.Sprintf("INSERT IGNORE INTO %s(guid, name, value) VALUE(?, ?, ?)", tableName)
	addSqlWithIdPg = fmt.Sprintf("INSERT INTO %s(guid, name, value) VALUES($1, $2, $3)", tableName)
	deleteVarSql = fmt.Sprintf("DELETE FROM %s WHERE name = ?", tableName)
	deleteVarSqlPg = fmt.Sprintf("DELETE FROM %s WHERE name = $1", tableName)
	settingsSql = fmt.Sprintf("SELECT 1 FROM %s", tableName)
	settingVar = "1"
	createSettingsTableSql = fmt.Sprintf("CREATE TABLE `%s` ( `guid` int(64) unsigned NOT NULL AUTO_INCREMENT COMMENT 'pk', `name` varchar(200) COLLATE utf8_bin NOT NULL COMMENT 'varname', `value` text COLLATE utf8_bin NOT NULL, PRIMARY KEY (`guid`), UNIQUE KEY `name_UNIQUE` (`name`), KEY `idx_name` (`name`)) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_bin", tableName)
	descSettingsSql = fmt.Sprintf("SELECT * FROM %s LIMIT 1", tableName)
	nameCol = "name"
	valCol = "value"
}

func InitV2(dbConn *sql.DB, autoGenerateId, pg bool) error {
	return InitV3(dbConn, autoGenerateId, pg, "settings")
}

func InitV3(dbConn *sql.DB, autoGenerateId, pg bool, tableName string) error {
	intSql(tableName)
	err := Init(dbConn)
	if err != nil {
		return err
	}
	withId = autoGenerateId
	postgres = pg
	return nil
}

func Init(dbConn *sql.DB) error {
	intSql("settings")
	if dbConn == nil {
		return errors.New("db connection should not be nil")
	}
	if db != nil {
		return errors.New("already initialized")
	}
	var lock sync.Mutex
	lock.Lock()
	defer lock.Unlock()

	// if table not exists create the table for them
	if !tableExists(dbConn) {
		// table not exists
		_, err := dbConn.Exec(createSettingsTableSql)
		if err != nil {
			return err
		}
	}
	rows, descError := dbConn.Query(descSettingsSql)
	if descError != nil {
		return descError
	}
	cols, _ := rows.Columns()
	if !contains(cols, nameCol, valCol) {
		return errors.New("table structure for settings is not supported")
	}
	db = dbConn
	return nil
}

func tableExists(dbConn *sql.DB) bool {
	row := dbConn.QueryRow(settingsSql)
	var re string
	err := row.Scan(&re)
	if err != nil && strings.EqualFold(err.Error(), "sql: no rows in result set") {
		return true
	}
	return err == nil && strings.EqualFold(re, settingVar)
}

func GetVar(name string) string {
	if v := cache.Get(name); v != nil {
		return v.(string)
	} else {
		var x string
		if postgres {
			r := db.QueryRow(getSqlPg, name)
			err := r.Scan(&x)
			if err != nil {
				return ""
			}
		} else {
			r := db.QueryRow(getSql, name)
			err := r.Scan(&x)
			if err != nil {
				return ""
			}
		}

		var lock sync.Mutex
		lock.Lock()
		cache.Set(name, x, 1000*60*30)
		lock.Unlock()
		return x
	}
}

// set variable and update cache
func SetVar(name, value string) {
	var lock sync.Mutex
	lock.Lock()
	defer lock.Unlock()
	cache.Set(name, value, 1000*60*30)
	var row *sql.Row
	if postgres {
		row = db.QueryRow(existSqlPg, name)
	} else {
		row = db.QueryRow(existSql, name)
	}

	exists := 0
	if err := row.Scan(&exists); err == nil && exists == 0 {
		if withId {
			id := uid.NextID()
			if postgres {
				stmt, err := db.Prepare(addSqlWithIdPg)
				if err != nil || stmt == nil {
					logger.Error(err)
				}
				defer stmt.Close()
				_, execErr := stmt.Exec(id, name, value)
				if execErr != nil {
					logger.Error(err)
				}
			} else {
				_, execErr := db.Exec(addSqlWithId, id, name, value)
				if execErr != nil {
					logger.Error(err)
				}
			}

		} else {
			if postgres {
				stmt, err2 := db.Prepare(addSqlPg)
				if err2 != nil || stmt == nil {
					return
				}
				_, err = stmt.Exec(name, value)
				defer stmt.Close()
			} else {
				_, err = db.Exec(addSql, uid.NextID(), name, value)
			}
		}
		if err != nil {
			logger.Error(err)
		}
	} else {
		if postgres {
			stmt, err := db.Prepare(updateSqlPg)
			if err != nil {
				logger.Error(err)
				return
			}
			_, execErr := stmt.Exec(value, name)
			if execErr != nil {
				logger.Error(execErr)
			}
			defer stmt.Close()

		} else {
			_, err = db.Exec(updateSql, value, name)
			if err != nil {
				logger.Error(err)
			}
		}
	}
}

func DelVar(name string) {
	var lock sync.Mutex
	lock.Lock()
	defer lock.Unlock()
	cache.Set(name, nil, 100)
	if postgres {
		stmt, err := db.Prepare(deleteVarSqlPg)
		if err != nil {
			logger.Error(err)
			return
		}
		defer stmt.Close()
		_, execErr := stmt.Exec(name)
		if execErr != nil {
			logger.Error(execErr)
		}
	} else {
		_, err := db.Exec(deleteVarSql, name)
		if err != nil {
			logger.Error(err)
		}
	}
}

func contains(collection []string, elements ...string) bool {
	if len(elements) == 0 {
		return true
	}
	if len(collection) == len(elements) && len(collection) == 0 {
		return true
	}
	if len(elements) == 0 && len(collection) != 0 {
		return true
	}
	if len(collection) == 0 && len(elements) != 0 {
		return false
	}
	// put elements to map
	elementMap := make(map[string]bool, len(elements))
	for _, v := range elements {
		elementMap[v] = true
	}
	count := 0
	for _, v := range collection {
		if elementMap[v] {
			count++
		}
	}
	return count >= len(elementMap)
}
