package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	home   = os.Getenv("HOME")
	user   = os.Getenv("USER")
	gopath = os.Getenv("GOPATH")
)

func init() {
	if user == "" {
		log.Fatal("$USER not set")
	}
	if home == "" {
		home = "/home/" + user
	}
	if gopath == "" {
		gopath = home + "/go"
	}

	fmt.Println(user, home, gopath)

}

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(c, link)
	}

	/*for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}*/

	for l := range c {
		go func() {
			time.Sleep(5 * time.Second)
			checkLink(c, l)
		}()
	}

}

func checkLink(c chan string, link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}
	fmt.Println(link, "is up!")
	c <- link
}
