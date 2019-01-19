package main

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var (
	connType string
	username string
	password string
	host     string

	dbName    string
	tableName string
)

var rootCmd = &cobra.Command{
	Use:   "sql-bigquery-schema",
	Short: "Generate Google BigQuery schema from SQL database tables",
	Long:  "Generate Google BigQuery schema from SQL database tables",
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := NewConnection(connType, username, password, host)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()

		schema, err := conn.GetBQSchema(dbName, tableName)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(schema)
	},
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	rootCmd.Flags().StringVar(&connType, "type", "", "connection type [mysql]")
	rootCmd.MarkFlagRequired("type")

	rootCmd.Flags().StringVar(&username, "username", "", "username")
	rootCmd.MarkFlagRequired("username")

	rootCmd.Flags().StringVar(&password, "password", "", "password")
	rootCmd.MarkFlagRequired("password")

	rootCmd.Flags().StringVar(&host, "host", "", "host")
	rootCmd.MarkFlagRequired("host")

	rootCmd.Flags().StringVar(&dbName, "database", "", "database name")
	rootCmd.MarkFlagRequired("database")

	rootCmd.Flags().StringVar(&tableName, "table", "", "table name")
	rootCmd.MarkFlagRequired("table")
}
