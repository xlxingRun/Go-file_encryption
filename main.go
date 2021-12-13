package main

import (
	"os"
	"strconv"
	"time"
)

func writeTime(fileName string, level int) error {
	f, err := os.Open(fileName)
	if err != nil {
		f, err = os.Create(fileName)
	}
	defer f.Close()
	now := time.Now()
	addDate := []int{0, 0, 0}
	addDate[level] = 1
	due := now.AddDate(addDate[0], addDate[1], addDate[2])
	_, err = f.WriteString(strconv.FormatInt(due.Unix(), 10))
	return err
}

func toast(level int) error {
	var fileName string
	switch level {
	case 2:
		fileName = "已激活：天卡"
	case 1:
		fileName = "已激活：月卡"
	case 0:
		fileName = "已激活：年卡"
	}
	_, err := os.Create(fileName)
	return err
}
func main() {
	fileName := "documents/secret"
	level := 0
	_ = writeTime(fileName, level)
	_ = toast(level)
}
