@echo off
TITLE AutoGo Updater 64bit
COLOR 03
:MENU
taskkill /I autogo.exe
ECHO.
ECHO ###############################################################
ECHO #                 AutoGo Helper: AUTO UPDATER 1.0             #
ECHO #                      WRITTEN BY: PROXY                      #
ECHO #          You can view what files are being updated          #
ECHO #    Check out the updates.auf file for specific details!     #
ECHO ###############################################################
ECHO.
cd System\helper
updater.exe
TIMEOUT 3
helper.exe -upd1
cd ..
cd ..
TIMEOUT 2
autogo.exe
