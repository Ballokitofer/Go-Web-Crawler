package main

import (
	"fmt"
	"net/http"
)

func main(){
	resp, err := http.Get("https://example.com")
	if err != nil {
		fmt.Println("Error :",err)
		return
	}
	
	defer resp.Body.Close()
	successCode := 0

	if resp.StatusCode == 200 {
			successCode++
		}
		
	fmt.Println("Status Code : ",resp.StatusCode)	
	fmt.Println("Sucessful requests : ",successCode)
	fmt.Println("Finish Checking URL")
}