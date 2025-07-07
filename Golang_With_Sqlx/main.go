package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// Product struct with db tags
type Product struct {
	ID    int     `db:"id"`
	Name  string  `db:"name"`
	Price float64 `db:"price"`
}

func main() {
	// 1. DSN (Data Source Name)
	dsn := "user:password@tcp(localhost:3306)/database_name"

	// 2. Connect to MySQL using sqlx
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatalln("Failed to connect to the database:", err)
	}
	defer db.Close()

	// 3. Ping to verify connection
	if err := db.Ping(); err != nil {
		log.Fatalln("Failed to ping the database:", err)
	}
	fmt.Println("Connected to the database!")

	// 4. SELECT multiple rows
	var products []Product
	err = db.Select(&products, "SELECT id, name, price FROM products")
	if err != nil {
		log.Println("Failed to fetch products:", err)
	} else {
		fmt.Println("Products:")
		for _, p := range products {
			fmt.Printf("→ ID: %d | Name: %s | Price: ₹%.2f\n", p.ID, p.Name, p.Price)
		}
	}

	// 5. SELECT a single product
	var product Product
	err = db.Get(&product, "SELECT id, name, price FROM products WHERE id = ?", 1)
	if err != nil {
		log.Println("Failed to fetch product by ID:", err)
	} else {
		fmt.Println("Single Product:", product)
	}

	// 6. INSERT a new product
	newProduct := Product{
		Name:  "New Product",
		Price: 350,
	}

	result, err := db.NamedExec(
		"INSERT INTO products (name, price) VALUES (:name, :price)",
		&newProduct,
	)
	if err != nil {
		log.Println("Failed to insert product:", err)
	} else {
		lastID, _ := result.LastInsertId()
		rowsAffected, _ := result.RowsAffected()
		fmt.Printf("Inserted! Last ID: %d | Rows affected: %d\n", lastID, rowsAffected)
	}
}
