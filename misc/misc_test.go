package misc

import (
	"fmt"
	"testing"
)

func TestIsDateJp(t *testing.T) {
	date := "20201105"
	fmt.Println("Date", date)
	isjp := IsDateJp(date)

	if !isjp {
		t.Errorf("Date %s should be returned as true", date)
	}
}

func TestDateJpToEs(t *testing.T) {
	date := "20201105"
	fmt.Println("Date:", date)
	esdate, _ := DateJpToEs(date, "/")

	if esdate != "05/11/2020" {
		t.Errorf("Date %s, returned as %s, should be %s", date, esdate, "05/11/2020")
	}
}
