import os.path
import threading

from flask import Flask

# set the project root directory as the static folder, you can set others.

server = Flask(__name__)


def start():
    thr = threading.Thread(target=server.run, kwargs={'port': 2998})
    thr.start()


# @server.route('/code', methods=['PUT'])
# def run():
#     from flask import request
#     res, ok = CodeRunner().run_code(request.json["is"], False, )
#     return


@server.route('/', methods=['GET'])
def index():
    with open(os.path.join(os.getcwd(), 'index.html')) as f:
        return f.read()


if __name__ == '__main__':
    server.run(port=2998)
