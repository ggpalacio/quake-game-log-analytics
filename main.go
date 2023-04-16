package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/ggpalacio/quake-game-log-analytics/logfile"
	"github.com/ggpalacio/quake-game-log-analytics/report"
	"os"
	"strings"
)

const (
	argLogFile    = "logFile"
	argReportFile = "reportFile"
)

func main() {
	arguments := parseArguments()
	logFilePath := arguments[argLogFile]
	reportFilePath := arguments[argReportFile]

	if logFilePath == "" {
		fmt.Printf("command-line error: log file path argument must be provided")
		return
	}

	log, err := logfile.New(logFilePath)
	if err != nil {
		fmt.Printf("reading file error: %v", err)
		return
	}

	games := report.New(log)
	reportJson, _ := json.MarshalIndent(games, "", "  ")
	if reportFilePath != "" {
		output, err := os.Create(reportFilePath)
		if err != nil {
			fmt.Printf("creating file error: %v", err)
			return
		}
		defer output.Close()
		output.WriteString(string(reportJson))
	} else {
		fmt.Println(string(reportJson))
	}
}

func parseArguments() map[string]string {
	var reportFile string
	flag.StringVar(&reportFile, "o", "", "")
	flag.Parse()

	return map[string]string{
		argLogFile:    strings.TrimSpace(flag.Arg(0)),
		argReportFile: strings.TrimSpace(reportFile),
	}
}
