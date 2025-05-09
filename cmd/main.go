// ✅ ReconEngine Main Controller (main.go)
package main

import (
	"ReconEngine/modules"
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func runRecon(domain string) {
	outputDir := filepath.Join("output", domain)
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		fmt.Printf("❌ Failed to create output directory: %v\n", err)
		return
	}

	fmt.Printf("🚀 Starting recon on: %s\n", domain)

	modules.RunSubfinder(domain, outputDir)
	modules.RunSubdominator(domain, outputDir)
	modules.RunHttpx(domain, outputDir)
	modules.RunSubzy(domain, outputDir)
	//modules.RunWappalyzer(domain, outputDir)
	//modules.RunGoLinkFinder(domain, outputDir)
	//modules.RunWaybackUrls(domain, outputDir)
	//modules.RunWaymore(domain, outputDir)
	modules.RunGau(domain, outputDir)
	modules.RunOTXFetcher(domain, outputDir)
	modules.RunHakrawler(domain, outputDir)
	//modules.RunSecretScanner(domain, outputDir) // stealth mode ON
	modules.RunURLPostProcessor(domain, outputDir)
	modules.RunGFScanner(domain, outputDir)
	modules.RunRegexFlagger(domain, outputDir)
	modules.RunDisclo(domain, outputDir)
	modules.RunHakCheckURL(domain, outputDir)
	modules.RunJSSecretScanner(domain, outputDir)
	modules.RunSecretFinder(domain, outputDir)
	modules.RunParamSpider(domain, outputDir)
	modules.RunArjun(domain, outputDir)
	modules.RunGitleaks(domain, outputDir)
	modules.RunWhatWeb(domain, outputDir)
	modules.RunGitHound(domain, outputDir)
	//modules.RunCorsy(domain, outputDir)    // ✅ Added Corsy

	//modules.RunFfuf(domain, outputDir)

	modules.RunSummaryWriter(domain, outputDir)
	modules.RunAIPlanner(domain, outputDir)
	modules.RunAINucleiScan(domain, outputDir)

	fmt.Println("✅ Recon complete.")
}

func main() {
	domain := flag.String("d", "", "Target domain (e.g., example.com)")
	list := flag.String("list", "", "Path to file with list of domains")
	flag.Parse()

	if *domain == "" && *list == "" {
		fmt.Println("Usage:")
		fmt.Println("  ./reconengine -d example.com")
		fmt.Println("  ./reconengine --list domains.txt")
		os.Exit(1)
	}

	if *list != "" {
		file, err := os.Open(*list)
		if err != nil {
			fmt.Printf("❌ Failed to open domain list file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())
			if line != "" && !strings.HasPrefix(line, "#") {
				runRecon(line)
			}
		}

		if err := scanner.Err(); err != nil {
			fmt.Printf("❌ Error reading domain list: %v\n", err)
		}
	} else {
		runRecon(*domain)
	}
}
