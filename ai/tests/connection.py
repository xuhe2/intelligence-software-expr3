import http.client

from base import print_test_result

def get_status_code(url):
    try:
        # 解析 URL
        conn = http.client.HTTPConnection("localhost", 3000)
        conn.request("GET", "/")
        response = conn.getresponse()
        status_code = response.status
        conn.close()
    except Exception as e:
        status_code = None
    return status_code

# 调用函数并打印状态码
url = "http://localhost:3000"
status_code = get_status_code(url)

if status_code == 200:
    print_test_result(True, "The server is running.")
else:
    print_test_result(False, f"The server is not running. Status code: {status_code}")
