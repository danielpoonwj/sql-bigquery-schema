package main

import (
	"cloud.google.com/go/bigquery"
)

// SQLQueryMap : Mapping to queries for table information by connection type
var SQLQueryMap = map[string]string{
	"mysql": `
		SELECT column_name, UPPER(data_type)
		FROM information_schema.columns
		WHERE table_schema=? AND table_name=?;
	`,
}

// TypeMap : Mapping of db type to BigQuery type
type TypeMap map[string]bigquery.FieldType

// SQLTypeMap : Mapping to TypeMap by connection type
var SQLTypeMap = map[string]TypeMap{
	"mysql": {
		"INTEGER":   bigquery.IntegerFieldType,
		"INT":       bigquery.IntegerFieldType,
		"SMALLINT":  bigquery.IntegerFieldType,
		"TINYINT":   bigquery.IntegerFieldType,
		"MEDIUMINT": bigquery.IntegerFieldType,
		"BIGINT":    bigquery.IntegerFieldType,

		"DECIMAL": bigquery.FloatFieldType,
		"NUMERIC": bigquery.FloatFieldType,
		"FLOAT":   bigquery.FloatFieldType,
		"DOUBLE":  bigquery.FloatFieldType,

		"DATETIME":  bigquery.TimestampFieldType,
		"TIMESTAMP": bigquery.TimestampFieldType,

		"CHAR":    bigquery.StringFieldType,
		"VARCHAR": bigquery.StringFieldType,
		"BLOB":    bigquery.StringFieldType,
		"TEXT":    bigquery.StringFieldType,
	},
}
