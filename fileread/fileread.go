package fileread

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Driver struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	License string `json:"license_number"`
}
type Metric struct {
	MetricName string  `json:"metric_name"`
	Value      string  `json:"value"`
	Lon        float64 `json:"lon"`
	Lat        float64 `json:"lat"`
	Timestamp  int64   `json:"timestamp"`
	DriverId   string  `json:"driver_id"`
}

func main() {

	fmt.Println("Reading Drivers")

	drivers := GetDrivers("src/fileread/drivers.json")

	for _, p := range drivers {

		fmt.Println("{id = , Name = " + p.Name + ", lic = " + p.License)
	}

	ProcessMetrics("src/fileread/metrics.json")
}

func GetDrivers(fileName string) []Driver {
	fmt.Println("reading " + fileName + "  file on current path")
	raw, err := ioutil.ReadFile("./" + fileName)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	var drivers []Driver

	json.Unmarshal(raw, &drivers)

	return drivers

}

func ProcessMetrics(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines

}
