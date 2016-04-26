package main


import (
  "math/rand"
  "time"
  "os"
  "bufio"
  "strings"
  "github.com/bwmarrin/discordgo"
)

















func AutoResponseSystem(GuildID string, isMaster bool, prefix string, s *discordgo.Session, m *discordgo.MessageCreate) string {
reply := ""

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
      dokick := false
      doban := false
  //  fmt.Println("RAW: " + ars)
    ardat := strings.Split(ars, "=")
    trigger := ardat[0]
    response := strings.Replace(ars, trigger+"=", "", -1)
  //  response := ardat[1]
    response = strings.Replace(response, "{user}", "<@"+m.Author.ID+">", -1)
    response = strings.Replace(response, "{/user}", m.Author.Username, -1)
    
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

  if isMaster == false {
    if strings.Contains(response, "{kick}") {
      dokick = true
    }

    if strings.Contains(response, "{ban}") {
      doban = true
    }
  }


    if strings.HasPrefix(trigger, "&") {
      cn++
      isfind = true
  //    fmt.Println("Found: &")
      trigger = strings.Replace(trigger, "&", "", -1)
    }

  if strings.HasPrefix(trigger, prefix) {
  //    fmt.Println("Found: "+in.Prefix)
    //  isfind = true
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

    if m.Content == trigger {
      if ispm == false {
        s.ChannelTyping(m.ID)
        time.Sleep(1000 * time.Millisecond)
      // s.ChannelMessageSend(m.ChannelID, response)
        reply = response
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

      s.ChannelTyping(m.ID)
      time.Sleep(1000 * time.Millisecond)
      // s.ChannelMessageSend(m.ChannelID, response)
      reply = response
    }
    } // end of cntrole == 0
    if dokick == true {
      response = strings.Replace(response, "{kick}", "", -1)
      s.GuildMemberDelete(GuildID, m.Author.ID)
    }

    if doban == true {
      response = strings.Replace(response, "{ban}", "", -1)
      s.GuildBanCreate(GuildID, m.Author.ID, 10)
    }

    } // end of dont == false
  } // end of for loop
} // check to see if they have autoresponse.txt file in bot dir.
}


return reply
}











// displays random integer
func random(min, max int) int {
    rand.Seed(time.Now().Unix())
    return rand.Intn(max - min) + min
}





func guildInfo(param string, s *discordgo.Session) (string, error) {
  var opt string
  guilds, err := s.UserGuilds()
    if err != nil {
      return opt, err
    }
   if err == nil {

    // name param was triggered.
    if strings.ToLower(param) == "name" {
      opt = guilds[0].ID
    }

    // ID param was triggered
    if strings.ToLower(param) == "id" {
      opt = guilds[0].Name
    }


   }
   return opt, err
}











// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
  file, err := os.Open(path)
  if err != nil {
    return nil, err
  }
  defer file.Close()

  var lines []string
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }
  return lines, scanner.Err()
}







// readLines reads a whole file into memory
// and returns a slice of its lines.
func countLines(path string) int {
  counter := 0

  file, err := os.Open(path)
  if err != nil {

  }
  defer file.Close()
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    counter++
  }
  return counter
}












func GetRoleID(s *discordgo.Session, guildID string, role string) string {
  var re string
  roles, err := s.GuildRoles(guildID)
  if err == nil {
    for _, v := range roles {
      if v.Name == role {
        re = v.ID
      }
    }
  }
  return re
}










/* caused a memory error :( will work into it more later.
func ServerID(s *discordgo.Session, ChannelID string) string {
  c, err := s.State.Channel(ChannelID)
if err != nil {
 // channel not found
}
  return c.GuildID
}
*/












func isMemberRole(s *discordgo.Session, GuildID string, AuthorID string, role string) bool {
  var opt bool
  opt = false
z, err := s.State.Member(GuildID, AuthorID) 
if err != nil {
z, err = s.GuildMember(GuildID, AuthorID)
}

if err == nil {
  var l []string
  l = z.Roles
  for r := range z.Roles {
    if strings.Contains(l[r], GetRoleID(s, GuildID, role)) {
 //     fmt.Println("Found the role!"+l[r])
      opt = true
    }
  }
}
  return opt
}