import sys
from pyqode.qt import QtWidgets
from pyqode.python.backend import server
from pyqode.python.widgets import PyCodeEdit
from pyqode.python.widgets import code_edit

if __name__ == '__main__':
    app = QtWidgets.QApplication(sys.argv)
    window = QtWidgets.QMainWindow()
    editor = PyCodeEdit(server_script=server.__file__)
    # show the PyCodeEdit module in the editor
    editor.file.open(code_edit.__file__.replace('.pyc', '.py'))
    window.setCentralWidget(editor)
    window.resize(1024, 720)
    window.show()
    app.exec_()
