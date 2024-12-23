from typing import Union
import os

from fastapi import FastAPI
from fastapi import File, UploadFile

from model import predict_img

app = FastAPI()

# cors
from fastapi.middleware.cors import CORSMiddleware
app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

@app.get("/")
def root():
    return {"Hello": "World"}

@app.post("/handle_picture/")
async def handle_picture(file: UploadFile = File(...)):
    # Do something with the file
    # 使用/tmp的历时文件保存
    with open(f"/tmp/{file.filename}", "wb") as buffer:
        buffer.write(file.file.read())
    
    kind = predict_img(f"/tmp/{file.filename}")
    kind = int(kind)

    # 删除临时文件
    os.remove(f"/tmp/{file.filename}")
    
    return {"filename": file.filename, "kind": kind}