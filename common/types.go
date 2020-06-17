package common

//CovidUnit represents object for comparing
type CovidUnit struct {
	Province    string
	Country     string
	Latitude    string
	Longitude   string
	DeathByDate int // Amount of deaths, happened in that date according to https://data.humdata.org/
}
