# editor.py

from PyQt5 import QtWidgets
import syntax

if __name__ == '__main__':
    app = QtWidgets.QApplication([])
    editor = QtWidgets.QPlainTextEdit()
    editor.setStyleSheet("""QPlainTextEdit{
	font-family:'Consolas'; 
	color: #ccc; 
	background-color: #2b2b2b;}""")
    highlight = syntax.PythonHighlighter(editor.document())
    editor.resize(1024, 720)
    editor.show()

    # Load syntax.py into the editor for demo purposes
    infile = open('syntax.py', 'r')
    editor.setPlainText(infile.read())

    app.exec_()
