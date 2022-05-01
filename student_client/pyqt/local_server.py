import os.path
import threading

from flask import Flask

from code_runner import CodeRunner
from config import INDEX_HTML_PATH, FLASK_PORT

server = Flask(__name__)


def start():
    thr = threading.Thread(target=server.run, kwargs={'port': FLASK_PORT})
    thr.daemon = True  # local server will exit automatically after main thread exits
    thr.start()


@server.route('/code', methods=['PUT'])
def run():
    from flask import request
    res, ok = CodeRunner().run_code(request.json["is_input_from_file"], request.json["is_output_to_file"],
                                    request.json["input"], request.json["output"], request.json["code"])
    return {'console_output': res, 'pass': ok}


@server.route('/', methods=['GET'])
def index():
    with open(INDEX_HTML_PATH) as f:
        return f.read()


if __name__ == '__main__':
    server.run(port=2998)
