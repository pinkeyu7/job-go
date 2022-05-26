package helper

import (
	"job-go/pkg/er"
	"time"
)

const (
	YmFormat = "200601"
)

func GetMonthLength(startYm, endYm string) (length int, err error) {
	// 驗證時間格式以及起始順序
	startTime, err := time.Parse(YmFormat, startYm)
	if err != nil {
		formatErr := er.NewAppErr(400, er.ErrorParamInvalid, "start ym format error.", err)
		return 0, formatErr
	}

	endTime, err := time.Parse(YmFormat, endYm)
	if err != nil {
		formatErr := er.NewAppErr(400, er.ErrorParamInvalid, "end ym format error.", err)
		return 0, formatErr
	}

	timeCheck := startTime.After(endTime)
	if timeCheck {
		timeCheckErr := er.NewAppErr(400, er.ErrorParamInvalid, "end ym earlier than start ym error.", nil)
		return 0, timeCheckErr
	}

	// 計算月份長度
	month := startTime.Month()
	for startTime.Before(endTime) {
		startTime = startTime.Add(time.Hour * 24)
		nextMonth := startTime.Month()
		if nextMonth != month {
			length++
		}
		month = nextMonth
	}

	// 最後一個月也要加入
	length++

	return length, nil
}

func GetQueryYm(startYm, endYm string) (queryYm []string, err error) {
	// 取得月份長度
	monthLen, err := GetMonthLength(startYm, endYm)
	if err != nil {
		return queryYm, err
	}

	// 初始化參數
	queryYm = make([]string, monthLen)

	// 將月份加入 Query
	startTime, _ := time.Parse(YmFormat, startYm)
	for i := 0; i < monthLen; i++ {
		queryYm[i] = startTime.AddDate(0, i, 0).Format(YmFormat)
	}

	return queryYm, nil
}
