import sys

from PyQt5 import QtGui
from PyQt5.QtCore import QUrl
from PyQt5.QtGui import QFont
from PyQt5.QtWebEngineWidgets import QWebEngineView
from PyQt5.QtWidgets import *

from config import FLASK_PORT

import local_server

if __name__ == "__main__":
    try:
        f = open("debug.log", 'w')
        sys.stdout = f
        sys.stderr = sys.stdout
    except Exception:
        pass

    local_server.start()

    app = QApplication(sys.argv)
    app.setFont(QFont("YaHei", 25))

    browser = QWebEngineView()
    browser.setWindowIcon(QtGui.QIcon('favicon.ico'))
    browser.setWindowTitle("Easy Testing")
    browser.load(QUrl(f"http://localhost:{FLASK_PORT}"))
    browser.setZoomFactor(1.8)
    browser.showMaximized()

    code = app.exec_()
    sys.exit(code)
