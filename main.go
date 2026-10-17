package main

import (
	"fmt"
	"net/http"
)

func main(){
	urls := []string{
		"https://example.com",
		"https://google.com",
		"https://github.com",
		"https://httpbin.org/status/404",
	}
	sucessCode :=0
	for i:=0;i<len(urls);i++{
		if checkURL(urls[i]){
			sucessCode++
		}
	}	
	fmt.Println("Sucessful Code :",sucessCode)
	fmt.Println("Finish Checking URL")
}

func checkURL(url string)bool{
	resp,err := http.Get(url)
		if err != nil {
			fmt.Println("Error :",err)
			return false
		}	

		fmt.Println("URL : ",url)
		fmt.Println("Status Code : ",resp.StatusCode)	
		resp.Body.Close()
		return resp.StatusCode == 200
}
