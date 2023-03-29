package src

import (
	"bufio"
	md "exporter/models"
	"fmt"
	"os"
	"strings"
	"time"
)

func FindErr(logFiles []string, conf md.ConfigExprt) int {
	counter := 0

	startPeriodStr := make([]string, conf.Exporter.Grep_period)
	beginTimestampUnix := timestatpBuild(int(conf.Exporter.Grep_period))

	for i := 0; i < conf.Exporter.Grep_period; i++ {

		tmpStartPeriodInt := beginTimestampUnix + (60 * int64(i+1))
		timeStampStr := time.Unix(tmpStartPeriodInt, 0)
		startPeriodStr[i] = timeStampStr.Format("15:04")
		fmt.Println(startPeriodStr[i])

	}

	for _, f := range logFiles {

		pathF := "../../" + conf.Exporter.Log_Path + f
		openFile, err := os.Open(pathF)

		if err != nil {
			logger("OPEN_FILE_ERR ", err)
			continue
		}
		readFile := bufio.NewScanner(openFile)

		for readFile.Scan() {
			for _, timestamp := range startPeriodStr {

				for _, errStr := range conf.TemplateErr {
					fmt.Println("Check", timestamp, errStr)
					if strings.Contains(readFile.Text(), timestamp+":") && strings.Contains(readFile.Text(), errStr) {
						counter++
					}
				}

			}
		}

		openFile.Close()
	}
	return counter
}
