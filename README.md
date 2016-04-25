## You need to pm the bot to gain ownership.
> Utilizes the DiscordGO Library: https://github.com/bwmarrin/discordgo<Br>
> AutoGo Written by: Proxy<br>
> Appmaker written by: bwmarrin<br>

### NEWLY ADDED
> Custom Responses in **responses.json**<br>
> Added **{kick}** key in Auto Response System<br>
> Added **{ban}** key in Auto Response System<br>

<br><br>
First You need to make a bot account by visiting the link below<br>
https://discordapp.com/developers/applications/me<br>
<br>
Or you can use bwmarrin's ```appmaker``` to convert user account to bot account.<br>
You can also use ```appmaker``` to list/delete/view and add new OAuth applications.<br>

<Br>

### OPEN CONFIG.JSON AND FOLLOW INSTRUCTIONS BELOW:

### BOT TOKEN
Once you have completed your application And have recieved your new Token<br>
Replace "YOUR-BOT-TOKEN-HERE" with your bot's token.<br>

<Br>

### STATUS
This allows you to set his startup status.<br>
You can later change this using --status command.<br>

<br>

### PREFIX
The default prefix is -- You can change this manually or via command --prefix<br>
This is how AutoGo can tell if it's a command or not. Examples: +, -, #, !<br>

<br>

### GREETMSG
You can change this manually or via command --greet<br>
typing {user} will mention (@User) the new member.<br>
type {/user} will just say the new members name.<br>

<br>

### BYEMSG
You can change this manually or via command --bye<br>
This will alert everyone when someone leaves the server.<br>

<br>

### ROLESYS
Leave this Alone. use the command --autorole Role Name<br>
This determines what role to give someone when they join.<br>
Setting this manually could cause for typos, etc = No Auto role response.<br>
Using the command will alert you if the role exists or not.<br>

<br>

### ACTION
You can change this manually or via command --setpunish<br>
There are three options: **kick**, **ban** and **warn**.<br>
This set's the punishment for the Anti Link System. Default = kick<br>

<br>

### SILENT
This can be set by typing --autorole -s Role Name<br>
true: Autogo will say something like "I have given @User the role Blah"<br>
false: Silently adds the member to the specified autorole.<br>


### HELPCMD
Here you can customize what they will see when --help command is triggered.
<br>

### BOTAUTOROLE
Autorole for Bot Accounts. This will detect OAuth bots<br>
and than give tham whatever role you choose!<br>

<Br>

### CUSTOM COMMAND NAMES (YOU HAVE TO EDIT THESE MANUALLY IN COMMANDS.JSON)
<br>
Let's look at the ORIGINAL commands.json file
<pre>
{
	"Greet": 		"greet",
	"Bye":			"bye",
	"Prefix":		"prefix",
	"Kick":			"kick",
	"Ban":			"ban",
	"Autorole":		"autorole",
	"SetPunish":	"setpunish",
	"AllowLinks":	"allowlinks",
	"DenyLinks":	"denylinks",
	"AddMaster":	"addmaster",
	"DelMaster":	"delmaster",
	"Invites":		"invites",
	"Meme":			"meme",
	"Joke":			"joke",
	"Give":			"give",
	"Take":			"take",
	"Giveme":		"giveme",
	"Mute":			"mute",
	"Unmute":		"unmute",
	"BotAutoRole":	"botrole"
}
</pre>
<br>
Let's change the addmaster and delmaster to givemaster and takemaster
<pre>
{
	"Greet": 		"greet",
	"Bye":			"bye",
	"Prefix":		"prefix",
	"Kick":			"kick",
	"Ban":			"ban",
	"Autorole":		"autorole",
	"SetPunish":	"setpunish",
	"AllowLinks":	"allowlinks",
	"DenyLinks":	"denylinks",
	"AddMaster":	"givemaster",
	"DelMaster":	"takemaster",
	"Invites":		"invites",
	"Meme":			"meme",
	"Joke":			"joke",
	"Give":			"give",
	"Take":			"take",
	"Giveme":		"giveme",
	"Mute":			"mute",
	"Unmute":		"unmute",
	"BotAutoRole":	"botrole"
}
</pre>
Now you will need to type --givemaster and --takemaster.<br>
Notice how i changed the Values not the Keys "AddMaster": or "DelMaster":<br>
It's very important you follow this structure.<br>
The key names can't change. Or Autogo will fail.<br>

<br>

### HELP COMMAND
You can completely customize what Autogo says when people type --help<br>
Just open the config.json file and change the HelpCmd value to what you want.<br>

<br>

### MAKE A BAT FILE THAT WILL AUTO RESTART IF IT CRASHES.
Create a new .txt file<br>
Name it StartBot.txt<br>
Open the file up and place the code below<br>
Now, rename the file to StartBot.bat<br>
Now you can run from the bat file.<br>

<br>

### STARTBOT.BAT SOURCE CODE
```bat
:start
TITLE AutoGo Personal Bot
COLOR 03
autogo.exe
goto start
```

<br>

### ..LOG FILE
If you want to log all the server text in the terminal the bot runs from<br>
than keep the ..log file, if you don't want to log any text simply delete the log file<br>
or rename it.

<br>

### AUTO RESPONSE SYSTEM
AutoGo has an Auto Response System! However it's new<br>
So it's very basic at the moment. if you open autoresponse.txt<br>
you will get ideas of what you need to do to add an auto response!<br>

<br>

### RESPONSES.JSON
You can customize what AutoGo says in every command or action.<br>
There are a few things that aren't customizable. But that's very few!<br>
Just check out **responses.json** for more info!<br>

<br>

### AUTOGO UPDATES AND SERVER
Visit our server to keep  updated on new features and more.<br>
https://discord.gg/0pTKzt2BDIpVzWgx
