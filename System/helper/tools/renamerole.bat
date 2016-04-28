@echo off
TITLE AutoGo Helper System
COLOR 03
:MENU
ECHO.
ECHO ###############################################################
ECHO #                 AutoGo Helper: ROLE RENAME TOOL             #
ECHO #                      WRITTEN BY: PROXY                      #
ECHO ###############################################################
ECHO.

SET /P R=What's the name of the role?:
SET /P C=Now what do you want the new name to be?:
cd ../
helper.exe -r1 "%R%" -r2 "%C%"


:error1
ECHO.
ECHO [ERROR] You need to type a valid role name.
ECHO.

:error
ECHO.
ECHO [ERROR] You need to type a valid color.
ECHO.