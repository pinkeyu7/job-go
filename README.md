# job-go

使用指令執行任務。

## Commands

### Hello world

```
go run main.go hello
```

### Seeding Data

1. Usages

   ```
   go run main.go seed::usage --seesaw [seesaw] --scale [scale]
   ```

   - seesaw: [up, down]

   - scale: [million]

2. Billing Usage Hour

   ```
   go run main.go seed::billing-usage-hour --seesaw [seesaw]
   ```

   - seesaw: [up, down]

### Billing Curl Post

1. Service

   ```
   go run main.go task::billing-curl --dataType service
   ```

2. Sku

   ```
   go run main.go task::billing-curl --dataType sku
   ```

3. Price
   ```
   go run main.go task::billing-curl --dataType price --startYm 202201 --endYm 202212
   ```

## GO

### 套件管理 Go Module

專案目前使用 Go Module 進行管理，Go 1.11 版本以上才有支援。

#### Go Module

先下指令 `go env` 確認 go module 環境變數是否為 `on`

如果不等於 `on` 的話，下指令

```
export GO111MODULE=on
```

即可打開 go module 的功能。

原則上專案編譯時會自行安裝相關套件，

但也可以先執行下列指令，安裝 module 套件。

```
go mod tidy
```
