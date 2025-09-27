package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "modernc.org/sqlite" // registers the "sqlite" driver
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// In-memory SQLite database
	db, err := sql.Open("sqlite", "file:memdb1?mode=memory&cache=shared")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			panic(err)
		}
	}()

	if err := db.PingContext(ctx); err != nil {
		panic(err)
	}

	// Schema
	_, _ = db.ExecContext(ctx, `CREATE TABLE users (id INTEGER PRIMARY KEY, name TEXT, age INT)`)

	// Exec with placeholders
	res, _ := db.ExecContext(ctx, `INSERT INTO users(name, age) VALUES (?, ?)`, "Alice", 30)
	id, _ := res.LastInsertId()
	fmt.Println("inserted id:", id)

	// Prepared statement (reused)
	stmt, _ := db.PrepareContext(ctx, `INSERT INTO users(name, age) VALUES (?, ?)`)
	defer func() {
		if err := stmt.Close(); err != nil {
			panic(err)
		}
	}()
	_, _ = stmt.ExecContext(ctx, "Bob", 42)
	_, _ = stmt.ExecContext(ctx, "Carol", 28)

	// QueryRow (single value)
	var count int
	_ = db.QueryRowContext(ctx, `SELECT COUNT(*) FROM users`).Scan(&count)
	fmt.Println("count:", count)

	// Query (multiple rows)
	rows, _ := db.QueryContext(ctx, `SELECT id, name, age FROM users ORDER BY id`)
	defer func() {
		if err := rows.Close(); err != nil {
			panic(err)
		}
	}()
	for rows.Next() {
		var id int64
		var name string
		var age int
		_ = rows.Scan(&id, &name, &age)
		fmt.Printf("row: id=%d name=%s age=%d\n", id, name, age)
	}
	_ = rows.Err()

	// Transaction
	tx, _ := db.BeginTx(ctx, nil)
	_, _ = tx.ExecContext(ctx, `INSERT INTO users(name, age) VALUES (?, ?)`, "Dave", 35)
	_ = tx.Commit()

	_ = db.QueryRowContext(ctx, `SELECT COUNT(*) FROM users`).Scan(&count)
	fmt.Println("count after tx:", count)
}
