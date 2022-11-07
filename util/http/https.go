package http

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"time"
)

func HttpsPost(url string, arg interface{}) (string, error) {
	var (
		response *http.Response
		body     []byte
		buf      *bytes.Buffer
	)
	var err error
	if arg != nil {
		if b, ok := arg.([]byte); !ok {
			if body, err = json.Marshal(arg); err != nil {
				return "", err
			}
		} else {
			body = b
		}
	}
	fmt.Println("[https-Post-request:]%s", string(body))
	buf = bytes.NewBuffer(body)
	tr := &http.Transport{
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
		DisableCompression: true,
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   time.Duration(10) * time.Minute,
	}
	//启用cookie
	client.Jar, _ = cookiejar.New(nil)
	response, err = client.Post(url, "application/json;charset=utf-8", buf)
	if err != nil {
		return "", err
	}

	if response != nil {
		defer response.Body.Close()
	} else {
		return "", err
	}
	if body, err = io.ReadAll(response.Body); err != nil {
		return "", err
	}
	return string(body), nil
}

func HttpsGet(url string) (string, error) {
	var (
		response *http.Response
	)
	tr := &http.Transport{
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
		DisableCompression: true,
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   time.Duration(10) * time.Minute,
	}
	//启用cookie
	client.Jar, _ = cookiejar.New(nil)
	///log.Println(url)
	response, err := client.Get(url)
	if err != nil {
		return "", err
	}
	if response != nil {
		defer response.Body.Close()
	} else {
		return "", err
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	//log4plus.Info("[Get]%#v", string(body))
	return string(body), nil
}
