package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"os"
) 
// Get a website specified by CLI
// Save to the file as index.html 

func main () {
	if len(os.Args) < 2 {
	fmt.Println("Please specify a url")
	os.Exit(1)
	}	
	must(getAndSaveURL(os.Args[1], "index.html"))
}

func must(err error) {
	if err != nil{
	panic(err)
	}
}

func getAndSaveURL(url, filename string) error {

	resp, err := http.Get(url)
	if err != nil {
	return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err 
	}

	return ioutil.WriteFile(filename, body, 0777)
}
