package main

import (
	"fmt"
	"time"

	"github.com/cheunn-panaa/eol-checker/configs"
	"github.com/cheunn-panaa/eol-checker/pkg/api"
)

func main() {

	fmt.Println("Loading configuration...")
	config, err := configs.NewConfig("config.yaml")
	if err != nil {
		panic(err)
	}
	client := api.NewHTTPClient(config)
	for _, product := range config.Products {
		CheckProductEol(client, &product, config)
	}

}

// CheckProductEol will call endoflife.date api to check wheter or not the product is readching its eof
func CheckProductEol(client *api.Client, product *configs.Product, config *configs.Config) (bool, error) {
	response, err := client.GetProjectCycle(product)
	if err != nil {
		panic(err)
	}
	if response.EOL.String != nil {
		alertDate := time.Now().AddDate(0, config.Default.Alerting.Month, 0)
		eol, _ := time.Parse("2006-01-02", *response.EOL.String)
		if eol.Before(alertDate) && response.IsSupportedAtDate(alertDate) {
			fmt.Printf("CEST LA PANIC %s, %s, %s \n ", product.Name,
				*response.EOL.String, alertDate)
		}
	}
	return false, err
}
