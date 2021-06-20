package timer

import (
	"time"
	"github.com/pkg/errors"
)


func GetCurrentTime() time.Time {
	return time.Now()
}

func GetCalculateTime(currentTimer time.Time, d string) (time.Time, error) {
	duration, err := time.ParseDuration(d)
	if err != nil {
		return time.Time{}, errors.Wrap(err, "get calculate error")
	}

	return currentTimer.Add(duration), nil
}


