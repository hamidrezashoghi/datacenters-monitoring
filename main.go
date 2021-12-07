package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

const (
	mci_tehran  int = 0
	mci_mashhad int = 0
	mci_tabriz  int = 0
	mci_esfehan int = 0
	mci_shiraz  int = 0
	afranet     int = 0
	irancell    int = 0
	mobinnet    int = 0
	shatel      int = 0

	// Keep datacenters state
	dc_mci_tehran_state  int = 0
	dc_mci_mashhad_state int = 0
	dc_mci_tabriz_state  int = 0
	dc_mci_esfehan_state int = 0
	dc_mci_shiraz_state  int = 0
	dc_afranet_state     int = 0
	dc_irancell_state    int = 0
	dc_mobinnet_state    int = 0
	dc_shatel_state      int = 0

	// Radar monitor below websites
	aparat    int = 0
	varzesh3  int = 0
	digikala  int = 0
	github    int = 0
	google    int = 0
	instagram int = 0
	wikipedia int = 0
	clubhouse int = 0
)

var data_dict map[string]string
var datacenters = []string{"mci", "mci-mashhad", "mci-tabriz", "mci-esfehan", "mci-shiraz", "afranet", "irancell", "mobinnet", "shatel"}

func main() {
	// now := time.Now()
	url := ""
	for _, isp := range datacenters {
		url = "https://radar.arvancloud.com/api/v1/internet-monitoring?isp=" + isp
		fmt.Println(isp)
		fmt.Println(getJSON(url))
	}
}

func getJSON(url string) interface{} {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("can't fetch URL %q: %v", url, err)
	}

	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return string(b)

}
