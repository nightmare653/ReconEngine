// ✅ modules/auth_requester.go
package modules

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func RunAuthenticatedRequester(domain, outputDir string) {
	fmt.Println("🔐 Running Authenticated API Requester...")

	headersPath := filepath.Join("config", "auth_headers.json")
	headersFile, err := os.ReadFile(headersPath)
	if err != nil {
		fmt.Printf("❌ Could not read auth headers: %v\n", err)
		return
	}

	headers := map[string]string{}
	if err := json.Unmarshal(headersFile, &headers); err != nil {
		fmt.Printf("❌ Invalid auth_headers.json format: %v\n", err)
		return
	}

	urlsPath := filepath.Join(outputDir, "auth_api_candidates.txt")
	urlList, err := os.ReadFile(urlsPath)
	if err != nil {
		fmt.Printf("❌ Could not read candidate URLs: %v\n", err)
		return
	}

	lines := strings.Split(string(urlList), "\n")
	for _, raw := range lines {
		url := strings.TrimSpace(raw)
		if url == "" {
			continue
		}

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Printf("❌ Request build failed: %v\n", err)
			continue
		}

		for k, v := range headers {
			req.Header.Set(k, v)
		}

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("❌ Request failed: %s → %v\n", url, err)
			continue
		}
		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)
		fmt.Printf("🔍 %s → %d bytes\n", url, len(body))
	}

	fmt.Println("✅ Authenticated requests complete.")
}
