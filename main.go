package main

import (
	"crypto/tls"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strings"

	"github.com/un4gi/skyscraper/models"
)

var Token = ""

func main() {
	token := flag.String("t", "", "Authorization Bearer Token")

	flag.Parse()

	if len(*token) == 0 {
		log.Println("You need to supply an Authorization: Bearer token.")
		os.Exit(0)
	} else {
		Token = *token
	}

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	req, err := http.NewRequest("GET", models.TARGETS, nil)
	if err != nil && err != io.EOF {
		log.Println(err)
	}
	req.Close = true
	SetHeaders(req)

	resp, err := http.DefaultClient.Do(req)
	if err != nil && err != io.EOF {
		log.Println(err)
	}

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println("Error grabbing response body:", err)
		}

		var registeredTargets models.TargetData
		err = json.Unmarshal(bodyBytes, &registeredTargets)
		if err != nil {
			log.Println(err)
		}
		bodyString := string(bodyBytes)

		if len(bodyString) > 3 {
			for i := 0; i < len(registeredTargets); i++ {
				slug := registeredTargets[i].Slug

				GetAcceptedLocations(slug)
				GetQueueLocations(slug)
				GetRejectedLocations(slug)
			}
		}
	}

}

func GetAcceptedLocations(slug string) {
	req, err := http.NewRequest("GET", models.SYNACKAPI+models.ANALYTICS+slug+models.STATUS_ACCEPTED, nil)
	if err != nil {
		log.Println("Error creating request:", err)
	}
	req.Close = true
	SetHeaders(req)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Error requesting accepted locations for "+slug+":", err)
	}

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println("Error reading response:", err)
		}

		var analyticData models.Analytics
		err = json.Unmarshal(bodyBytes, &analyticData)
		if err != nil {
			log.Println("Error unmarshalling JSON:", err)
		}

		for j, _ := range analyticData.Value {
			for k, _ := range analyticData.Value[j].ExploitableLocations {
				location := fmt.Sprint(analyticData.Value[j].ExploitableLocations[k].Value)
				location = strings.ReplaceAll(location, " ", "")
				test, err := url.Parse(location)
				if err != nil {
					log.Println("Error parsing URL:", location)
				}

				testslice := strings.Split(strings.TrimPrefix(test.EscapedPath(), "/"), "/")
				sortslice := sort.StringSlice(testslice)
				for l := len(sortslice) - 1; l >= 0; l-- {
					fmt.Printf("%s\n", sortslice[l])
				}
			}
		}
	}
}

func GetQueueLocations(slug string) {
	req, err := http.NewRequest("GET", models.SYNACKAPI+models.ANALYTICS+slug+models.STATUS_IN_QUEUE, nil)
	if err != nil {
		log.Println("Error creating request:", err)
	}
	req.Close = true
	SetHeaders(req)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Error requesting in-queue locations for "+slug+":", err)
	}

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println("Error reading response:", err)
		}

		var analyticData models.Analytics
		err = json.Unmarshal(bodyBytes, &analyticData)
		if err != nil {
			log.Println("Error unmarshalling JSON:", err)
		}

		for j, _ := range analyticData.Value {
			for k, _ := range analyticData.Value[j].ExploitableLocations {
				location := fmt.Sprint(analyticData.Value[j].ExploitableLocations[k].Value)
				location = strings.ReplaceAll(location, " ", "")
				test, err := url.Parse(location)
				if err != nil {
					log.Println("Error parsing URL:", location)
				}

				testslice := strings.Split(strings.TrimPrefix(test.EscapedPath(), "/"), "/")
				sortslice := sort.StringSlice(testslice)
				for l := len(sortslice) - 1; l >= 0; l-- {
					fmt.Printf("%s\n", sortslice[l])
				}
			}
		}
	}
}

func GetRejectedLocations(slug string) {
	req, err := http.NewRequest("GET", models.SYNACKAPI+models.ANALYTICS+slug+models.STATUS_REJECTED, nil)
	if err != nil {
		log.Println("Error creating request:", err)
	}
	req.Close = true
	SetHeaders(req)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Error requesting rejected locations for "+slug+":", err)
	}

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println("Error reading response:", err)
		}

		var analyticData models.Analytics
		err = json.Unmarshal(bodyBytes, &analyticData)
		if err != nil {
			log.Println("Error unmarshalling JSON:", err)
		}

		for j, _ := range analyticData.Value {
			for k, _ := range analyticData.Value[j].ExploitableLocations {
				location := fmt.Sprint(analyticData.Value[j].ExploitableLocations[k].Value)
				location = strings.ReplaceAll(location, " ", "")
				test, err := url.Parse(location)
				if err != nil {
					log.Println("Error parsing URL:", location)
				}

				testslice := strings.Split(strings.TrimPrefix(test.EscapedPath(), "/"), "/")
				sortslice := sort.StringSlice(testslice)
				for l := len(sortslice) - 1; l >= 0; l-- {
					fmt.Printf("%s\n", sortslice[l])
				}
			}
		}
	}
}

func SetHeaders(req *http.Request) {
  req.Header.Set("User-Agent", "Skyscraper (https://github.com/un4gi/skyscraper)")
	req.Header.Set("Authorization", "Bearer "+Token)
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Referer", "https://platform.synack.com/tasks/user/available")
	req.Header.Set("X-CSRF-Token", "xxxx")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Connection", "close")
}
