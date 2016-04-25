// This file provides a basic "quick start" example of using the Discordgo
// package to connect to Discord using the New() helper function.
package main

import (
	"fmt"
	"time"
	"github.com/bwmarrin/discordgo"
	"strings"
	"io/ioutil"
	"strconv"
	"encoding/json"
//	"github.com/zymtom/argconf"
)
	var err error
	var startTime time.Time
	var js obj
	var cmd commands
	var resp responses


type update struct {
	Version		int
}






// until i can figure out how to get the roles from every member
// i have to do the masters system this way (the noob way)
/*
func isMaster(server string, user string) bool {
	mas := true

	var in info
	vfile, err := ioutil.ReadFile("servers/" + server + "/main.json")
	if err != nil {
		// mas = false
	}

	json.Unmarshal(vfile, &in)

	if _, err := os.Stat("servers/" + server + "/" + user + ".json"); err != nil {
		mas = false
	}

	if user == in.Admin {
		mas = true
	}
	return mas
}
*/













func main() {

	file, err := ioutil.ReadFile("config.json")
	json.Unmarshal(file, &js)

    // Login to discord. You can use a token or email, password arguments.
	dg, err := discordgo.New(js.Bot)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Register messageCreate as a callback for the messageCreate events.
	dg.AddHandler(messageCreate)
	dg.AddHandler(GuildMemberAdd)
	dg.AddHandler(GuildMemberRemove)
	dg.AddHandler(onReady)
	// dg.AddHandler(GuildRoleUpdate)
	// Open the websocket and begin listening.
	dg.Open()

	// Simple way to keep program running until any key press.
	var input string
	fmt.Scanln(&input)
	return
}










// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated user has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	chkErr := true

  c, err := s.State.Channel(m.ChannelID)
if err != nil {
c, err = s.Channel(m.ChannelID)
}


// let's try to prevent any errors from happening. hahaha....
if err != nil {
	chkErr = false // couldn't determine the Guild ID
}


// ####### FUNCTIONS: Located in the funcs.go file #######
	var in obj

