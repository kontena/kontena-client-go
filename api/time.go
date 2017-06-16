package api

import "time"

// API JSON timestamp formats
//  "2017-06-14T13:38:03.785Z"
//  "2017-06-14T12:33:05.084+00:00"
type Time time.Time

func (t *Time) UnmarshalJSON(data []byte) error {
	if parseTime, err := time.Parse(`"2006-01-02T15:04:05.000Z"`, string(data)); err == nil {
		*t = Time(parseTime)
	} else if parseTime, err := time.Parse(`"2006-01-02T15:04:05.000+00:00"`, string(data)); err == nil {
		*t = Time(parseTime)
	} else {
		return err
	}

	return nil
}
