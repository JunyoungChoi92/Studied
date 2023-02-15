package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main(){
	req, err := http.NewRequest("GET", "https://www.naver.com", nil)
	if(err != nil){log.Fatalln(err)}

	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if(err != nil){log.Fatalln(err)}

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if(err != nil){log.Fatalln(err)}

	doch, err := doc.Html()
	if(err != nil){log.Fatalln(err)}
	
	fmt.Println("Query resp: ", doch)

}