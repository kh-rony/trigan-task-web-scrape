package main

const (
	numberOfMaleUsers   = 100
	numberOfFemaleUsers = 100
	baseURL             = "https://www.random-name-generator.com/"
)

func main() {
	app := App{}
	app.Initialize(numberOfMaleUsers, numberOfFemaleUsers, baseURL)
	app.LoadCountryNames()
	app.ScrapeUsersAndWriteToJSON()
}