if chkErr == true {



	// Load up the custom commands.
	cfile, err := ioutil.ReadFile("commands.json")
	if err != nil {
		return
	} else {
	json.Unmarshal(cfile, &cmd)
	}


	// Load up the custom responses.
	rfile, err := ioutil.ReadFile("responses.json")
	if err != nil {
		return
	} else {
	json.Unmarshal(rfile, &resp)
	}



	// Load up the server information
	vfile, err := ioutil.ReadFile("config.json")
	if err != nil {
		return
	} else {
	json.Unmarshal(vfile, &in)
	}


in.BotCommander = GetRoleID(s, c.GuildID, "Bot Commander")	// GetRoleID returns the role id.
in.BotMaster = isMemberRole(s, c.GuildID, m.Author.ID, "Bot Commander")	// isMemberRole returns true or false

if m.Author.ID == in.Admin {
	in.BotMaster = true
}


// ############## Status Attempt #1 ##################
/*
ticker := time.NewTicker(2 * time.Minute)
quit := make(chan struct{})

    for {
       select {
        case <- ticker.C:
			myrand := random(1, 9)
			var status []string
			status, err := readLines("status.txt")
			if err == nil {
				s.UpdateStatus(0, status[myrand])
			}
        case <- quit:
            ticker.Stop()
            return
        }
    }
*/




//	fmt.Println("Prefix: " + in.Prefix + "\nGreet: " + in.GreetMsg + "\nBye: " + in.ByeMsg + "\nAutorole: " + in.RoleSys)
	// Print message to stdout.

_, err = ioutil.ReadFile("..log")
if err == nil {
	if m.Author.ID != s.State.User.ID {
		fmt.Println("[" + time.Now().Format(time.Stamp) + "] " + " - " + m.Author.Username + ": " + m.Content)
	}
}








if m.Author.ID != s.State.User.ID {
 // -#$-
var auto []string
var cn int
cn = 0
auto, err = readLines("autoresponse.txt")
if err == nil {
	for _, ars := range auto {
		dont := false
		ispm := false
		isfind := false
		if strings.HasPrefix(ars, "//") {
			dont = true
		}
		if ars == "" {
			dont = true
		}



		if dont == false && strings.Contains(ars, "=") {
	//	fmt.Println("RAW: " + ars)
		ardat := strings.Split(ars, "=")
		trigger := ardat[0]
		response := strings.Replace(ars, trigger+"=", "", -1)
	//	response := ardat[1]
		response = strings.Replace(response, "{user}", "<@"+m.Author.ID+">", -1)
		response = strings.Replace(response, "{/user}", m.Author.Username, -1)
		
		if strings.Contains(response, "{kick}") {
			response = strings.Replace(response, "{kick}", "", -1)
			s.GuildMemberDelete(c.GuildID, m.Author.ID)
		}

		if strings.Contains(response, "{ban}") {
			response = strings.Replace(response, "{ban}", "", -1)
			s.GuildBanCreate(c.GuildID, m.Author.ID, 10)
		}


		cntrole := 0
		/*
		// lets work on excluding roles from triggers.
		if strings.Contains(response, "{exc=") {
			newdat := strings.Split(response, "{exc=")
			newdat = strings.Split(newdat[1], "}")
			exclude := newdat[0]
			// let's see if it's multiple roles or just one.
			if strings.Contains(exclude, ",") {
				excluded := strings.Split(exclude, ",")
				for _, vR := range excluded {
					if isMemberRole(s, c.GuildID, m.Author.ID, vR) == true {
						fmt.Println("Found the role "+vR)
						response = strings.Replace(response, "{exc="+exclude+"}", "", -1)
						cntrole++
					}
				} // checking for multiple ppl.
			} else {
				// only a single role is detected.
				if isMemberRole(s, c.GuildID, m.Author.ID, exclude) == true {
					fmt.Println("Found the single role: "+exclude)
					response = strings.Replace(response, "{exc="+exclude+"}", "", -1)
					cntrole++
				}
			} // end of excludes check
		} // end of excluding roles from triggers
		*/




		if strings.HasPrefix(trigger, "&") {
			cn++
			isfind = true
	//		fmt.Println("Found: &")
			trigger = strings.Replace(trigger, "&", "", -1)
		}

	if strings.HasPrefix(trigger, in.Prefix) {
	//		fmt.Println("Found: "+in.Prefix)
		//	isfind = true
	}

	// Let's detect if it was a PM or not.
	if strings.Contains(response, "{pm}") {
		response = strings.Replace(response, "{pm}", "", -1)
		ispm = true
	}

	if strings.Contains(response, ":br") {
		response = strings.Replace(response, ":br", "\n", -1)
	}
	if cntrole == 0 {
	//	fmt.Println("Trigger: " + trigger)
	//	fmt.Println("Response: " + response)
		// just a basic ARS trigger. Later i will code for {find=word}
		if m.Content == trigger {
			if ispm == false {
				s.ChannelTyping(m.ID)
				time.Sleep(1000 * time.Millisecond)
			s.ChannelMessageSend(m.ChannelID, response)
			} else {
				k, err := s.UserChannelCreate(m.Author.ID)
				if err == nil {
					s.ChannelTyping(k.ID)
					time.Sleep(1000 * time.Millisecond)
					s.ChannelMessageSend(k.ID, response)
				}
			} // check if it's a pm or a server request.
		} // end of basic trigger

		if strings.Contains(m.Content, trigger) && isfind == true && cn == 1 {
	//		fmt.Println("It has worked!")
			s.ChannelTyping(m.ID)
			time.Sleep(1000 * time.Millisecond)
			s.ChannelMessageSend(m.ChannelID, response)				
		}
		} // end of cntrole == 0
		} // end of dont == false
	} // end of for loop
} // check to see if they have autoresponse.txt file in bot dir.
}








if strings.HasPrefix(m.Content, in.Prefix) {
	in.CmdsRun++
	newConf := obj{
		Bot:			in.Bot,
		Admin:			in.Admin,
		Status:			in.Status,
		BotMaster:		in.BotMaster,
		BotCommander:	in.BotCommander,
		CmdsRun:		in.CmdsRun,
		Prefix:			in.Prefix,
		GreetMsg:		in.GreetMsg,
		ByeMsg:			in.ByeMsg,
		RoleSys:		in.RoleSys,
		Name:			in.Name,
		AntiLink:		in.AntiLink,
		Action:			in.Action,
		Silent:			in.Silent,
		HelpCmd:		in.HelpCmd,
		BotAutoRole:	in.BotAutoRole,
		}
	b, err := json.Marshal(newConf)
	if err == nil {
		ioutil.WriteFile("config.json", b, 0777)
	}
}








// fmt.Println(js.BotMaster)














	if strings.HasPrefix(m.Content, in.Prefix + "help") {
		s.ChannelTyping(m.ChannelID)
		time.Sleep(1000 * time.Millisecond)
		if m.Content == in.Prefix + "help" {
			bm := "False"

			if in.BotMaster == true {
				bm = "True"
			}
			ts := in.CmdsRun
			i := strconv.Itoa(ts)

			var _ = bm + i

	//		fmt.Println("Converted: " + i)
	//		fmt.Println(in.CmdsRun)
		
		s.ChannelMessageSend(m.ChannelID, in.HelpCmd)
		}


		if m.Content == in.Prefix + "help "+cmd.AddMaster {
			s.ChannelMessageSend(m.ChannelID, "```ruby\npermissions: Server Owner\nusage: "+in.Prefix+cmd.AddMaster+" @User\ninfo: gives user acces to mod commands.```")
		}

		if m.Content == in.Prefix + "help "+cmd.DelMaster {
			s.ChannelMessageSend(m.ChannelID, "```ruby\npermissions: Server Owner\nusage: "+in.Prefix+cmd.DelMaster+" @User\ninfo: removes access to mod commands.```")
		}

		if m.Content == in.Prefix + "help "+cmd.Greet {
			s.ChannelMessageSend(m.ChannelID, "```ruby\npermissions: Server Owner & Bot Commanders\nusage: "+in.Prefix+cmd.Greet+" Welcome {user} if you need any help ask!\ninfo: use {user} to mention the new member.\nuse {/user} to just say their username!\njust set the message to off for turning the greet message off```")
		}

		if m.Content == in.Prefix + "help "+cmd.Bye {
			s.ChannelMessageSend(m.ChannelID, "```ruby\npermissions: Server Owner & Bot Commanders\nusage: "+in.Prefix+cmd.Bye+" {user} has left the server.\nto turn off set the bye message to off```")
		}

		if m.Content == in.Prefix + "help "+cmd.DenyLinks {
			s.ChannelMessageSend(m.ChannelID, "```ruby\npermissions: Server Owner & Bot Commanders\nusage: "+in.Prefix+cmd.DenyLinks+"\ninfo: enables my anti link module. by default i kick offenders. you can use setpunish command to change to ban.```")
		}

		if m.Content == in.Prefix + "help "+cmd.AllowLinks {
			s.ChannelMessageSend(m.ChannelID, "```ruby\npermissions: Server Owner & Bot Commanders\nusage: "+in.Prefix+cmd.AllowLinks+"\ninfo: turns my anti link module off.```")
		}

		if m.Content == in.Prefix + "help "+cmd.Prefix {
			s.ChannelMessageSend(m.ChannelID, "```ruby\npermissions: Server Owner & Bot Commanders\nusage: "+in.Prefix+cmd.Prefix+" #\ninfo: sets my prefix in your server to #```")
		}

		if m.Content == in.Prefix + "help "+cmd.Autorole {
			s.ChannelMessageSend(m.ChannelID, "```ruby\npermissions: Server Owner & Bot Commanders\nusage: "+in.Prefix+cmd.Autorole+" Role Name\ninfo: automatically assign a role to new members.\nsilently add roles type "+in.Prefix+"autorole -s Role Name```")
		}

		if m.Content == in.Prefix + "help "+cmd.Invites {
			s.ChannelMessageSend(m.ChannelID, "```ruby\npermissions: Everyone\nusage: "+in.Prefix+"invites\ninfo: gives you a list of available invite codes for your channel.```")
		}

		if m.Content == in.Prefix + "help mkinvite" {
			s.ChannelMessageSend(m.ChannelID, "```ruby\npermissions: Server Owner & Bot Commanders\nusage: "+in.Prefix+"mkinvite\ninfo: creates a permenant invite code for your channel.```")
		}

		if m.Content == in.Prefix + "help "+cmd.Give {
			s.ChannelMessageSend(m.ChannelID, "```ruby\npermissions: Server Owner\nusage: "+in.Prefix+cmd.Give+" @User Role Name\ninfo: gives the user the specified role.```")
		}

		if m.Content == in.Prefix + "help "+cmd.Take {
			s.ChannelMessageSend(m.ChannelID, "```ruby\npermissions: Server Owner\nusage: "+in.Prefix+cmd.Take+" @User Role Name\ninfo: takes the user the specified role.```")
		}

		if m.Content == in.Prefix + "help "+cmd.Giveme {
			s.ChannelMessageSend(m.ChannelID, "```ruby\npermissions: Server Owner\nusage: "+in.Prefix+cmd.Giveme+" Role Name\ninfo: gives you the specified role.```")
		}

		if m.Content == in.Prefix + "help "+cmd.SetPunish {
			s.ChannelMessageSend(m.ChannelID, "```ruby\npermissions: Server Owner & Bot Commanders\nusage: "+in.Prefix+cmd.SetPunish+" kick or ban\ninfo: you can set the anti links module punishment to either kick or ban. by default its set to kick.```")
		}

		if m.Content == in.Prefix + "help "+cmd.Mute {
			s.ChannelMessageSend(m.ChannelID, "```ruby\npermissions: Server Owner & Bot Commanders\nusage: "+in.Prefix+cmd.Mute+" @User\ninfo: mutes the user. you need to make a role named muted and set the permissions to not speak and than add the role to your channels.```")
		}

		if m.Content == in.Prefix + "help "+cmd.Unmute {
			s.ChannelMessageSend(m.ChannelID, "```ruby\npermissions: Server Owner & Bot Commanders\nusage: "+in.Prefix+cmd.Unmute+" @User\ninfo: unmutes the user.```")
		}

		if m.Content == in.Prefix + "help "+cmd.BotAutoRole {
			s.ChannelMessageSend(m.ChannelID, "```ruby\npermissions: Server Owner & Bot Commanders\nusage: "+in.Prefix+cmd.BotAutoRole+" Role Name\ninfo: Will give any bot that joins your server a specific role.```")
		}
	}









if in.Admin == m.Author.ID && strings.HasPrefix(m.Content, in.Prefix + "rename") {
	str := strings.Replace(m.Content, in.Prefix + "rename ", "", -1)
	tr := strings.Split(str, " ")
	Before := tr[0]
	After := tr[1]

	if Before == cmd.Greet {
		cmd.Greet = After
	}
	if Before == cmd.Bye {
		cmd.Bye = After
	}
	if Before == cmd.Prefix {
		cmd.Prefix = After
	}
	if Before == cmd.Kick {
		cmd.Kick = After
	}
	if Before == cmd.Ban {
		cmd.Ban = After
	}
	if Before == cmd.Autorole {
		cmd.Autorole = After
	}
	if Before == cmd.SetPunish {
		cmd.SetPunish = After
	}
	if Before == cmd.AllowLinks {
		cmd.AllowLinks = After
	}
	if Before == cmd.DenyLinks {
		cmd.DenyLinks = After
	}
	if Before == cmd.AddMaster {
		cmd.AddMaster = After
	}
	if Before == cmd.DelMaster {
		cmd.DelMaster = After
	}
	if Before == cmd.Invites {
		cmd.Kick = After
	}
	if Before == cmd.Meme {
		cmd.Meme = After
	}
	if Before == cmd.Joke {
		cmd.Joke = After
	}
	if Before == cmd.Give {
		cmd.Give = After
	}
	if Before == cmd.Take {
		cmd.Take = After
	}
	if Before == cmd.Giveme {
		cmd.Giveme = After
	}
	if Before == cmd.Mute {
		cmd.Mute = After
	}
	if Before == cmd.Unmute {
		cmd.Unmute = After
	}

	newConf := commands{
		Greet: 			cmd.Greet,
		Bye:			cmd.Bye,
		Prefix:			cmd.Prefix,
		Kick:			cmd.Kick,
		Ban:			cmd.Ban,
		Autorole:		cmd.Autorole,
		SetPunish:		cmd.SetPunish,
		AllowLinks:		cmd.AllowLinks,
		DenyLinks:		cmd.DenyLinks,
		AddMaster:		cmd.AddMaster,
		DelMaster:		cmd.DelMaster,
		Invites:		cmd.Invites,
		Meme:			cmd.Meme,
		Joke:			cmd.Joke,
		Give:			cmd.Give,
		Take:			cmd.Take,
		Giveme:			cmd.Giveme,
		Mute:			cmd.Mute,
		Unmute:			cmd.Unmute,
		BotAutoRole:	cmd.BotAutoRole,
	}
	b, err := json.Marshal(newConf)
	if err == nil {
		ioutil.WriteFile("commands.json", b, 0777)
		s.ChannelTyping(m.ChannelID)
		time.Sleep(1000 * time.Millisecond)
		s.ChannelMessageSend(m.ChannelID, "I've renamed the command `"+Before+"` to `"+After+"`")
	}


}






























if in.Admin == m.Author.ID && strings.HasPrefix(m.Content, in.Prefix + cmd.AddMaster) {
str := strings.Replace(m.Content, in.Prefix + cmd.AddMaster+" ", "", -1)
str = strings.Replace(str, "<@", "", -1)
str = strings.Replace(str, ">", "", -1)

z, err := s.State.Member(c.GuildID, str) 
if err != nil {
z, err = s.GuildMember(c.GuildID, str)
}
	
	if err != nil {
		fmt.Println(err)
	}

		roles, err := s.GuildRoles(c.GuildID)
		if err == nil {
			for _, v := range roles {
    			if v.Name == "Bot Commander" {
    				z.Roles = append(z.Roles, v.ID)
    				s.GuildMemberEdit(c.GuildID, str, z.Roles)
					s.ChannelTyping(m.ChannelID)
					time.Sleep(1000 * time.Millisecond)
					newdata := strings.Replace(resp.AddMaster, "{user}", "<@"+str+">", -1)
					s.ChannelMessageSend(m.ChannelID, newdata)
    			}
			}
		}
	} // end of Add master command




















if in.Admin == m.Author.ID && strings.HasPrefix(m.Content, in.Prefix + cmd.DelMaster) {
str := strings.Replace(m.Content, in.Prefix + cmd.DelMaster+" ", "", -1)
str = strings.Replace(str, "<@", "", -1)
str = strings.Replace(str, ">", "", -1)

x, err := s.State.Member(c.GuildID, str) 
if err != nil {
x, err = s.GuildMember(c.GuildID, str)
}

	if err != nil {
		fmt.Println(err)
	}

	var mc []string
	mc = x.Roles
	for mr := range x.Roles {
		t := mc[mr]
		if strings.Contains(t, in.BotCommander) {
    		// z.Roles = append(z.Roles, t[:0])
    		x.Roles = append(x.Roles[:mr], x.Roles[mr+1:]...)
    		if err != nil {
    			return
    		}
    		s.GuildMemberEdit(c.GuildID, str, x.Roles)
			s.ChannelTyping(m.ChannelID)
			time.Sleep(1000 * time.Millisecond)
			newdata := strings.Replace(resp.DelMaster, "{user}", "<@"+str+">", -1)
			s.ChannelMessageSend(m.ChannelID, newdata)
		}
	}
} // end of Del master command





















if in.Admin == m.Author.ID && strings.HasPrefix(m.Content, in.Prefix + cmd.Take) {
str := strings.Replace(m.Content, in.Prefix + cmd.Take + " ", "", -1)
dat := strings.Split(str, " ")
usr := dat[0]
usr = strings.Replace(usr, "<@", "", -1)
usr = strings.Replace(usr, ">", "", -1)
role := strings.Replace(str, "<@"+usr+"> ", "", -1)

 var roleID string

mroles, err := s.GuildRoles(c.GuildID)
if err == nil {
 for _, v := range mroles {
    if v.Name == role {
//    	fmt.Println("Found the role: "+role+"\nID: "+v.ID)
    	roleID = v.ID
    }
  }
  }


x, err := s.State.Member(c.GuildID, str) 
if err != nil {
x, err = s.GuildMember(c.GuildID, str)
}

	if err != nil {
		fmt.Println(err)
	}

	var ms []string
	ms = x.Roles
	for mr := range x.Roles {
		t := ms[mr]
		if strings.Contains(t, roleID) {
	//		fmt.Println("Membert has role: "+t)
    		x.Roles = append(x.Roles[:mr], x.Roles[mr+1:]...)
    		s.GuildMemberEdit(c.GuildID, usr, x.Roles)
			s.ChannelTyping(m.ChannelID)
			time.Sleep(1000 * time.Millisecond)
			newdata := strings.Replace(resp.Take, "{user}", "<@"+usr+">", -1)
			newdata = strings.Replace(newdata, "{data}", role, -1)
			s.ChannelMessageSend(m.ChannelID, newdata)
    	}
	}
} // end of giveme command.

























if in.Admin == m.Author.ID && strings.HasPrefix(m.Content, in.Prefix + cmd.Give) {
str := strings.Replace(m.Content, in.Prefix + cmd.Give + " ", "", -1)
dat := strings.Split(str, " ")
usr := dat[0]
usr = strings.Replace(usr, "<@", "", -1)
usr = strings.Replace(usr, ">", "", -1)
role := strings.Replace(str, "<@"+usr+"> ", "", -1)

x, err := s.State.Member(c.GuildID, usr) 
if err != nil {
x, err = s.GuildMember(c.GuildID, usr)
}

if err != nil {
	fmt.Println(err)
}

		roles, err := s.GuildRoles(c.GuildID)
		if err == nil {
			for _, v := range roles {
    			if v.Name == role {
    				x.Roles = append(x.Roles, v.ID)
    				s.GuildMemberEdit(c.GuildID, usr, x.Roles)
					s.ChannelTyping(m.ChannelID)
					time.Sleep(1000 * time.Millisecond)
					newdata := strings.Replace(resp.Give, "{user}", "<@"+usr+">", -1)
					newdata = strings.Replace(newdata, "{data}", role, -1)
					s.ChannelMessageSend(m.ChannelID, newdata)
    			}
			}
		}
	} // end of giveme command.




























if in.BotMaster == true && strings.HasPrefix(m.Content, in.Prefix + cmd.Mute) {
str := strings.Replace(m.Content, in.Prefix + cmd.Mute + " ", "", -1)
str = strings.Replace(str, "<@", "", -1)
str = strings.Replace(str, ">", "", -1)

// fmt.Println("Adding Master: "+str)

z, err := s.State.Member(c.GuildID, str) 
if err != nil {
z, err = s.GuildMember(c.GuildID, str)
}

if err != nil {
	fmt.Println(err)
}

		roles, err := s.GuildRoles(c.GuildID)
		if err == nil {
			for _, v := range roles {
    			if strings.ToLower(v.Name) == "muted" {
    				z.Roles = append(z.Roles, v.ID)
    				s.GuildMemberEdit(c.GuildID, str, z.Roles)
					s.ChannelTyping(c.GuildID)
					time.Sleep(1000 * time.Millisecond)
					newdata := strings.Replace(resp.Mute, "{user}", "<@"+str+">", -1)
					newdata = strings.Replace(newdata, "{data}", "<#"+m.ChannelID+">", -1)
					s.ChannelMessageSend(m.ChannelID, newdata)
    			}
			}
		}
	} // end of giveme command.

























if in.BotMaster == true && strings.HasPrefix(m.Content, in.Prefix + cmd.Unmute) {
str := strings.Replace(m.Content, in.Prefix + cmd.Unmute + " ", "", -1)
usr := strings.Replace(str, "<@", "", -1)
usr = strings.Replace(usr, ">", "", -1)
	var roleID string

mroles, err := s.GuildRoles(c.GuildID)
if err == nil {

 for _, v := range mroles {
    if v.Name == "muted" {
    	roleID = v.ID
    }
  }
  }
x, err := s.State.Member(c.GuildID, usr) 
if err != nil {
x, err = s.GuildMember(c.GuildID, usr)
}

if err != nil {
	fmt.Println(err)
}

	var ms []string
	ms = x.Roles
	for mr := range x.Roles {
		t := ms[mr]
		if t == roleID {
    		x.Roles = append(x.Roles[:mr], x.Roles[mr+1:]...)
    		s.GuildMemberEdit(c.GuildID, usr, x.Roles)
			s.ChannelTyping(m.ChannelID)
			time.Sleep(1000 * time.Millisecond)
			newdata := strings.Replace(resp.Unmute, "{user}", "<@"+usr+">", -1)
			s.ChannelMessageSend(m.ChannelID, newdata)
			}
	}
} // end of giveme command.


























if in.BotMaster == true && strings.HasPrefix(m.Content, in.Prefix + cmd.Greet) {
		str := strings.Replace(m.Content, in.Prefix + cmd.Greet+" ", "", -1)

		
			newjs := obj{
		Bot:			in.Bot,
		Admin:			in.Admin,
		Status:			in.Status,
		BotMaster:		in.BotMaster,
		BotCommander:	in.BotCommander,
		CmdsRun:		in.CmdsRun,
		Prefix:			in.Prefix,
		GreetMsg:		str,
		ByeMsg:			in.ByeMsg,
		RoleSys:		in.RoleSys,
		Name:			in.Name,
		AntiLink:		in.AntiLink,
		Action:			in.Action,
		Silent:			in.Silent,
		HelpCmd:		in.HelpCmd,
		BotAutoRole:	in.BotAutoRole,
			}
			b, err := json.Marshal(newjs)
			if err == nil {
				ioutil.WriteFile("config.json", b, 0777)
			}
	if str != "off" {
		s.ChannelTyping(m.ChannelID)
		time.Sleep(1000 * time.Millisecond)
			newdata := strings.Replace(resp.Greet, "{data}", str, -1)
		s.ChannelMessageSend(m.ChannelID, newdata)
	} else {
		s.ChannelTyping(m.ChannelID)
		time.Sleep(1000 * time.Millisecond)
		s.ChannelMessageSend(m.ChannelID, resp.GreetOff)		
	}
}




















if in.BotMaster == true && strings.HasPrefix(m.Content, in.Prefix + cmd.Bye) {
	str := strings.Replace(m.Content, in.Prefix + cmd.Bye + " ", "", -1)
	newjs := obj{
		Bot:			in.Bot,
		Admin:			in.Admin,
		Status:			in.Status,
		BotMaster:		in.BotMaster,
		BotCommander:	in.BotCommander,
		CmdsRun:		in.CmdsRun,
		Prefix:			in.Prefix,
		GreetMsg:		in.GreetMsg,
		ByeMsg:			str,
		RoleSys:		in.RoleSys,
		Name:			in.Name,
		AntiLink:		in.AntiLink,
		Action:			in.Action,
		Silent:			in.Silent,
		HelpCmd:		in.HelpCmd,
		BotAutoRole:	in.BotAutoRole,
	}
		b, err := json.Marshal(newjs)
		if err == nil {
			ioutil.WriteFile("config.json", b, 0777)
		}
	if str != "off" {
		s.ChannelTyping(m.ChannelID)
		time.Sleep(1000 * time.Millisecond)
		newdata := strings.Replace(resp.Bye, "{data}", str, -1)
		s.ChannelMessageSend(m.ChannelID, newdata)
	} else {
		s.ChannelTyping(m.ChannelID)
		time.Sleep(1000 * time.Millisecond)
		s.ChannelMessageSend(m.ChannelID, resp.ByeOff)	
	}
}




















	if in.BotMaster == true && strings.HasPrefix(m.Content, in.Prefix + "mkinvite") {

		b := discordgo.Invite{
			MaxAge:		0,
			MaxUses:	0,
			Temporary:	false,
			XkcdPass:	false,
		}
	iv, err := s.ChannelInviteCreate(m.ChannelID, b)
	if err != nil {
		return
	}
		s.ChannelTyping(m.ChannelID)
		time.Sleep(1000 * time.Millisecond)
		s.ChannelMessageSend(m.ChannelID, "I've made the Invite Code: "+iv.Code)
		return
	}




















	if in.BotMaster == true && strings.HasPrefix(m.Content, in.Prefix + cmd.DenyLinks) {
		newjs := obj{
		Bot:			in.Bot,
		Admin:			in.Admin,
		Status:			in.Status,
		BotMaster:		in.BotMaster,
		BotCommander:	in.BotCommander,
		CmdsRun:		in.CmdsRun,
		Prefix:			in.Prefix,
		GreetMsg:		in.GreetMsg,
		ByeMsg:			in.ByeMsg,
		RoleSys:		in.RoleSys,
		Name:			in.Name,
		AntiLink:		true,
		Action:			in.Action,
		Silent:			in.Silent,
		HelpCmd:		in.HelpCmd,
		BotAutoRole:	in.BotAutoRole,
		}
		b, err := json.Marshal(newjs)
		if err == nil {
			ioutil.WriteFile("config.json", b, 0777)
		}

		s.ChannelTyping(m.ChannelID)
		time.Sleep(1000 * time.Millisecond)
		s.ChannelMessageSend(m.ChannelID, resp.DenyLinks)
		return
	}























	if in.BotMaster == true && strings.HasPrefix(m.Content, in.Prefix + cmd.AllowLinks) {
		newjs := obj{
		Bot:			in.Bot,
		Admin:			in.Admin,
		Status:			in.Status,
		BotMaster:		in.BotMaster,
		BotCommander:	in.BotCommander,
		CmdsRun:		in.CmdsRun,
		Prefix:			in.Prefix,
		GreetMsg:		in.GreetMsg,
		ByeMsg:			in.ByeMsg,
		RoleSys:		in.RoleSys,
		Name:			in.Name,
		AntiLink:		false,
		Action:			in.Action,
		Silent:			in.Silent,
		HelpCmd:		in.HelpCmd,
		BotAutoRole:	in.BotAutoRole,
		}
		b, err := json.Marshal(newjs)
		if err == nil {
			ioutil.WriteFile("config.json", b, 0777)
		}

		s.ChannelTyping(m.ChannelID)
		time.Sleep(1000 * time.Millisecond)
		s.ChannelMessageSend(m.ChannelID, resp.AllowLinks)
		return
	}



















	if in.BotMaster == true && strings.HasPrefix(m.Content, in.Prefix + cmd.Prefix) {
		str := strings.Replace(m.Content, in.Prefix + cmd.Prefix + " ", "", -1)

		newjs := obj{
		Bot:			in.Bot,
		Admin:			in.Admin,
		Status:			in.Status,
		BotMaster:		in.BotMaster,
		BotCommander:	in.BotCommander,
		CmdsRun:		in.CmdsRun,
		Prefix:			str,
		GreetMsg:		in.GreetMsg,
		ByeMsg:			in.ByeMsg,
		RoleSys:		in.RoleSys,
		Name:			in.Name,
		AntiLink:		in.AntiLink,
		Action:			in.Action,
		Silent:			in.Silent,
		HelpCmd:		in.HelpCmd,
		BotAutoRole:	in.BotAutoRole,
		}
		b, err := json.Marshal(newjs)
		if err == nil {
			ioutil.WriteFile("config.json", b, 0777)
		}

		s.ChannelTyping(m.ChannelID)
		time.Sleep(1000 * time.Millisecond)
		newdata := strings.Replace(resp.Prefix, "{data}", str, -1)
		s.ChannelMessageSend(m.ChannelID, newdata)
		return
	}





















	if in.BotMaster == true && strings.HasPrefix(m.Content, in.Prefix + cmd.SetPunish) {
		str := strings.Replace(m.Content, in.Prefix + cmd.SetPunish + " ", "", -1)

		check := false

		if strings.ToLower(str) == "kick" {
			check = true
		}
		if strings.ToLower(str) == "ban" {
			check = true
		}
		if strings.ToLower(str) == "warn" {
			check = true
		}
		if check == true {
			newjs := obj{
		Bot:			in.Bot,
		Admin:			in.Admin,
		Status:			in.Status,
		BotMaster:		in.BotMaster,
		BotCommander:	in.BotCommander,
		CmdsRun:		in.CmdsRun,
		Prefix:			in.Prefix,
		GreetMsg:		in.GreetMsg,
		ByeMsg:			in.ByeMsg,
		RoleSys:		in.RoleSys,
		Name:			in.Name,
		AntiLink:		in.AntiLink,
		Action:			str,
		Silent:			in.Silent,
		HelpCmd:		in.HelpCmd,
		BotAutoRole:	in.BotAutoRole,
			}
			b, err := json.Marshal(newjs)
			if err == nil {
				ioutil.WriteFile("config.json", b, 0777)
			}
			s.ChannelTyping(m.ChannelID)
			time.Sleep(1000 * time.Millisecond)
			newdata := strings.Replace(resp.SetPunish, "{data}", str, -1)
			s.ChannelMessageSend(m.ChannelID, newdata)
		}
		if check == false {
			s.ChannelTyping(m.ChannelID)
			time.Sleep(1000 * time.Millisecond)
			s.ChannelMessageSend(m.ChannelID, "You need to pick a proper punishment for Anti links type `"+in.Prefix+cmd.SetPunish+" kick` or `"+in.Prefix+cmd.SetPunish+" ban` or `"+in.Prefix+cmd.SetPunish+" warn`")
		}
		return
	}




















if in.BotMaster == true && strings.HasPrefix(m.Content, in.Prefix + cmd.Autorole) {
str := strings.Replace(m.Content, in.Prefix + cmd.Autorole + " ", "", -1)
var br bool
br = false
if strings.Split(str, " ")[0] == "-s" {
br = true
str = strings.Replace(m.Content, in.Prefix + cmd.Autorole + " -s ", "", -1)
}

cnt := 0
roles, err := s.GuildRoles(c.GuildID)
if err != nil {
fmt.Println(err)
}

if str != "off" {
 for _, v := range roles {
    if v.Name == str {
    	cnt++
		newjs := obj{
		Bot:			in.Bot,
		Admin:			in.Admin,
		Status:			in.Status,
		BotMaster:		in.BotMaster,
		BotCommander:	in.BotCommander,
		CmdsRun:		in.CmdsRun,
		Prefix:			in.Prefix,
		GreetMsg:		in.GreetMsg,
		ByeMsg:			in.ByeMsg,
		RoleSys:		str,
		Name:			in.Name,
		AntiLink:		in.AntiLink,
		Action:			in.Action,
		Silent:			br,
		HelpCmd:		in.HelpCmd,
		BotAutoRole:	in.BotAutoRole,
			}
		b, err := json.Marshal(newjs)
		if err == nil {
			ioutil.WriteFile("config.json", b, 0777)
		}
		s.ChannelTyping(m.ChannelID)
		time.Sleep(1000 * time.Millisecond)
		newdata := strings.Replace(resp.AutoRole, "{data}", str, -1)
		s.ChannelMessageSend(m.ChannelID, newdata)
		return
	}
}
} // make sure they don't want to turn the autorole off


if str == "off" {
	cnt = 1
		newjs := obj{
		Bot:			in.Bot,
		Admin:			in.Admin,
		Status:			in.Status,
		BotMaster:		in.BotMaster,
		BotCommander:	in.BotCommander,
		CmdsRun:		in.CmdsRun,
		Prefix:			in.Prefix,
		GreetMsg:		in.GreetMsg,
		ByeMsg:			in.ByeMsg,
		RoleSys:		"off",
		Name:			in.Name,
		AntiLink:		in.AntiLink,
		Action:			in.Action,
		Silent:			in.Silent,
		HelpCmd:		in.HelpCmd,
		BotAutoRole:	in.BotAutoRole,
			}
		b, err := json.Marshal(newjs)
		if err == nil {
			ioutil.WriteFile("config.json", b, 0777)
		}
		s.ChannelTyping(m.ChannelID)
		time.Sleep(1000 * time.Millisecond)
		s.ChannelMessageSend(m.ChannelID, resp.AutoRoleOff)
}




	if cnt == 0 {
		s.ChannelTyping(m.ChannelID)
		time.Sleep(1000 * time.Millisecond)
		newdata := strings.Replace(resp.NoRole, "{data}", str, -1)
		s.ChannelMessageSend(m.ChannelID, newdata)
	}
}
































if in.BotMaster == true && strings.HasPrefix(m.Content, in.Prefix + cmd.BotAutoRole) {
str := strings.Replace(m.Content, in.Prefix + cmd.BotAutoRole + " ", "", -1)

cnt := 0
roles, err := s.GuildRoles(c.GuildID)
if err != nil {
fmt.Println(err)
}

if str != "off" {
 for _, v := range roles {
    if v.Name == str {
    	cnt++
		newjs := obj{
		Bot:			in.Bot,
		Admin:			in.Admin,
		Status:			in.Status,
		BotMaster:		in.BotMaster,
		BotCommander:	in.BotCommander,
		CmdsRun:		in.CmdsRun,
		Prefix:			in.Prefix,
		GreetMsg:		in.GreetMsg,
		ByeMsg:			in.ByeMsg,
		RoleSys:		in.RoleSys,
		Name:			in.Name,
		AntiLink:		in.AntiLink,
		Action:			in.Action,
		Silent:			in.Silent,
		HelpCmd:		in.HelpCmd,
		BotAutoRole:	str,
			}
		b, err := json.Marshal(newjs)
		if err == nil {
			ioutil.WriteFile("config.json", b, 0777)
		}
		s.ChannelTyping(m.ChannelID)
		time.Sleep(1000 * time.Millisecond)
		newdata := strings.Replace(resp.BotAutoRole, "{data}", str, -1)
		s.ChannelMessageSend(m.ChannelID, newdata)
		return
	}
}
} // make sure they don't want to turn the autorole off


if str == "off" {
	cnt = 1
		newjs := obj{
		Bot:			in.Bot,
		Admin:			in.Admin,
		Status:			in.Status,
		BotMaster:		in.BotMaster,
		BotCommander:	in.BotCommander,
		CmdsRun:		in.CmdsRun,
		Prefix:			in.Prefix,
		GreetMsg:		in.GreetMsg,
		ByeMsg:			in.ByeMsg,
		RoleSys:		in.RoleSys,
		Name:			in.Name,
		AntiLink:		in.AntiLink,
		Action:			in.Action,
		Silent:			in.Silent,
		HelpCmd:		in.HelpCmd,
		BotAutoRole:	"off",
			}
		b, err := json.Marshal(newjs)
		if err == nil {
			ioutil.WriteFile("config.json", b, 0777)
		}
		s.ChannelTyping(m.ChannelID)
		time.Sleep(1000 * time.Millisecond)
		s.ChannelMessageSend(m.ChannelID, resp.AutoBotRoleOff)
}




	if cnt == 0 {
		s.ChannelTyping(m.ChannelID)
		time.Sleep(1000 * time.Millisecond)
		newdata := strings.Replace(resp.NoRole, "{data}", str, -1)
		s.ChannelMessageSend(m.ChannelID, newdata)
	}
}
























	if strings.HasPrefix(m.Content, in.Prefix + "invites") {
		o, err := s.ChannelInvites(m.ChannelID)
		if err == nil {
			data := ""
			s.ChannelTyping(m.ChannelID)
			time.Sleep(1000 * time.Millisecond)
			s.ChannelMessageSend(m.ChannelID, "Invites for: `"+c.GuildID+"`\n```ruby\nGrabbing Results..```")
			for _, v := range o {
			data = data + "\nissuer: "+v.Inviter.Username+"\ncode: "+v.Code
			// s.ChannelMessageDelete(m.ChannelID, theid)
			//s.ChannelMessageEdit(m.ChannelID, m.ID, "Invites for: `"+c.GuildID+"` "+v.Inviter.Username + "\n" + v.Code)
  		}
			s.ChannelTyping(m.ChannelID)
			time.Sleep(1000 * time.Millisecond)
  		s.ChannelMessageSend(m.ChannelID, "```ruby\n"+data+"```")
		return
		}
	}

















dblchk := isMemberRole(s, c.GuildID, m.Author.ID, "Bot Commander")	// isMemberRole returns true or false
if m.Author.ID == in.Admin {
	dblchk = true
}
// let's see if they want advertising disabled
if in.AntiLink == true && in.BotMaster == false {

	if dblchk == false {
		var deny [5]string
		deny[0] = "https://"
		deny[1] = "http://"
		deny[2] = ".com"
		deny[3] = ".net"
		deny[4] = "www."
		for i := 0; i <= 4; i++ {
			if strings.Contains(strings.ToLower(strings.ToLower(m.Content)), deny[i]) && in.Admin != m.Author.ID {
				s.ChannelMessageDelete(m.ChannelID, m.ID)
				s.ChannelTyping(m.ChannelID)
				time.Sleep(1000 * time.Millisecond)
				if in.Action == "kick" {
					newkick := strings.Replace(resp.AntiLinkKick, "{data}", "<@"+m.Author.ID+">", -1)
					s.ChannelMessageSend(m.ChannelID, newkick)
					s.GuildMemberDelete(c.GuildID, m.Author.ID)
				} // if they want a kick

				if in.Action == "ban" {
					newban := strings.Replace(resp.AntiLinkBan, "{data}", "<@"+m.Author.ID+">", -1)
					s.ChannelMessageSend(m.ChannelID, newban)
					s.GuildBanCreate(c.GuildID, m.Author.ID, 10)
				} // if they want a kick

				if in.Action == "warn" {
					newwarn := strings.Replace(resp.AntiLinkWarn, "{data}", "<@"+m.Author.ID+">", -1)
					s.ChannelMessageSend(m.ChannelID, newwarn)
					// s.GuildBanCreate(c.GuildID, m.Author.ID, 10)
				} // if they want a kick


			}
		} // end of for loop
	} // end of dbl check
} // end of anti link system















if strings.Contains(strings.ToLower(m.Content), "--name") && in.Admin == m.Author.ID {
	str := strings.Replace(m.Content, "--name ", "", -1)

	if str != "" {
		s.UserUpdate(s.State.User.Email, in.Bot, str, s.State.User.Avatar, "")
	}
}




if strings.Contains(strings.ToLower(m.Content), "--avatar") && in.Admin == m.Author.ID {
	str := strings.Replace(m.Content, "--avatar ", "", -1)
	if str != "" {
		s.UserUpdate(s.State.User.Email, in.Bot, s.State.User.Username, str, "")
	}
}



	if strings.HasPrefix(m.Content, in.Prefix + cmd.Meme) {
		var meme []string
		cnt := 0
		cnt = countLines("memes.txt")
	//	fmt.Println("Lines: " + strconv.Itoa(cnt))
		meme, err := readLines("memes.txt")
		if err == nil {
	//	fmt.Println("memes.txt has ")
		myrand := random(1, cnt)
		s.ChannelTyping(m.ChannelID)
		time.Sleep(1000 * time.Millisecond)
		s.ChannelMessageSend(m.ChannelID, meme[myrand])
		} // make sure err == nil
		if err != nil {
			fmt.Println(err)
		}
	}




	if strings.HasPrefix(m.Content, in.Prefix + cmd.Joke) {
		var meme []string
		cnt := 0
		cnt = countLines("jokes.txt")
		// fmt.Println("Lines: " + strconv.Itoa(cnt))
		meme, err := readLines("jokes.txt")
		if err == nil {
		myrand := random(1, cnt)
		s.ChannelTyping(m.ChannelID)
		time.Sleep(1000 * time.Millisecond)
		s.ChannelMessageSend(m.ChannelID, meme[myrand])
		}
		if err != nil {
			fmt.Println(err)
		}
	}




	if strings.HasPrefix(m.Content, in.Prefix + cmd.Kick) && in.BotMaster == true {
		str := strings.Replace(m.Content, in.Prefix + cmd.Kick + " ", "", -1)
		str = strings.Replace(str, "<@", "", -1)
		str = strings.Replace(str, ">", "", -1)
		// fmt.Println("the"+str+"string")
		s.ChannelTyping(m.ID)
		time.Sleep(1000 * time.Millisecond)
		newdata := strings.Replace(resp.Kick, "{data}", "<@"+str+">", -1)
		s.ChannelMessageSend(m.ChannelID, newdata)
		s.GuildMemberDelete(c.GuildID, str)
	}









	if strings.HasPrefix(m.Content, in.Prefix + cmd.Ban) && in.BotMaster == true {
		str := strings.Replace(m.Content, in.Prefix + cmd.Ban + " ", "", -1)
		str = strings.Replace(str, "<@", "", -1)
		str = strings.Replace(str, ">", "", -1)
		s.ChannelTyping(m.ID)
		time.Sleep(1000 * time.Millisecond)
		newdata := strings.Replace(resp.Ban, "{data}", "<@"+str+">", -1)
		s.ChannelMessageSend(m.ChannelID, newdata)
		s.GuildBanCreate(c.GuildID, str, 10)
	}











	if m.Author.ID == in.Admin && strings.HasPrefix(m.Content, in.Prefix + "status") {
		str := strings.Replace(m.Content, in.Prefix + "status ", "", -1)
		s.ChannelTyping(m.ChannelID)
		time.Sleep(1000 * time.Millisecond)
		s.ChannelMessageSend(m.ChannelID, resp.Status)
		s.UpdateStatus(0, str)
		time.Sleep(500 * time.Millisecond)
		s.UpdateStatus(0, str)
	}









if in.Admin == m.Author.ID && strings.HasPrefix(m.Content, in.Prefix + cmd.Giveme) {
str := strings.Replace(m.Content, in.Prefix + cmd.Giveme + " ", "", -1)
z, err := s.State.Member(c.GuildID, m.Author.ID) 
if err != nil {
return
}
		roles, err := s.GuildRoles(c.GuildID)
		if err == nil {
			for _, v := range roles {
    			if v.Name == str {
    				z.Roles = append(z.Roles, v.ID)
    				s.GuildMemberEdit(c.GuildID, m.Author.ID, z.Roles)
					s.ChannelTyping(c.GuildID)
					time.Sleep(1000 * time.Millisecond)
					newdata := strings.Replace(resp.Giveme, "{data}", "<@"+str+">", -1)
					s.ChannelMessageSend(c.GuildID, newdata)
    			}
			}
		}
	} // end of giveme command.





// ############## Pm's??

if c.GuildID == "" && in.Admin == "" {
	k, err := s.UserChannelCreate(m.Author.ID)
	if err == nil {
		s.ChannelTyping(k.ID)
		time.Sleep(1000 * time.Millisecond)

	fmt.Println("Install Worked! Author ID: " + m.Author.ID)
	in.Admin = m.Author.ID
	newConf := obj{
		Bot:			in.Bot,
		Admin:			in.Admin,
		Status:			in.Status,
		BotMaster:		in.BotMaster,
		BotCommander:	in.BotCommander,
		CmdsRun:		in.CmdsRun,
		Prefix:			in.Prefix,
		GreetMsg:		in.GreetMsg,
		ByeMsg:			in.ByeMsg,
		RoleSys:		in.RoleSys,
		Name:			in.Name,
		AntiLink:		in.AntiLink,
		Action:			in.Action,
		Silent:			in.Silent,
		HelpCmd:		in.HelpCmd,
		}
	b, err := json.Marshal(newConf)
	if err == nil {
		ioutil.WriteFile("config.json", b, 0777)
	}
	if err != nil {
		fmt.Println(err)
	}
s.ChannelMessageSend(k.ID, "You have sucessfully installed `AutoGo` check out `commands.json` to customize the command names. and check out `autoresponse.txt` to add mew auto responses!")

}
}



} // end of chkErr


} // ##########   END OF messageCreate










