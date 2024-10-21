package output

import "testing"

var dateParseStructs = []struct {
	input string
	day   int
	month int
	year  int
}{
	{"2021-02-28", 28, 2, 2021},
	{"2019-07-04", 4, 7, 2019},
	{"2025-12-19", 19, 12, 2025},
	{"2030-01-01", 1, 1, 2030},
	{"1980-11-15", 15, 11, 1980},
	{"1965-05-25", 25, 5, 1965},
	{"2028-08-09", 9, 8, 2028},
	{"2022-10-31", 31, 10, 2022},
	{"2026-03-14", 14, 3, 2026},
	{"2015-09-05", 5, 9, 2015},
	{"1930-02-28", 28, 2, 1930},
	{"1960-12-31", 31, 12, 1960},
	{"1999-07-01", 1, 7, 1999},
	{"2012-02-29", 29, 2, 2012}, // Leap year case
	{"2050-02-28", 28, 2, 2050},
	{"2100-02-28", 28, 2, 2100}, // Century non-leap year
	{"2200-02-28", 28, 2, 2200}, // Century non-leap year
	{"2400-02-29", 29, 2, 2400}, // Century leap year
	{"8000-08-15", 15, 8, 8000},
	{"3000-01-01", 1, 1, 3000},
	{"1800-07-04", 4, 7, 1800},
	{"1776-07-04", 4, 7, 1776},
	{"2540-11-11", 11, 11, 2540},
	{"1900-02-28", 28, 2, 1900}, // Century non-leap year
	{"2023-03-01", 1, 3, 2023},
	{"2055-05-30", 30, 5, 2055},
	{"2301-12-25", 25, 12, 2301},
	{"9999-11-11", 11, 11, 9999},
	{"2101-03-15", 15, 3, 2101},
	{"2052-02-29", 29, 2, 2052}, // Leap year case
}

func TestPateParse(t *testing.T) {

	for _, i := range dateParseStructs {
		output1, output2, output3 := DateParse(i.input)
		if output1 != i.day || output2 != i.month || output3 != i.year {
			t.Error("Test failed: input {} day {} month {} year{}, output1{}, output2{}, output3{}", i.input, i.day, i.month, i.year, output1, output2, output3)
		}

	}
}
