import os
import shutil
import subprocess
import sys
import tempfile


def run_test(is_input_from_file: bool, is_output_to_file: bool, input: str, output: str, code: str) -> (str, bool):
    dir_path = tempfile.mkdtemp()

    if is_input_from_file:
        with open(os.path.join(dir_path, 'input.txt'), 'w') as f:
            f.write(input.strip())
    std_in = None if is_input_from_file else input

    error_header = "\n Error Running Your Code: "
    result = None
    # sys.executable is the absolute path to the Python executable that your program was originally invoked with,
    # for example, sys.executable might be a path like /usr/local/bin/python
    # The -c component is a python command line option that allows passing a string to execute.
    # We also change the working directory
    # Max execution time is 3 seconds
    # Stdout and stderr will be captured into result, and we combine both streams into one.
    # A possible stdin is supplied
    try:
        result = subprocess.run([sys.executable, "-c", code], cwd=dir_path, timeout=3,
                                stdout=subprocess.PIPE, stderr=subprocess.STDOUT, input=std_in, text=True)
    except subprocess.TimeoutExpired:
        return error_header + "timeout, check if there's any endless loop or recursion", False

    # is answer right? a.k.a. does `code` passed the test
    right = False
    if is_output_to_file:
        try:
            with open(os.path.join(dir_path, 'output.txt'), 'r') as f:
                right = f.read().strip() == output.strip()
        except Exception:
            # the `code` didn't do it correctly, for example, didn't create file
            return result.stdout + error_header + "you didn't create output file properly", False
    else:
        right = result.stdout.strip() == output.strip()

    # Clean up
    shutil.rmtree(dir_path)

    return result.stdout, right


if __name__ == "__main__":
    std_out, ok = run_test(is_input_from_file=False, is_output_to_file=False, input="Jack", output="Hello Jack", code="""
print('Hello '+input())
    """)
    print(std_out, ok)

    std_out, ok = run_test(is_input_from_file=False, is_output_to_file=False,
                           input="1 4 5 6 0 -2 -8 25 1000000000000000", output="1000000000000031", code="""
li = [int(x) for x in input().split()]
print(sum(li))
    """)
    print(std_out, ok)

    std_out, ok = run_test(is_input_from_file=True, is_output_to_file=True,
                           input="1 4 5 6 0 -2 -8 25 1000000000000000", output="1000000000000031", code="""
import os
with open(os.path.join(os.getcwd(), 'input.txt'), 'r') as f:
    li = [int(x) for x in f.read().split()]
    print(li) # DEBUG
    with open(os.path.join(os.getcwd(), 'output.txt'), 'w') as f:
        f.write(str(sum(li)))
    """)
    print(std_out, ok)

    std_out, ok = run_test(is_input_from_file=True, is_output_to_file=False,
                           input="Red Black White", output="", code="""
import time
time.sleep(5) # Sleep for 5 seconds
    """)
    print(std_out, ok)
