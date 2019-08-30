package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main(){
	resp,err:=http.Get("http://localhost:8080/test")
	if err!=nil{
		log.Println("Error in http")
		return
	}
	jsons,err:=ioutil.ReadAll(resp.Body)
	if err!=nil{
		log.Println("Error in read response")
		return
	}
	log.Println(string(jsons))
}
