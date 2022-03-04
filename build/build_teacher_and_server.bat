@ECHO OFF
TITLE Build Student Client
SET client_src_path=%~dp0
FOR %%a IN ("%client_src_path:~0,-1%") DO SET root_path=%%~dpa
SET client_src_path=%root_path%\teacher_client
SET server_src_path=%root_path%\server

SET dist_path="%root_path%\easy_testing_server_windows"
:: Delete Older Dist Folder
IF EXIST %dist_path% RMDIR /S /Q %dist_path%

CD "%client_src_path%"
ECHO ::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
ECHO NPM Building %client_src_path%
ECHO ::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
CALL npm run build

ECHO ::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
ECHO Collecting Dist Of Vue
ECHO ::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
:: Copy Folder
ROBOCOPY /E "%client_src_path%\dist\static" "%dist_path%\static"
:: Copy With Overwriting
ROBOCOPY /IS /IT /IM "%client_src_path%\dist" "%dist_path%" "index.html" "favicon.ico"

CD "%server_src_path%"
ECHO ::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
ECHO Go Building %server_src_path%
ECHO ::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
CALL go build

ECHO ::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
ECHO Collecting Dist Of Server
ECHO ::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
:: Copy Folder
ROBOCOPY /E "%server_src_path%\sql" "%dist_path%\sql"
ROBOCOPY "%server_src_path%" "%dist_path%" "EasyTesting.exe" "server-config.yaml"

ECHO ::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
ECHO %dist_path% IS READY FOR DEPLOYMENT
ECHO ::::::::::::::::::::::::::::::::::::::::::::::::::::::::::

PAUSE