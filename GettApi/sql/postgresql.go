package postgresql

import (
	"database/sql"
	"fmt"
	"log"
)

func AddDriver(id int, name string, license string) {

	connStr := "user=postgres dbname=Moshe-gett password=storm sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("1 ", err)
	}
	fmt.Println("connection opened")

	sqlStatement := `  
	INSERT INTO "Gett"."Drivers" (id, name, license)   
	VALUES ($1, $2, $3)`

	_, err = db.Exec(sqlStatement, id, name, license)
	if err != nil {
		panic(err)
	}

	fmt.Println("inserted: = , Name = " + name + ", lic = " + license)
}
