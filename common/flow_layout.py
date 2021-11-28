from PyQt5.QtCore import *
from PyQt5.QtGui import *
from PyQt5.QtWidgets import *


class FlowLayout(QLayout):
    def __init__(self, orientation=Qt.Horizontal, parent=None, margin=0, spacing=-1):
        super().__init__(parent)
        self.orientation = orientation

        if parent is not None:
            self.setContentsMargins(margin, margin, margin, margin)

        self.setSpacing(spacing)

        self.itemList = []

    def __del__(self):
        item = self.takeAt(0)
        while item:
            item = self.takeAt(0)

    def addItem(self, item):
        self.itemList.append(item)

    def count(self):
        return len(self.itemList)

    def itemAt(self, index):
        if index >= 0 and index < len(self.itemList):
            return self.itemList[index]

        return None

    def takeAt(self, index):
        if index >= 0 and index < len(self.itemList):
            return self.itemList.pop(index)

        return None

    def expandingDirections(self):
        return Qt.Orientations(Qt.Orientation(0))

    def hasHeightForWidth(self):
        return self.orientation == Qt.Horizontal

    def heightForWidth(self, width):
        return self.doLayout(QRect(0, 0, width, 0), True)

    def hasWidthForHeight(self):
        return self.orientation == Qt.Vertical

    def widthForHeight(self, height):
        return self.doLayout(QRect(0, 0, 0, height), True)

    def setGeometry(self, rect):
        super().setGeometry(rect)
        self.doLayout(rect, False)

    def sizeHint(self):
        return self.minimumSize()

    def minimumSize(self):
        size = QSize()

        for item in self.itemList:
            size = size.expandedTo(item.minimumSize())

        margin, _, _, _ = self.getContentsMargins()

        size += QSize(2 * margin, 2 * margin)
        return size

    def doLayout(self, rect, testOnly):
        x = rect.x()
        y = rect.y()
        lineHeight = columnWidth = heightForWidth = 0

        for item in self.itemList:
            wid = item.widget()
            spaceX = self.spacing() + wid.style().layoutSpacing(QSizePolicy.PushButton, QSizePolicy.PushButton,
                                                                Qt.Horizontal)
            spaceY = self.spacing() + wid.style().layoutSpacing(QSizePolicy.PushButton, QSizePolicy.PushButton,
                                                                Qt.Vertical)
            if self.orientation == Qt.Horizontal:
                nextX = x + item.sizeHint().width() + spaceX
                if nextX - spaceX > rect.right() and lineHeight > 0:
                    x = rect.x()
                    y = y + lineHeight + spaceY
                    nextX = x + item.sizeHint().width() + spaceX
                    lineHeight = 0

                if not testOnly:
                    item.setGeometry(QRect(QPoint(x, y), item.sizeHint()))

                x = nextX
                lineHeight = max(lineHeight, item.sizeHint().height())
            else:
                nextY = y + item.sizeHint().height() + spaceY
                if nextY - spaceY > rect.bottom() and columnWidth > 0:
                    x = x + columnWidth + spaceX
                    y = rect.y()
                    nextY = y + item.sizeHint().height() + spaceY
                    columnWidth = 0

                heightForWidth += item.sizeHint().height() + spaceY
                if not testOnly:
                    item.setGeometry(QRect(QPoint(x, y), item.sizeHint()))

                y = nextY
                columnWidth = max(columnWidth, item.sizeHint().width())

        if self.orientation == Qt.Horizontal:
            return y + lineHeight - rect.y()
        else:
            return heightForWidth - rect.y()
