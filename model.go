package main

import "time"

type Record struct {
	Nome   string
	Ruolo  string
	Voto   float32
	Evento string
	Id     int
}

type TimedRecords struct {
	Records   []Record
	Timestamp time.Time
}
