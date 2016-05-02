@echo off
TITLE AutoGo Helper System
COLOR 03
:MENU
ECHO.
ECHO ###############################################################
ECHO #                 AutoGo Helper System 1.5.2                  #
ECHO #                      WRITTEN BY: PROXY                      #
ECHO #-------------------------------------------------------------#
ECHO #              1. Update AutoGo Files                         #
ECHO #              2. Make A Bot Application                      #
ECHO #              3. Get Your Bot's Invite link                  #
ECHO #              4. Initiate AutoGo (Might have to pm him)      #
ECHO #              5. Create Bot Commander and muted Role         #
ECHO #              6. Howto: post msg every x hours               #
ECHO #              7. Rename a Role.                              #
ECHO #              8. Changelog for 1.8.5                         #
ECHO ###############################################################
ECHO.

SET /P M=What would you like me to do?:
IF %M%==1 GOTO runupdate
IF %M%==2 GOTO botapp
IF %M%==3 GOTO botinv
IF %M%==4 GOTO initbot
IF %M%==5 GOTO botrole
IF %M%==6 GOTO botinfo
IF %M%==7 GOTO rename
IF %M%==8 GOTO changelog

:runupdate
SET /P M=Are you on 64bit? (yes\no):
IF %M%==yes GOTO upd64
IF %M%==no GOTO upd32

:upd64
cd System/helper
updater.exe
TIMEOUT 3
helper.exe -upd1
GOTO MENU

:upd32
cd System/helper
updater32.exe
TIMEOUT 3
helper32.exe -upd1
GOTO MENU

:botapp
SET /P M=Are you on 64bit? (yes\no):
IF %M%==yes GOTO botapp64
IF %M%==no GOTO botapp32

:botapp64
cd System/helper
helper.exe -m
GOTO MENU

:botapp32
cd System/helper
helper32.exe -m
GOTO MENU

:botinv
SET /P M=Are you on 64bit? (yes\no):
IF %M%==yes GOTO botinv64
IF %M%==no GOTO botinv32

:botinv64
SET /P M=Ok In the developer website you have a CLIENT ID Type here:
cd System/helper
helper.exe -i %m%
GOTO MENU

:botinv32
SET /P M=Ok In the developer website you have a CLIENT ID Type here:
cd System/helper
helper32.exe -i %m%
GOTO MENU

:initbot
SET /P M=Are you on 64bit? (yes\no):
IF %M%==yes GOTO initbot64
IF %M%==no GOTO initbot32

:initbot64
cd System/helper
helper.exe -o
GOTO MENU

:initbot32
cd System/helper
helper32.exe -o
GOTO MENU

:botrole
SET /P M=Are you on 64bit? (yes\no):
IF %M%==yes GOTO botrole64
IF %M%==no GOTO botrole32

:botrole64
cd System/helper
helper.exe -b
GOTO MENU

:botrole32
cd System/helper
helper32.exe -b
GOTO MENU

:botinfo
ECHO.
ECHO AutoGo's helper can post text every 1 to 10 hours. OR once every day!
ECHO.
ECHO browse to System/custom folder and edit notice.txt
ECHO.
ECHO cd to /System/helper folder and run helper.exe -a general -h 1
ECHO --------------------------------------
ECHO The above code is for 64bit users. this will loop the notice.txt once an hour in your main room
ECHO this room doesn't have to be named general. general is just how we determine to use your main room.
ECHO /helper32.exe -a general -h 1 for 32bit users!
ECHO --------------------------------------
ECHO You can also specify a channel by it's id: helper.exe -a CHANNELID -h 5
ECHO.
ECHO Looping every 24 hours: helper.exe -a general -h 24
ECHO.
ECHO Accepted -h(hours): 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 24
SET /P M=Would you like to go back to the main menu? (yes\no):
IF %M%==yes GOTO MENU
IF %M%==no GOTO exit

:rename
cd System/helper/tools
renamerole.bat

:changelog
ECHO.
ECHO ===================================
ECHO - AutoGo Change log: 1.8.5
ECHO - Monday May 02, 2016 12:00 AM(EST)
ECHO - Suggested by: Dog2Puppy
ECHO - ---------------------------------
ECHO - Fixed the meme and joke command.
ECHO - Added the {del} key to autoresponse system.
ECHO - Took away rolecolor from helper.exe due to abuse.
ECHO - Will add back when i find a way to limit the system.
ECHO - 
ECHO - Added a way to Enable\Disable commands.
ECHO - 
ECHO - Huge updates for updater.exe, helper.exe
ECHO - 
ECHO - config.json Added PC and CommanderRole
ECHO - 
ECHO - PC: will help autogo figure out if you're on 32bit or 64bit.
ECHO - Set this accordingly for proper updates and functions.
ECHO - 
ECHO - CommanderRole: let's you choose the name of your Bot Commander role.
ECHO - 
ECHO - Will be adding a lot more keys for the auto response system.
ECHO - Fix Credits: EdibleDerpy, Doorstop, Peter, Dog2Puppy, SkyWielder
ECHO - Tore, scattrbrain
ECHO ===================================
ECHO.
SET /P M=Would you like to go back to the main menu? (yes\no):
IF %M%==yes GOTO MENU
IF %M%==no GOTO exit

:exit
CLS
ECHO.
ECHO ===================================
ECHO.
ECHO Exiting, have a wonderful day!
ECHO.
ECHO Goodbye!
ECHO.
ECHO ===================================
ECHO.
TIMEOUT 2
CLS