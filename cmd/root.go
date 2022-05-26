package cmd

import (
	"fmt"
	"job-go/cmd/seed_cmd"
	"job-go/cmd/task_cmd"
	"job-go/pkg/logr"
	"job-go/pkg/valider"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	_ "go.uber.org/automaxprocs"
)

var rootCmd = &cobra.Command{
	Use:   "/servers/task-service",
	Short: "執行各式 task ",
	Long:  `執行各式 task，輸入 /servers/task-service [cmd] [args]`,
	Run: func(cmd *cobra.Command, args []string) {
		// @TODO：做個輸出
		fmt.Println("")
	},
	TraverseChildren: true,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	// 全域參數
	rootCmd.PersistentFlags().String("env", "dev", "global env string, can be dev / development / stag / staging / prod / production")

	// 範例及測試
	rootCmd.AddCommand(task_cmd.CmdHello)
	// Seed - Invoice
	rootCmd.AddCommand(seed_cmd.CmdSeedUsage)
	seed_cmd.CmdSeedUsage.Flags().String("scale", "", "數量[million]")
	seed_cmd.CmdSeedUsage.Flags().String("seesaw", "", "要執行的操作[up, down]")
	// Task - PriceCurl
	rootCmd.AddCommand(task_cmd.CmdTaskBillingCurl)
	task_cmd.CmdTaskBillingCurl.Flags().String("dataType", "", "資料類型[service, sku, price]")
	task_cmd.CmdTaskBillingCurl.Flags().String("startYm", "", "開始月份[YYYYMM]")
	task_cmd.CmdTaskBillingCurl.Flags().String("endYm", "", "結束月份[YYYYMM]")
}

func initConfig() {
	fmt.Println("==== initialize ====")
	// init logger
	logr.InitLogger()

	// load env
	envFileName := ".env"
	err := godotenv.Load(envFileName)
	if err != nil {
		log.Panicln(err)
		os.Exit(1)
	}

	// validation init
	valider.Init()
}
