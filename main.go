package main

import (
	"fmt"
	"strings"
	"github.com/robfig/cron/v3"
)

type First struct {

}

func foo(i interface{}){
}
func boo(i *interface{}){
}

func main() {

	str:="Hello & World!"
	lowercase:= strings.ToLower(str)

	fmt.Println(lowercase)

	uppercase:= strings.ToUpper(str)
	fmt.Println(uppercase)

	c := cron.New()
	//							min  hour day month day-week
	//                                   of-m
	// every day at 4 am 		"0   4    *     *     *"
	// 00:05 in August   		"5   0    *     8     *"
	// 14:15 on day-of-month 1 	"15  14   1     *     *"
	// 22:00 every day from		"0   22   *     *	 1-5"
	// 		monday to friday
	// 04:05 on sunday			"5   4    *     *    sun"

	// every hour
	_, err := c.AddFunc("0 * * * *", func(){
		// do something
	})

	if err != nil{
		// return err
	}

}
