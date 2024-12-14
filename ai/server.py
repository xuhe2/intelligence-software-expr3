from typing import Union

from fastapi import FastAPI
from fastapi import File, UploadFile

app = FastAPI()


@app.get("/")
def root():
    return {"Hello": "World"}

@app.post("/handle_picture/")
async def handle_picture(file: UploadFile = File(...)):
    # Do something with the file
    return {"filename": file.filename}