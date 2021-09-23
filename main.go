package main

import (
	"os"
	"time"

	datadog "github.com/GlobalFreightSolutions/logrus-datadog-hook"
	"github.com/sirupsen/logrus"
)

func main() {
	apiKey := os.Getenv("DD_API_KEY")

	host, _ := os.Hostname()

	service := "sandbox-golang-logrus"

	batching := false

	options := &datadog.Options{
		ApiKey:                &apiKey,
		Service:               &service,
		Host:                  &host,
		ClientBatchingEnabled: &batching,
	}

	hook, err := datadog.New(options)

	if err != nil {
		panic(err.Error())
	}

	logger := logrus.New()
	logger.AddHook(hook)

	// This ensures that the logger exits gracefully and all buffered logs are sent before closing down
	logrus.DeferExitHandler(hook.Close)

	sum := 0
	for i := 1; i < 5; i++ {
		logger.Infof("This is a log sent to datadog - %d", i)
		time.Sleep(1 * time.Second)
		sum += i
	}
}
