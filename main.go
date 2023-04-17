package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/ggpalacio/quake-game-log-analytics/game"
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

	logFile, err := game.NewLogFile(logFilePath)
	if err != nil {
		fmt.Printf("reading file error: %v", err)
		return
	}

	report := game.NewReport(logFile)
	reportJson, _ := json.MarshalIndent(report, "", "  ")
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
