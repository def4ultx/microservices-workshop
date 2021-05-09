from fastapi import FastAPI

app = FastAPI()

@app.get("/")
async def root():
    return {"message": "Hello World"}

@app.post("/user/register")
async def register():
    return {"message": "OK"}

@app.post("/user/login")
async def login():
    return {"token": "dj0yJmk9N2pIazlsZk1iTzIx"}

@app.get("/user/profile")
async def view_profile():
    profile = {
        "email": "john@futureskill.com",
        "name": "John"
    }
    return profile

@app.put("/user/profile")
async def update_profile():
    profile = {
        "email": "john@futureskill.com",
        "name": "Richard"
    }
    return profile