import sys

from PyQt5.QtCore import QUrl
from PyQt5.QtGui import QFont
from PyQt5.QtWebEngineWidgets import QWebEngineView
from PyQt5.QtWidgets import *

import local_server

if __name__ == "__main__":
    local_server.start()

    app = QApplication(sys.argv)
    app.setFont(QFont("YaHei", 25))

    browser = QWebEngineView()
    browser.load(QUrl("http://localhost:2998"))
    browser.showMaximized()

    code = app.exec_()
    sys.exit(code)
