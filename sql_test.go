package belajar_golang_database

import (
	"context"
	"fmt"
	"testing"
)

func TestExecSQL(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO customer(idcustomer, name) VALUES('kevin', 'Kevin')"
	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer")
}

func TestQuerySQL(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	rows, err := db.QueryContext(ctx, "SELECT idcustomer, name FROM customer")

	for rows.Next() {
		var id, name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("ID:", id)
		fmt.Println("Name:", name)
		fmt.Println("===============")
	}

	if err != nil {
		panic(err)
	}
	defer rows.Close()
}
