from typing import Callable

from PyQt5 import QtCore
from PyQt5.QtGui import *
from PyQt5.QtWidgets import *


class CardWidget(QWidget):
    def __init__(self, parent, image: str, text: str, on_click: Callable):
        super().__init__(parent)
        layout = QVBoxLayout(self)
        layout.setSpacing(20)
        label = QLabel(text)
        label.setAlignment(QtCore.Qt.AlignCenter)
        layout.addWidget(label)
        img = QLabel(self)
        pixmap = QPixmap(image)
        pixmap.scaled(640, 640)
        img.setPixmap(pixmap)
        layout.addWidget(img)
        btn = QPushButton("查看")
        btn.clicked.connect(on_click)
        layout.addWidget(btn)
