package main

import(
	"fmt"
	"io/ioutil"
	"net/http"
	"bytes"

	// "training-go-api/client"
)

func main() {
		// ファイルパス
		filePath := "input/sample.tar.gz"

		// ファイルを読み込み
		fileData, err := ioutil.ReadFile(filePath)
		if err != nil {
			fmt.Println("Failed to read file:", err)
			return
		}
	
		// POST先のURL
		url := "http://localhost:18080/v1/upload"
	
		// リクエストを作成
		resp, err := http.Post(url, "application/octet-stream", bytes.NewBuffer(fileData))
		if err != nil {
			fmt.Println("Failed to send POST request:", err)
			return
		}
		defer resp.Body.Close()
	
		// レスポンスを表示
		fmt.Println("Response Status:", resp.Status)
		fmt.Println("Response Body:")
		responseBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Failed to read response body:", err)
			return
		}
		fmt.Println(string(responseBody))	
}
