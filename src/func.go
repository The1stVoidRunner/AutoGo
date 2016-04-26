package main


import (
  "math/rand"
  "time"
  "os"
  "bufio"
  "strings"
  "github.com/bwmarrin/discordgo"
)










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