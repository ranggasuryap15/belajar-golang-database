package belajar_golang_database

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
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

func TestSQLInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	// user input
	username := "admin'; #"
	password := "admin"

	script := "SELECT username FROM user WHERE username='" + username + "' AND password='" + password + "' LIMIT 1"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err = rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Print("=====")
		fmt.Println("Berhasil login sebagai", username)
	} else {
		fmt.Println("Gagal login", username)
	}
}

func TestSQLInject(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	// user input
	username := "admin'; #"
	password := "admin"

	script := "SELECT username FROM user WHERE username=? AND password=? LIMIT 1" // sebagai mengatasi sql injection
	rows, err := db.QueryContext(ctx, script, username, password)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err = rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Print("=====")
		fmt.Println("Berhasil login sebagai", username)
	} else {
		fmt.Println("Gagal login", username)
	}
}

func TestLastInsertID(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	email := "ranggasuryap15"
	comment := "Mantap gan"

	ctx := context.Background()
	sqlQuery := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	result, err := db.ExecContext(ctx, sqlQuery, email, comment)
	if err != nil {
		panic(err)
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	fmt.Println("Last Insert ID:", insertId)
}

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	//email := "ranggasuryap15"
	//comment := "Sayang revina wkwk"
	sqlQuery := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	stmt, err := db.PrepareContext(ctx, sqlQuery)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	// untuk melakukan insert data ke database
	for i := 0; i < 10; i++ {
		email := "rangga" + strconv.Itoa(i) + "@gmail.com"
		comment := "ini komen ke " + strconv.Itoa(i)

		result, err := stmt.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println("Comment ID:", id)
	}
}

func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	// do transaction
	
	ctx := context.Background()
	sqlQuery := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	for i := 0; i < 10; i++ {
		email := "ranggas" + strconv.Itoa(i) + "@gmail.com"
		comment := "Ini komen ke " + strconv.Itoa(i)
		_, err = tx.ExecContext(ctx, sqlQuery, email, comment)
		if err != nil {
			panic(err)
		}
	}

	err = tx.Commit() // akan dicommit ke dataabsenya 
	if err != nil {
		panic(err)
	}
}