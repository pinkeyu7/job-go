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

var CmdSeedUsage = &cobra.Command{
	Use:   "seed::usage",
	Short: "usage 相關測試資料",
	Long:  `usage 相關測試資料`,
	Run:   jobSeedUsage,
}

func jobSeedUsage(cmd *cobra.Command, args []string) {
	seesaw := cmd.Flag("seesaw").Value.String() // up, down
	scale := cmd.Flag("scale").Value.String()   // thousand, million, tem-million

	switch seesaw {
	case "up":
		seedUsageUp(scale)
	case "down":
		seedUsageDown(scale)
	default:
		logr.L.Info("seed::usage seesaw not found. (scale:" + scale + ",seesaw:" + seesaw + ")")
		return
	}
}

func seedUsageUp(scale string) {
	switch scale {
	case "million":
		seedUsageUpMillion()
	default:
		logr.L.Info("seed::usage scale not found. (scale:" + scale + ",seesaw:up)")
		return
	}
	logr.L.Info("seed::usage scale:" + scale + " seesaw:up finished")
}

func seedUsageDown(scale string) {
	switch scale {
	case "million":
		seedUsageDownMillion()
	default:
		logr.L.Info("seed::usage scale not found. (scale:" + scale + ",seesaw:down)")
		return
	}
	logr.L.Info("seed::usage scale:" + scale + " seesaw:down finished")
}

func seedUsageUpMillion() {
	orm, _ := driver.NewPostgresql()
	defer orm.Close()

	keyAmount := 10
	dayAmount := 10
	hourAmount := dayAmount * 24

	// 10 key * 25 day * 24 hour * 2 service * 75 row per hour(average) = 900,000 rows

	total := 0

	packRandomArray := []int{1, 3, 8, 11, 20, 33, 41, 50}
	packRandomArrayRange := len(packRandomArray) - 1

	rowRandomArray := []int{2, 20, 352, 64, 125, 1, 487, 50, 200, 367, 41, 106, 267, 84, 307, 147, 191, 409, 453, 365, 213}
	rowRandomArrayRange := len(rowRandomArray) - 1

	startDate, _ := time.Parse("20060102", "20220101")

	usageFields := getUsageFields()

	skuIdArray := []string{"1", "2", "3", "4"}
	skuIdArrayRange := len(skuIdArray) - 1

	for i := 1; i < keyAmount; i++ {
		key := fmt.Sprintf("test_%d", i)
		fmt.Println("key:", key)
		for j := 0; j < hourAmount; j++ {
			tempDate := startDate.Add(time.Duration(j) * time.Hour)

			random, _ := rand.Int(rand.Reader, big.NewInt(int64(packRandomArrayRange)))
			packAmount := packRandomArray[int(random.Int64())]
			for k := 0; k < packAmount; k++ {
				// 每包產生的資料亂數決定筆數
				random, _ := rand.Int(rand.Reader, big.NewInt(int64(rowRandomArrayRange)))
				rowAmount := rowRandomArray[int(random.Int64())]
				total += rowAmount

				random2, _ := rand.Int(rand.Reader, big.NewInt(int64(skuIdArrayRange)))
				skuId := skuIdArray[int(random2.Int64())]

				stmt := "INSERT INTO usages ( " + usageFields + " ) VALUES "
				for l := 1; l <= rowAmount; l++ {
					// ASR
					invStmt := fmt.Sprintf(`( '%s', '%s', '%s', null, '%s', 1 )`,
						"test_session_id",
						skuId,
						key,
						tempDate.Format("2006-01-02T15:00:00Z"),
					)

					if l < rowAmount {
						stmt = stmt + invStmt + ", "
					}

					if l == rowAmount {
						stmt = stmt + invStmt + ";"
					}
				}
				row := orm.QueryRow(stmt)
				_ = row.Scan()
			}
		}
	}

	fmt.Println("keyAmount:", keyAmount)
	fmt.Println("dayAmount:", dayAmount)
	fmt.Println("hourAmount:", hourAmount)
	fmt.Println("total:", total)
}

func seedUsageDownMillion() {

}

func getUsageFields() string {
	fields := []string{
		"session_id",
		"sku_id",
		"api_key",
		"metadata",
		"ctime",
		"count",
	}

	return strings.Join(fields, ", ")
}
