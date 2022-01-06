package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type websites struct {
	Aparat       []float32     `json:"aparat"`
	Clubhouseapi []float32     `json:"clubhouseapi"`
	Digikala     []float32     `json:"digikala"`
	Github       []float32     `json:"github"`
	Google       []float32     `json:"google"`
	Instagram    []float32     `json:"instagram"`
	Varzesh3     []float32     `json:"varzesh3"`
	Wikipedia    []interface{} `json:"wikipedia"`
}

var website websites

// Keep datacenters(ISPs)
var datacenters = []string{"mci", "mci-mashhad", "mci-tabriz", "mci-esfehan", "mci-shiraz", "afranet", "irancell", "mobinnet", "shatel"}
var wbs = []string{"aparat", "varzesh3", "digikala", "github", "google", "instagram", "wikipedia", "clubhouse"}
var isp_state = make(map[string]float32)

// Radar monitor below websites
var (
	aparat    int
	varzesh3  int
	digikala  int
	github    int
	google    int
	instagram int
	wikipedia int
	clubhouse int
)

func main() {
	url := ""

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

	// Radar monitor Internet and Intranet states by these ISPs
	mci_tehran := 0
	mci_mashhad := 0
	mci_tabriz := 0
	mci_esfehan := 0
	mci_shiraz := 0
	afranet := 0
	irancell := 0
	mobinnet := 0
	shatel := 0

	for {
		current_time := time.Now()
		for _, dc := range datacenters {
			url = "https://radar.arvancloud.com/api/v1/internet-monitoring?isp=" + dc
			// fmt.Println("=======>", dc, "<=======")

			data, err := GetJSON(url)

			if err != nil {
				// log.Fatal(err)
				fmt.Println(err)
			} else {
				// fmt.Println(data)
				json.Unmarshal([]byte(data), &website)
				if dc == "mci" {
					dc = "mci-tehran"
				}
			}
		}

		for k, v := range isp_state {
			if strings.Contains(k, "aparat") && v == 0.5 {
				aparat += 1
				// fmt.Println("aparat:", aparat, "=>", k)
			} else if strings.Contains(k, "varzesh3") && v == 0.5 {
				varzesh3 += 1
				// fmt.Println("varzesh3:", varzesh3, "->", k)
			} else if strings.Contains(k, "digikala") && v == 0.5 {
				digikala += 1
				// fmt.Println("digikala:", digikala, "->", k)
			} else if strings.Contains(k, "github") && v == 0.5 {
				github += 1
				// fmt.Println("github:", github, "->", k)
			} else if strings.Contains(k, "google") && v == 0.5 {
				google += 1
				// fmt.Println("google:", google, "->", k)
			} else if strings.Contains(k, "instagram") && v == 0.5 {
				instagram += 1
				fmt.Println("instagram:", instagram, "->", k)
			} else if strings.Contains(k, "wikipedia") && v == 0.5 {
				wikipedia += 1
				// fmt.Println("wikipedia:", wikipedia, "->", k)
			} else if strings.Contains(k, "clubhouse") && v == 0.5 {
				clubhouse += 1
				fmt.Println("clubhouse:", clubhouse, "->", k)
			}
		}
		fmt.Println("==========================")

		// fmt.Println(isp_state)
		NetworkChecker(current_time, aparat, varzesh3, digikala, github, google, instagram, wikipedia, clubhouse)

		for k, v := range isp_state {
			if strings.Contains(k, "mci-tehran") && v == 0.5 {
				mci_tehran += 1
				if mci_tehran >= 2 {
					dc_mci_tehran_state += 1
				} else {
					dc_mci_tehran_state = 0
				}
			}

			if strings.Contains(k, "mci-mashhad") && v == 0.5 {
				mci_mashhad += 1
				if mci_mashhad >= 2 {
					dc_mci_mashhad_state += 1
				} else {
					dc_mci_mashhad_state = 0
				}
			}

			if strings.Contains(k, "mci-tabriz") && v == 0.5 {
				mci_tabriz += 1
				if mci_tabriz >= 2 {
					dc_mci_tabriz_state += 1
				} else {
					dc_mci_tabriz_state = 0
				}
			}

			if strings.Contains(k, "mci-esfehan") && v == 0.5 {
				mci_esfehan += 1
				if mci_esfehan >= 2 {
					dc_mci_esfehan_state += 1
				} else {
					dc_mci_esfehan_state = 0
				}
			}

			if strings.Contains(k, "mci-shiraz") && v == 0.5 {
				mci_shiraz += 1
				if mci_shiraz >= 2 {
					dc_mci_shiraz_state += 1
				} else {
					dc_mci_shiraz_state = 0
				}
			}

			if strings.Contains(k, "afranet") && v == 0.5 {
				afranet += 1
				if afranet >= 2 {
					dc_afranet_state += 1
				} else {
					dc_afranet_state = 0
				}
			}

			if strings.Contains(k, "irancell") && v == 0.5 {
				irancell += 1
				if irancell >= 2 {
					dc_irancell_state += 1
				} else {
					dc_irancell_state = 0
				}
			}

			if strings.Contains(k, "mobinnet") && v == 0.5 {
				mobinnet += 1
				if mobinnet >= 2 {
					dc_mobinnet_state += 1
				} else {
					dc_mobinnet_state = 0
				}
			}

			if strings.Contains(k, "shatel") && v == 0.5 {
				shatel += 1
				if shatel >= 2 {
					dc_shatel_state += 1
				} else {
					dc_shatel_state = 0
				}
			}
		}

		if dc_mci_tehran_state >= 4 {
			// send_alert("Tehran MCI data center", 1, 0, 0)
			fmt.Println(current_time.String() + " Based on radar.arvan.com Tehran MCI data center has problem.")
			dc_mci_tehran_state = 0
		}
		if dc_mci_mashhad_state >= 4 {
			// send_alert("Mashhad MCI data center", 1, 0, 0)
			fmt.Println(current_time.String() + " Based on radar.arvan.com Mashhad MCI data center has problem.")
			dc_mci_mashhad_state = 0
		}
		if dc_mci_tabriz_state >= 4 {
			// send_alert("Tabriz MCI data center", 1, 0, 0)
			fmt.Println(current_time.String() + " Based on radar.arvan.com Tabriz MCI data center has problem.")
			dc_mci_tabriz_state = 0
		}
		if dc_mci_esfehan_state >= 4 {
			// send_alert("Esfehan MCI data center", 1, 0, 0)
			fmt.Println(current_time.String() + " Based on radar.arvan.com Esfehan MCI data center has problem.")
			dc_mci_esfehan_state = 0
		}
		if dc_mci_shiraz_state >= 4 {
			// send_alert("Shiraz MCI data center", 1, 0, 0)
			fmt.Println(current_time.String() + " Based on radar.arvan.com Shiraz MCI data center has problem.")
			dc_mci_shiraz_state = 0
		}
		if dc_afranet_state >= 4 {
			// send_alert("Afranet data center", 1, 0, 0)
			fmt.Println(current_time.String() + " Based on radar.arvan.com Afranet data center in Tehran has problem.")
			dc_afranet_state = 0
		}
		if dc_irancell_state >= 4 {
			// send_alert("Irancell data center", 1, 0, 0)
			fmt.Println(current_time.String() + " Based on radar.arvan.com Irancell data center in Tehran has problem.")
			dc_irancell_state = 0
		}
		if dc_mobinnet_state >= 4 {
			// send_alert("Mobinnet data center", 1, 0, 0)
			fmt.Println(current_time.String() + " Based on radar.arvan.com Mobinnet data center in Tehran has problem.")
			dc_mobinnet_state = 0
		}
		if dc_shatel_state >= 4 {
			// send_alert("Shatel data center", 1, 0, 0)
			fmt.Println(current_time.String() + " Based on radar.arvan.com Shatel data center in Tehran has problem.")
			dc_shatel_state = 0
		}

		fmt.Println("Waiting 5 seconds ...... ")
		time.Sleep(5 * time.Second)

		aparat = 0
		varzesh3 = 0
		digikala = 0
		github = 0
		google = 0
		instagram = 0
		wikipedia = 0
		clubhouse = 0

		// Radar monitor Internet and Intranet states by these ISPs
		mci_tehran = 0
		mci_mashhad = 0
		mci_tabriz = 0
		mci_esfehan = 0
		mci_shiraz = 0
		afranet = 0
		irancell = 0
		mobinnet = 0
		shatel = 0

		isp_state = make(map[string]float32)
	}
}

