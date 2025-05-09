# 🛡️ ReconEngine - Automated Web Reconnaissance Framework

ReconEngine is a powerful and extensible reconnaissance engine built for bug bounty hunters, penetration testers, and red teamers. It automates the full recon stack from subdomain enumeration to AI-powered vulnerability detection, offering both CLI and dashboard UI.

---

## 🚀 Features

- 🔍 **Subdomain Enumeration**: `subfinder`, `Subdominator`, `gau`, `waybackurls`, `hakrawler`
- 🌐 **Host Probing**: `httpx`, `WhatWeb`, `Wappalyzer`
- 🧠 **AI + Nuclei Integration**: Natural language prompt-based scans using Nuclei's `-ai` mode
- 🔐 **Secret Detection**: Advanced multi-pattern secret detection via JS & HTML scan, plus Git, config & environment leaks
- 🧬 **JS Analysis**: `goLinkFinder`, `ParamSpider`, `Arjun`, Wayback inspection
- 🧩 **Vulnerability Mapping**: GF pattern matching, regex flagging, sensitive file detection
- 🧰 **Tool Integration**: Gitleaks, Git-Hound, Corsy, Subzy, Disclo PDF scanner
- 📊 **Streamlit Dashboard**: Launch scans, track progress, and view summaries
- 📄 **Structured Reports**: `recon_summary.json`, secrets, endpoints, parameters, headers

---

## 🧱 Project Structure

ReconEngine/
├── cmd/
│ └── main.go # Entrypoint for the Go engine
├── modules/ # Go scanning modules
│ ├── subfinder.go
│ ├── httpx.go
│ ├── ai_nuclei_runner.go
│ └── ...
├── utils/
│ └── helper.go # Host extraction, file utils
├── tools/
│ └── secret_detector/ # Python-based JS/HTML/Git secret scanner
├── dashboard.py # Streamlit UI
├── domains.txt # List of targets
├── run_recon.sh # CLI batch launcher
└── README.md


---

## ⚙️ Installation

### 1. Clone the repository

```bash
git clone https://github.com/yourusername/ReconEngine.git
cd ReconEngine
```

2. Install prerequisites
🧪 Golang Tools
```
go install github.com/projectdiscovery/subfinder/v2/cmd/subfinder@latest
go install github.com/projectdiscovery/httpx/cmd/httpx@latest
go install github.com/projectdiscovery/nuclei/v3/cmd/nuclei@latest
```
🐍 Python Tools
```
pip install -r tools/secret_detector/requirements.txt
pip install streamlit
```

🐧 Optional External Tools

Install and add to $PATH:
```
    git-hound

    gitleaks

    corsy

    subzy
```
🔍 Usage
🔸 Single Scan (CLI)
```
./reconengine -d example.com
```
🔸 Batch Scan

Add domains to domains.txt:
```
example.com
target.org
staging.site
```
Run with:
```
./run_recon.sh --list domains.txt
```
🖥️ Dashboard (Streamlit)

Launch:
```
streamlit run dashboard.py
```

Features

    ✅ Select single or all domains to scan

    📡 View real-time logs

    📊 View results in structured UI

    🧠 Uses AI prompt from ai_plan.txt

  📂 Output Structure

Each scan produces a directory like:

```
output/example.com/
├── subdomains.txt
├── live_hosts.txt
├── recon_summary.json
├── nuclei_ai_*.json
├── js_secrets.txt
├── gf_xss.txt
├── secretfinder.txt
└── ...

```

🧠 AI Planner + Nuclei Integration

You can define an ai_plan.txt with any of these prompts:

```
Extract email addresses from web pages
Find exposed .env files leaking credentials
Identify admin login endpoints, filter 404 response code

```

This prompt is passed into nuclei -ai "<prompt>" automatically.

📑 Built-in Modules
Module	Purpose
subfinder	Passive subdomain enumeration
Subdominator	API-powered subdomain scanner
httpx	Live host probing & tech detection
goLinkFinder	JS endpoint extraction
gau, hakrawler	Historical & JS crawling
secret_detector	Regex-based secrets detection
ParamSpider	Param discovery
Arjun	More param guessing
Gitleaks	Secrets in repo/code
Git-Hound	GitHub secret scraping
Subzy	Subdomain takeover scanner
Corsy	CORS misconfig detection
WhatWeb	Stack fingerprinting
Disclo	PDF disclosure detection

⚠️ Legal Disclaimer

This project is intended for authorized security testing and educational use only. Do not scan systems you do not own or have explicit permission to test.


📜 License

MIT License © 2025 ReconEngine Authors @nigthmare653
