package main

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/nukopy/intro-sqlc/app/gen/sqlc/db"
)

func main() {
	// prepare database connection
	driver := "pgx"
	user := "myuser"
	pass := "mypassword"
	host := "localhost"
	port := 15432
	dbName := "intro_sqlc_db"
	dataSourceName := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", user, pass, host, port, dbName)

	pgx, err := sql.Open(driver, dataSourceName)
	if err != nil {
		panic(err)
	}
	defer pgx.Close()

	ctx := context.Background()
	q := db.New(pgx)
	defer q.ResetTables(ctx)

	// create authors
	authorParamsList := []db.CreateAuthorParams{
		{
			ID:   101,
			Name: "Thorsten Ball",
		},
		{
			ID:   103,
			Name: "Mat Ryer",
		},
		{
			ID:   105,
			Name: "Yosuke Okuwaki",
		},
	}

	fmt.Println("Creating authors...")
	for _, param := range authorParamsList {
		a, err := q.CreateAuthor(ctx, param)
		if err != nil {
			panic(err)
		}
		fmt.Println(a)
	}

	// create book
	bookParamsList := []db.CreateBookParams{
		{
			ID:       201,
			Title:    "Go言語でつくるインタプリタ",
			AuthorID: 101,
			Price:    3740,
		},
		{
			ID:       202,
			Title:    "Go言語によるWebアプリケーション開発",
			AuthorID: 103,
			Price:    3520,
		},
	}

	fmt.Println("Creating books...")
	for _, param := range bookParamsList {
		b, err := q.CreateBook(ctx, param)
		if err != nil {
			panic(err)
		}
		fmt.Println(b)
	}

	// get author
	var targetAuthorID int32 = 101
	fmt.Println("Getting author by ID:", targetAuthorID)
	a, err := q.GetAuthor(ctx, targetAuthorID)
	if err != nil {
		panic(err)
	}
	fmt.Println(a)

	// delete author
	targetAuthorID = 100
	fmt.Println("Deleting author by ID:", targetAuthorID)
	if err := q.DeleteAuthor(ctx, targetAuthorID); err != nil {
		panic(err)
	}

	// list book
	fmt.Println("Listing books over price 3500...")
	var priceThreshold int32 = 3500
	ls, err := q.ListBookOverPrice(ctx, priceThreshold)
	if err != nil {
		panic(err)
	}
	for _, l := range ls {
		fmt.Println(l)
	}
}