// GetJSON function get one before the last v (real v is -2 in the slice)
func GetJSON(url string) (string, error) {
	var err error
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("can't fetch URL %q: %v", url, err)
	}

	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("++++ %v", err)
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("---- %v", err)
	}

	return string(b), err

}

// intranetChecker check if we have network problem inside country
func WebsitesCheckers(dc string, wb string) {
	if wb == "instagram" {
		if website.Instagram[len(website.Instagram)-2] == 0.5 {
			instagram++
		}
		isp_state[dc+"-"+wb] = 0.5
	}

	if wb == "aparat" {
		if website.Aparat[len(website.Aparat)-2] == 0.5 {
			aparat++
		}
		isp_state[dc+"-"+wb] = 0.5
	}

	if wb == "clubhouse" {
		if website.Clubhouseapi[len(website.Clubhouseapi)-2] == 0.5 {
			clubhouse++
		}
		isp_state[dc+"-"+wb] = 0.5
	}

	if wb == "digikala" {
		if website.Digikala[len(website.Digikala)-2] == 0.5 {
			digikala++
		}
		isp_state[dc+"-"+wb] = 0.5
	}

	if wb == "github" {
		if website.Github[len(website.Github)-2] == 0.5 {
			github++
		}
		isp_state[dc+"-"+wb] = 0.5
	}

	if wb == "google" {
		if website.Google[len(website.Google)-2] == 0.5 {
			google++
		}
		isp_state[dc+"-"+wb] = 0.5
	}

	if wb == "wikipedia" {
		if website.Wikipedia[len(website.Wikipedia)-2] == 0.5 {
			wikipedia++
		}
		isp_state[dc+"-"+wb] = 0.5
	}

	if wb == "varzesh3" {
		if website.Varzesh3[len(website.Varzesh3)-2] == 0.5 {
			varzesh3++
		}
		isp_state[dc+"-"+wb] = 0.5
	}
}

