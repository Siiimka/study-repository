package main

import (
	"fmt"
	"study/feature1"
	simple_connection "study/feature_postgres"
)

func main() {
	fmt.Println("Hello, GIT!")
	feature1.Feature1()
	simple_connection.CheckConnection()
}
