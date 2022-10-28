package main

import (
  "fmt"
)

func main()  {
  countries := map[string]string {
    "Costa Rica": "San Jose",
    "United States": "Washington D. C.",
    "Nicaragua": "Managua",
    "Italy": "Rome",
  }

  fmt.Println(countries["Italy"])

  countries["France"] = "Paris"

  fmt.Println(countries)

  delete(countries, "France")

  // Delete a key
  fmt.Println(countries)

  _, exists := countries["France"]

  fmt.Println(exists)

}
