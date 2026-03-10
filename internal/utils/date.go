package utils

import (
	"time"
)

type JSONDate time.Time

func (d JSONDate) MarshalJSON() ([]byte, error) {
	t := time.Time(d)
	formatted := t.Format("2006-01-02")
	return []byte(`"` + formatted + `"`), nil
}
