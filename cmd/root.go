package cmd

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	_ "go.uber.org/automaxprocs"
	"job-go/cmd/task_cmd"
	"job-go/pkg/logr"
	"job-go/pkg/valider"
	"os"
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
