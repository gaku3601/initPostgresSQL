package main

import (
	"flag"
	"fmt"
)

func main() {
	// 左から順にオプション名、デフォルトの値、helpテキストが引数に入る
	sqlFolder := flag.String("s", "", "SqlFolder Path")
	host := flag.String("h", "", "host名")
	port := flag.Int("p", 0, "port番号")
	dbName := flag.String("d", "", "database名")
	userName := flag.String("u", "", "user name")
	password := flag.String("pw", "", "password")
	flag.Parse()
	fmt.Println(*host)
	fmt.Println(*port)
	fmt.Println(*dbName)
	fmt.Println(*sqlFolder)
	fmt.Println(*userName)
	fmt.Println(*password)

	o := NewOperateFile(*sqlFolder)
	fmt.Print(o)
}
