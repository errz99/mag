package misc

import (
    "fmt"
	"errors"
	"strconv"
	"strings"
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
// 'english' MM/DD/YYYY or 'spanish' DD/MM/YYYY (separator must be included)
func DateJpToEnEs(date, c, sep string) (string, error) {
	if IsDateJp(date) {
		year := date[:4]
		month := date[4:6]
		day := date[6:]

		switch strings.ToLower(c) {
		case "es":
			return day + sep + month + sep + year, nil
		case "en":
			return month + sep + day + sep + year, nil
		default:
			return year + sep + month + sep + day, nil
		}
	}

	return "", errors.New("Incorrect date: " + date)
}

// ProcessColor processes a string like 'rgb(255,0,0)' or 'rgba(255,0,0,0.50000)'
// and returns a string like '#ff0000
func ProcessColor(temp string) string {
		color := "#000000"

		temp = strings.Replace(temp, "rgb", "", -1)
		temp = strings.Replace(temp, "a", "", -1)
		temp = strings.Replace(temp, "(", "", -1)
		temp = strings.Replace(temp, ")", "", -1)
		colors := strings.Split(temp, ",")

		if len(colors) >= 3 {
			red, _ := strconv.Atoi(colors[0])
			green, _ := strconv.Atoi(colors[1])
			blue, _ := strconv.Atoi(colors[2])
			color = fmt.Sprintf("#%0.2x%0.2x%0.2x", red, green, blue)
		}
		return color
	}
}	
