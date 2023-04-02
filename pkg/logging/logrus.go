package logging

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
)

func StartLogging() {
	logFile, err := os.OpenFile("log.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(fmt.Sprintf("[Error]: %s", err))
	}

	log.SetFormatter(&log.TextFormatter{FullTimestamp: true, TimestampFormat: "02.01 15:04:05"})

	// Output to log file and console
	log.SetOutput(io.MultiWriter(logFile, os.Stdout))

	//logrus show line number
	log.SetReportCaller(true)
}
