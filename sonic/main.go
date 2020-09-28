package main

import (
	"fmt"

	"github.com/expectedsh/go-sonic/sonic"
)

func main() {

	ingester, err := sonic.NewIngester("localhost", 1491, "SecretPassword")
	if err != nil {
		panic(err)
	}

	// I will ignore all errors for demonstration purposes

	_ = ingester.BulkPush("movies", "general", 3, []sonic.IngestBulkRecord{
		{"id:6ab56b4kk3", "Lan đi du học ở Mỹ Tho"},
		{"id:5hg67f8dg5", "Minh học mỹ thuật"},
		{"id:1m2n3b4vf6", "Tương lai AI có thể làm ra các bản hit giỏi hơn nghệ sĩ"},
		{"id:68d96h5h9d0", "khôi Thanh không thích đẩy Khánh đi"},
	})

	search, err := sonic.NewSearch("localhost", 1491, "SecretPassword")
	if err != nil {
		panic(err)
	}

	results, _ := search.Query("movies", "general", "có làm ra", 10, 0)

	fmt.Println(results)
}
