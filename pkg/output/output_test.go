package output

import "testing"

var dateParseStructs = []struct {
	name  string
	input string
	day   int
	month int
	year  int
}{
	{"date 1", "2021-02-28", 28, 2, 2021},
	{"date 2", "2019-07-04", 4, 7, 2019},
	{"date 3", "2025-12-19", 19, 12, 2025},
	{"date 4", "2030-01-01", 1, 1, 2030},
	{"date 5", "1980-11-15", 15, 11, 1980},
	{"date 6", "1965-05-25", 25, 5, 1965},
	{"date 7", "2028-08-09", 9, 8, 2028},
	{"date 8", "2022-10-31", 31, 10, 2022},
	{"date 9", "2026-03-14", 14, 3, 2026},
	{"date 10", "2015-09-05", 5, 9, 2015},
	{"date 11", "1930-02-28", 28, 2, 1930},
	{"date 12", "1960-12-31", 31, 12, 1960},
	{"date 13", "1999-07-01", 1, 7, 1999},
	{"date 14 leap year case", "2012-02-29", 29, 2, 2012},
	{"date 15", "2050-02-28", 28, 2, 2050},
	{"date 16 century non-leap year", "2100-02-28", 28, 2, 2100},
	{"date 17 century non-leap year", "2200-02-28", 28, 2, 2200},
	{"date 18 century leap year", "2400-02-29", 29, 2, 2400},
	{"date 19", "8000-08-15", 15, 8, 8000},
	{"date 20", "3000-01-01", 1, 1, 3000},
	{"date 21", "1800-07-04", 4, 7, 1800},
	{"date 22", "1776-07-04", 4, 7, 1776},
	{"date 23", "2540-11-11", 11, 11, 2540},
	{"date 24 century non-leap year", "1900-02-28", 28, 2, 1900},
	{"date 25", "2023-03-01", 1, 3, 2023},
	{"date 26", "2055-05-30", 30, 5, 2055},
	{"date 27", "2301-12-25", 25, 12, 2301},
	{"date 28", "9999-11-11", 11, 11, 9999},
	{"date 29", "2101-03-15", 15, 3, 2101},
	{"date 30 leap year case", "2052-02-29", 29, 2, 2052},
	// {"date 31", "2222-02-29", 29, 2, 2222, true}, // not a leap year, you can uncomment it to test if it will be detected
}

func TestDateParse(t *testing.T) {

	for _, i := range dateParseStructs {
		output1, output2, output3 := DateParse(i.input)
		if output1 != i.day || output2 != i.month || output3 != i.year {
			t.Error("Test failed: input {} day {} month {} year{}, output1{}, output2{}, output3{}", i.input, i.day, i.month, i.year, output1, output2, output3)
		}

	}
}
