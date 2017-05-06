package main

import (
	"NSUbot/weather"
	"log"
	"net/http"
	"os"
	"time"
)

var loggerAll *log.Logger

func handler(w http.ResponseWriter, r *http.Request) {
	loggerAll.Print("URL: ", r.URL)
	w.Write([]byte("<page version=\"2.0\"><div>" + weather.CurrentWeather + "</div></page>"))
}

func main() {
	file, err := os.OpenFile(time.Now().Format("2006-01-02T15-04")+".txt", os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	loggerAll = log.New(file, "", log.LstdFlags)

	go func() {
		for {
			err := weather.SearchWeather()
			if err != nil {
				loggerAll.Print(err)
			}

			time.Sleep(79 * time.Second)
		}
	}()

	loggerAll.Print("Пошла жара")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":33642", nil)
}
