package main

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"io/ioutil"
	"log"
	"strings"
	"time"
	"trigan-task-web-scrape/models"
)

func ScrapeCountryNames(URL string) []models.Country {
	var countryList []models.Country

	collector := colly.NewCollector()

	collector.OnError(func(response *colly.Response, err error) {
		log.Println("Error: ", err)
	})

	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting: ", request.URL.String())
	})

	collector.OnResponse(func(response *colly.Response) {
		fmt.Println("Visited: ", response.Request.URL)
	})

	// scraping logic
	collector.OnHTML("select.country-select", func(element *colly.HTMLElement) {
		element.ForEach("option", func(_ int, innerElement *colly.HTMLElement) {
			newCountry := models.Country{}
			newCountry.Name = innerElement.Text
			countryList = append(countryList, newCountry)
		})
	})

	collector.OnScraped(func(response *colly.Response) {
		fmt.Println("Scraped: ", response.Request.URL)
	})

	// visiting the target page
	err := collector.Visit(URL)
	if err != nil {
		fmt.Println("Could not visit provided URL")
		fmt.Println("Error: ", err)
	}

	return countryList
}

func ScrapeUsers(URL string) []models.User {
	var userList []models.User

	collector := colly.NewCollector()

	collector.OnError(func(response *colly.Response, err error) {
		log.Println("Error: ", err)
	})

	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting: ", request.URL.String())
	})

	collector.OnResponse(func(response *colly.Response) {
		fmt.Println("Visited: ", response.Request.URL)
	})

	// scraping logic
	collector.OnHTML("div.media-body", func(element *colly.HTMLElement) {
		// initializing a new User instance
		newUser := models.User{}

		/**
		scraping data starts from here
		*/
		// image
		newUser.Image = element.ChildAttr("img", "data-ezsrc")
		// name and gender
		nameAndGender := strings.Split(element.ChildText("dd.col-12"), " (")
		newUser.Name = nameAndGender[0]
		newUser.Gender = strings.ReplaceAll(nameAndGender[1], ")", "")

		/**
		already found Image, Name, Gender
		so the value of userDataFragmentCounter is set to 3
		*/
		userDataFragmentCounter := 3

		element.ForEach("dd.col-sm-8", func(_ int, innerElement *colly.HTMLElement) {
			userDataFragment := strings.Join(strings.Fields(strings.TrimSpace(innerElement.Text)), " ")
			/**
			each time while iterating over dd.col-sm-8 elements
			a new userDataFragment is accessed
			so the value of userDataFragmentCounter incremented each time accordingly
			*/
			userDataFragmentCounter++

			if userDataFragmentCounter == 4 {
				newUser.Address = userDataFragment
			} else if userDataFragmentCounter == 5 {
				newUser.PhoneNumber = userDataFragment
			} else if userDataFragmentCounter == 6 {
				newUser.Email = userDataFragment
			} else if userDataFragmentCounter == 7 {
				newUser.IP = userDataFragment
			} else if userDataFragmentCounter == 8 {
				newUser.Username = userDataFragment
			} else if userDataFragmentCounter == 9 {
				newUser.Password = userDataFragment
			} else if userDataFragmentCounter == 10 {
				newUser.CreditCardNumber = userDataFragment
			} else if userDataFragmentCounter == 11 {
				newUser.ExpirationDate = userDataFragment
			} else if userDataFragmentCounter == 12 {
				newUser.IBAN = userDataFragment
			} else if userDataFragmentCounter == 13 {
				newUser.SwiftBicNumber = userDataFragment
			} else if userDataFragmentCounter == 14 {
				newUser.Company = userDataFragment
			}
		})
		userList = append(userList, newUser)
	})

	collector.OnScraped(func(response *colly.Response) {
		fmt.Println("Scraped: ", response.Request.URL)
	})

	// visiting the target page
	err := collector.Visit(URL)
	if err != nil {
		fmt.Println("Could not visit provided URL")
		fmt.Println("Error: ", err)
	}

	return userList
}

func WriteToJSON(userList []models.User) {
	file, err := json.MarshalIndent(userList, "", "    ")
	if err != nil {
		log.Println("Unable to create json file")
		return
	}

	currentTimeString := time.Now().Format("2006_01_02 15_04_05")
	_ = ioutil.WriteFile("users "+currentTimeString+".json", file, 0644)
}
