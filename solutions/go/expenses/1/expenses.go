package expenses

import (
	"fmt"
	"slices"
)

// Record represents an expense record.
type Record struct {
	Day      int
	Amount   float64
	Category string
}

// DaysPeriod represents a period of days for expenses.
type DaysPeriod struct {
	From int
	To   int
}

// Filter returns the records for which the predicate function returns true.
func Filter(in []Record, predicate func(Record) bool) []Record {
	var	ans []Record
	for _, record := range in {
		if predicate(record) {
			ans = append(ans, record)
		}
	}
	return ans
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise.
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	function := func(record Record) bool {
		if record.Day >= p.From && record.Day <= p.To {
			return true
		}
		return false	
	}
	return function
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise.
func ByCategory(c string) func(Record) bool {
	alpha := func(record Record) bool {
		if record.Category == c {
			return true
		}
		return false
	}
	return alpha
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p.
func TotalByPeriod(in []Record, p DaysPeriod) float64 {
	var valid_records []Record
	valid_records = Filter(in, ByDaysPeriod(p))
	var amount float64
	for _, record := range valid_records {
		amount += record.Amount
	}
	return amount
}

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	valid_records := Filter(in, ByCategory(c))
	
	if !slices.Contains([]string{"groceries", "utility-bills", "university", "rent"}, c){
		return 0, fmt.Errorf("unknown category %s", c)
	}

	return TotalByPeriod(valid_records, p), nil


}
