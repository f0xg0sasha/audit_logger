package audit

import (
	"time"
)

type LogItem struct {
	Entity    string    `son:"entity"`
	Action    string    `son:"action"`
	EntityID  int64     `son:"entity_id"`
	Timestamp time.Time `son:"timestamp"`
}
