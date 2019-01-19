# sql-bigquery-schema

## Introduction

This project aims to do one thing - generate BigQuery JSON schema files from SQL databases. The most basic use case involves mirroring an existing SQL table to BigQuery. This project does not facilitate the rest of the transfer process, granting freedom for users to choose the approach for extracting and uploading data.

## Compatibility

Currently only MySQL is supported, but it can be extended to other dialects in the future.

## Development Requirements

`go` >= 1.11

## Examples

Connect to MySQL instance on localhost and generate the JSON schema file for the `users`.`contacts` table.

```
./sql-bigquery-schema \
    --type="mysql" \
    --username="foo" \
    --password="bar" \
    --host="127.0.0.1" \
    --port="3306" \
    --database="users" \
    --table="contacts" \
    --output="./schemas/users_contacts.schema"
```
