package commands

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/cheunn-panaa/eol-checker/configs"
	"github.com/cheunn-panaa/eol-checker/internal/utils"
	"github.com/cheunn-panaa/eol-checker/pkg/api"
	_plugins "github.com/cheunn-panaa/eol-checker/pkg/plugins"
	"github.com/spf13/viper"
	"golang.org/x/exp/slices"
)

// Run is a basic impl
func Run() {
	config := configs.GetConfig()
	client := api.NewHTTPClient()

	var productList []_plugins.PluginsMessage
	for _, product := range config.Products {
		val, err := checkProductEol(client, &product)
		if err != nil {
			panic(err)
		}
		if val != nil {
			message := _plugins.MessageBuilder(val, &product)
			productList = append(productList, message)
		}
	}
	if !viper.GetBool("disable-message") {
		sendMessages(config, productList)
	}

}
func sendMessages(config *configs.Configuration, messages []_plugins.PluginsMessage) {
	plugins, _ := _plugins.Load(config)

	pluginList, _ := plugins.GetAllPlugins()
	for plugin := range pluginList {
		i, _ := plugins.GetPlugin(plugin)
		i.SendMessage(messages)
	}
}

func handleSpecific(products []configs.Product, specific string) []configs.Product {
	specificSlice := strings.Split(specific, ",")
	newProductList := []configs.Product{}
	for _, product := range products {
		if slices.Contains(specificSlice, product.Name) {
			newProductList = append(newProductList, product)
		}
	}
	return newProductList
}

// checkProductEol will call endoflife.date api to check wheter or not the product is readching its eof
func checkProductEol(client *api.Client, product *configs.Product) (*api.ProjectCycle, error) {
	alertingPeriod, err := strconv.Atoi(viper.GetString("default.alerting.month"))
	if err != nil {
		alertingPeriod = 12
	}
	response, err := client.GetProjectCycle(product)
	if err != nil {
		panic(err)
	}
	if response.EOL.String != "" {
		alertDate := time.Now().AddDate(0, alertingPeriod, 0)
		eol, _ := time.Parse("2006-01-02", response.EOL.String)
		if eol.Before(alertDate) {
			fmt.Printf("%s version %s, end's of life is on %s \n ", product.Name, product.Version,
				utils.DateFormat(eol, "02-Jan-2006"))

			latest, _ := client.GetProjectLatestRelease(product)
			response.LatestCycle = latest.Cycle
			return response, nil
		}
	} else if response.EOL.Bool {
		return response, nil
	}
	return nil, nil
}
