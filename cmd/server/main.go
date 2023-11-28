package main

import (
	"fmt"

	"github.com/riaz/go-rest-api/internal/db"
)

// Run - is going to be responsible of instantiation
// and startup of go application
func Run() error {
	fmt.Println("Stating up an application")

	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to the database")
		return err
	}

	if err := db.MigrateDB(); err != nil {
		fmt.Println("failed to migrate the database")
		return err
	}

	fmt.Println("successfully connected and pinged database")

	return nil
}

func main() {
	fmt.Println("REST API")
	// This avoids the main function from panicing
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
