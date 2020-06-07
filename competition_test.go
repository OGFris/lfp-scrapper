package lfp_scrapper

import (
	"fmt"
	"testing"
)

func Test_Journees(t *testing.T) {
	competition := Competition{
		Id:       1,
		Div:      1,
		Category: 1,
	}

	journees, err := competition.Journees()

	if err != nil {
		t.Error(err)
	}

	fmt.Println(journees)
}
