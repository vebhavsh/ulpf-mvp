#  ULPF - Universal Log Pre-processing Framework (MVP)

A blazing-fast, real-time log ingestion and parsing engine built for SIEMs. 
This MVP ingests unstructured raw firewall logs, extracts critical indicators using a custom regex engine, normalizes them into the **OCSF (Open Cybersecurity Schema Framework)** standard, and visualizes the data live using Grafana.

---

## ⚙️ Prerequisites (What you need installed)

To run this project on a brand new machine, you will need:
1. **[Git](https://git-scm.com/downloads)** - To clone this repository.
2. **[Go (Golang)](https://go.dev/doc/install)** - To run the backend parsing engine.
3. **[Grafana](https://grafana.com/grafana/download)** - For the frontend visualization dashboard.

---

##  Step-by-Step Setup Guide

### Phase 1: Start the Go Backend
1. Open your terminal/command prompt and clone this repository:
   ```bash
   git clone https://github.com/vebhavsh/ulpf-mvp.git
   cd ulpf-mvp
Start the Golang processing server:

Bash
go run cmd/server/main.go
Success! You should see a message saying the server is running on port 8080. The engine will automatically create a local ulpf.db SQLite database in your folder to store the logs safely. Keep this terminal open.

Phase 2: Setup the Grafana Dashboard
Because Grafana doesn't support SQLite out-of-the-box, we need to add a quick plugin.

Open a new terminal and install the SQLite plugin for Grafana:

Bash
grafana-cli plugins install frser-sqlite-datasource
Restart your Grafana service on your PC so the plugin loads.

Open your browser and go to http://localhost:3000 (Login with default username admin, password admin).

Go to Connections > Data Sources > Add data source. Search for SQLite.

In the path section, provide the absolute path to the ulpf.db file located in your project folder (e.g., C:\Users\YourName\ulpf-mvp\ulpf.db). Click Save & Test.

Phase 3: Import the UI
In Grafana, click the + (Plus) icon in the top right or go to Dashboards > New > Import.

Upload the ulpf-dashboard.json file provided in this repository.

Select your SQLite data source from the dropdown and click Import.

🔥 Live Action Demo (How to Test)
Now that both your backend and frontend are running, let's inject a fake firewall log to see the real-time processing.

Open a new terminal window.

Run this cURL command to send a raw log to the Go engine:

Bash
curl -X POST -d "srcip=192.168.9.99 dstip=1.1.1.1 action=deny" http://localhost:8080/ingest
Open your Grafana dashboard and hit the Refresh button at the top right.

You will instantly see the new log parsed into OCSF JSON format in the data table, and the Action Pie Chart will update dynamically!

Built for Smart India Hackathon 2026 by Team Hydra 🧠
