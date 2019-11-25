package main

import (
	"encoding/csv"
	"github.com/pinkumohikan/java_sparrow/batch/business_day"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"log"
	"net/http"
	"time"
)

func main() {
	holidays, err := business_day.FetchPublicHoliday()
	if err != nil {
		log.Fatalln(err)
	}

	for _, h := range holidays {
		log.Printf("%#v", h)
	}
}

func FetchPublicHoliday() ([]PublicHoliday, error) {
	const PublicHolidayCsvUrl = "https://www8.cao.go.jp/chosei/shukujitsu/syukujitsu.csv"
	const IndexHeaderLine = 0

	r, err := http.Get(PublicHolidayCsvUrl)
	if err != nil {
		return nil, err
	}

	sjisReader := transform.NewReader(r.Body, japanese.ShiftJIS.NewDecoder())
	records, err := csv.NewReader(sjisReader).ReadAll()
	if err != nil {
		return nil, err
	}

	var holidays []PublicHoliday

	for i, v := range records {
		if i == IndexHeaderLine {
			continue
		}

		d, err := time.Parse("2006/1/2", v[0])
		if err != nil {
			return nil, err
		}

		holidays = append(holidays, PublicHoliday{
			Date: d,
			Name: v[1],
		})
	}

	return holidays, nil
}

type PublicHoliday struct {
	Date time.Time
	Name string
}
