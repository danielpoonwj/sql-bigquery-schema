package main

import (
	"encoding/json"
	"fmt"

	"cloud.google.com/go/bigquery"
	bq "google.golang.org/api/bigquery/v2"
)

// TableSchema : Convenience wrapper for interfacing with DB and BQ schemas
type TableSchema struct {
	TypeMap  TypeMap
	BQSchema bigquery.Schema
}

// AddColumn : Add and convert DB to BQ schema column
func (t *TableSchema) AddColumn(c *ColumnSchema) error {
	fieldSchema, err := c.toBQ(t.TypeMap)
	if err != nil {
		return err
	}

	t.BQSchema = append(t.BQSchema, fieldSchema)
	return nil
}

// ToJSON : Generate BQ schema JSON
func (t *TableSchema) ToJSON() ([]byte, error) {
	bqSchema := schemaToBQ(t.BQSchema)
	return json.MarshalIndent(bqSchema, "", "    ")
}

// NewTableSchema : New TableSchema
func NewTableSchema(typeMap TypeMap) *TableSchema {
	return &TableSchema{
		TypeMap: typeMap,
	}
}

// ColumnSchema : Representation of DB column
type ColumnSchema struct {
	Name string
	Type string
}

func (c *ColumnSchema) toBQ(fieldMap TypeMap) (*bigquery.FieldSchema, error) {
	mappedType, ok := fieldMap[c.Type]
	if !ok {
		err := fmt.Errorf("Field %s has unrecognized type: %s", c.Name, c.Type)
		return nil, err
	}

	return &bigquery.FieldSchema{
		Name:     c.Name,
		Repeated: false,
		Required: false,
		Type:     mappedType,
	}, nil
}

// fieldSchemaToBQ : Convert between internal formats - adapted from private method
// https://github.com/googleapis/google-cloud-go/blob/105f0564f8d67e66e7ea5ecc5a6e46dad440aa09/bigquery/schema.go#L53-L71
func fieldSchemaToBQ(fs *bigquery.FieldSchema) *bq.TableFieldSchema {
	tfs := &bq.TableFieldSchema{
		Description: fs.Description,
		Name:        fs.Name,
		Type:        string(fs.Type),
	}

	if fs.Repeated {
		tfs.Mode = "REPEATED"
	} else if fs.Required {
		tfs.Mode = "REQUIRED"
	} // else leave as default, which is interpreted as NULLABLE.

	for _, f := range fs.Schema {
		tfs.Fields = append(tfs.Fields, fieldSchemaToBQ(f))
	}

	return tfs
}

// schemaToBQ : Convert between internal formats - adapted from private method
// https://github.com/googleapis/google-cloud-go/blob/105f0564f8d67e66e7ea5ecc5a6e46dad440aa09/bigquery/schema.go#L73-L79
func schemaToBQ(s bigquery.Schema) []*bq.TableFieldSchema {
	var fields []*bq.TableFieldSchema
	for _, f := range s {
		fields = append(fields, fieldSchemaToBQ(f))
	}
	return fields
}
