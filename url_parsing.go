package main

import (
	"fmt"
	"net/url"
)

func main() {
	//[scheme://[userinfo]@host[:port][?query][#fragment]]

	rawUrl := "https://example.com:8080/path?query=param#fragment"

	parseURL, err := url.Parse(rawUrl)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return

	}
	fmt.Println("Scheme:", parseURL.Scheme)
	fmt.Println("Host:", parseURL.Host)
	fmt.Println("Port:", parseURL.Port())
	fmt.Println("Path:", parseURL.Path)
	fmt.Println("RawQuery:", parseURL.RawQuery)
	fmt.Println("Fragment:", parseURL.Fragment)
	// fmt.Println("Scheme:", parseURL.Scheme)

	rawUrl1 := "http://example.com/path?name=John&age=30"
	parseURL1, err := url.Parse(rawUrl1)
	if err != nil {
		fmt.Println("Error parsing the URL:", err)
		return
	}

	queryParams := parseURL1.Query()
	fmt.Println(queryParams)
	fmt.Println("Name:", queryParams.Get("name"))
	fmt.Println("Age:", queryParams.Get("age"))

	//Building an URL
	basedUrl := &url.URL{
		Scheme: "https",
		Host:   "example.com",
		Path:   "/path",
	}

	query := basedUrl.Query()
	query.Set("name", "John")
	query.Set("age", "30")
	basedUrl.RawQuery = query.Encode()

	fmt.Println("Build URL:", basedUrl.String())

	//the ultimate way to build an URL
	values := url.Values{}
	// Add key-values pairs to the values object
	values.Add("name", "Eddy")
	values.Add("age", "40")
	values.Add("city", "LBV")
	values.Add("country", "Gabon")

	//Encoding thee key value
	encodedQuery := values.Encode()
	fmt.Println(encodedQuery)

	//Adding encoded query to the url
	basedUrl1 := "https://example.com/search"
	fullURL := basedUrl1 + "?" + encodedQuery

	fmt.Println(fullURL)

}
