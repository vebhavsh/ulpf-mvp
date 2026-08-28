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
