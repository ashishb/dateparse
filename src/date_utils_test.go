package dateparse

import (
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"
)

var testDataFilename = filepath.Join("testdata", "date_entries.txt")

func TestGetParsedDate(t *testing.T) {
	data, err := ioutil.ReadFile(testDataFilename)
	if err != nil {
		t.Errorf("Failed to read file %s: %v", testDataFilename, err)
		t.Fail()
	}
	for _, dateEntry := range strings.Split(string(data), "\n") {
		if len(dateEntry) == 0 {
			continue
		}
		_, err := GetParsedDate(dateEntry)
		if err != nil {
			t.Errorf("Failed to parse date from \"%s\": %v", dateEntry, err)
		}
		// t.Logf(`Parsed date is "%v"`, parsedDate)
	}
}
