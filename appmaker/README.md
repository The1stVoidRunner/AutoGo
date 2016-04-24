# bwmarrin's Bot Account Maker
This was entirely built by **bwmarrin** i simply compiled for use with AutoGo
<br>
<br>
### How to use:
This tutorial was copied from bwmarrin's github<br>
<br>


### Usage:
<pre>
Usage of ./appmaker:
  -a string
        App/Bot Name
  -c string
        Token of account to convert.
  -d string
        Application ID to delete
  -e string
        Account Email
  -l    List Applications Only
  -p string
        Account Password
  -t string
        Account Token
        </pre><br>
        <br>


Account Email and Password or Token are required. The account provided with these fields will be the "owner" of any bot applications created.
<br>
If you provide the ```-l``` flag than appmaker will only display a list of applications on the provided account.
<br>
If you provide a ```-d``` flag with a valid application ID then that application will be deleted.
<br>
If you provide a ```-c``` flag with a valid user token then than user account will be converted into a Bot account instead of creating a new Bot account for an application.
<Br><br>
Below example will create a new Bot Application under the given Email/Password account. The Bot will be named DiscordGoRocks
<Br>
<pre>
./appmaker -e Email -p Password -a DiscordGoRocks
</pre>
