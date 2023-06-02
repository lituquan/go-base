package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
)

type User struct {
	UserId string
	Link   string
}

func (u *User) format() []string {
	return []string{u.UserId, u.Link}
}

// 链接头
const HELP = ``
const MAIN_URL = `http://www.baidu.com?user=%s`
const OUTPUT_CSV_NAME = `output.csv`

// init log
var (
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
)

func init() {
	InfoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// 输入--生成分享链接
func main() {
	log.Println(MAIN_URL, "xxx")
	log.Println(os.Args[1])
	input := os.Args[1]
	userIds, err := readFile(input)
	if err != nil {
		log.Println("Error:", err)
		return
	}
	results := make([]User, 0)
	// add csv header
	results = append(results, User{"UserId", "Link"})
	for _, v := range userIds {
		results = append(results, User{UserId: v, Link: trans2Short(fmt.Sprintf(MAIN_URL, v))})
	}
	// write Users
	datas, _ := format(results)
	writeFile(OUTPUT_CSV_NAME, datas)
}

// 读取文件--批量生成
func readFile(input string) (result []string, err error) {
	file, err := os.Open(input)
	if err != nil {
		log.Println("Error:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		log.Println("Error:", err)
		return
	}
	result = make([]string, 0)
	for _, record := range records {
		log.Println(record)
		result = append(result, record[0])
	}
	return
}

// 转成短链接
func trans2Short(link string) string {
	return link
}

// 序列化
func format(users []User) (result [][]string, err error) {
	if users == nil || len(users) == 0 {
		return [][]string{}, errors.New("users 为空")
	}
	result = make([][]string, 0)
	for _, v := range users {
		result = append(result, v.format())
	}
	return
}

// 输出
func writeFile(output string, records [][]string) {
	// Open the output file
	file, err := os.Create(output)
	if err != nil {
		log.Println("Error:", err)
		return
	}
	defer file.Close()

	// Create a new CSV writer and write some data
	writer := csv.NewWriter(file)

	for _, record := range records {
		err := writer.Write(record)
		if err != nil {
			log.Println("Error:", err)
			return
		}
	}

	// Flush the CSV writer and check for errors
	writer.Flush()
	if err := writer.Error(); err != nil {
		log.Println("Error:", err)
		return
	}

}
