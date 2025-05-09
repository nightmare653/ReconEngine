import streamlit as st
import subprocess
import os
import json
from glob import glob

st.set_page_config(page_title="ReconEngine Dashboard", layout="wide")

st.sidebar.title("🛠️ ReconEngine Controls")
menu = st.sidebar.radio("Select Action", ["🛱 Launch Scan", "📊 View Results"])

# Paths
BASE_OUTPUT_DIR = "output"
DOMAINS_FILE = "domains.txt"

if menu == "🛱 Launch Scan":
    st.header("🛱 Launch Recon Scan")

    if not os.path.exists(DOMAINS_FILE):
        st.error("❌ domains.txt not found.")
    else:
        with open(DOMAINS_FILE) as f:
            domain_list = [line.strip() for line in f if line.strip()]

        col1, col2 = st.columns([3, 1])
        selected_domain = col1.selectbox("Select a domain:", domain_list)
        scan_all = col2.checkbox("🔁 Scan All Domains")

        if st.button("🚀 Start Scan"):
            log_output = st.empty()

            def run_scan(domain):
                log_output.markdown(f"**🔍 Scanning `{domain}`...**")
                process = subprocess.Popen(["./reconengine", "-d", domain],
                                           stdout=subprocess.PIPE, stderr=subprocess.STDOUT, text=True)
                output = ""
                for line in process.stdout:
                    output += line
                    log_output.code(output, language="bash")
                process.wait()
                return process.returncode

            if scan_all:
                st.info("🔁 Running batch scan for all domains...")
                failed = []

                for domain in domain_list:
                    ret = run_scan(domain)
                    if ret != 0:
                        failed.append(domain)
                        st.error(f"❌ Failed: {domain}")
                    else:
                        st.success(f"✅ Completed: {domain}")

                if failed:
                    st.warning(f"⚠️ Scan failed for: {', '.join(failed)}")
                else:
                    st.success("🎉 All scans completed successfully.")
            else:
                ret = run_scan(selected_domain)
                if ret != 0:
                    st.error(f"❌ Scan failed for {selected_domain}")
                else:
                    st.success(f"✅ Completed scan for {selected_domain}")

elif menu == "📊 View Results":
    st.header("📊 Recon Results Viewer")

    domains = sorted([os.path.basename(d) for d in glob(f"{BASE_OUTPUT_DIR}/*") if os.path.isdir(d)])
    if not domains:
        st.warning("⚠️ No scanned domains found yet.")
    else:
        selected_domain = st.selectbox("Select domain to view results:", domains)
        summary_file = os.path.join(BASE_OUTPUT_DIR, selected_domain, "recon_summary.json")

        if not os.path.exists(summary_file):
            st.warning(f"⚠️ No summary found for {selected_domain}")
        else:
            with open(summary_file) as f:
                data = json.load(f)

            st.subheader("📌 Recon Summary")
            st.json(data, expanded=False)
