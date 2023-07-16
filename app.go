package main

import (
	"math/rand"
	"strconv"
	"strings"
	"trigan-task-web-scrape/models"
)

type App struct {
	numberOfMaleUsers   int
	numberOfFemaleUsers int
	baseURL             string
	countryList         []models.Country
	userList            []models.User
}

func (app *App) Initialize(numberOfMaleUsers int, numberOfFemaleUsers int, baseURL string) {

	app.numberOfMaleUsers = numberOfMaleUsers
	app.numberOfFemaleUsers = numberOfFemaleUsers
	app.baseURL = baseURL
}

func (app *App) LoadCountryNames() {

	app.countryList = ScrapeCountryNames(baseURL)
}

func (app *App) ScrapeUsersAndWriteToJSON() {
	var numberOfCountries = len(app.countryList)
	var formattedCountryName string
	var numberOfMaleUsersToBeGenerated int
	var numberOfFemaleUsersToBeGenerated int
	var urlForMaleUsers string
	var urlForFemaleUsers string

	for i := 0; i < numberOfCountries; i++ {
		formattedCountryName = strings.ToLower(strings.ReplaceAll(app.countryList[i].Name, " ", "-"))

		for numberOfMaleUsersToBeGenerated = app.numberOfMaleUsers; numberOfMaleUsersToBeGenerated > 10; numberOfMaleUsersToBeGenerated -= 10 {
			urlForMaleUsers = baseURL + formattedCountryName + "?s=" + strconv.Itoa(900+rand.Intn(100)) + "&search_terms=&gender=male&search_terms=&n=10"
			app.userList = append(app.userList, ScrapeUsers(urlForMaleUsers)...)
		}
		if numberOfMaleUsersToBeGenerated > 0 {
			urlForMaleUsers = baseURL + formattedCountryName + "?s=" + strconv.Itoa(900+rand.Intn(100)) + "&search_terms=&gender=male&search_terms=&n=" + strconv.Itoa(numberOfMaleUsersToBeGenerated)
			app.userList = append(app.userList, ScrapeUsers(urlForMaleUsers)...)
		}

		for numberOfFemaleUsersToBeGenerated = app.numberOfFemaleUsers; numberOfFemaleUsersToBeGenerated > 10; numberOfFemaleUsersToBeGenerated -= 10 {
			urlForFemaleUsers = baseURL + formattedCountryName + "?s=" + strconv.Itoa(900+rand.Intn(100)) + "&search_terms=&gender=female&search_terms=&n=10"
			app.userList = append(app.userList, ScrapeUsers(urlForFemaleUsers)...)
		}
		if numberOfFemaleUsersToBeGenerated > 0 {
			urlForFemaleUsers = baseURL + formattedCountryName + "?s=" + strconv.Itoa(900+rand.Intn(100)) + "&search_terms=&gender=female&search_terms=&n=" + strconv.Itoa(numberOfFemaleUsersToBeGenerated)
			app.userList = append(app.userList, ScrapeUsers(urlForFemaleUsers)...)
		}
	}

	WriteToJSON(app.userList)
}
