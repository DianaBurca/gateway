package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func doRequest(url string, data io.Reader, method string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, data)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func InfoHandler(c *gin.Context) {
	params := c.Request.URL.Query()
	city := params["city"][0]

	go func(city string) {
		jsonData := map[string]string{"city_name": city}
		bytes, _ := json.Marshal(jsonData)
		payloadStr := string(bytes)
		data := strings.NewReader(payloadStr)
		doRequest("http://cron/store", data, http.MethodPut)
	}(city)

	url := fmt.Sprintf("http://info-reader/read?city_name=%s", city)
	resp, err := doRequest(url, nil, http.MethodGet)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	var bodyBytes []byte
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		payload := map[string]string{}
		json.Unmarshal(bodyBytes, &payload)
		c.JSON(http.StatusOK, payload)
	}
	defer resp.Body.Close()
}
