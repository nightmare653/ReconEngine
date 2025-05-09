package modules

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"ReconEngine/utils"
)

var aiPrompts = []string{
	"Extract page title, detect tech and versions",
	"Extract email addresses from web pages",
	"Extract all subdomains referenced in web pages",
	"Find sensitive information in HTML comments (debug notes, API keys, credentials)",
	"Find exposed .env files leaking credentials, API keys, and database passwords",
	"Find exposed configuration files such as config.json, config.yaml, config.php, application.properties containing API keys and database credentials.",
	"Find exposed database configuration files such as database.yml, db_config.php, .pgpass, .my.cnf leaking credentials.",
	"Find exposed Docker and Kubernetes configuration files such as docker-compose.yml, kubeconfig, .dockercfg, .docker/config.json containing cloud credentials and secrets.",
	"Find exposed SSH keys and configuration files such as id_rsa, authorized_keys, and ssh_config.",
	"Find exposed WordPress configuration files (wp-config.php) containing database credentials and authentication secrets.",
	"Identify open directory listings exposing sensitive files",
	"Find exposed .git directories allowing full repo download",
	"Find exposed .svn and .hg repositories leaking source code",
	"Identify open FTP servers allowing anonymous access",
	"Detect debug endpoints revealing system information",
	"Identify test and staging environments exposed to the internet",
	"Find admin login endpoints, filter 404 response code",
	"Detect exposed stack traces in error messages",
	"Identify default credentials on login pages",
	"Find misconfigured Apache/Nginx security headers",
	"Scan for exposed environment files (.env) containing credentials",
	"Find open directory listings and publicly accessible files",
	"Detect exposed .git repositories and sensitive files",
	"Identify publicly accessible backup and log files (.log, .bak, .sql, .dump)",
	"Detect exposed .htaccess and .htpasswd files",
	"Check for SSH private keys leaked in web directories",
	"Find exposed API keys and secrets in responses and URLs",
	"Identify API endpoints leaking sensitive data",
	"Find leaked database credentials in JavaScript files",
}

func RunAINucleiScan(domain, outputDir string) {
	rawFile := filepath.Join(outputDir, "live_hosts.txt")
	filteredFile := filepath.Join(outputDir, "filtered_hosts.txt")

	if _, err := os.Stat(rawFile); os.IsNotExist(err) {
		fmt.Printf("[!] live_hosts.txt not found: %s\n", rawFile)
		return
	}

	if err := utils.Extract200OKURLs(rawFile, filteredFile); err != nil {
		fmt.Printf("[✗] Failed to extract 200 OK URLs: %v\n", err)
		return
	}

	for i, prompt := range aiPrompts {
		cleanPrompt := sanitizeForCLI(prompt)
		timestamp := time.Now().Format("20060102-150405")
		outFile := filepath.Join(outputDir, fmt.Sprintf("nuclei_ai_%02d_%s.json", i+1, timestamp))

		cmd := exec.Command(
			"nuclei",
			"-list", filteredFile,
			"-ai", cleanPrompt,
			"rate-limit", "2",
			"c", "2",
			"-json-export", outFile,
		)

		fmt.Printf("[+] Running Nuclei AI scan (%d/%d) with prompt: \"%s\"\n", i+1, len(aiPrompts), cleanPrompt)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			fmt.Printf("[✗] Scan failed for prompt %d: %v\n", i+1, err)
		} else {
			fmt.Printf("[✓] Scan %d complete. Output saved to: %s\n", i+1, outFile)
		}
	}
}

func sanitizeForCLI(input string) string {
	replacer := strings.NewReplacer(
		"\n", " ", "\r", " ", "\t", " ",
		"\"", "", "'", "", "`", "", ";", "",
		"&", "and", "|", "", "$", "", "!", "",
	)
	cleaned := replacer.Replace(input)
	cleaned = strings.TrimLeft(cleaned, "-")
	return strings.Join(strings.Fields(cleaned), " ")
}
