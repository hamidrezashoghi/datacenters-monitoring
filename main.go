package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

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

// Keep datacenters(ISPs)
var datacenters = []string{"mci", "mci-mashhad", "mci-tabriz", "mci-esfehan", "mci-shiraz", "afranet", "irancell", "mobinnet", "shatel"}

var isp_state = make(map[string]float64)

// Radar monitor below websites
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

var websiteCounter websiteCounters

func main() {
	const waitTime = 5
	url := ""
	websiteList := []string{"aparat", "varzesh3", "digikala", "github", "google", "instagram", "wikipedia", "clubhouse"}

	// Keep datacenters state
	dc_mci_tehran_state := 0
	dc_mci_mashhad_state := 0
	dc_mci_tabriz_state := 0
	dc_mci_esfehan_state := 0
	dc_mci_shiraz_state := 0
	dc_afranet_state := 0
	dc_irancell_state := 0
	dc_mobinnet_state := 0
	dc_shatel_state := 0

	for {
		current_time := time.Now()
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
			websitesChecker(dc, websiteList)
		}

		fmt.Println("-----------------")

		networkChecker(current_time, websiteCounter.Aparat, websiteCounter.Varzesh3, websiteCounter.Digikala,
			websiteCounter.Github, websiteCounter.Google, websiteCounter.Instagram, websiteCounter.Wikipedia, websiteCounter.Clubhouse)

		for k, v := range isp_state {
			if strings.Contains(k, "mci-tehran") && v == 0.5 {
				dc_mci_tehran_state += 1
				if dc_mci_tehran_state >= 4 {
					// send_alert("Tehran MCI data center", 1, 0, 0)
					fmt.Println(current_time.String() + " Based on radar.arvan.com Tehran MCI data center has problem.")
				}
			}

			if strings.Contains(k, "mci-mashhad") && v == 0.5 {
				dc_mci_mashhad_state += 1
				if dc_mci_mashhad_state >= 4 {
					// send_alert("Mashhad MCI data center", 1, 0, 0)
					fmt.Println(current_time.String() + " Based on radar.arvan.com Mashhad MCI data center has problem.")
				}
			}

			if strings.Contains(k, "mci-tabriz") && v == 0.5 {
				dc_mci_tabriz_state += 1
				if dc_mci_tabriz_state >= 4 {
					// send_alert("Tabriz MCI data center", 1, 0, 0)
					fmt.Println(current_time.String() + " Based on radar.arvan.com Tabriz MCI data center has problem.")
				}
			}

			if strings.Contains(k, "mci-esfehan") && v == 0.5 {
				dc_mci_esfehan_state += 1
				if dc_mci_esfehan_state >= 4 {
					// send_alert("Esfehan MCI data center", 1, 0, 0)
					fmt.Println(current_time.String() + " Based on radar.arvan.com Esfehan MCI data center has problem.")
				}
			}

			if strings.Contains(k, "mci-shiraz") && v == 0.5 {
				dc_mci_shiraz_state += 1
				if dc_mci_shiraz_state >= 4 {
					// send_alert("Shiraz MCI data center", 1, 0, 0)
					fmt.Println(current_time.String() + " Based on radar.arvan.com Shiraz MCI data center has problem.")
				}
			}

			if strings.Contains(k, "afranet") && v == 0.5 {
				dc_afranet_state += 1
				if dc_afranet_state >= 4 {
					// send_alert("Afranet data center", 1, 0, 0)
					fmt.Println(current_time.String() + " Based on radar.arvan.com Afranet data center has problem.")
				}
			}

			if strings.Contains(k, "irancell") && v == 0.5 {
				dc_irancell_state += 1
				if dc_irancell_state >= 4 {
					// send_alert("Irancell data center", 1, 0, 0)
					fmt.Println(current_time.String() + " Based on radar.arvan.com Irancell data center has problem.")
				}
			}

			if strings.Contains(k, "mobinnet") && v == 0.5 {
				dc_mobinnet_state += 1
				if dc_mobinnet_state >= 4 {
					// send_alert("Mobinnet data center", 1, 0, 0)
					fmt.Println(current_time.String() + " Based on radar.arvan.com Mobinnet data center has problem.")
				}
			}

			if strings.Contains(k, "shatel") && v == 0.5 {
				dc_shatel_state += 1
				if dc_shatel_state >= 4 {
					// send_alert("Shatel data center", 1, 0, 0)
					fmt.Println(k, v)
					fmt.Println(current_time.String() + " Based on radar.arvan.com Shatel data center has problem.")
				}
			}
		}

		time.Sleep(waitTime * time.Second)

		websiteCounter.Aparat = 0
		websiteCounter.Varzesh3 = 0
		websiteCounter.Digikala = 0
		websiteCounter.Github = 0
		websiteCounter.Google = 0
		websiteCounter.Instagram = 0
		websiteCounter.Wikipedia = 0
		websiteCounter.Clubhouse = 0

		// Radar monitor Internet and Intranet states by these ISPs
		dc_mci_tehran_state = 0
		dc_mci_tabriz_state = 0
		dc_mci_shiraz_state = 0
		dc_mci_esfehan_state = 0
		dc_mci_mashhad_state = 0
		dc_afranet_state = 0
		dc_shatel_state = 0
		dc_mobinnet_state = 0
		dc_irancell_state = 0

		// Reset(clear) isp_state
		isp_state = make(map[string]float64)
	}
}

