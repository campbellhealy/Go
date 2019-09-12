package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/tidwall/gjson"
)

var location1x, location1y, location2x, location2y, location3x, location3y string
var location1, location2, location3 string
var revGeo, geoRev string

// Building the URL
func join(strs ...string) string {
	var revGeo strings.Builder
	for _, str := range strs {
		revGeo.WriteString(str)
	}
	return revGeo.String()
}

func builderURL(x, y string) string {
	geoRev = (join(
		"https://eu1.locationiq.com/v1/reverse.php?key=",
		"****************&lat=",
		x,
		"&lon=",
		y,
		"&format=json",
	))

	return geoRev
}

func locator() {
	location1x, location1y = "48.853106", "2.384202"
	location2x, location2y = "48.858539", "2.294438"
	location3x, location3y = "48.860754", "2.337632"

	location1 = builderURL(location1x, location1y)
	location2 = builderURL(location2x, location2y)
	location3 = builderURL(location3x, location3y)
}

func main() {
	// This loop allows each of the three Locations from the locator to be searched
	// If there were a lot more locations a construction of the variable i may  require a bit more

	locator()
	var loc string
	var i = 1
	for i <= 3 {
		// Changed from If/Else to Switch as it is cleaner to read
		switch i {
		case 1:
			loc = location1
		case 2:
			loc = location2
		case 3:
			loc = location3
		}
		i++
		println()
		println(loc) //Print the constructed URL

		// GET Request
		time.Sleep(1 * time.Second) // Added this as it was hitting the request limit for a second
		request, err := http.Get(loc)
		if err != nil {
			log.Fatalln(err)
		}
		defer request.Body.Close()

		// GET Reply
		body, err := ioutil.ReadAll(request.Body)
		if err != nil {
			log.Fatalln(err)
		}

		//Print the output of the required

		json := string(body)
		road := gjson.Get(json, "address.road")
		city := gjson.Get(json, "address.city")
		postcode := gjson.Get(json, "address.postcode")
		country := gjson.Get(json, "address.country")
		countryCode := gjson.Get(json, "address.country_code")

		println("Road: ", road.String())
		println("City: ", city.String())
		println("Postcode: ", postcode.String())
		println("Country: ", country.String())
		println("Country Code: ", countryCode.String())
		log.Println() // Adding a date stamp to detail when the test ran
	}
}
