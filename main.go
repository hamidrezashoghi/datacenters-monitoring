package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/hamidrezashoghi/datacenters-monitoring/internal/config"
)

type websiteCounters struct {
	Aparat    int
	Varzesh3  int
	Digikala  int
	Github    int
	Google    int
	Instagram int
	Wikipedia int
	Clubhouse int
}

type websites struct {
	Aparat       []float64     `json:"aparat"`
	Clubhouseapi []float64     `json:"clubhouseapi"`
	Digikala     []float64     `json:"digikala"`
	Github       []float64     `json:"github"`
	Google       []float64     `json:"google"`
	Instagram    []float64     `json:"instagram"`
	Varzesh3     []float64     `json:"varzesh3"`
	Wikipedia    []interface{} `json:"wikipedia"`
}

var website websites

func main() {
	const waitTime = 5
	url := ""
	var datacenters = []string{"mci", "mci-mashhad", "mci-tabriz", "mci-esfehan", "mci-shiraz", "afranet", "irancell", "mobinnet"} // , "shatel"}

	for {
		current_time := time.Now()
		var websiteCounter websiteCounters
		dcState := config.New()
		var dc_info = make(map[string]float64)

		for _, dc := range datacenters {
			url = "https://radar.arvancloud.com/api/v1/internet-monitoring?isp=" + dc

			data, err := getJson(url)
			if err != nil {
				log.Println(err)
			} else {
				json.Unmarshal([]byte(data), &website)
				if dc == "mci" {
					dc = "mci-tehran"
				}
			}
			dc_info = websiteCounter.websitesChecker(dc, dc_info)
		}

		fmt.Println("-----------------")

		websiteCounter.networkChecker(current_time)

		for k, v := range dc_info {
			if strings.Contains(k, "mci-tehran") && v == 0.5 {
				dcState.MCITehran += 1
				if dcState.MCITehran >= 4 {
					// send_alert("Tehran MCI data center", 1, 0, 0)
					fmt.Println(current_time.String() + " Based on radar.arvan.com Tehran MCI data center has problem.")
				}
			}

			if strings.Contains(k, "mci-mashhad") && v == 0.5 {
				dcState.MCIMashhad += 1
				if dcState.MCIMashhad >= 4 {
					// send_alert("Mashhad MCI data center", 1, 0, 0)
					fmt.Println(current_time.String() + " Based on radar.arvan.com Mashhad MCI data center has problem.")
				}
			}

			if strings.Contains(k, "mci-tabriz") && v == 0.5 {
				dcState.MCITabriz += 1
				if dcState.MCITabriz >= 4 {
					// send_alert("Tabriz MCI data center", 1, 0, 0)
					fmt.Println(current_time.String() + " Based on radar.arvan.com Tabriz MCI data center has problem.")
				}
			}

			if strings.Contains(k, "mci-esfehan") && v == 0.5 {
				dcState.MCIEsfehan += 1
				if dcState.MCIEsfehan >= 4 {
					// send_alert("Esfehan MCI data center", 1, 0, 0)
					fmt.Println(current_time.String() + " Based on radar.arvan.com Esfehan MCI data center has problem.")
				}
			}

			if strings.Contains(k, "mci-shiraz") && v == 0.5 {
				dcState.MCIShiraz += 1
				if dcState.MCIShiraz >= 4 {
					// send_alert("Shiraz MCI data center", 1, 0, 0)
					fmt.Println(current_time.String() + " Based on radar.arvan.com Shiraz MCI data center has problem.")
				}
			}

			if strings.Contains(k, "afranet") && v == 0.5 {
				dcState.Afranet += 1
				if dcState.Afranet >= 4 {
					// send_alert("Afranet data center", 1, 0, 0)
					fmt.Println(current_time.String() + " Based on radar.arvan.com Afranet data center has problem.")
				}
			}

			if strings.Contains(k, "irancell") && v == 0.5 {
				dcState.Irancell += 1
				if dcState.Irancell >= 4 {
					// send_alert("Irancell data center", 1, 0, 0)
					fmt.Println(current_time.String() + " Based on radar.arvan.com Irancell data center has problem.")
				}
			}

			if strings.Contains(k, "mobinnet") && v == 0.5 {
				dcState.Mobinnet += 1
				if dcState.Mobinnet >= 4 {
					// send_alert("Mobinnet data center", 1, 0, 0)
					fmt.Println(current_time.String() + " Based on radar.arvan.com Mobinnet data center has problem.")
				}
			}

			if strings.Contains(k, "shatel") && v == 0.5 {
				dcState.Shatel += 1
				if dcState.Shatel >= 4 {
					// send_alert("Shatel data center", 1, 0, 0)
					fmt.Println(k, v)
					fmt.Println(current_time.String() + " Based on radar.arvan.com Shatel data center has problem.")
				}
			}
		}
		time.Sleep(waitTime * time.Second)
	}
}

