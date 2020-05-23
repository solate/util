package cron

import (
	"time"

	"github.com/gorhill/cronexpr"
)

// 验证cron表达式字符串是否正确，并判断下一个节点的时间是否为零，如果是零返回false
//   Seconds Minutes Hours DayofMonth Month DayofWeek Year
//   Seconds Minutes Hours DayofMonth Month DayofWeek
//
// Note: 这个包要改一下源码，添加DayofMonth和DayofWeek验证，必须有一个是 "?",
// 删除 secondList add []int{0}，这个会在前面加一位，正常应该不加这位。
func Valid(cron string) (b bool) {
	expr, err := cronexpr.Parse(cron)
	if err != nil {
		return
	}

	now := time.Now()
	nextTime := expr.Next(now)
	if nextTime.IsZero() {
		return
	}
	return true
}
