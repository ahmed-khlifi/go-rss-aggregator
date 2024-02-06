# Multithread RSS posts scrapper with GO language

This little GO project is a small server with `go-chi` & `sqlc` that make users post RSS links as `feeds`, and then we have a worker that scrape the content of these feed URLs and save theme as `Posts` in a `Postgresql` database.
We use `GoRoutines` to handle `Concurrency` and we can specify how much time to wait before each request etc..

## How to use this project?

First You must install GO from [here](https://golang.org/dl/).

`Postgersql` is The database used for this project

Create a `.env` file in the root directory and add your own values like this :
PORT=[your_port]
DB_URL=[database_connection_string]

To handle DB Migration We used `Goose`, [see more](https://github.com/pressly/goose)

```
go install github.com/pressly/goose/v3/cmd/goose@latest
```

To interact with DB we used `sqlc`, [see more](https://github.com/sqlc-dev/sqlc)

```
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```

## Run The project

run `go run .`

## Build the executable

run `go build`

To run the migration in this file, go to sql/schema & run :

```
postgres://Pepolls:ahmed@localhost:5432/rssagg
```

Generate queries, In home folder Run :

```
sqlc generate
```
