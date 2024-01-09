package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"
)

type LogEntry struct {
	IP              string `json:"ip"`
	RemoteLogName   string `json:"remoteLogName"`
	User            string `json:"user"`
	Timestamp       string `json:"timestamp"`
	Method          string `json:"method"`
	URL             string `json:"url"`
	ProtocolVersion string `json:"protocolVersion"`
	Status          int    `json:"status"`
	Size            int    `json:"size"`
	Referer         string `json:"referer"`
	UserAgent       string `json:"userAgent"`
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		entry, err := parseLogEntry(line)
		if err != nil {
			fmt.Println("Error parsing line:", err)
			continue
		}

		jsonData, err := json.Marshal(entry)
		if err != nil {
			fmt.Println("Error marshaling to JSON:", err)
			continue
		}

		fmt.Println(string(jsonData))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from stdin:", err)
	}
}

func parseLogEntry(line string) (*LogEntry, error) {
	re := regexp.MustCompile(`(\S+) (\S+) (\S+) \[([^\]]+)\] "(\S+) (\S+) (\S+)" (\d+) (\d+) "([^"]+)" "([^"]+)"`)
	matches := re.FindStringSubmatch(line)

	if len(matches) != 12 {
		return nil, fmt.Errorf("invalid log entry format")
	}

	timestamp, err := formatTimestamp(matches[4])
	if err != nil {
		return nil, err
	}

	status, _ := strconv.Atoi(matches[8])
	size, _ := strconv.Atoi(matches[9])

	entry := LogEntry{
		IP:              matches[1],
		RemoteLogName:   matches[2],
		User:            matches[3],
		Timestamp:       timestamp,
		Method:          matches[5],
		URL:             matches[6],
		ProtocolVersion: matches[7],
		Status:          status,
		Size:            size,
		Referer:         matches[10],
		UserAgent:       matches[11],
	}

	return &entry, nil
}

func formatTimestamp(timestamp string) (string, error) {
	t, err := time.Parse("02/Jan/2006:15:04:05 -0700", timestamp)
	if err != nil {
		return "", err
	}
	return t.Format(time.RFC3339), nil
}
