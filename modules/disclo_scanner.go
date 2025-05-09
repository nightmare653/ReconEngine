package modules

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func RunDisclo(domain, outputDir string) {
	fmt.Println("🔍 Running Disclo PDF scanner...")

	inputFile := filepath.Join(outputDir, "all_urls.txt")
	outputFile := filepath.Join(outputDir, "pdf_keywords.txt")

	// Ensure input exists
	if _, err := os.Stat(inputFile); os.IsNotExist(err) {
		fmt.Println("❌ all_urls.txt not found, skipping Disclo.")
		return
	}

	cmd := exec.Command("./tools/disclo/disclo.sh", inputFile, outputFile)
	cmd.Dir = "." // ensure relative path
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("❌ Disclo error: %v\n", err)
		return
	}

	fmt.Println("✅ Disclo done.")
	fmt.Println(string(output))
}
