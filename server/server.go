package main

import (
	"NSUbot/weather"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

var loggerAll *log.Logger

func handler(w http.ResponseWriter, r *http.Request) {
	loggerAll.Print("URL: ", r.URL)
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		loggerAll.Print(err)
	} else {
		loggerAll.Print(string(b))
	}
	loggerAll.Print(r.Header)
	loggerAll.Print(r.Form)
	loggerAll.Print(r.RequestURI)
	loggerAll.Print(r.Trailer)
	loggerAll.Print("Конец\n\n")
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
