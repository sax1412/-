package util

import (
	"github.com/tealeg/xlsx"
	"time"
	"fmt"
)


func init() {
	file := xlsx.NewFile()
	sheet, _ := file.AddSheet("Sheet1")
	row := sheet.AddRow()
	cell := row.AddCell()
	cell.Value = "姓名"
	cell = row.AddCell()
	cell.Value = "标题"
	cell = row.AddCell()
	cell.Value = "链接"
	cell = row.AddCell()
	cell.Value = "时间"
	err := file.Save("star.xlsx")
	if err != nil {
		panic(err)
	}
}

func Substr(str string, start int, length int) string {
	rl := len(str)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}

	return string(str[start:end])
}

func Str_delete(s string) string {
	status := 0
	var str []rune
	for _, r := range s {
		if r == rune('<') {
			status = 1
			continue
		}
		if r == rune('>') {
			status = 0
			continue
		}
		if status == 0 {
			str = append(str, r)
		}
	}
	return string(str)
}

func Excel(keys, title, link string) {
	file, _ := xlsx.OpenFile("star.xlsx")
	sheet := file.Sheet["Sheet1"]
	row := sheet.AddRow()
	cell := row.AddCell()
	cell.Value = keys
	cell = row.AddCell()
	cell.Value = title
	cell = row.AddCell()
	cell.Value = link
	cell = row.AddCell()
	cell.Value = time.Now().Format("2006-01-02 15:04:05")
	err := file.Save("star.xlsx")
	if err != nil {
		fmt.Printf(err.Error())
	}
}
