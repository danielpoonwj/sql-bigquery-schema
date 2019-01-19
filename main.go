package main

import (
	"io/ioutil"
	"log"

	"github.com/spf13/cobra"
)

var (
	connType string
	username string
	password string
	host     string
	port     string

	dbName    string
	tableName string

	outputPath string
)

var rootCmd = &cobra.Command{
	Use:   "sql-bigquery-schema",
	Short: "Generate Google BigQuery schema from SQL database tables",
	Long:  "Generate Google BigQuery schema from SQL database tables",
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := NewConnection(connType, username, password, host, port)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()

		schema, err := conn.GetBQSchema(dbName, tableName)
		if err != nil {
			log.Fatal(err)
		}

		err = ioutil.WriteFile(outputPath, schema, 0644)
		if err != nil {
			log.Fatal(err)
		}
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

	rootCmd.Flags().StringVar(&port, "port", "3306", "host")

	rootCmd.Flags().StringVar(&dbName, "database", "", "database name")
	rootCmd.MarkFlagRequired("database")

	rootCmd.Flags().StringVar(&tableName, "table", "", "table name")
	rootCmd.MarkFlagRequired("table")

	rootCmd.Flags().StringVarP(&outputPath, "output", "o", "", "output file path")
	rootCmd.MarkFlagRequired("output")
}
