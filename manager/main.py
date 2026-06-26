from flask import Flask, render_template
from utils import fetch

app = Flask(__name__)

@app.route("/")
def home():
    # Return all agent
    data = fetch("/agents")
    return render_template("home.html", agents=data["agents"])

@app.route("/api/agents")
def api_agents():
    data = fetch("/agents")
    return data

@app.route("/result/<uuid>")
def get_result(uuid: str):  # ← nom identique au url_for() du template
    data = fetch(f"/results/{uuid}")
    return render_template("results.html", agent_id=uuid, results=data)



if __name__ == "__main__":
    app.run(debug=True)