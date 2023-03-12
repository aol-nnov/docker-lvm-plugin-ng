package lvm

import (
	"strings"
	"time"
)

type CreatedTime time.Time

// https://stackoverflow.com/a/45304122
func (c *CreatedTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02 15:04:05 -0700", s)
	if err != nil {
		return err
	}
	*c = CreatedTime(t)
	return nil
}

func (c *CreatedTime) Format(s string) string {
	t := time.Time(*c)
	return t.Format(s)
}
