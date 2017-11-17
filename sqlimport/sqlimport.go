package main

import (
	"database/sql"
	"encoding/json"
	reader "fileread"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {

	fmt.Println("Reading Drivers")

	drivers := reader.GetDrivers("src/fileread/drivers.json")

	connStr := "user=postgres dbname=Moshe-gett password=storm sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("1 ", err)
	}
	fmt.Println("connection opened")

	sqlStatement := `  
	INSERT INTO "Gett"."Drivers" (id, name, license)   
	VALUES ($1, $2, $3)`

	for _, p := range drivers {

		_, err = db.Exec(sqlStatement, p.ID, p.Name, p.License)
		if err != nil {
			panic(err)
		}

		fmt.Println("inserted: = , Name = " + p.Name + ", lic = " + p.License)
	}

	lines := reader.ProcessMetrics("src/fileread/metrics.json")

	// due to problem inserting the timestamp... I skipped this.

	sqlMetric := `INSERT INTO "Gett"."Metrics" (metric_name, value, lat, lon, driver_id)   
	VALUES ($1, $2, $3, $4, $5)`
	var metr reader.Metric
	counter := 0

	//parsing metrics, row by row...
	for _, l := range lines {

		err := json.Unmarshal([]byte(l), &metr)

		fmt.Println("metric = ", metr.MetricName, metr.Value, metr.Lat, metr.Lon, metr.Timestamp, metr.DriverId)
		if metr.DriverId != "" && metr.DriverId != "0" {

			// when succeeded, insert row to db
			_, err = db.Exec(sqlMetric, metr.MetricName, metr.Value, metr.Lat, metr.Lon, metr.DriverId)
			if err != nil {
				panic(err)
			}
			counter++
		}
	}
	fmt.Println(counter, " rows inserted!")

}
