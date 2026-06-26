import requests
from flask import abort

SERVER_URL = "http://127.0.0.1:8000" # DO NOT FORGET TO UPDATE GUYS


def fetch(path: str): 
    try:
        r = requests.get(f"{SERVER_URL}{path}", timeout=10)
        r.raise_for_status()
        return r.json()
    except requests.exceptions.HTTPError as e:
        abort(r.status_code)
    except requests.exceptions.ConnectionError:
        abort(503)


