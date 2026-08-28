 ULPF — Universal Log Pre-processing Framework (MVP)

A blazing-fast, real-time log ingestion and parsing engine built for SIEMs.

ULPF ingests unstructured raw firewall logs, extracts critical security indicators using a custom regex engine, normalizes them into the OCSF (Open Cybersecurity Schema Framework) standard, and visualizes the processed data in Grafana.

🧩 Architecture

Raw Firewall Logs
       │
       ▼
┌──────────────────┐
│   Go Backend     │
│  Log Ingestion   │
└────────┬─────────┘
         │
         ▼
┌──────────────────┐
│ Custom Regex     │
│ Parsing Engine   │
└────────┬─────────┘
         │
         ▼
┌──────────────────┐
│ OCSF Normalizer  │
└────────┬─────────┘
         │
         ▼
┌──────────────────┐
│ SQLite Database  │
└────────┬─────────┘
         │
         ▼
┌──────────────────┐
│     Grafana      │
│    Dashboard     │
└──────────────────┘

🛠️ Prerequisites

To run this project on a brand-new machine, make sure you have:

Tool

Purpose

Git

Clone the repository

Go (Golang)

Run the backend parsing engine

Grafana

Frontend visualization dashboard

⚙️ Step-by-Step Setup Guide

Phase 1 — Start the Go Backend

1. Clone the repository

Open your terminal / Command Prompt:

git clone https://github.com/vebharsh/ulpf-mvp.git
cd ulpf-mvp

2. Start the Go server

Run:

go run cmd/server/main.go

You should see a message similar to:

Server is running on port 8080

The engine will automatically create a local SQLite database:

ulpf.db

Keep this terminal running while using the application.

Phase 2 — Set Up Grafana

Grafana does not support SQLite out-of-the-box, so we use the frser-sqlite-datasource plugin.

1. Install the SQLite plugin

Open a new terminal and run:

grafana-cli plugins install frser-sqlite-datasource

Restart Grafana after installation so the plugin loads correctly.

2. Open Grafana

Open your browser and visit:

http://localhost:3000

Default credentials:

Username: admin
Password: admin

Grafana may ask you to change the default password on first login.

3. Add the SQLite data source

In Grafana:

Connections
   → Data Sources
      → Add data source
         → SQLite

In the Path field, enter the absolute path to the ulpf.db file inside your project folder.

Example:

C:\Users\YourName\ulpf-mvp\ulpf.db

Then click:

Save & Test

Phase 3 — Import the Dashboard

In Grafana, click the ➕ (Plus) icon in the top-right.

Select Dashboards → New → Import.

Upload:

ulpf-dashboard.json

Select your SQLite data source from the dropdown.

Click Import.

Your dashboard should now be ready.

🔥 Live Action Demo — How to Test

Once both the Go backend and Grafana dashboard are running, you can inject a fake firewall log and watch the system process it in real time.

1. Open a new terminal

Run:

curl -X POST -d "src=192.168.99.99 dstip=1.1.1.1 action=deny" http://localhost:8080/ingest

2. Refresh Grafana

Open the dashboard and click Refresh.

You should see:

✅ The new firewall log appear in the data table

✅ The log parsed into OCSF JSON

✅ The Action Pie Chart update dynamically

📊 What ULPF Does

ULPF follows a simple processing pipeline:

Ingest → Parse → Normalize → Store → Visualize

1️⃣ Ingest

Receives raw, unstructured firewall logs.

2️⃣ Parse

Uses a custom regex-based parsing engine to extract useful fields such as:

Source IP

Destination IP

Action

Other security indicators

3️⃣ Normalize

Converts the extracted information into the OCSF standard.

4️⃣ Store

Stores the processed events locally in SQLite.

5️⃣ Visualize

Grafana provides a real-time dashboard for monitoring and analysis.

🎯 MVP Highlights

Feature

Description

⚡ Real-Time Ingestion

Processes logs as they arrive

🔎 Custom Regex Engine

Extracts useful indicators from raw logs

🌐 OCSF Normalization

Converts events into a standardized security schema

💾 SQLite Storage

Lightweight local persistence

📊 Grafana Dashboard

Provides live visualization

🧪 Easy Demo

Test the entire pipeline with a single curl request

📁 Project Structure

ulpf-mvp/
├── cmd/
│   └── server/
│       └── main.go
├── ulpf.db
├── ulpf-dashboard.json
└── README.md

The exact project structure may vary as the MVP evolves.

🚀 Quick Start

For experienced users, the entire setup can be summarized as:

# Clone
git clone https://github.com/vebharsh/ulpf-mvp.git
cd ulpf-mvp

# Start backend
go run cmd/server/main.go

# Install Grafana SQLite plugin
grafana-cli plugins install frser-sqlite-datasource

Then:

Open Grafana → Add SQLite Data Source → Import Dashboard

Test the pipeline:

curl -X POST -d "src=192.168.99.99 dstip=1.1.1.1 action=deny" http://localhost:8080/ingest

🧠 Processing Flow

Firewall Log
     │
     ▼
[ Ingestion API ]
     │
     ▼
[ Regex Parser ]
     │
     ▼
[ OCSF Normalizer ]
     │
     ▼
[ SQLite ]
     │
     ▼
[ Grafana ]
     │
     ▼
Security Visualization

🏆 Built For

Smart India Hackathon 2026

👨‍💻 Team Hydra 🧠

Turning raw security logs into structured, actionable intelligence.
