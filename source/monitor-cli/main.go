/*
   Copyright 2018 the original author or authors

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.

   See: https://github.com/sworisbreathing/go-ibbq/blob/master/examples/datalogger/main.go

*/
package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-ble/ble"
	"github.com/sworisbreathing/go-ibbq/ibbq"
)

//var logger = log.New("main")
type reading struct {
	probe0 float64 `json:"probe0"`
	probe1 float64 `json:"probe1"`
	probe2 float64 `json:"probe2"`
	probe3 float64 `json:"probe3"`
}

func temperatureReceived(temperatures []float64) {
	//(0°C × 9/5) + 32 = 32°F

	recordReadings(temperatures)

	for i := 0; i < len(temperatures); i++ {
		fmt.Printf("probe%d - %v ", i, (temperatures[i]*9/5 + 32))
	}

	fmt.Printf("\n")

	//fmt.Println("Received temperature data", "temperatures", temperatures)
}
func batteryLevelReceived(batteryLevel int) {
	fmt.Println("Received battery data", "batteryPct", strconv.Itoa(batteryLevel))
}
func statusUpdated(status ibbq.Status) {
	fmt.Println("Status updated", "status", status)
}

func disconnectedHandler(cancel func(), done chan struct{}) func() {
	return func() {
		fmt.Println("Disconnected")
		cancel()
		close(done)
	}
}

// Code modified to remove hard coded things... obviously there's work here to be done to make it not-dumb.
func recordReadings(temps []float64) {
	url := "http://localhost:21337/v1/development/data/temperature/0aa73e9b-f978-43cc-bc8b-ef8bc42467f9"

	var tempReading reading

	tempReading.probe0 = temps[0]
	tempReading.probe1 = temps[1]
	tempReading.probe2 = temps[2]
	tempReading.probe3 = temps[3]

	data := map[string]float64{"probe0": (temps[0]*9/5 + 32), "probe1": (temps[1]*9/5 + 32), "probe2": (temps[2]*9/5 + 32), "probe3": (temps[3]*9/5 + 32)}

	tempReadingJson, err := json.Marshal(data)

	if err != nil {
		fmt.Println(err)
	}

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(tempReadingJson))

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJodHRwczovL2JicS5rOHMuc3NhcmdlbnQubmV0LyIsImV4cCI6MTU1ODY2ODk2NywiZm4iOiJTY290dCBTYXJnZW50IiwiaWF0IjoxNTU4NTY4OTY3LCJpc3MiOiJodHRwczovL2JicS5rOHMuc3NhcmdlbnQubmV0LyIsImxvZ2luIjoic2NvdHQiLCJzdWIiOiJkY2U0YjI0Yy1mYzcwLTQwYzctOTJlZi1jYmNhYmNiYmFmMGQiLCJ0ZW5hbnQiOiJjMjhmMTVhNy0yNGYzLTVlYWQtODQwMy1iNGQwODMxMjgwMWUifQ.LPKzI4MhAj1x6OJrbhad8LJD9n6A8YxPR4HC6EOtT80")
	req.Header.Add("cache-control", "no-cache")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	//body, _ := ioutil.ReadAll(res.Body)
}

// Flesh this out more.. it should log in and grab a bearer token.
func doLogin(loginname string, password string) (string, error) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJodHRwczovL2JicS5rOHMuc3NhcmdlbnQubmV0LyIsImV4cCI6MTU1NzExOTgwOCwiZm4iOiJTY290dCBTYXJnZW50IiwiaWF0IjoxNTU3MDE5ODA4LCJpc3MiOiJodHRwczovL2JicS5rOHMuc3NhcmdlbnQubmV0LyIsImxvZ2luIjoic2NvdHQiLCJzdWIiOiJkY2U0YjI0Yy1mYzcwLTQwYzctOTJlZi1jYmNhYmNiYmFmMGQifQ.soc3hiRpoTGD1RI3vnC9ohJ_D5nkd9hXvs2O6asguV4"

	return token, nil
}

func findActiveSession(monitorAddress string) (string, error) {

	session := "39ae436c-61d3-4e17-826c-56ed9ef33c30"
	return session, nil
}

func main() {

	//bearerToken, err := doLogin("someUserName", "somePassword")

	/*
		if err != nil {
			fmt.Println("Login to bbq.k8s.ssargent.net failed")
			return
		}

		//sessionId, err := findActiveSession("GetAddressFirstAndUseHere")

		if err != nil {
			// we'll need to do something more interesting here. perhaps poll and wait for a session... but for now let's exit.
			fmt.Println("Please create a session and rerun this")
			return
		}*/

	var err error

	//logger.Debug("initializing context")
	ctx1, cancel := context.WithCancel(context.Background())
	defer cancel()
	registerInterruptHandler(cancel)
	ctx := ble.WithSigHandler(ctx1, cancel)
	//logger.Debug("context initialized")
	var bbq ibbq.Ibbq
	//logger.Debug("instantiating ibbq struct")
	done := make(chan struct{})
	var config ibbq.Configuration
	if config, err = ibbq.NewConfiguration(60*time.Second, 5*time.Minute); err != nil {
		fmt.Println("Error creating configuration", "err", err)
		//logger.Fatal("Error creating configuration", "err", err)
	}
	if bbq, err = ibbq.NewIbbq(ctx, config, disconnectedHandler(cancel, done), temperatureReceived, batteryLevelReceived, statusUpdated); err != nil {
		fmt.Println("Error creating iBBQ", "err", err)
	}
	//	logger.Debug("instantiated ibbq struct")
	//	logger.Info("Connecting to device")
	if err = bbq.Connect(); err != nil {
		fmt.Println("Error connecting to device", "err", err)
	}
	//	logger.Info("Connected to device")
	<-ctx.Done()
	<-done
	//	logger.Info("Exiting")
}
