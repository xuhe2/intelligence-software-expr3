PYTHON=python3
SERVER_FILE=server
TEST_FILE=test.py

.PHONY: server test

server: $(SERVER_FILE)
	uvicorn $(SERVER_FILE):app --reload --host 0.0.0.0 --port 3000

test: $(TEST_FILE)
	$(PYTHON) $(TEST_FILE)

clean:
	rm -rf *.pyc *.pyo __pycache__

