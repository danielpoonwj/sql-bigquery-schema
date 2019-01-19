package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// Connection : Wrapper for managing interactions with DB
type Connection struct {
	Type        string
	DB          *sql.DB
	QueryString string
	TypeMap     TypeMap
}

// Close : Close Connection
func (c *Connection) Close() error {
	return c.DB.Close()
}

// GetBQSchema : Get BigQuery schema for specific table
func (c *Connection) GetBQSchema(dbName, tableName string) ([]byte, error) {
	tableSchema := NewTableSchema(c.TypeMap)

	rows, err := c.DB.Query(c.QueryString, dbName, tableName)
	if err != nil {
		return nil, err
	}

	isEmpty := true

	for rows.Next() {
		var columnSchema ColumnSchema

		err = rows.Scan(&columnSchema.Name, &columnSchema.Type)
		if err != nil {
			return nil, err
		}

		err = tableSchema.AddColumn(&columnSchema)
		if err != nil {
			return nil, err
		}

		isEmpty = false
	}

	if isEmpty {
		return nil, fmt.Errorf("No columns found for %s.%s", dbName, tableName)
	}

	return tableSchema.ToJSON()
}

// NewConnection : Create new Connection
func NewConnection(connType, username, password, host, port string) (*Connection, error) {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/", username, password, host, port)

	db, err := sql.Open(connType, connStr)
	if err != nil {
		return nil, err
	}

	queryStr, ok := SQLQueryMap[connType]
	if !ok {
		err := fmt.Errorf("Unable to find query for %s", connType)
		return nil, err
	}

	typesMap, ok := SQLTypeMap[connType]
	if !ok {
		err := fmt.Errorf("Unable to find types mapping for %s", connType)
		return nil, err
	}

	return &Connection{
		Type:        connType,
		DB:          db,
		QueryString: queryStr,
		TypeMap:     typesMap,
	}, nil
}
