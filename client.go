package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func GetVotiLive(giornata, codiceSquadra, magicNumber int) TimedRecords {
	url := fmt.Sprintf("https://www.fantacalcio.it/api/live/%d?g=%d&i=%d",
		codiceSquadra, giornata, magicNumber)

	// make http request to fantacalcio server
	resp, err := http.Get(url)
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()

	// read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
	}

	// unmarshal json
	var records TimedRecords
	err = json.Unmarshal(body, &records.Records)
	if err != nil {
		log.Panic(err)
	}
	records.Timestamp = time.Now()

	return records
}
