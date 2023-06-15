package monitors

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/go-ble/ble"
	"github.com/sworisbreathing/go-ibbq/v2"
)

var logger = log.Default()

func temperatureReceived(temperatures []float64) {
	fmt.Printf("Received temperature data %s, %+v\n", "temperatures", temperatures)
}
func batteryLevelReceived(batteryLevel int) {
	fmt.Printf("Received battery data %s, %+v\n", "batteryPct", strconv.Itoa(batteryLevel))
}
func statusUpdated(status ibbq.Status) {
	fmt.Printf("Status updated %s, %+v\n", "status", status)
}

func disconnectedHandler(cancel func(), done chan struct{}) func() {
	return func() {
		fmt.Printf("Disconnected\n")
		cancel()
		close(done)
	}
}

func BbqRunMain() {
	var err error
	fmt.Printf("initializing context\n")
	ctx1, cancel := context.WithCancel(context.Background())
	defer cancel()
	registerInterruptHandler(cancel)
	ctx := ble.WithSigHandler(ctx1, cancel)
	fmt.Printf("context initialized\n")
	var bbq ibbq.Ibbq
	fmt.Printf("instantiating ibbq struct\n")
	done := make(chan struct{})
	var config ibbq.Configuration
	if config, err = ibbq.NewConfiguration(60*time.Second, 5*time.Minute); err != nil {
		fmt.Printf("Error creating configuration %s, %+v\n", "err", err)
	}
	fmt.Printf("config created\n")
	if bbq, err = ibbq.NewIbbq(ctx, config, disconnectedHandler(cancel, done), temperatureReceived, batteryLevelReceived, statusUpdated); err != nil {
		fmt.Printf("Error creating iBBQ %s, %+v\n", "err", err)
	}
	fmt.Printf("instantiated ibbq struct\n")
	fmt.Printf("Connecting to device\n")
	if err = bbq.Connect(); err != nil {
		fmt.Printf("Error connecting to device %s, %+v\n", "err", err)
	}

	fmt.Printf("Connected to device\n")
	<-ctx.Done()
	<-done
	fmt.Printf("Exiting\n")
}

func registerInterruptHandler(cancel context.CancelFunc) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		cancel()
	}()
}
