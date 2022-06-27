@ECHO OFF
TITLE Build Student Client
SET client_src_path=%~dp0
FOR %%a IN ("%client_src_path:~0,-1%") DO SET root_path=%%~dpa
SET client_src_path=%root_path%\teacher_client
SET server_src_path=%root_path%\server

SET dist_path_windows="%root_path%\easy_testing_server_windows"
SET dist_path_linux="%root_path%\easy_testing_server_linux"
:: Delete Older Dist Folder
IF EXIST %dist_path_windows% RMDIR /S /Q %dist_path_windows%
IF EXIST %dist_path_linux% RMDIR /S /Q %dist_path_linux%

CD "%client_src_path%"
ECHO ::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
ECHO NPM Building %client_src_path%
ECHO ::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
CALL npm run build

ECHO ::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
ECHO Collecting Dist Of Vue
ECHO ::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
:: Copy Folder
ROBOCOPY /E "%client_src_path%\dist\static" "%dist_path_windows%\static"
ROBOCOPY /E "%client_src_path%\dist\static" "%dist_path_linux%\static"
:: Copy With Overwriting
ROBOCOPY /IS /IT /IM "%client_src_path%\dist" "%dist_path_windows%" "index.html" "favicon.ico"
ROBOCOPY /IS /IT /IM "%client_src_path%\dist" "%dist_path_linux%" "index.html" "favicon.ico"

CD "%server_src_path%"
ECHO ::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
ECHO Go Building %server_src_path%
ECHO ::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
CALL go env -w GOOS=linux
CALL go build
CALL go env -w GOOS=windows
CALL go build

ECHO ::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
ECHO Collecting Dist Of Server
ECHO ::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
:: Copy Folder
ROBOCOPY /E "%server_src_path%\sql" "%dist_path_windows%\sql"
ROBOCOPY /E "%server_src_path%\sql" "%dist_path_linux%\sql"
ROBOCOPY "%server_src_path%" "%dist_path_windows%" "EasyTesting.exe" "server-config.yaml"
ROBOCOPY "%server_src_path%" "%dist_path_linux%" "EasyTesting" "server-config.yaml"

:: Copy README.md And LICENSE
ROBOCOPY %root_path% %dist_path_windows% README.md LICENSE
ROBOCOPY %root_path% %dist_path_linux% README.md LICENSE
ECHO ::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
ECHO %dist_path_windows% AND %dist_path_linux% IS READY FOR DEPLOYMENT
ECHO ::::::::::::::::::::::::::::::::::::::::::::::::::::::::::

PAUSE