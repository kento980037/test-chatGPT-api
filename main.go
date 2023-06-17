package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"kento980037/test-chatGPT-api/model"
	"net/http"
	"os"
)

const openaiURL = "https://api.openai.com/v1/chat/completions"

func main() {
	//　まず最初にapikeyを環境変数に入れる
	// export CHATGPT_APIKEY='xxxxxx'
	apiKey := os.Getenv("CHATGPT_APIKEY")

	requestBody := model.RequestBody{
		Model: "gpt-3.5-turbo",
		Messages: []model.Message{
			{
				Role:    "system",
				Content: "You are a helpful assistant.",
			},
			{
				Role:    "user",
				Content: "Hello!",
			},
		},
	}

	// JSONにエンコード
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Println("JSONのエンコードに失敗しました:", err)
		return
	}

	// リクエストの作成
	req, err := http.NewRequest("POST", openaiURL, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("リクエストの作成に失敗しました:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	// HTTPリクエストの送信
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("リクエストの送信に失敗しました:", err)
		return
	}
	defer resp.Body.Close()

	// レスポンスボディの読み取り
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("レスポンスの読み取りに失敗しました:", err)
		return
	}

	// レスポンスデータのパース
	var responseBody model.ResponseBody
	err = json.Unmarshal(body, &responseBody)
	if err != nil {
		fmt.Println("レスポンスデータのパースに失敗しました:", err)
		return
	}

	// レスポンスボディの表示
	fmt.Println(responseBody)
	fmt.Println(responseBody.Choices[0].Message.Content)

}
