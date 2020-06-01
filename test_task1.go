package main

import (
    "bufio"
    "encoding/csv"
    "fmt"
    "io"
	"os"
	"strconv"
	"regexp"
	"time"
)

func main() {
    csvFile, _ := os.Open("publicholiday.UA.2020.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	currentTime := time.Now()
	layout := "2006-01-02"
    Timenow := currentTime.Format(layout)
    for {
        line, err := reader.Read()
        if err == io.EOF {
            break
        } else if err != nil {
            fmt.Println(err)
		}
		if Timenow <= line[0] && line[0] != "Date"{
			if Timenow == line[0] {
				fmt.Print("Today is ", line[2])
			} else {
				fmt.Print("The next holiday is ", line[2])
			}
			t, err := time.Parse(layout, line[0])
			if err != nil  {
				fmt.Print(err)
			}
			fmt.Print(", ", t.Format("Jan 2")) 
			if t.Format("Mon") == "Fri" {
				re := regexp.MustCompile(`-`)
				dashes := re.FindAllStringSubmatchIndex(line[0], -1)
				day,err := strconv.Atoi(line[0][dashes[1][1]:])
				if err != nil  {
					fmt.Print(err)
				}
				fmt.Print(", and the weekend will last 3 days: ")
				fmt.Print(t.Format("Jan 2"), " - ", t.Format("Jan "), day + 2)
			}
			break
		}
	}
}
