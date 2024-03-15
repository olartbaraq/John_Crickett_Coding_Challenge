import json, glob
import os
import sys

filenames = glob.glob(os.path.join("tests/step1", "*.json"))
exit_code = 0
for filename in filenames:
    with open(filename, "r") as f:
        content = f.read()

        try:
            parsed_value = json.loads(content)
            if isinstance(parsed_value, dict):
                print(f"{filename} is a json")
            else:
                print(f"{filename} is not a json")
                exit_code = 2
        except json.JSONDecodeError:
            print(f"{filename} is not a json")
            exit_code = 2
sys.exit(exit_code)