func GuildMemberAdd(s *discordgo.Session, m *discordgo.GuildMemberAdd) {

	var in obj
	var resp responses

	vfile, err := ioutil.ReadFile("config.json")
	json.Unmarshal(vfile, &in)


	// Load up the custom responses.
	rfile, err := ioutil.ReadFile("responses.json")
	if err != nil {
		return
	} else {
	json.Unmarshal(rfile, &resp)
	}



	// fmt.Println(in.RoleSys)
	roles, err := s.GuildRoles(m.GuildID)

		if in.GreetMsg != "" && in.GreetMsg != "off" {
			s.ChannelTyping(m.GuildID)
			time.Sleep(1000 * time.Millisecond)
			data := strings.Replace(in.GreetMsg, "{user}", "<@"+m.User.ID+">", -1)
			data = strings.Replace(data, "{/user}", m.User.Username, -1)
			s.ChannelMessageSend(m.GuildID, data)
		}


if m.User.Bot == false {
	if err == nil {
 		for _, v := range roles {
    		if v.Name == in.RoleSys {
    			if in.RoleSys != "" && in.RoleSys != "off" {
    				m.Roles = append(m.Roles, v.ID)
    				s.GuildMemberEdit(m.GuildID, m.User.ID, m.Roles)
    				if in.Silent == false {
    					newdata := strings.Replace(resp.AutoRoleMsg, "{data}", in.RoleSys, -1)
    					newdata = strings.Replace(newdata, "{user}", "<@"+m.User.ID+">", -1)
						s.ChannelMessageSend(m.GuildID, newdata)
					}
				}
    		}
  		}
	} // end of err == nil
}



if m.User.Bot == true {
	if err == nil {
 		for _, v := range roles {
    		if v.Name == in.BotAutoRole {
    			if in.BotAutoRole != "" && in.BotAutoRole != "off" {
    				m.Roles = append(m.Roles, v.ID)
    				s.GuildMemberEdit(m.GuildID, m.User.ID, m.Roles)
    				if in.Silent == false {
    					newdata := strings.Replace(resp.BotAutoRoleMsg, "{data}", in.RoleSys, -1)
    					newdata = strings.Replace(newdata, "{user}", "<@"+m.User.ID+">", -1)
						s.ChannelMessageSend(m.GuildID, newdata)
					}
				}
    		}
  		}
	} // end of err == nil
}





} // end of GuildMemberAdd















func GuildMemberRemove(s *discordgo.Session, m *discordgo.GuildMemberRemove) {
	var in obj
	vfile, err := ioutil.ReadFile("config.json")
	if err == nil {
		json.Unmarshal(vfile, &in)
	}

if in.ByeMsg != "" && in.ByeMsg != "off" {
// fmt.Println(m.GuildID, m.User)
s.ChannelTyping(m.GuildID)
time.Sleep(1000 * time.Millisecond)
data := strings.Replace(in.ByeMsg, "{user}", m.User.Username, -1)
s.ChannelMessageSend(m.GuildID, data)
}

}











func onReady(s *discordgo.Session, event *discordgo.Ready) {
	var in obj
	vfile, err := ioutil.ReadFile("config.json")
	if err == nil {
		json.Unmarshal(vfile, &in)
	}

	s.UpdateStatus(0, in.Status)
}









/* disabled until i make this a toggle feature.
func GuildRoleUpdate(s *discordgo.Session, m *discordgo.GuildRoleUpdate) {
s.ChannelTyping(m.GuildID)
time.Sleep(1000 * time.Millisecond)
s.ChannelMessageSend(m.GuildID, "Someone has edited the role: `"+m.Role.Name+"`")
}
*/

