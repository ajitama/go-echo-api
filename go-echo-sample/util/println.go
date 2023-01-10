package util

import (
	"fmt"
	"path/filepath"
	"runtime"
	"time"
)

func Println(a ...interface{}) (n int, err error) {
	// 出力フォーマット
	const MilliFormat = "2006/01/02 15:04:05.000"
	// Timezoneの設定
	time.Local = time.FixedZone("JST", 9*60*60)
	time.LoadLocation("JST")
	now := time.Now()
	_, file, line, _ := runtime.Caller(1)
	fname := filepath.Base(file)
	str := fmt.Sprintf("file:%s:%d", fname, line)
	return fmt.Println(append([]interface{}{now.Format(MilliFormat), str}, a...)...)
}
