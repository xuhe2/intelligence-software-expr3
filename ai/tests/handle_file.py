import requests as req

from base import print_test_result

url:str = "http://localhost:3000/handle_picture"

# 发送图片
file_name:str = "0.png"
file_path:str = "./tests/pictures/" + file_name

with open(file_path, "rb") as file:
    files = {"file": file}
    response = req.post(url, files=files)
    if response.json()['filename'] == file_name:
        print_test_result(True)
    else:
        print_test_result(False, "filename is not correct")