import os

tests_dir:str = "./tests"

for test_file in os.listdir(tests_dir):
    if test_file.endswith(".py") and test_file != "base.py":
        print(f"Running test: {test_file}")
        os.system(f"python {tests_dir}/{test_file}")