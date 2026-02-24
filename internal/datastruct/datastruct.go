package datastruct

import (
	"fmt"
	"shopapi/internal/supports"
	"strings"
	"time"
)

type Status struct {
	Message string `json:"status,omitempty" example:"status message"`
}

func (s Status) GetStatus() string {
	return s.Message
}

type CachedStatus struct {
	Cached bool `json:"cached" schema:"cached" example:"false"`
}

func (r *CachedStatus) SetCached(is bool) {
	r.Cached = is
}

type AvoidCacheFlag struct {
	Flag bool `schema:"avoid_cache" json:"avoid_cache" example:"true"`
}

func (a *AvoidCacheFlag) AvoidCache() bool {
	return a.Flag
}

const (
	StatusNotFound      = "resource not found"
	StatusServiceError  = "service failed exec request"
	StatusAlreadyExists = "resource already exists"
	StatusOK            = "Success"

	OffsetParam        = "offset"
	LimitParam         = "limit"
	ClientNameParam    = "client_name"
	ClientSurnameParam = "client_surname"
)

var dateFormats = []string{
	"2006-01-02",
	"2006/01/02",
	"2006.01.02",
	"02-01-2006",
	"02.01.2006",
	"02/01/2006",
}

type DateOnly time.Time

// Panics in case string has wrong format
func DateOnlyFromString(t string) (do DateOnly) {
	err := do.UnmarshalJSON([]byte(t))
	if err != nil {
		panic(err)
	}
	return do
}

func (d *DateOnly) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")

	for _, f := range dateFormats {
		t, err := time.Parse(f, s)
		if err != nil {
			continue
		}

		*d = DateOnly(t)
		return nil
	}

	return fmt.Errorf("incorrect date format: '%s'", s)
}

func (d *DateOnly) MarshalJSON() ([]byte, error) {
	return []byte(supports.Concat("\"", time.Time(*d).Format(time.DateOnly), "\"")), nil
}
