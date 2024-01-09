package main

import (
	"testing"
)

func TestFormatTimestamp(t *testing.T) {
	input := "03/Jan/2024:09:20:32 +0100"
	expected := "2024-01-03T09:20:32+01:00"
	if output, _ := formatTimestamp(input); output != expected {
		t.Errorf("formatTimestamp(\"%s\") returned \"%s\", expected \"%s\".", input, output, expected)
	}
}

func TestParseLogEntryParses(t *testing.T) {
	input := `192.168.1.10 - 0hLNPuxTM7JscWD [08/Jan/2024:11:44:16 +0100] "GET /foo/bar?baz=no HTTP/1.1" 200 33 "http://example.com/" "Mozilla/5.0 (X11; Linux x86_64; rv:102.0) Gecko/20100101 Firefox/102.0"`
	_, err := parseLogEntry(input)
	if err != nil {
		t.Errorf("parseLogEntry(\"%s\") returned an error: \"%s\"", input, err)
	}
}
