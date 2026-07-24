package client

import "time"

func BusinessDays() float64 {

	var businessDays []time.Time
	now := time.Now()

	// Get the first and last day of the current month
	// time.Date rolls over naturally, so day '0' of the next month is the last day of this month
	firstDay := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	lastDay := time.Date(now.Year(), now.Month()+1, 0, 0, 0, 0, 0, now.Location())
	saved_days := 0.0

	// Loop through every day of the month
	for d := firstDay; !d.After(lastDay); d = d.AddDate(0, 0, 1) {
		// Filter out Saturday and Sunday
		if d.Weekday() != time.Saturday && d.Weekday() != time.Sunday {
			businessDays = append(businessDays, d)
			saved_days++
		}
	}

	return saved_days

}

func CalculateValue(found float64, businessdays float64) float64 {

	const daily = 152.0

	salary_gbp := daily * businessdays
	salary_brl := salary_gbp * found

	return salary_brl

}
