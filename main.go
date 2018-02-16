package main

import (
	"fmt"
	"net/http"
	"time"
)

type Resp struct {
	id       int
	url      string
	success  bool
	spend    float64 //time.Duration
	finished bool
	err      string
}

func main() {
	links := []Resp{
		{1, "http://google.com", false, 0, false, ""},
		{2, "http://mail.ru", false, 0, false, ""},
		{3, "http://sdgdggsdgh.org", false, 0, false, ""},
		{4, "http://amazon.com", false, 0, false, ""},
		{5, "http://ya.ru", false, 0, false, ""},
		{6, "http://stackoverflow.com", false, 0, false, ""},
		{7, "http://github.com", false, 0, false, ""},
		{8, "http://vk.com", false, 0, false, ""},
		{9, "http://vkerror.com", false, 0, false, ""},
		{10, "http://gmailgdsdg.com", false, 0, false, ""},
	}

	for i := 0; i < len(links); i++ {
		go sendRequest(&links[i])
	}

	wait(links)

	//fmt.Println(links)

	showResult(links)

	fmt.Printf("Average load time: %v", countAverage(links))

}

func sendRequest(r *Resp) {
	start := time.Now()

	response, err := http.Get(r.url)

	if err != nil {

		r.finished = true
		r.success = false
		r.err = err.Error()
		fmt.Println(err.Error())

	} else {

		if response.Status == "200" {
			fmt.Println(200)
		}

		r.finished = true
		r.success = true
		fmt.Printf("%v succeed\n", r.url)
	}

	end := time.Now()

	r.spend = end.Sub(start).Seconds()
}

func wait(l []Resp) {
	for i := 0; i < len(l); i++ {
		if !l[i].finished {
			time.Sleep(500)
			wait(l)
		}
	}
}

func showResult(l []Resp) {
	for i := 0; i < len(l); i++ {
		fmt.Printf("======================\n")
		fmt.Printf("Url: %v \n", l[i].url)
		fmt.Printf("Loaded: %v \n", l[i].success)

		if l[i].success {
			fmt.Printf("Load time: %v \n", l[i].spend)
		} else {
			fmt.Printf("Load error: %v \n", l[i].err)
		}

	}
	fmt.Printf("======================\n")
}

func countAverage(l []Resp) float64 {

	loaded := 0
	var sum float64

	for i := 0; i < len(l); i++ {
		if l[i].success {
			loaded++
			sum = sum + l[i].spend
		}
	}

	return sum / float64(loaded)

}
