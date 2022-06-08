package seed_cmd

import (
	"crypto/rand"
	"fmt"
	"job-go/driver"
	"job-go/pkg/logr"
	"math/big"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var CmdSeedBillingUsageHour = &cobra.Command{
	Use:   "seed::billing-usage-hour",
	Short: "billing usage hour 相關測試資料",
	Long:  `billing usage hour 相關測試資料`,
	Run:   jobSeedBillingUsageHour,
}

func jobSeedBillingUsageHour(cmd *cobra.Command, args []string) {
	seesaw := cmd.Flag("seesaw").Value.String() // up, down

	switch seesaw {
	case "up":
		seedBillingUsageHourUp()
	case "down":
		seedBillingUsageHourDown()
	default:
		logr.L.Info("seed::billing-usage-hour seesaw not found. (seesaw:" + seesaw + ")")
		return
	}
}

func seedBillingUsageHourUp() {
	orm, _ := driver.NewXorm()
	defer orm.Close()

	month := 6

	start, _ := time.Parse("2006-01-02", fmt.Sprintf("2022-0%d-01", month))
	end, _ := time.Parse("2006-01-02", fmt.Sprintf("2022-0%d-20", month))

	fmt.Println("start:", start, "end:", end)

	rowRandomArray := []int{2, 20, 352, 64, 125, 1, 487, 50, 200, 367, 41, 106, 267, 84, 307, 147, 191, 409, 453, 365, 213}
	rowRandomArrayRange := len(rowRandomArray)

	keyList := []string{"key_1", "key_2"}

	skuPrices := []SkuPrice{
		{1, 1, month, 240, 0.2},
		{2, 2, month + 12, 1000, 0.72},
		{3, 3, month + 24, 2000, 10},
	}

	hourDiff := int(end.Sub(start).Hours())
	for _, key := range keyList {
		for _, price := range skuPrices {
			stmt := fmt.Sprintf("INSERT INTO billing_usage_hour ( %s ) VALUES", getBillingUsageHourFields())

			for i := 1; i <= hourDiff; i++ {
				tempTime := start.Add(time.Duration(i) * time.Hour)

				random, _ := rand.Int(rand.Reader, big.NewInt(int64(rowRandomArrayRange)))
				rowAmount := rowRandomArray[int(random.Int64())]

				cost := price.Price * float64(rowAmount)

				invStmt := fmt.Sprintf("( '%s', '%s', %d, %d, %d, '%.3f', '%d', '%.3f')",
					key,
					tempTime.Format("2006-01-02 15:04:05"),
					price.ServiceId,
					price.SkuId,
					price.PriceId,
					price.Price,
					rowAmount,
					cost,
				)

				if i < hourDiff {
					stmt = stmt + invStmt + ","
				}

				if i == hourDiff {
					stmt = stmt + invStmt + ";"
				}
			}

			_, err := orm.Query(stmt)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
		}
	}
}

func seedBillingUsageHourDown() {

}

func getBillingUsageHourFields() string {
	fields := []string{
		"`key`",
		"time_by_hour",
		"service_id",
		"sku_id",
		"price_id",
		"price",
		"`usage`",
		"cost",
	}

	return strings.Join(fields, ", ")
}

type SkuPrice struct {
	ServiceId int     `json:"serviceId"`
	SkuId     int     `json:"skuId"`
	PriceId   int     `json:"priceId"`
	Moq       int     `json:"moq"`
	Price     float64 `json:"price"`
}