func NetworkChecker(current_time time.Time, aparat, varzesh3, digikala, github, google, instagram, wikipedia, clubhouse int) {
	// Intranet has problem - if 3 ISP have problem in same time
	if (aparat >= 3) && (varzesh3 >= 3) && (digikala >= 3) {
		//send_alert('', 0, 1, 0)
		print(current_time.String() + " Based on radar.arvan.com interanet has problem.")
	}

	// Internet has problem - if 4 data centers have problem in samte time
	if (instagram >= 4) && (google >= 4) && (github >= 4) {
		print(current_time.String() + " 1. Based on radar.arvan.com internet has problem.")
		// send_alert('', 0, 0, 1)
	} else if (instagram >= 4) && (google >= 4) && (clubhouse >= 4) {
		print(current_time.String() + " 2. Based on radar.arvan.com internet has problem.")
		// send_alert('', 0, 0, 1)
	} else if (instagram >= 4) && (github >= 4) && (clubhouse >= 4) {
		// send_alert('', 0, 0, 1)
		print(current_time.String() + " 3. Based on radar.arvan.com internet has problem.")
	} else if (instagram >= 4) && (github >= 4) && (wikipedia >= 4) {
		// send_alert('', 0, 0, 1)
		print(current_time.String() + " 4. Based on radar.arvan.com internet has problem.")
	} else if (instagram >= 4) && (google >= 4) && (wikipedia >= 4) {
		// send_alert('', 0, 0, 1)
		print(current_time.String() + " 5. Based on radar.arvan.com internet has problem.")
	} else if (instagram >= 4) && (clubhouse >= 4) && (wikipedia >= 4) {
		// send_alert('', 0, 0, 1)
		print(current_time.String() + " 6. Based on radar.arvan.com internet has problem.")
	} else if (google >= 4) && (github >= 4) && (clubhouse >= 4) {
		// send_alert('', 0, 0, 1)
		print(current_time.String() + " 7. Based on radar.arvan.com internet has problem.")
	} else if (google >= 4) && (github >= 4) && (wikipedia >= 4) {
		// send_alert('', 0, 0, 1)
		print(current_time.String() + " 8. Based on radar.arvan.com internet has problem.")
	}
}
