package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"gopl.io/ch4/github"
)

var m = flag.Int("m", 0, "positive: query issues created during recently m months;\nnegative: query issues created before -m months")

func main() {
	// -m 1 一月以内：created:2019-01-20...2019-02-20
	// -m 12 一年以内：created:2018-02-20...2019-02-20
	// -m -12 超过一年：created:<2018-02-20
	flag.Parse()
	q := os.Args[1:]

	for i, v := range q {
		if v == "-m" {
			copy(q[i:], q[i+2:])
			q = q[0 : len(q)-2]
			break
		}
	}
	if *m > 0 {
		end := time.Now()
		start := end.AddDate(0, -*m, 0)
		log.Println(start)
		q = append(q, "created:"+start.Format("2006-01-02")+".."+end.Format("2006-01-02"))
	} else if *m < 0 {
		end := time.Now().AddDate(0, *m, 0)
		q = append(q, "created:<"+end.Format("2006-01-02"))
	}
	log.Println(q)

	result, err := github.SearchIssues(q)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}