
package main

import "fmt"
import "net/url"

func getUrlInfos(_url string) (string, url.Values) {
	urlInfo, _ := url.Parse(_url)	
	baseUrl := fmt.Sprintf("%s://%s", urlInfo.Scheme, urlInfo.Host) 
	queryVars := urlInfo.Query()
	return baseUrl, queryVars
}

func main() {

	baseUrl, queryVars := getUrlInfos("https://www.tanguay.info?category=howtos&search=react&showAll=off")
	fmt.Printf("URL is %s\n", baseUrl)
	for key,value := range queryVars {
		fmt.Printf("%s = %v\n", key, value)
	}

}
