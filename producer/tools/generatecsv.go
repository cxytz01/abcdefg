package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Create("recipients.csv")
	if err != nil {
		fmt.Println("Error creating CSV file:", err)
		return
	}
	defer file.Close()

	// 创建CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush() // 确保所有缓冲的数据被写入

	// 写入CSV头部
	header := []string{"phone", "name"}
	if err := writer.Write(header); err != nil {
		fmt.Println("Error writing header to CSV:", err)
		return
	}

	// 写入CSV内容
	records := make([][]string, 100000)

	for i := 0; i < 100000; i++ {
		records = append(records, []string{strconv.Itoa(i), "aaa" + strconv.Itoa(i)})
	}

	for _, record := range records {
		if err := writer.Write(record); err != nil {
			fmt.Println("Error writing record to CSV:", err)
			return
		}
	}

	fmt.Println("CSV file generated successfully.")
}
