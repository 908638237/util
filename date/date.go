package date

import (
	"fmt"
	"github.com/golang-module/carbon/v2"
	"strconv"
	"strings"
	"time"
)

func GetYTD(typ string) (YTD int) {
	now := time.Now()
	switch typ {
	case "year":
		YTD = now.Year()
	case "month":
		YTD = int(now.Month())
	case "day":
		YTD = now.Day()
	case "yearMonth":
		YTD, _ = strconv.Atoi(now.Format("200601"))
	default:
		YTD, _ = strconv.Atoi(now.Format("20060102"))
	}
	return
}

func GetDateByDay(year, month, day int) (rs int) {
	carbon.Now().AddYears(year).AddMonths(month).AddDays(day)
	rs, _ = strconv.Atoi(time.Now().AddDate(year, month, day).Format("20060102"))
	return
}

func GetYtdOutArray(strings ...string) map[string]int {
	data := make(map[string]int)
	now := time.Now()
	cur, _ := time.Parse("20060102", now.Format("20060102"))
	if len(strings) > 0 {
		for _, v := range strings {
			switch v {
			case "year":
				data["year"] = int(cur.Year())
			case "month":
				data["month"] = int(cur.Month())
			case "day":
				data["day"] = int(cur.Day())
			case "week":
				_, week := now.ISOWeek()
				data["week"] = week
			}
		}
	}
	return data
}

func GetUnixByTime(hour, min, sec int) int64 {
	t := time.Now()
	zeroTime := time.Date(t.Year(), t.Month(), t.Day(), hour, min, sec, 0, t.Location())
	return zeroTime.Unix()
}

func GetDate(timeUnix time.Time) string {
	date := []string{
		strconv.Itoa(timeUnix.Hour()),   //时
		strconv.Itoa(timeUnix.Minute()), //分
		strconv.Itoa(timeUnix.Second()), //秒
	}
	return strings.Join(date, ":")
}

// GetDay date 2006-01-02
func GetDay(date string) (dayNum int) {
	dayNum = 1
	oD, err := time.ParseInLocation("2006-01-02", date, time.Local)
	if err != nil {
		return dayNum
	}
	//获取当前天数
	nowDate := time.Now().Format("2006-01-02")
	nD, nErr := time.ParseInLocation("2006-01-02", nowDate, time.Local)
	if nErr != nil {
		return dayNum
	}
	dayNum += int(nD.Sub(oD).Hours() / 24)
	return dayNum
}

func GetWeekDay() int {
	weekDay := time.Now().Weekday()
	return int(weekDay)
}

func GetZeroUnix() int64 {
	currentTime := time.Now()
	return time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, currentTime.Location()).Unix()
}

func GetMinutesByUnix(n1, n2 int64) float64 {
	if n1 < n2 {
		return 0
	}
	return time.Unix(n1, 0).Sub(time.Unix(n2, 0)).Minutes()
}

func CheckFestival() bool {
	nowDate := carbon.Now()
	// 获取农历的月日
	_, lunarMonth, lunarDay := nowDate.Lunar().Date()
	lunar := fmt.Sprintf("%v-%v", lunarMonth, lunarDay)
	if nowDate.DayOfWeek() >= 5 {
		return true
	}
	if nowDate.IsSameDay(nowDate.SetMonth(1).SetDay(1)) {
		// 元旦节
		return true
	}
	if nowDate.IsSameDay(nowDate.SetMonth(5).SetDay(1)) {
		// 劳动节
		return true
	}
	if nowDate.BetweenIncludedBoth(nowDate.SetMonth(10).SetDay(1), nowDate.SetMonth(10).SetDay(3)) {
		// 国庆节三天
		return true
	}
	if lunar == "5-5" || lunar == "8-15" || lunar == "1-1" || lunar == "1-2" || lunar == "1-3" {
		// 端午 中秋 春节
		return true
	}
	return false
}
