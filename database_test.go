package belajar_golang_database

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func TestEmpty(t *testing.T) {

}

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/belajar-golang-database")
	if err != nil {
		panic(err)
	}

	err = db.Close()
	if err != nil {
		return
	}
}

func TestInsertSQL(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	_, err := db.ExecContext(ctx, "INSERT INTO customer(id, name, email, balance, rating, birth_date, married) VALUES('kevin354', 'Kevin', 'kvin@gmail.com', 100000, '5.0', '2002-05-15', false);")
	if err != nil {
		panic(err)
	}
	fmt.Println("Success insert data to database")
}

func TestQuerySQLComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customer"
	rows, err := db.QueryContext(ctx, script)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float64
		var birthDate sql.NullTime
		var createdAt time.Time
		var married bool

		err = rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)

		if err != nil {
			panic(err)
		}

		fmt.Println("============")
		fmt.Println("ID:", id, "Name:", name, "balance:", balance, "rating:", rating, "married:", married, "creted at:", createdAt)
		if email.Valid {
			fmt.Println("Email", email.String)
		}
		if birthDate.Valid {
			fmt.Println("Birth date:", birthDate.Time)
		}
	}
}
