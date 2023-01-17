package main

import (
	"postgres/gormPostgres/connection"
	"postgres/gormPostgres/handler"
)

func main() {
	connection.Connection()
	handler.Handler()

}
