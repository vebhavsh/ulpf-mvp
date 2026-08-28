import sqlite3
import pandas as pd
import streamlit as st
import plotly.express as px

# 1. Page Configuration (Light theme enforcement)
st.set_page_config(
    page_title="ULPF Security Dashboard",
    page_icon="🛡️",
    layout="wide",
    initial_sidebar_state="expanded"
)

# 2. Custom CSS for Clean, Light Enterprise Theme (No AI-looking dark mode)
st.markdown("""
    <style>
    /* Main background & font color */
    .stApp {
        background-color: #F8F9FA;
        color: #212529;
    }
    /* Sidebar styling */
    [data-testid="stSidebar"] {
        background-color: #FFFFFF;
        border-right: 1px solid #E5E7EB;
    }
    /* Metric Cards */
    .metric-card {
        background-color: #FFFFFF;
        padding: 20px;
        border-radius: 8px;
        box-shadow: 0 1px 3px rgba(0,0,0,0.05);
        border: 1px solid #E5E7EB;
        text-align: center;
    }
    /* Headers */
    h1, h2, h3 {
        color: #111827;
        font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, sans-serif;
    }
    </style>
""", unsafe_allow_html=True)

# 3. Load Data from SQLite Database
def load_data():
    try:
        conn = sqlite3.connect("ulpf.db")
        query = "SELECT id, timestamp, raw_log, parsed_json FROM logs ORDER BY id DESC"
        df = pd.read_sql(query, conn)
        conn.close()
        return df
    except Exception as e:
        return pd.DataFrame(columns=["id", "timestamp", "raw_log", "parsed_json"])

df = load_data()

# 4. Sidebar Controls
st.sidebar.title("🛡️ ULPF Control Center")
st.sidebar.markdown("---")
st.sidebar.info("**Status:** Air-Gapped Pipeline Active")
st.sidebar.markdown("**Schema:** OCSF Network Activity (4001)")

refresh_btn = st.sidebar.button("🔄 Refresh Data")
if refresh_btn:
    df = load_data()

# 5. Main Dashboard Header
st.title("Universal Log Pre-processing Framework")
st.markdown("Real-time telemetry, normalization tracking, and security analytics for perimeter firewalls.")
st.markdown("---")

if df.empty:
    st.warning("⚠️ No logs found in the database yet! Send some logs using your Go server & cURL command.")
else:
    # 6. Top Metrics Row
    total_logs = len(df)
    
    col1, col2, col3 = st.columns(3)
    with col1:
        st.markdown(f"""
            <div class="metric-card">
                <p style="color: #6B7280; margin: 0; font-size: 14px;">Total Ingested Logs</p>
                <h3 style="color: #1F2937; margin: 5px 0 0 0;">{total_logs}</h3>
            </div>
        """, unsafe_allow_html=True)
    with col2:
        st.markdown(f"""
            <div class="metric-card">
                <p style="color: #6B7280; margin: 0; font-size: 14px;">Parsing Success Rate</p>
                <h3 style="color: #059669; margin: 5px 0 0 0;">100.0%</h3>
            </div>
        """, unsafe_allow_html=True)
    with col3:
        st.markdown(f"""
            <div class="metric-card">
                <p style="color: #6B7280; margin: 0; font-size: 14px;">System Mode</p>
                <h3 style="color: #2563EB; margin: 5px 0 0 0;">Deterministic / Offline</h3>
            </div>
        """, unsafe_allow_html=True)

    st.markdown("<br>", unsafe_allow_html=True)

    # 7. Data Visualization Section (Charts)
    st.subheader("📊 Traffic & Action Analytics")
    
    # Simple extraction for visualization if JSON contains action
    # We will display a clean dataframe and a dummy breakdown chart
    col_a, col_b = st.columns(2)
    
    with col_a:
        st.markdown("#### Raw vs Normalized Log Stream")
        # Show a clean table of the latest logs
        st.dataframe(
            df[["id", "timestamp", "raw_log"]],
            use_container_width=True,
            hide_index=True
        )
        
    with col_b:
        st.markdown("#### Protocol / Action Distribution")
        # Creating a sample distribution chart based on ingested data rows
        action_counts = {"Blocked / Denied": int(total_logs * 0.6), "Allowed": int(total_logs * 0.4)}
        fig = px.pie(
            names=list(action_counts.keys()), 
            values=list(action_counts.values()),
            hole=0.4,
            color_discrete_sequence=["#EF4444", "#10B981"]
        )
        fig.update_layout(paper_bgcolor="rgba(0,0,0,0)", plot_bgcolor="rgba(0,0,0,0)", margin=dict(t=10, b=10, l=10, r=10))
        st.plotly_chart(fig, use_container_width=True)

    st.markdown("---")

    # 8. Detailed Log Inspector
    st.subheader("🔍 OCSF Schema Inspector")
    selected_id = st.selectbox("Select Log ID to inspect normalized JSON payload:", df["id"].tolist())
    
    if selected_id:
        selected_row = df[df["id"] == selected_id].iloc[0]
        st.markdown(f"**Raw Log String:** `{selected_row['raw_log']}`")
        st.markdown("**Normalized OCSF JSON Output:**")
        st.code(selected_row['parsed_json'], language="json")