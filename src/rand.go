package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/fatih/color"
)

var ASCI = "QWERTYUIOPLKJHGFDSAZXCVBNMmnbvcxzasdfghjklopiuytrewq7869543210"
var Pass = []string{}

var Sign = "#!@&()_-][><"

func strung() []string {
	for i := 0; i < 10; i++ {
		Pass = append(Pass, string(ASCI[rand.Intn(len(ASCI))]))
	}
	for i := 0; i < rand.Intn(3); i++ {
		Pass[rand.Intn(10)] = string(Sign[rand.Intn(len(Sign))])
	}
	return Pass
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(strings.Join(strung(), ""))
	fmt.Println(len(strung()))
	fmt.Println(strung())
	color.Red("")
}
