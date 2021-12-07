from datetime import datetime

from PyQt5.QtWidgets import *

_STATUS_BAR: QStatusBar = None
_CNT = 0

def success(text: str):
    _show_message(text, "green")


def failure(text: str):
    _show_message(text, "red")


def _show_message(msg: str, color):
    global _CNT
    _CNT += 1
    cnt_str = f"({_CNT}) "
    time_str = f"[{datetime.now()}]  "
    m = {"red": ("gainsboro", "darkred"),
         "green": ("snow", "yellowgreen")
         }[color]
    _STATUS_BAR.setStyleSheet("QStatusBar{background:" + m[0] + ";color:" + m[1] + ";font-weight:bold;}")
    _STATUS_BAR.showMessage(cnt_str+time_str + msg)
