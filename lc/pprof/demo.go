package main

import (
	"lc/pprof/data"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main(){
	go func(){
		for{
			log.Println(data.Add("https://github.com/EDDYCJY"))
		}
	}()

	http.ListenAndServe(":6060",nil)
}