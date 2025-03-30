package main

import (
	"ecommerce-product/cmd"
	"ecommerce-product/helpers"
)

func main() {
	// Load config
	helpers.SetupConfig()

	// Load log 
	helpers.SetupLogger()

	// load db
	helpers.SetupMySQL()

	// load redis 
	// helpers.SetupRedis()

	// run kafka consumer
	// cmd.ServeKafkaConsumer()

	// run http
	cmd.ServeHTTP()
}