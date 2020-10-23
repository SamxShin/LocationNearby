package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type location struct {
	primaryKey int
	streetName string
	country    string
	latitude   float64
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "5936"
	dbname   = "testData"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	//this is the connection string to database
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//needed to add data in correct collumns in table
	statement := "INSERT INTO public.places(country, streetname, latitude) VALUES ($1, $2, $3)"
	stmt, err := db.Prepare(statement)
	if err != nil {
		fmt.Print(err)
	}
	db.Exec(statement)

	//create an instance of location
	place := location{}
	//for loop to add data into database
	for i := 0; i < 3; i++ {
		fmt.Print("\nEnter the street name: ")
		scanner.Scan()
		srt := scanner.Text()
		place.streetName = srt
		fmt.Print("Enter the country name: ")
		scanner.Scan()
		crty := scanner.Text()
		place.country = crty
		fmt.Print("enter the latitude: ")
		fmt.Scanln(&place.latitude)
		stmt.QueryRow(place.streetName, place.country, place.latitude)
	}

	defer stmt.Close()
	rows, err := db.Query("Select streetName, country, latitude from inputs")
	if err != nil {
		fmt.Print(err)
	}

	defer rows.Close()

	for rows.Next() {
		var streetName string
		var countryName string
		var latitude float64

		err := rows.Scan(&streetName, &countryName, &latitude)
		if err != nil {
			fmt.Print(err)
		}
		fmt.Printf("%s %s ", streetName, countryName)
	}
	defer rows.Close()

	for rows.Next() {
		var streetName string
		var countryName string
		var latitude float64

		err := rows.Scan(&streetName, &countryName, &latitude)
		if err != nil {
			fmt.Print(err)
		}
		fmt.Printf("%s %s ", streetName, countryName)
	}

}
