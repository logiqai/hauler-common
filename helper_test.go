package hauler_common

import (
	"testing"
	"time"
)

var (
	fixture1 = []byte(`{"event":{"message":"test","severity":"INFO","timestamp":"2021-03-31T20:00:00.000Z"}}`)
	fixture2 = []byte(`{"event":{"message":"test","severity":"INFO","timestamp":1617219200000}}`)
	fixture3 = []byte(`{"event":{"message":"test","severity":"INFO","timestamp":1617219200}}`)
)

func TestTimeParse(t *testing.T) {
	tn := time.Now()
	tStr := tn.Format(time.RFC3339)
	tInt := tn.UnixMilli()
	tFloat := float64(tInt)
	tMap := map[string]interface{}{}

	timeParse(tStr, tMap)
	if tMap["timestamp"] != tStr {
		t.Errorf("Expected %s, got %s", tStr, tMap["timestamp"])
	}
	timeParse(tInt, tMap)
	if tMap["timestamp"] != tStr {
		t.Errorf("Expected %s, got %s", tStr, tMap["timestamp"])
	}
	timeParse(tFloat, tMap)
	if tMap["timestamp"] != tStr {
		t.Errorf("Expected %s, got %s", tStr, tMap["timestamp"])
	}
}

func TestParseEventFixture1(t *testing.T) {
	m, err := ParseEvent(fixture1, "default")
	if err != nil {
		t.Errorf("Error parsing event: %v", err)
	}
	if m["message"] != "test" {
		t.Errorf("Expected message to be 'test', got %s", m["message"])
	}
	if m["severity"] != "INFO" {
		t.Errorf("Expected severity to be 'INFO', got %s", m["severity"])
	}
	if m["timestamp"] != "2021-03-31T20:00:00.000Z" {
		t.Errorf("Expected timestamp to be '2021-03-31T20:00:00.000Z', got %s", m["timestamp"])
	}
}

func TestParseEventFixture2(t *testing.T) {
	m, err := ParseEvent(fixture2, "default")
	if err != nil {
		t.Errorf("Error parsing event: %v", err)
	}
	if m["message"] != "test" {
		t.Errorf("Expected message to be 'test', got %s", m["message"])
	}
	if m["severity"] != "INFO" {
		t.Errorf("Expected severity to be 'INFO', got %s", m["severity"])
	}
	if m["timestamp"] != "2021-03-31T12:33:20-07:00" {
		t.Errorf("Expected timestamp to be '2021-03-31T12:33:20-07:00', got %s", m["timestamp"])
	}
}

func TestParseEventFixture3(t *testing.T) {
	m, err := ParseEvent(fixture3, "default")
	if err != nil {
		t.Errorf("Error parsing event: %v", err)
	}
	if m["message"] != "test" {
		t.Errorf("Expected message to be 'test', got %s", m["message"])
	}
	if m["severity"] != "INFO" {
		t.Errorf("Expected severity to be 'INFO', got %s", m["severity"])
	}
	if m["timestamp"] != "2021-03-31T12:33:20-07:00" {
		t.Errorf("Expected timestamp to be '2021-03-31T12:33:20-07:00', got %s", m["timestamp"])
	}
}
