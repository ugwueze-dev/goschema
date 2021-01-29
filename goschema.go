package goschema

import (
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql" // import mysql driver
	"github.com/jmoiron/sqlx"
)

// DBConfig is the database configuration struct
type DBConfig struct {
	Name     string // database name
	Host     string // database server host
	Username string // database server username
	Password string // database server password
}

type GoSchema struct {
	tx *sqlx.Tx

	dbConfig      DBConfig
	purgeDatabase bool
	tables        []*table
}

const (
	databaseDriver = "mysql"
)

// New initializes and returns an instance of GoSchema
func NewSchema(config DBConfig, purgeDatabase bool) *GoSchema {
	return &GoSchema{
		dbConfig:      config,
		purgeDatabase: purgeDatabase,
	}
}

// NewTable returns a database table instance
func (s *GoSchema) NewTable(tableName string) *table {
	table := newTable(tableName)
	s.tables = append(s.tables, table)
	return table
}

// Create
func (s *GoSchema) Create() error {
	db, err := s.connect()
	if err != nil {
		return err
	}

	s.tx, err = db.Beginx()
	if err != nil {
		return fmt.Errorf("error starting transaction: %s", err.Error())
	}

	err = s.create()
	if err != nil {
		s.tx.Rollback()
	}

	return err
}

func (s *GoSchema) connect() (*sqlx.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", s.dbConfig.Username, s.dbConfig.Password, s.dbConfig.Host, s.dbConfig.Name)
	db, err := sqlx.Open(databaseDriver, dsn)
	if err != nil {
		return nil, fmt.Errorf("error opening database connection: %s", err.Error())
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error pinging database server: %s", err.Error())
	}

	return db, nil
}

func (s *GoSchema) createDatabase() error {
	str := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s`", s.dbConfig.Name)
	if _, err := s.tx.Exec(str); err != nil {
		return fmt.Errorf("error creating database '%s': %s", s.dbConfig.Name, err.Error())
	}

	return nil
}

func (s *GoSchema) dropDatabase() error {
	str := fmt.Sprintf("DROP DATABASE IF EXISTS `%s`", s.dbConfig.Name)
	if _, err := s.tx.Exec(str); err != nil {
		return fmt.Errorf("error dropping database '%s': %s", s.dbConfig.Name, err.Error())
	}
	return nil
}

func (s *GoSchema) useDatabase() error {
	str := fmt.Sprintf("USE  `%s`", s.dbConfig.Name)
	if _, err := s.tx.Exec(str); err != nil {
		return fmt.Errorf("error using database '%s': %s", s.dbConfig.Name, err.Error())
	}

	return nil
}

func (s *GoSchema) dropTable(tableName string) error {
	str := fmt.Sprintf("DROP TABLE IF EXISTS `%s`", tableName)
	if _, err := s.tx.Exec(str); err != nil {
		return fmt.Errorf("error dropping table '%s': %s", tableName, err.Error())
	}

	return nil
}

func (s *GoSchema) create() error {
	var err error

	if s.purgeDatabase {
		err = s.dropDatabase()
		if err != nil {
			return err
		}

		err = s.createDatabase()
		if err != nil {
			return err
		}
	}

	if err = s.useDatabase(); err != nil {
		return err
	}

	if err = s.createTables(); err != nil {
		return err
	}

	if err := s.tx.Commit(); err != nil {
		return fmt.Errorf("error committing database schema creation transaction: %s", err.Error())
	}

	return nil
}

func (s *GoSchema) createTables() error {
	var err error
	for _, table := range s.tables {
		if err = s.dropTable(table.name); err != nil {
			return err
		}

		queryStr := fmt.Sprintf("CREATE TABLE `%s` (%s) ENGINE=InnoDB DEFAULT CHARSET=latin1;", table.name, s.getColumnQuery(table.columns))
		if _, err = s.tx.Exec(queryStr); err != nil {
			return fmt.Errorf("error executing create table transaction: %s", err.Error())
		}
	}

	return nil
}

func (s *GoSchema) getColumnQuery(columns []*column) string {
	var primaryKeyStr, queryString string
	var keys, references []string

	for _, column := range columns {
		nullStatus := "NOT NULL"
		defaultValue := ""
		autoIncrement := ""

		if column.nullable {
			nullStatus = "NULL"
		}

		if column.isPrimaryKey {
			autoIncrement = "AUTO_INCREMENT"
			primaryKeyStr = fmt.Sprintf("PRIMARY KEY (`%s`),\n", column.name)
		}

		if column.defaultValue != "" {
			defaultValue = fmt.Sprintf("DEFAULT %s", column.defaultValue)
		}

		if column.key != "" {
			key := fmt.Sprintf("KEY `%s` (`%s`),\n", column.key, column.name)
			keys = append(keys, key)
		}

		queryString += s.getColumnQueryString(column, nullStatus, defaultValue, autoIncrement)

		for _, reference := range column.references {
			refStr := fmt.Sprintf("FOREIGN KEY (`%s`) REFERENCES `%s` (`%s`)", column.name, reference.column.tableName, reference.column.name)

			if reference.onDelete != "" {
				refStr += " ON DELETE " + reference.onDelete.String()
			}

			if reference.onUpdate != "" {
				refStr += " ON UPDATE " + reference.onUpdate.String()
			}

			refStr += ",\n"
			references = append(references, refStr)
		}

	}

	if primaryKeyStr != "" {
		queryString += primaryKeyStr
	}

	for _, key := range keys {
		queryString += key
	}

	for _, reference := range references {
		queryString += reference
	}

	return strings.TrimSuffix(queryString, ",\n")
}

func (s *GoSchema) getColumnQueryString(column *column, nullStatus, defaultValue, autoIncrement string) string {
	var queryString string
	vals := []interface{}{
		column.name,
		column.dataType.String(),
	}

	if column.setSize {
		queryString = "%s %s(%d) %s %s %s,\n"
		vals = append(vals, column.size)
	} else if column.length != 0 || column.numDecimals != 0 {
		queryString = "%s %s(%d,%d) %s %s %s,\n"
		vals = append(vals, column.length)
		vals = append(vals, column.numDecimals)
	} else {
		queryString = "%s %s %s %s %s,\n"
	}

	vals = append(vals, nullStatus)
	vals = append(vals, defaultValue)
	vals = append(vals, autoIncrement)
	return fmt.Sprintf(queryString, vals...)
}
