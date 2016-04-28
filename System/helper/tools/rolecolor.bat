@echo off
TITLE AutoGo Helper System
COLOR 03
:MENU
ECHO.
ECHO ###############################################################
ECHO #                 AutoGo Helper: ROLECOLOR EDITOR             #
ECHO #                      WRITTEN BY: PROXY                      #
ECHO ###############################################################
ECHO.

SET /P R=What's the name of the role?:
SET /P C=What color? example: #000000 type now:
cd ../
helper.exe -role "%R%" -color %C%


:error1
ECHO.
ECHO [ERROR] You need to type a valid role name.
ECHO.

:error
ECHO.
ECHO [ERROR] You need to type a valid color.
ECHO.