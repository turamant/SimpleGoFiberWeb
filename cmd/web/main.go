package main

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func main()  {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		_, err := w.Write([]byte("hello ome page"))
		if err != nil {
			logrus.WithError(err).Error("write hello home oage")
		}
	})
	mux.HandleFunc("/course", func(w http.ResponseWriter, r *http.Request){
		page := r.URL.Query().Get("page")
		_, err := w.Write([]byte(page))
		if err != nil {
			logrus.WithError(err).Error("Ошибка печати")
		}
	})

    port := "4000"
	logrus.WithFields(logrus.Fields{
		"port": port,
	}).Info("Starting a web-serwer on port")

	logrus.Fatal(http.ListenAndServe(":" + port, mux))
	
	
}