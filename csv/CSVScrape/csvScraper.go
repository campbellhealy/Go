package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	//Open the csv file
	psfile, err := os.Open("data/nfl-2019-pittsburgh-steelers-GMTStandardTime.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	//Read the csv file
	home := csv.NewReader(psfile)
	println("Home Game with Away Team on Date")
	// Go through the file records
	for {
		record, err := home.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if "Heinz Field" == record[2] {
			fmt.Printf("Home Game with %s on %s \n", record[4], record[1])
		}
	}
	//Open the csv file
	psafile, err := os.Open("data/nfl-2019-pittsburgh-steelers-GMTStandardTime.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	//Read the csv file
	away := csv.NewReader(psafile)
	println("")
	println("")

	// Go through the file records
	for {
		record, err := away.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if "Heinz Field" != record[2] {
			fmt.Printf("Away Game to %s at %s on %s \n", record[3], record[2], record[1])
		}
	}
}
