package models

import (
	"database/sql"
	"errors"
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

var (
	Drivers map[string]*Driver
)

type Driver struct {
	DriverId int
	Name     string
	License  string
}

func init() {
	Drivers = make(map[string]*Driver)
}

func AddOneDriver(Driver Driver) (DriverId int) {
	//Driver.DriverId = DriverId
	//Drivers[Driver.DriverId] = &Driver
	AddDriver(Driver.DriverId, Driver.Name, Driver.License)
	return Driver.DriverId
}

func GetOneDriver(DriverId string) (Driver *Driver, err error) {
	if v, ok := Drivers[DriverId]; ok {
		return v, nil
	}
	return nil, errors.New("DriverId Not Exist")
}

func GetAllDriver() map[string]*Driver {
	return Drivers
}

// func UpdateDriver(DriverId string, Score int64) (err error) {
// 	if ok := Drivers[DriverId]; ok != nil {
// 		//		v.Score = Score
// 		return nil
// 	}
// 	return errors.New("DriverId Not Exist")
// }

func DeleteDriver(DriverId string) {
	delete(Drivers, DriverId)
}
