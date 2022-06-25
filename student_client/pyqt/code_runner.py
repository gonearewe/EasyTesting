import os
import shutil
import subprocess
import sys
import tempfile

from PyQt5.QtCore import QObject
from config import PYTHON_RUNNER_EXE_PATH


class CodeRunner(QObject):
    # run_code prepares the test environment based on the first three parameters,
    # and runs the provided Python code. It returns the console output and the test output.
    # Depending on `is_output_to_file`, the test output is either from the output file or from the console,
    # and in the second case, the test output equals to the console output (but with leading and tailing spaces
    # stripped) for sure.
    # However, when the code causes Error (Timeout, etc.), it's possible for the test output to be empty.
    def run_code(self, is_input_from_file: bool, is_output_to_file: bool, input: str, code: str) -> (str, str):
        dir_path = tempfile.mkdtemp()

        if is_input_from_file:
            with open(os.path.join(dir_path, 'input.txt'), 'w') as f:
                f.write(input.strip())
        std_in = None if is_input_from_file else input

        error_header = "\n Error Running Your Code: "
        result = None
        # The -c component is a python command line option that allows passing a string to execute.
        # We also change the working directory
        # Max execution time is 3 seconds
        # Stdout and stderr will be captured into result, and we combine both streams into one.
        # A possible stdin is supplied
        try:
            result = subprocess.run([PYTHON_RUNNER_EXE_PATH, "-c", code], cwd=dir_path, timeout=3,
                                    stdout=subprocess.PIPE, stderr=subprocess.STDOUT, input=std_in, text=True)
        except subprocess.TimeoutExpired:
            return error_header + "timeout, check if there's any endless loop or recursion", ''

        # collect test output from the file or stdout
        output = ''
        if is_output_to_file:
            try:
                with open(os.path.join(dir_path, 'output.txt'), 'r') as f:
                    output = f.read().strip()
            except Exception:
                # the `code` didn't do it correctly, for example, didn't create file
                return result.stdout + error_header + "you didn't create output file properly", ''
        else:
            output = result.stdout.strip()

        # Clean up
        shutil.rmtree(dir_path)

        return result.stdout, output

# 用于测试。使用时可能要把 PYTHON_RUNNER_EXE_PATH 改成实际的 Python 执行程序路径；
# 或者它在 PATH 中的话，直接改成 "python"。
if __name__ == "__main__":
    obj = CodeRunner()

    std_out, output = obj.run_code(is_input_from_file=False, is_output_to_file=False, input="Jack",
                               code="""
print('Hello '+input())
    """)
    # print(std_out, output)
    if std_out.strip() == output == "Hello Jack":
        print("PASS")
    else:
        print('FAIL')

    std_out, output = obj.run_code(is_input_from_file=False, is_output_to_file=False,
                               input="1 4 5 6 0 -2 -8 25 1000000000000000", code="""
li = [int(x) for x in input().split()]
print(sum(li))
    """)
    # print(std_out, output)
    if std_out.strip() == output == "1000000000000031":
        print("PASS")
    else:
        print('FAIL')

    # test output 为文件中的结果时，stdout 不影响它
    std_out, output = obj.run_code(is_input_from_file=True, is_output_to_file=True,
                               input="1 4 5 6 0 -2 -8 25 1000000000000000", code="""
import os
with open(os.path.join(os.getcwd(), 'input.txt'), 'r') as f:
    li = [int(x) for x in f.read().split()]
    print(li) # DEBUG
    with open(os.path.join(os.getcwd(), 'output.txt'), 'w') as f:
        f.write(str(sum(li)))
    """)
    if std_out.strip() != output and output == "1000000000000031":
        print("PASS")
    else:
        print('FAIL')

    # 超时时不管结果正确与否，test output 一律为空
    std_out, output = obj.run_code(is_input_from_file=True, is_output_to_file=False,
                               input="Red Black White", code="""
import time
time.sleep(5) # Sleep for 5 seconds
print("a correct result but after lots of time")
    """)
    if "Error" in std_out and output == "":
        print("PASS")
    else:
        print('FAIL')
