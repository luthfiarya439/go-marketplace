package main

import "go-marketplace/config"

func init() {
	config.LoadEnv()
	config.ConnectDatabase()
}

func main() {
	// router := gin.Default()
}
