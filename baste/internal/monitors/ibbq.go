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

func BbqRunMain(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	registerInterruptHandler(cancel)

	ctx = ble.WithSigHandler(ctx, cancel)
	done := make(chan struct{})

	config, err := ibbq.NewConfiguration(60*time.Second, 5*time.Minute)
	if err != nil {
		return fmt.Errorf("ibbq.NewConfiguration: %w", err)
	}

	bbq, err := ibbq.NewIbbq(ctx, config, disconnectedHandler(cancel, done), temperatureReceived, batteryLevelReceived, statusUpdated)
	if err != nil {
		return fmt.Errorf("ibbq.NewIbbq (hint: run with sudo): %w", err)
	}

	if err = bbq.Connect(); err != nil {
		// this error will be context.Canceled if its a timeout.
		// refactor this to connect in a loop.
		return fmt.Errorf("bbq.Connect: %w", err)
	}

	<-ctx.Done()
	<-done

	return nil
}

func registerInterruptHandler(cancel context.CancelFunc) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		cancel()
	}()
}
