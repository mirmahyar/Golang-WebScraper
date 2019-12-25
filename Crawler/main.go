package main

import (
	_"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	_"os"
	_ "strings"
	"regexp"
)

//This function makes life easier for me ! Error catcher !
func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	

	fmt.Println("Please enter the URL:  ")
	var address string
	fmt.Scanln(&address)
	fmt.Println(address)

	resp, err := http.Get(address)
	check(err)

	fmt.Printf("%T", resp)
	text, _ := ioutil.ReadAll(resp.Body)
	str := fmt.Sprintf("%s", text)

	d1 := []byte(str)
	err = ioutil.WriteFile(`C:\Users\Mahyar\Desktop\Go\src\github.com\mirmahyar\Crawler\string.txt`, d1, 0644)
	check(err)

	var word string
	fmt.Println("\nPlease enter the word you are looking in this website: ")
	fmt.Scanln(&word)
	//fmt.Println(word)

	file, err := ioutil.ReadFile(`C:\Users\Mahyar\Desktop\Go\src\github.com\mirmahyar\Crawler\string.txt`)
	check(err)
	

	isExist, err := regexp.Match(word, file)
	check(err)

	if isExist {
		fmt.Println("Found the word! :)")
	}else{
		fmt.Println("Not found the word :(")
	}
	

}
