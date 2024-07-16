
package main

import "fmt"
import "net/url"

func getUrlInfos(_url string) (string, map[string]string) {
	urlInfo, _ := url.Parse(_url)	
	baseUrl := fmt.Sprintf("%s://%s", urlInfo.Scheme, urlInfo.Host) 
	rawQueryVars := urlInfo.Query()
	queryVars := make(map[string]string) 
	for k,v := range rawQueryVars{
		queryVars[k] = v[0]
	}
	return baseUrl, queryVars
}

func main() {

	baseUrl, queryVars := getUrlInfos("https://www.tanguay.info?category=howtos&search=react&showAll=off")
	fmt.Printf("URL is %s\n", baseUrl)
	for key,value := range queryVars {
		fmt.Printf("%s = %v\n", key, value)
	}

}
