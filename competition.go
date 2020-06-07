package lfp_scrapper

import (
	"fmt"
	"github.com/gocolly/colly"
	"strconv"
)

type Competition struct {
	Id       int
	Div      int
	Category int
}

func (c Competition) Journees() (journees []Journee, err error) {
	scrap := colly.NewCollector()

	scrap.OnHTML("tbody tr a", func(e *colly.HTMLElement) {
		values, err2 := URLParameters(e.Attr("href"))
		if err2 != nil {
			err = err2

			return
		}

		id, err2 := strconv.Atoi(values.Get("id"))
		if err2 != nil {
			err = err2

			return
		}

		div, err2 := strconv.Atoi(values.Get("div"))
		if err2 != nil {
			err = err2

			return
		}

		cat, err2 := strconv.Atoi(values.Get("cat"))
		if err2 != nil {
			err = err2

			return
		}

		journees = append(journees, Journee{
			Name:     e.Text,
			Id:       id,
			Div:      div,
			Category: cat,
		})
	})

	scrap.OnError(func(response *colly.Response, foundErr error) {
		err = foundErr
	})

	err = scrap.Visit("https://lfp.dz/programme/index?cat=" + fmt.Sprint(c.Category) + "&div=" + fmt.Sprint(c.Div) + "&competition=" + fmt.Sprint(c.Id))

	return journees, err
}
