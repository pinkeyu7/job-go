package task_cmd

import (
	"fmt"
	"job-go/pkg/helper"
	"job-go/pkg/logr"

	"github.com/spf13/cobra"
)

var CmdTaskBillingCurl = &cobra.Command{
	Use:   "task::billing-curl",
	Short: "usage 相關測試資料",
	Long:  `usage 相關測試資料`,
	Run:   jobBillingCurl,
}

const (
	ServerUrl = "http://localhost:3000"

	UrlCreateService = "/services"
	UrlCreateSku     = "/services/%d/skus"
	UrlCreatePrice   = "/services/%d/skus/%d/prices"
)

type Sku struct {
	ServiceId int    `json:"serviceId"`
	Resource  string `json:"resource"`
	Operation string `json:"operation"`
	UsageType string `json:"usageType"`
}

type SkuPrice struct {
	ServiceId    int     `json:"serviceId"`
	SkuId        int     `json:"skuId"`
	Moq          int     `json:"moq"`
	Price        float64 `json:"price"`
	PriceOverMoq float64 `json:"priceOverMoq"`
	Description  string  `json:"description"`
}

func jobBillingCurl(cmd *cobra.Command, args []string) {
	dataType := cmd.Flag("dataType").Value.String() // monthYm

	services := []string{"asr", "tts", "news"}
	skus := []Sku{
		{1, "asr_standard", "streaming-15secs-bytes", "standard"},
		{2, "tts_16K", "1M-characters", "standard"},
		{3, "news_3M", "query-1-time", "standard"},
	}

	skuPrices := []SkuPrice{
		{1, 1, 240, 0.2, 0.2, "asr_standard"},
		{2, 2, 1000, 0.72, 0.72, "tts_16K"},
		{3, 3, 2000, 10, 20, "news_3M"},
	}

	switch dataType {
	case "service":
		createService(services)
	case "sku":
		createSku(skus)
	case "price":
		startYm := cmd.Flag("startYm").Value.String() // startYm
		endYm := cmd.Flag("endYm").Value.String()     // endYm

		queryYm, err := helper.GetQueryYm(startYm, endYm)
		if err != nil {
			logr.L.Error(fmt.Sprintf("task::billing-curl time error. (data type: %s)(startYm: %s)(endYm: %s)", dataType, startYm, endYm))
			return
		}
		createPrice(skuPrices, queryYm)
	default:
		logr.L.Error(fmt.Sprintf("task::billing-curl dataType error. (data type: %s)", dataType))
	}
	logr.L.Info(fmt.Sprintf("task::billing-curl finished. (data type: %s)", dataType))
}

func createPrice(skuPrices []SkuPrice, queryYm []string) {
	logr.L.Info("task::billing-curl create price start.")
	for _, price := range skuPrices {
		for _, monthYm := range queryYm {
			url := fmt.Sprintf(fmt.Sprintf("%s%s", ServerUrl, UrlCreatePrice), price.ServiceId, price.SkuId)

			payload := map[string]interface{}{
				"monthYm":      monthYm,
				"moq":          price.Moq,
				"price":        price.Price,
				"priceOverMoq": price.PriceOverMoq,
				"description":  price.Description}

			err := helper.CreatePost(url, payload)
			if err != nil {
				logr.L.Error("task::billing-curl create price error.")
				return
			}
		}
	}
	logr.L.Info("task::billing-curl create sku finished.")
}

func createSku(skus []Sku) {
	logr.L.Info("task::billing-curl create sku start.")
	for _, sku := range skus {
		url := fmt.Sprintf(fmt.Sprintf("%s%s", ServerUrl, UrlCreateSku), sku.ServiceId)

		payload := map[string]interface{}{
			"resource":  sku.Resource,
			"operation": sku.Operation,
			"usageType": sku.UsageType}

		err := helper.CreatePost(url, payload)
		if err != nil {
			logr.L.Error("task::billing-curl create sku error.")
			return
		}
	}
	logr.L.Info("task::billing-curl create sku finished.")
}

func createService(services []string) {
	logr.L.Info("task::billing-curl create service start.")
	for _, service := range services {
		url := fmt.Sprintf("%s%s", ServerUrl, UrlCreateService)
		payload := map[string]interface{}{
			"name": service}

		err := helper.CreatePost(url, payload)
		if err != nil {
			logr.L.Error("task::billing-curl create service error.")
			return
		}
	}
	logr.L.Info("task::billing-curl create service finished.")
}
