package main

import (
	"testing"
)

func TestSingleLineParsing(t *testing.T) {
	parsedEntities := make([]Entity, 0)
	parsedEntities = outputPdfText("resources/test/single_line_record.pdf", parsedEntities)
	if len(parsedEntities) == 0 {
		t.Errorf("Single record is expected to parsed correctly")
	}
	en := parsedEntities[0]
	if en.ID != "5.1" {
		t.Errorf("expected item id not found")
	}
}
