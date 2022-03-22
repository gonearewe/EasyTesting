@ECHO OFF
TITLE Build Student Client
SET script_path=%~dp0
FOR %%a IN ("%script_path:~0,-1%") DO SET root_path=%%~dpa
SET client_src_path=%root_path%\student_client

CD "%client_src_path%\vue"
ECHO ::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
ECHO NPM Building %client_src_path%\vue
ECHO ::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
CALL npm run build

ECHO ::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
ECHO Copying Dist Of Vue To PyQt
ECHO ::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
:: Copy Folder
ROBOCOPY /E "%client_src_path%\vue\dist\static" "%client_src_path%\pyqt\static"
:: Copy With Overwriting
ROBOCOPY /IS /IT /IM "%client_src_path%\vue\dist" "%client_src_path%\pyqt" "index.html" "favicon.ico"

CD "%client_src_path%\pyqt"
ECHO ::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
ECHO PyInstaller Packaging %client_src_path%\pyqt
ECHO ::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
CALL pipenv run pyinstaller main.spec

ECHO ::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
ECHO Cleaning
ECHO ::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
RMDIR /S /Q "%client_src_path%\pyqt\static"
DEL         "%client_src_path%\pyqt\index.html"
DEL         "%client_src_path%\pyqt\favicon.ico"

SET dist_path="%root_path%\easy_testing_student_client_windows"
:: Delete Older Dist Folder
IF EXIST %dist_path% RMDIR /S /Q %dist_path%
:: Rename And Move Dist To Root Path
MOVE /Y "%client_src_path%\pyqt\dist\main" %dist_path%
:: Copy README.md And LICENSE
ROBOCOPY %root_path% %dist_path% README.md LICENSE
ECHO ::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
ECHO %dist_path% IS READY FOR DEPLOYMENT
ECHO ::::::::::::::::::::::::::::::::::::::::::::::::::::::::::

PAUSE