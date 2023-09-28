package contents

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetDroneYamlFile(repo string, tag string) (respdata []byte, err error) {
	url := "https://127.0.0.1:9090/build/v1/appcc/drone"
	// 构建 POST 请求的数据
	data := map[string]string{
		"repo": repo,
		"tag":  tag,
	}

	// 将数据编码为 JSON 格式
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	// 发送 POST 请求
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 检查响应状态码
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP request failed with status code: %d", resp.StatusCode)
	}

	// 读取响应体
	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var jsonDataStruct struct {
		Code    int    `json:"code"`
		Data    string `json:"data"`
		Message string `json:"message"`
	}
	err = json.Unmarshal([]byte(responseData), &jsonDataStruct)
	if err != nil {
		fmt.Println("解析 JSON 数据时出错:", err)
		return
	}
	dataValue := jsonDataStruct.Data
	return []byte(dataValue), err
}