// getJson function get one before the last v (real v is -2 in the slice)
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
func websitesChecker(dc string, websiteList []string) {
	if website.Instagram[len(website.Instagram)-2] == 0.5 {
		websiteCounter.Instagram++
		isp_state[dc+"-instagram"] = 0.5
		fmt.Println(dc + "-instagram")
	}

	if website.Aparat[len(website.Aparat)-2] == 0.5 {
		websiteCounter.Aparat++
		isp_state[dc+"-aparat"] = 0.5
		fmt.Println(dc + "-aparat")
	}

	if website.Clubhouseapi[len(website.Clubhouseapi)-2] == 0.5 {
		websiteCounter.Clubhouse++
		isp_state[dc+"-clubhouse"] = 0.5
		fmt.Println(dc + "-clubhouse")
	}

	if website.Digikala[len(website.Digikala)-2] == 0.5 {
		websiteCounter.Digikala++
		isp_state[dc+"-digikala"] = 0.5
		fmt.Println(dc + "-digikala")
	}

	if website.Github[len(website.Github)-2] == 0.5 {
		websiteCounter.Github++
		isp_state[dc+"-github"] = 0.5
		fmt.Println(dc + "-github")
	}

	if website.Google[len(website.Google)-2] == 0.5 {
		websiteCounter.Google++
		isp_state[dc+"-google"] = 0.5
		fmt.Println(dc + "-google")
	}

	if website.Wikipedia[len(website.Wikipedia)-2] == 0.5 {
		websiteCounter.Wikipedia++
		isp_state[dc+"-wikipedia"] = 0.5
		fmt.Println(dc + "-wikipedia")
	}

	if website.Varzesh3[len(website.Varzesh3)-2] == 0.5 {
		websiteCounter.Varzesh3++
		isp_state[dc+"-varzesh3"] = 0.5
		fmt.Println(dc + "-varzesh3")
	}
}

func networkChecker(current_time time.Time, aparat, varzesh3, digikala, github, google, instagram, wikipedia, clubhouse int) {
	// Intranet has problem - if 3 ISP have problem in same time
	if (aparat >= 3) && (varzesh3 >= 3) && (digikala >= 3) {
		//send_alert('', 0, 1, 0)
		fmt.Println(fmt.Sprintf("%s Based on radar.arvan.com interanet has problem.", current_time.String()))
	}

	// Internet has problem - if 4 data centers have problem in samte time
	// I show local timezone because prmotheus use utc timezone
	if (instagram >= 4) && (google >= 4) && (github >= 4) {
		fmt.Println(fmt.Sprintf("%s Based on radar.arvan.com internet has problem.", current_time.String()))
		// send_alert('', 0, 0, 1)
	} else if (instagram >= 4) && (google >= 4) && (clubhouse >= 4) {
		fmt.Println(fmt.Sprintf("%s Based on radar.arvan.com internet has problem.", current_time.String()))
		// send_alert('', 0, 0, 1)
	} else if (instagram >= 4) && (github >= 4) && (clubhouse >= 4) {
		// send_alert('', 0, 0, 1)
		fmt.Println(fmt.Sprintf("%s Based on radar.arvan.com internet has problem.", current_time.String()))
	} else if (instagram >= 4) && (github >= 4) && (wikipedia >= 4) {
		// send_alert('', 0, 0, 1)
		fmt.Println(fmt.Sprintf("%s Based on radar.arvan.com internet has problem.", current_time.String()))
	} else if (instagram >= 4) && (google >= 4) && (wikipedia >= 4) {
		// send_alert('', 0, 0, 1)
		fmt.Println(fmt.Sprintf("%s Based on radar.arvan.com internet has problem.", current_time.String()))
	} else if (instagram >= 4) && (clubhouse >= 4) && (wikipedia >= 4) {
		// send_alert('', 0, 0, 1)
		fmt.Println(fmt.Sprintf("%s Based on radar.arvan.com internet has problem.", current_time.String()))
	} else if (google >= 4) && (github >= 4) && (clubhouse >= 4) {
		// send_alert('', 0, 0, 1)
		fmt.Println(fmt.Sprintf("%s Based on radar.arvan.com internet has problem.", current_time.String()))
	} else if (google >= 4) && (github >= 4) && (wikipedia >= 4) {
		// send_alert('', 0, 0, 1)
		fmt.Println(fmt.Sprintf("%s Based on radar.arvan.com internet has problem.", current_time.String()))
	}
}
