package karl_ding

import (
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func Notice(pWebHook string, pKeyword string, pMsg string) (string, error) {
	webHook := pWebHook
	content := `{"msgtype": "text",
		"text": {"content": "` + pKeyword + `:\n` + pMsg + `"}
	}`
	//创建一个请求
	req, err := http.NewRequest("POST", webHook, strings.NewReader(content))
	if err != nil {
		return "", err
	}

	client := &http.Client{
		Timeout: time.Duration(10) * time.Minute,
	}
	//设置请求头
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-agent", "firefox")
	//发送请求
	resp, err := client.Do(req)
	//关闭请求
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return strconv.Itoa(resp.StatusCode) + "--" + string(body), nil
}
