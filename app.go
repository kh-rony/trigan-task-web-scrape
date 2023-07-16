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
	var urlForMaleUsers string
	var urlForFemaleUsers string
	var maleUsers int
	var femaleUsers int
	var iterationLimit = len(app.countryList)

	for i := 0; i < iterationLimit; i++ {
		for maleUsers = app.numberOfMaleUsers; maleUsers > 10; maleUsers -= 10 {
			urlForMaleUsers = baseURL + strings.ToLower(app.countryList[i].Name) + "?s=" + strconv.Itoa(900+rand.Intn(100)) + "&search_terms=&gender=male&search_terms=&n=10"
			app.userList = append(app.userList, ScrapeUsers(urlForMaleUsers)...)
		}
		urlForMaleUsers = baseURL + strings.ToLower(app.countryList[i].Name) + "?s=" + strconv.Itoa(900+rand.Intn(100)) + "&search_terms=&gender=male&search_terms=&n=" + strconv.Itoa(maleUsers)
		app.userList = append(app.userList, ScrapeUsers(urlForMaleUsers)...)

		for femaleUsers = app.numberOfFemaleUsers; femaleUsers > 10; femaleUsers -= 10 {
			urlForFemaleUsers = baseURL + strings.ToLower(app.countryList[i].Name) + "?s=" + strconv.Itoa(900+rand.Intn(100)) + "&search_terms=&gender=female&search_terms=&n=10"
			app.userList = append(app.userList, ScrapeUsers(urlForFemaleUsers)...)
		}
		urlForFemaleUsers = baseURL + strings.ToLower(app.countryList[i].Name) + "?s=" + strconv.Itoa(900+rand.Intn(100)) + "&search_terms=&gender=female&search_terms=&n=" + strconv.Itoa(femaleUsers)
		app.userList = append(app.userList, ScrapeUsers(urlForFemaleUsers)...)
	}

	WriteToJSON(app.userList)
}
