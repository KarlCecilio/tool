package http

import (
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func HttpGet(url string) (string, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}
	//设置请求头
	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	req.Header.Set("User-agent", "firefox")
	client := &http.Client{
		Timeout: time.Duration(10) * time.Minute,
	}
	//发送请求
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	//关闭请求
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	return strconv.Itoa(resp.StatusCode) + "--" + string(body), nil
}

func HttpPost(url string, reqbody string) (string, error) {
	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(reqbody))
	if err != nil {
		return "", err
	}
	//设置请求头
	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	req.Header.Set("User-agent", "firefox")

	client := &http.Client{
		Timeout: time.Duration(10) * time.Minute,
	}
	//发送请求
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	//关闭请求
	defer resp.Body.Close()
	resbody, _ := io.ReadAll(resp.Body)
	return strconv.Itoa(resp.StatusCode) + "--" + string(resbody), nil
}
