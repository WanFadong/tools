package service

import (
	"fmt"
	"github.com/pkg/errors"
	"time"
)

const MinUnix int64 = 1e8  // 1976-05-04 03:33:20
const MaxUnix int64 = 1e10 // 2033-05-18 11:33:20，不包含
var CSTLoc = time.FixedZone("CST", 8*60*60)

var (
	AllowedDayLayout = []string{"2006-01-02", "20060102", "06-02-02", "060102"}
)

var (
	ErrInvalidTimestamp      = errors.New(fmt.Sprintf("invalid timestamp, not in [%v, %v)", MinUnix, MaxUnix))
	ErrUnsupportedTimeString = errors.New(fmt.Sprintf("unrecognized time string, supported list: %+v", AllowedDayLayout))
)

func ParseTimeStringWithLayout(timeStr string, layout string) (t time.Time, err error) {
	t, err = time.ParseInLocation(layout, timeStr, CSTLoc)
	return
}

func ParseTimeString(timeStr string) (t time.Time, err error) {
	for _, layout := range AllowedDayLayout {
		if t, err = ParseTimeStringWithLayout(timeStr, layout); err == nil {
			return
		}
	}

	err = ErrUnsupportedTimeString
	return
}

func ParseTimestamp(unixnano int64) (t time.Time) {
	t = time.Unix(unixnano/1e9, unixnano%1e9)
	return
}

// 支持三种格式：
// 1551024000 约定9-10位，unix
// 1551024000 0000000 +7 位，七牛格式
// 1551024000 000000000 +9 位，unixnano
func ParseTimestampAdaptive(val int64) (t time.Time, err error) {

	divs := []int64{1, 1e7, 1e9}
	var sec int64
	var nsec int64
	var ok bool
	for _, div := range divs {
		sec = val / div
		if sec >= MinUnix && sec < MaxUnix {
			nsec = val % div * (1e9 / div)
			ok = true
			break
		}
	}
	if !ok {
		err = ErrInvalidTimestamp
		return
	}

	t = time.Unix(sec, nsec)
	return
}

func LastDays(days int) (ts []time.Time) {
	if days <= 0 {
		panic("invalid days")
	}

	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, CSTLoc)
	for i := 0; i< days; i++ {
		day := today.Add(- time.Hour * 24 * time.Duration(i))
		ts = append(ts, day)
	}
	return
}
