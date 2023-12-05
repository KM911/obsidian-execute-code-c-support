package util

import (
	"fmt"
	"time"
	// "github.com/gookit/color"
)

type timer struct {
	start time.Time
}

func TimerStart() *timer {
	t := timer{}
	t.start = time.Now()
	return &t
}

func (t *timer) End() {
	end := time.Now()
	// fmt.Println("花费时间为：", color.Red.Render(end.Sub(t.start)))
	fmt.Println("const time : ", end.Sub(t.start))
}

func UnixTime() int64 {
	return time.Now().Unix()
}

func Timestamp() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
