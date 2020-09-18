package main

import (
	"fmt"
)

func main() {
	var publisher, writer, artist, title string
	var year, pageNumber int
	var grade float32

	title = "Mr. GoToSleep"
	writer = "Tracey Hatchet"
	artist = "Jewel Tampson"
	publisher = "DizzyBook Publishing Inc."
	year = 1997
	pageNumber = 14
	grade = 6.5

	fmt.Println(title, "written by", writer, "drawn by", artist, "year:", year, "Number of pages:", pageNumber)

	title = "Epic Vol. 1"
	writer = "Ryan N. Shawn"
	artist = "Phoebe Paperclips"
	publisher = "DizzyBook Publishing Inc."
	year = 2013
	pageNumber = 169
	grade = 9.0

	fmt.Println(title, "written by", writer, "drawn by", artist, "year:", year)

	title = "Epic Vol. 2"
	writer = "Tracey Hatchet"
	artist = "Jewel Tampson"
	publisher = "DizzyBook Publishing Inc."
	year = 2017
	pageNumber = 130
	grade = 9.5

	fmt.Println(title, "written by", writer, "drawn by", artist, "year:", year)

	title = "The Comic Chief"
	writer = "Enna William"
	artist = "Mordecai Egbe"
	publisher = "Flying men Publishing Inc."
	year = 2020
	pageNumber = 4
	grade = 8.9

	fmt.Println(title, "written by", writer, "drawn by", artist, "year:", year, "publisher:", publisher, "grade:", grade)

}
