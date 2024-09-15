package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

type TestCase struct {
	Name                 string                 `yaml:"name"`
	Method               string                 `yaml:"method"`
	Endpoint             string                 `yaml:"endpoint"`
	Headers              map[string]string      `yaml:"headers"`
	Body                 map[string]interface{} `yaml:"body,omitempty"`
	ExpectedStatus       int                    `yaml:"expected_status"`
	ExpectedBodyContains string                 `yaml:"expected_body_contains"`
}

type TestSuite struct {
	Tests []TestCase `yaml:"tests"`
}

func main() {
	// --file オプションを定義
	filePath := flag.String("file", "", "Path to the YAML test file")
	flag.Parse()

	// ファイルパスが指定されているか確認
	if *filePath == "" {
		log.Fatal("Please provide a YAML test file using --file option")
	}

	// YAMLファイルの読み込み
	data, err := os.ReadFile(*filePath)
	if err != nil {
		log.Fatalf("Failed to read YAML file: %v", err)
	}

	// YAMLのパース
	var testSuite TestSuite
	err = yaml.Unmarshal(data, &testSuite)
	if err != nil {
		log.Fatalf("Failed to parse YAML: %v", err)
	}

	// テストケースの実行

	for _, testCase := range testSuite.Tests {
		fmt.Printf("Running test: %s\n", testCase.Name)

		// リクエスト作成
		var req *http.Request
		if testCase.Method == "POST" || testCase.Method == "PUT" {
			jsonBody, err := json.Marshal(testCase.Body)
			if err != nil {
				log.Fatalf("Failed to marshal request body: %v", err)
			}
			req, _ = http.NewRequest(testCase.Method, testCase.Endpoint, bytes.NewBuffer(jsonBody))
		} else {
			req, _ = http.NewRequest(testCase.Method, testCase.Endpoint, nil)
		}

		// ヘッダーを設定
		for key, value := range testCase.Headers {
			req.Header.Set(key, value)
		}

		// HTTPクライアントを使ってリクエストを送信
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("Test failed: %s - Error: %v\n", testCase.Name, err)
			continue
		}
		defer resp.Body.Close()

		// ステータスコードの検証
		if resp.StatusCode != testCase.ExpectedStatus {
			fmt.Printf("Test failed: %s - Expected status %d but got %d\n",
				testCase.Name, testCase.ExpectedStatus, resp.StatusCode)
			continue
		}

		// レスポンスボディの検証
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Failed to read response body: %v\n", err)
			continue
		}
		bodyString := string(bodyBytes)
		if !strings.Contains(bodyString, testCase.ExpectedBodyContains) {
			fmt.Printf("Test failed: %s - Expected body to contain '%s'\n",
				testCase.Name, testCase.ExpectedBodyContains)
			continue
		}

		// 成功
		fmt.Printf("Test passed: %s\n", testCase.Name)
	}
}
