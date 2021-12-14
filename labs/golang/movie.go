package main

type movie struct {
	title         string
	studio        string
	yearPublished int
}

func newMovie(title string, studio string, year int) *movie {
	return &movie{title: title, studio: studio, yearPublished: year}
}

func (movie *movie) compareTo(other *movie) int {
	return 0
}
