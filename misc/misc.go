package misc

import (
	"errors"
	"strconv"
	"time"
)

// IsDateJp returns true if the passed string has a 'japanese' YYYYMMDD format date
func IsDateJp(date string) bool {
	if len(date) != 8 {
		return false
	}

	year, er := strconv.Atoi(date[:4])
	if er != nil {
		return false
	} else if year < 1000 {
		return false
	}

	month, er := strconv.Atoi(date[4:6])
	if er != nil {
		return false
	} else if month < 1 || month > 12 {
		return false
	}

	day, er := strconv.Atoi(date[6:])
	if er != nil {
		return false
	} else if day < 1 || day > 31 {
		return false
	} else if day == 29 && month == 02 && year%4 != 0 {
		return false
	}

	return true
}

// DateTodayJp returns a string with today's date in 'japanese' YYYYMMDD format
func DateTodayJp() string {
	today := time.Now().String()
	today = today[:10]
	year := today[:4]
	month := today[5:7]
	day := today[8:10]

	return year + month + day
}

// DateJpToEs returns a string with the passed 'japanese' YYYYMMDD date to a
// 'spanish' DD/MM/YYYY (separator must be included)
func DateJpToEs(date, sep string) (string, error) {
	if IsDateJp(date) {
		ye := date[:4]
		mo := date[4:6]
		da := date[6:]
		return da + sep + mo + sep + ye, nil
	}

	return "", errors.New("Incorrect date: " + date)
}
