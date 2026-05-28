package main

import (
	"context"
	"fmt"
	simple_connection "study/feature_postgres"
	"study/simple_sql"
)

func main() {
	ctx := context.Background()
	conn, err := simple_connection.CheckConnection(ctx)
	if err != nil {
		panic(err)
	}

	// if err := simple_sql.CreateTable(ctx, conn); err != nil {
	// 	panic(err)
	// }

	// if err := simple_sql.InsertRow(ctx, conn); err != nil {
	// 	panic(err)
	// }

	// if err := simple_sql.UpdateRow(ctx, conn); err != nil {
	// 	panic(err)
	// }

	// if err := simple_sql.DeleteRow(ctx, conn); err != nil {
	// 	panic(err)
	// }

	if err := simple_sql.SelectRows(ctx, conn); err != nil {
		panic(err)
	}
	fmt.Println("success")
}
