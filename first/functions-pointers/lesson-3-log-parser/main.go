package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const formatError = "you have an incorrect format in the input"

type log struct {
	domain       string
	timesVisited int
}

type foundLog struct {
	log
	index int
}

func main() {
	var lines []string

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	logs, err := getAllLogs(lines)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%-3s %-20s %20s\n", "👇", "DOMAIN", "TIME VISITED")
	fmt.Println(strings.Repeat("-", 46))

	for i, _log := range logs {
		fmt.Printf("%-3d %-20s %20d\n", i+1, _log.domain, _log.timesVisited)
	}
}

func getAllLogs(lines []string) ([]log, error) {
	if len(lines) == 0 {
		return nil, nil
	}

	logs := make([]log, 0, len(lines))

	for _, line := range lines {
		if parts := strings.Split(line, " "); len(parts) == 2 {
			newLog, err := getLog(parts, logs)

			if err == nil {
				if len(logs) > 0 && newLog.domain == logs[newLog.index].domain{
					logs[newLog.index] = newLog.log
				} else {
					logs = append(logs, newLog.log)
				}
			} else {
				return nil, errors.New(formatError)
			}
		} else {
			return nil, errors.New(formatError)
		}
	}

	return logs, nil
}

func getLog(stringParts []string, logs []log) (foundLog, error) {
	domain := stringParts[0]
	__foundLog := findLogWithDomain(logs, domain)
	var newLog foundLog

	if __foundLog == nil {
		newLog.domain = domain
	} else {
		newLog = __foundLog[0]
	}

	if num, convErr := strconv.Atoi(stringParts[1]); convErr == nil {
		newLog.timesVisited += num
	} else {
		return foundLog{}, convErr
	}

	return newLog, nil
}

func findLogWithDomain(logs []log, domain string) []foundLog {
	if len(logs) == 0 {
		return nil
	}

	for i, _log := range logs {
		if _log.domain == domain {
			return []foundLog{
				{
					log:   _log,
					index: i,
				},
			}
		}
	}

	return nil
}
