@echo off
TITLE AutoGo Helper System
COLOR 03
:MENU
ECHO.
ECHO ###############################################################
ECHO #                 AutoGo Helper: AUTO UPDATER 1.0             #
ECHO #                      WRITTEN BY: PROXY                      #
ECHO ###############################################################
ECHO.
cd ../
updater.exe
TIMEOUT 3
helper.exe -upd1
