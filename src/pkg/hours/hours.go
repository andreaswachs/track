package hours

import "time"

func TimeDifference(t1 string, t2 string) (time.Duration, error) {
	layout := "15:04"
	time1, err := time.Parse(layout, t1)
	if err != nil {
		return 0, err
	}

	time2, err := time.Parse(layout, t2)
	if err != nil {
		return 0, err
	}

	diff := time2.Sub(time1)
	if diff < 0 {
		// this handles situations where the first timestamp is later in the day than the second one
		diff += 24 * time.Hour
	}

	return diff, nil
}
