import sys, os
import configparser

# from https://stackoverflow.com/questions/404744/determining-application-path-in-a-python-exe-generated-by-pyinstaller

APP_PATH = ''
# the PyInstaller bootloader extends the sys module by a flag frozen=True
if getattr(sys, 'frozen', False):
    # If the application is run as a bundle, the PyInstaller bootloader also
    # sets the app path into variable '_MEIPASS'.
    APP_PATH = sys._MEIPASS
    # for a --onefile executable, the path to the application is given by 
    # APP_PATH = os.path.dirname(sys.executable)
else:
    APP_PATH = os.path.dirname(os.path.abspath(__file__))

INDEX_HTML_PATH = os.path.join(APP_PATH,'index.html')

__exe_name = 'python.exe' if os.name == 'nt' else 'python'
PYTHON_RUNNER_EXE_PATH = os.path.join(APP_PATH,'runner',__exe_name)

_config = configparser.ConfigParser()
_config.read("config.ini")
FLASK_PORT = int(_config["flask"]["port"])