import sys

from PyQt5.QtCore import Qt, QUrl
from PyQt5.QtGui import QFont
from PyQt5.QtWebChannel import QWebChannel
from PyQt5.QtWebEngineWidgets import QWebEngineView
from PyQt5.QtWidgets import *

import local_server
from code_runner import CodeRunner

if __name__ == "__main__":
    local_server.start()

    app = QApplication(sys.argv)
    app.setFont(QFont("YaHei", 25))

    browser = QWebEngineView()
    browser.load(QUrl("http://localhost:2998"))
    browser.show()
    channel = QWebChannel()
    channel.registerObject('code_runner', CodeRunner())
    browser.page().setWebChannel(channel)

    code = app.exec_()
    sys.exit(code)
