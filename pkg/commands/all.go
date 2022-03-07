package commands

import (
	"fmt"
	"time"

	"github.com/cheunn-panaa/eol-checker/configs"
	"github.com/cheunn-panaa/eol-checker/pkg/api"
	_plugins "github.com/cheunn-panaa/eol-checker/pkg/plugins"
)

// RunAll is a basic impl
func RunAll() {
	config := configs.GetConfig()
	client := api.NewHTTPClient(config)

	plugins, _ := _plugins.Load(config)
	var productList []_plugins.PluginsMessage
	for _, product := range config.Products {
		val, err := checkProductEol(client, &product, config)
		if err != nil {
			panic(err)
		}
		if val != nil {
			message := _plugins.MessageBuilder(val, &product)
			productList = append(productList, message)
		}
	}

	slack, _ := plugins.GetPlugin("slack")
	slack.SendMessage(productList)

}

// checkProductEol will call endoflife.date api to check wheter or not the product is readching its eof
func checkProductEol(client *api.Client, product *configs.Product, config *configs.Configuration) (*api.ProjectCycle, error) {
	response, err := client.GetProjectCycle(product)
	if err != nil {
		panic(err)
	}
	if response.EOL.String != "" {
		alertDate := time.Now().AddDate(0, config.Default.Alerting.Month, 0)
		eol, _ := time.Parse("2006-01-02", response.EOL.String)
		if eol.Before(alertDate) {
			fmt.Printf("CEST LA PANIC %s, %s, %s \n ", product.Name,
				eol, alertDate)

			latest, _ := client.GetProjectLatestRelease(product)
			response.LatestCycle = latest.Cycle
			return response, nil
		}
	} else if response.EOL.Bool {
		return response, nil
	}
	return nil, nil
}
