package main

import (
"github.com/go-test/deep"
"testing"
"time"
)

func TestFetchPublicHoliday(t *testing.T) {
	// 実際にリクエスト飛んじゃうけど、少々ええやろ...w
	fetched, err := FetchPublicHoliday()
	if err != nil {
		t.Error(err)
	}

	d, _ := time.Parse("2006/1/2", "1955/1/1")
	expected := PublicHoliday{
		Date: d,
		Name: "元日",
	}

	if err := deep.Equal(expected, fetched[0]); err != nil {
		t.Error(err)
	}
}
