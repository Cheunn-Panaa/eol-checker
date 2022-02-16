package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/cheunn-panaa/eol-checker/configs"
)

//Response type from endoflife.date Interface when its booleanstring ???
type Response struct {
	Cycle          string      `json:"cycle"`
	Release        string      `json:"release"`
	Eol            interface{} `json:"eol,omitempty"`
	Latest         string      `json:"latest"`
	Link           string      `json:"link,omitempty"`
	Lts            string      `json:"lts"`
	Support        interface{} `json:"support,omitempty"`
	CycleShortHand string      `json:"cycleShortHand,omitempty"`
	Discontinued   interface{} `json:"discontinued,omitempty"`
}

func main() {

	fmt.Println("Loading configuration...")
	config, err := configs.NewConfig("config.yaml")
	if err != nil {
		panic(err)
	}
	for _, product := range config.Products {
		boo, _ := CheckProductEof(&product, config)
		fmt.Println(boo)
	}

}

// CheckProductEof will call endoflife.date api to check wheter or not the product is readching its eof
func CheckProductEof(product *configs.Product, config *configs.Config) (bool, error) {

	//fmt.Println("Calling API...")
	client := &http.Client{}

	req, err := http.NewRequest("GET", generateUrl(config, product), nil)
	if err != nil {
		fmt.Print(err.Error())
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}

	var response Response

	json.Unmarshal(bodyBytes, &response)
	//fmt.Printf("API Response as struct %+v\n", response)
	if response.Eol != nil {

		switch eolT := response.Eol.(type) {
		case string:
			alertDate := time.Now().AddDate(0, config.Default.Alerting.Month, 0)
			eol, _ := time.Parse("2006-01-02", eolT)
			if eol.Before(alertDate) && response.Support != true {
				fmt.Println(fmt.Sprintf("CEST LA PANIC %s, %s, %s", product.Name, eolT, alertDate))
			}
		}
	}
	return false, err
}

func generateUrl(config *configs.Config, product *configs.Product) string {
	url := fmt.Sprintf("%s/api/%s.json",
		config.Default.Url, product.Name)
	if product.Version != "" {
		url = fmt.Sprintf("%s/api/%s/%s.json",
			config.Default.Url, product.Name, product.Version)

		//fmt.Println("Checking End of Life for %s with version %s", product.Name, product.Version)
	}
	return url
}
