package entity

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type Event struct {
	UserID    string      `db:"user_id"`
	Data      DataRequest `db:"data"`
	EventDate time.Time   `db:"event_date"`
}

type DataRequest struct {
	Headers map[string]any `json:"headers" db:"headers"`
	Body    map[string]any `json:"body" db:"body"`
}

func NewEvent(headers, body map[string]any) *Event {
	return &Event{
		UserID:    headers["X-Tantum-Authorization"].([]string)[0],
		Data:      DataRequest{Body: body, Headers: headers},
		EventDate: time.Now().UTC(),
	}
}

func (e *Event) GetEventDBRecord() map[string]any {
	return map[string]any{
		"user_id":    e.UserID,
		"event_date": e.EventDate,
		"data":       e.Data,
	}
}

func (a DataRequest) Value() (driver.Value, error) {
	return json.Marshal(a)
}

func (a *DataRequest) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}

func Validate(headers, body map[string]any) error {
	v, ok := headers["X-Tantum-Authorization"].([]string)
	if !ok || len(v) == 0 {
		return errors.New("not have user_id in headers")
	}

	if len(body) == 0 {
		return errors.New("request not have body")
	}

	return nil
}
