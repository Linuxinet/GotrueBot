package truecaller

import (
	"bytes"
	"io"
	"log"
	"net/http"
)

const TOKEN = "Bearer a2i07--aK29_z-v-MGjMpctUve_wfqc4-zlWUk-M4ZMoAZd_MZLDIUouO6USj649"

type Response struct {
	Data []struct {
		Name       string       `json:"name"`
		Address    Addresses    `json:"addresses"`
		Phone      Phones       `json:"phones"`
		INTaddress INTaddresses `json:"internetaddresses"`
	} `json:"data"`
}

type Addresses struct {
	City string `json:"city"`
}

type Phones struct {
	Mobile  string `json:"e164Format"`
	Carrier string `json:"carrier"`
}

type INTaddresses struct {
	Email string `json:"id"`
}

func Auth() {

	auth_url := "https://account-asia-south1.truecaller.com/v2.1/credentials/check?encoding=json"

	data := []byte(`{"reason": "restored_from_account_manager"}`)

	r, err := http.NewRequest("POST", auth_url, bytes.NewBuffer(data))
	if err != nil {
		log.Panic(err)
	}

	r.Header.Add("Host", "account-asia-south1.truecaller.com")
	r.Header.Add("authorization", TOKEN)
	r.Header.Add("content-type", "application/json; charset=UTF-8")
	r.Header.Add("content-length", "42")
	r.Header.Add("accept-encoding", "gzip")
	r.Header.Add("user-agent", "Truecaller/11.5.7 (Android;10)")

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		log.Panic(err)
	}
	defer res.Body.Close()
}

func Search_num(mobilenumber string) []byte {

	search_api := "https://search5-noneu.truecaller.com/v2/search?q=" + mobilenumber + "&countryCode=IN&type=4&locAddr=&placement=SEARCHRESULTS%2CHISTORY%2CDETAILS&encoding=json"

	s, err := http.NewRequest("GET", search_api, nil)
	if err != nil {
		log.Println(err)
	}

	s.Header.Add("Host", "search5-noneu.truecaller.com")
	s.Header.Add("authorization", TOKEN)
	s.Header.Add("user-agent", "Truecaller/11.5.7 (Android;10)")

	client := &http.Client{}
	search_res, err := client.Do(s)
	if err != nil {
		log.Println(err)
	}

	body, _ := io.ReadAll(search_res.Body)
	if err != nil {
		log.Println(err)
	}
	defer search_res.Body.Close()

	return body

}
