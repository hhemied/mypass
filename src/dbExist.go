package main

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/fatih/color"
)

var DBDir = filepath.Join(os.Getenv("HOME"), ".mypass")
var DBFile = filepath.Join(DBDir, ".db")

func PassGenerator() []string {
	// -> FUNC to generate a strong Password and return it as string.
	//    Use ASCI
	var ASCI = "QWERTYUIOPLKJHGFDSAZXCVBNMmnbvcxzasdfghjklopiuytrewq7869543210"
	var Sign = "#!@&()_-][><~;{}"
	var Pass = []string{}
	for i := 0; i < 12; i++ {
		Pass = append(Pass, string(ASCI[rand.Intn(len(ASCI))]))
	}
	for i := 0; i < rand.Intn(3); i++ {
		Pass[rand.Intn(10)] = string(Sign[rand.Intn(len(Sign))])
	}
	return Pass
}

func DBExist() {
	// -> FUNC to check if the DB file "~/.mypass/.db" is exist or not
	//    is not exist, go to another FUNC to create the DB structure.
	if _, err := os.Stat(DBDir); os.IsNotExist(err) {
		color.Red("This program made for you ❤️ ...")
		color.Green("Starting for the first time")
		color.Green("Creating [ %v ] ...", DBDir)
		os.Mkdir(DBDir, 0700)

	}
	if _, err := os.Stat(DBFile); os.IsNotExist(err) {
		color.Green("Creating [ %v ] ...", DBFile)
		os.Create(DBFile)
		os.Chmod(DBFile, 0600)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	DBExist()
	fmt.Println(strings.Join(PassGenerator(), ""))
}
