PYTHON=python3
SERVER_FILE=server

.PHONY: server

server: $(SERVER_FILE)
	uvicorn $(SERVER_FILE):app --reload --host 0.0.0.0 --port 3000