package task_cmd

import (
	"github.com/spf13/cobra"
	"job-go/pkg/logr"
	"time"
)

var CmdHello = &cobra.Command{
	Use:   "hello",
	Short: "簡單的測試 task",
	Long:  `測試 task 會執行 10 秒確認功能`,
	Run:   hello,
}

// 只是個測試用的 task，會做個簡單輸出測試功能
func hello(cmd *cobra.Command, args []string) {
	for i := 0; i < 10; i++ {
		logr.L.Info("Hello Job output: " + string(i))
		time.Sleep(1 * time.Second)
	}
}