// getJson get one before the last value (real value is -2 in the slice)
func getJson(url string) (string, error) {
	var err error
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("can't fetch URL %q: %v", url, err)
	}

	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("%v", err)
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("%v", err)
	}

	return string(b), err

}

// intranetChecker check if we have network problem inside country
func (wC *websiteCounters) websitesChecker(dc string, dc_info map[string]float64) map[string]float64 {
	if website.Instagram[len(website.Instagram)-2] == 0.5 {
		wC.Instagram++
		dc_info[dc+"-instagram"] = 0.5
		fmt.Println(dc + "-instagram")
	}

	if website.Aparat[len(website.Aparat)-2] == 0.5 {
		wC.Aparat++
		dc_info[dc+"-aparat"] = 0.5
		fmt.Println(dc + "-aparat")
	}

	if website.Clubhouseapi[len(website.Clubhouseapi)-2] == 0.5 {
		wC.Clubhouse++
		dc_info[dc+"-clubhouse"] = 0.5
		fmt.Println(dc + "-clubhouse")
	}

	if website.Digikala[len(website.Digikala)-2] == 0.5 {
		wC.Digikala++
		dc_info[dc+"-digikala"] = 0.5
		fmt.Println(dc + "-digikala")
	}

	if website.Github[len(website.Github)-2] == 0.5 {
		wC.Github++
		dc_info[dc+"-github"] = 0.5
		fmt.Println(dc + "-github")
	}

	if website.Google[len(website.Google)-2] == 0.5 {
		wC.Google++
		dc_info[dc+"-google"] = 0.5
		fmt.Println(dc + "-google")
	}

	if website.Wikipedia[len(website.Wikipedia)-2] == 0.5 {
		wC.Wikipedia++
		dc_info[dc+"-wikipedia"] = 0.5
		fmt.Println(dc + "-wikipedia")
	}

	if website.Varzesh3[len(website.Varzesh3)-2] == 0.5 {
		wC.Varzesh3++
		dc_info[dc+"-varzesh3"] = 0.5
		fmt.Println(dc + "-varzesh3")
	}

	return dc_info
}

func (wC websiteCounters) networkChecker(current_time time.Time) {
	// Intranet has problem - if 3 ISP have problem
	if (wC.Aparat >= 3) && (wC.Varzesh3 >= 3) && (wC.Digikala >= 3) {
		//send_alert('', 0, 1, 0)
		fmt.Println(fmt.Sprintf("%s Based on radar.arvan.com interanet has problem.", current_time.String()))
	}

	// Internet has problem - if 4 data centers have problem
	// I show local timezone because prmotheus use utc timezone
	if (wC.Instagram >= 4) && (wC.Google >= 4) && (wC.Github >= 4) {
		fmt.Println(fmt.Sprintf("%s Based on radar.arvan.com internet has problem.", current_time.String()))
		// send_alert('', 0, 0, 1)
	} else if (wC.Instagram >= 4) && (wC.Google >= 4) && (wC.Clubhouse >= 4) {
		fmt.Println(fmt.Sprintf("%s Based on radar.arvan.com internet has problem.", current_time.String()))
		// send_alert('', 0, 0, 1)
	} else if (wC.Instagram >= 4) && (wC.Github >= 4) && (wC.Clubhouse >= 4) {
		// send_alert('', 0, 0, 1)
		fmt.Println(fmt.Sprintf("%s Based on radar.arvan.com internet has problem.", current_time.String()))
	} else if (wC.Instagram >= 4) && (wC.Github >= 4) && (wC.Wikipedia >= 4) {
		// send_alert('', 0, 0, 1)
		fmt.Println(fmt.Sprintf("%s Based on radar.arvan.com internet has problem.", current_time.String()))
	} else if (wC.Instagram >= 4) && (wC.Google >= 4) && (wC.Wikipedia >= 4) {
		// send_alert('', 0, 0, 1)
		fmt.Println(fmt.Sprintf("%s Based on radar.arvan.com internet has problem.", current_time.String()))
	} else if (wC.Instagram >= 4) && (wC.Clubhouse >= 4) && (wC.Wikipedia >= 4) {
		// send_alert('', 0, 0, 1)
		fmt.Println(fmt.Sprintf("%s Based on radar.arvan.com internet has problem.", current_time.String()))
	} else if (wC.Google >= 4) && (wC.Github >= 4) && (wC.Clubhouse >= 4) {
		// send_alert('', 0, 0, 1)
		fmt.Println(fmt.Sprintf("%s Based on radar.arvan.com internet has problem.", current_time.String()))
	} else if (wC.Google >= 4) && (wC.Github >= 4) && (wC.Wikipedia >= 4) {
		// send_alert('', 0, 0, 1)
		fmt.Println(fmt.Sprintf("%s Based on radar.arvan.com internet has problem.", current_time.String()))
	}
}
