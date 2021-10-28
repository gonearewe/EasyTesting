import sys

from PyQt5 import QtWidgets
from PyQt5.QtWidgets import QMainWindow
from pyqode.python.backend import server
from pyqode.python.widgets import PyCodeEdit, code_edit

if __name__ == '__main__':
    app = QtWidgets.QApplication(sys.argv)
    window = QMainWindow()
    editor = PyCodeEdit(server_script=server.__file__)
    # show the PyCodeEdit module in the editor
    editor.file.open(code_edit.__file__.replace('.pyc', '.py'))
    window.setCentralWidget(editor)
    window.showFullScreen()
    app.exec_()
