package main

import (
	"net/http"
	"github.com/tcnksm/go-httpstat"
	"fmt"
	"io/ioutil"
)

func main() {
	req()
	req()
	req()
	req()
	req()
}

func req() {
	client := http.Client{}
	req, _ := http.NewRequest("GET", "https://niwatori-antenna.com/", nil)

	result := new(httpstat.Result)
	ctx := httpstat.WithHTTPStat(req.Context(), result)
	req = req.WithContext(ctx)

	res, _ := client.Do(req)
	defer res.Body.Close()

	ioutil.ReadAll(res.Body)

	fmt.Printf("%+v\n", result)
}
