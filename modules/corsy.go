package modules

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func RunCorsy(domain, outputDir string) {
	fmt.Println("🔍 Running Corsy...")

	inputFile := filepath.Join(outputDir, "all_urls.txt")
	outputFile := filepath.Join(outputDir, "corsy.json")

	// Check if input file exists
	if _, err := os.Stat(inputFile); os.IsNotExist(err) {
		fmt.Println("❌ all_urls.txt not found. Skipping Corsy.")
		return
	}

	// Create output file handle
	outFile, err := os.Create(outputFile)
	if err != nil {
		fmt.Printf("❌ Failed to create corsy.json: %v\n", err)
		return
	}
	defer outFile.Close()

	// Read URLs from file
	f, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("❌ Failed to read all_urls.txt: %v\n", err)
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var urls []string
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "http://") || strings.HasPrefix(line, "https://") {
			urls = append(urls, line)
		}
	}
	if len(urls) == 0 {
		fmt.Println("❌ No valid URLs found for Corsy.")
		return
	}

	// Write URLs to temporary input file for Corsy
	tmpFile := filepath.Join(outputDir, "corsy_input.txt")
	if err := os.WriteFile(tmpFile, []byte(strings.Join(urls, "\n")), 0644); err != nil {
		fmt.Printf("❌ Failed to write temporary input for Corsy: %v\n", err)
		return
	}

	// Run Corsy
	cmd := exec.Command("python3", "tools/Corsy/corsy.py", "-i", tmpFile, "-o", outputFile)
	cmd.Stdout = outFile
	cmd.Stderr = outFile

	if err := cmd.Run(); err != nil {
		fmt.Printf("❌ Corsy error: %v\n", err)
		return
	}

	fmt.Printf("✅ Corsy results saved to: %s\n", outputFile)
}
