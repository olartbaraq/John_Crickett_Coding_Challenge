import json, glob
import os
import sys
from types import NoneType

filenames = glob.glob(os.path.join("tests/step3", "*.json"))
exit_code = 0

for filename in filenames:
    with open(filename, "r") as f:
        content = f.read()

        try:
            parsed_value = json.loads(content)
            dict_length = len(parsed_value)
            count = 0
            for keys, values in parsed_value.items():
                if isinstance(keys, str) and isinstance(
                    values, (bool, int, str, NoneType)
                ):
                    count += 1
            if dict_length == count:
                print(f"{filename} is a json")
            else:
                print(f"{filename} is not a json")
                exit_code = 2
        except json.JSONDecodeError:
            print(f"{filename} is not a json")
            exit_code = 2
sys.exit(exit_code)
