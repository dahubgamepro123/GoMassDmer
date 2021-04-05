package main
import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os/signal"
	"syscall"
	"github.com/fatih/color"
	"os"
	"strings"
	"time"

)

func banner() {
	color.Red(` 
			███╗   ███╗ █████╗ ███████╗███████╗██████╗ ███╗   ███╗███████╗██████╗ 
			████╗ ████║██╔══██╗██╔════╝██╔════╝██╔══██╗████╗ ████║██╔════╝██╔══██╗
			██╔████╔██║███████║███████╗███████╗██║  ██║██╔████╔██║█████╗  ██████╔╝
			██║╚██╔╝██║██╔══██║╚════██║╚════██║██║  ██║██║╚██╔╝██║██╔══╝  ██╔══██╗
			██║ ╚═╝ ██║██║  ██║███████║███████║██████╔╝██║ ╚═╝ ██║███████╗██║  ██║
			╚═╝     ╚═╝╚═╝  ╚═╝╚══════╝╚══════╝╚═════╝ ╚═╝     ╚═╝╚══════╝╚═╝  ╚═╝
						Made by lxi#1400
	`)
}

func Clear() {
	fmt.Print("\033[H\033[2J")

}


func main() {
	Clear()
	fmt.Print("Insert Bot Token > ")
	fmt.Scanln(&token)
	fmt.Print("\nInsert a ID to be able to run the dmall command > ")
	fmt.Scanln(&whitelistedID)
	Clear()
	lxiontop, err := discordgo.New(fmt.Sprintf("Bot %s", token))
	if err != nil {
		fmt.Println("[ERROR]", err)
		return
	}
	lxiontop.AddHandler(CommandHandler)

	
	lxiontop.AddHandler(func(s *discordgo.Session, event *discordgo.Ready) {
		banner()
		fmt.Printf("[STATUS] Connected to: %s%s, and ready to mass-dm!\nType .dmall [message] to start!\n", s.State.User.Username, s.State.User.Discriminator)
		lxiontop.UpdateStreamingStatus(2, "made by lxi <3", "https://twitch.tv/wizzedyou")
	})

	lxiontop.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsAll)
	err = lxiontop.Open()

	if err != nil {
		fmt.Println("[ERROR]", err)
		return
	}
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	lxiontop.Close()
}





func CommandHandler(s *discordgo.Session, m *discordgo.MessageCreate) { 
	if m.Author.ID == s.State.User.ID {
		return
	}
	if strings.Contains(m.Content, ".dmall") {
		if m.Author.ID != whitelistedID {
			return
		}
		message := strings.Trim(m.Content, ".dmall")
		members, fetchmemberserror := s.GuildMembers(m.GuildID, "0", 1000)
		if fetchmemberserror != nil {
			fmt.Println(fetchmemberserror)
			return
		}
		for _, member := range members {
			if member.User.ID == s.State.User.ID {
				return
			}
			channel, CreateChannelError := s.UserChannelCreate(member.User.ID)
			if CreateChannelError != nil {
				fmt.Println(CreateChannelError)
				continue;
			}
			_, SendMessageError := s.ChannelMessageSend(channel.ID, message)
			if SendMessageError != nil {
				color.Red(fmt.Sprintf("Error: %s", SendMessageError))
				continue;
			}
			color.Green(fmt.Sprintf("Messaged %s#%s", member.User.Username, member.User.Discriminator))
			time.Sleep(1)
		}
	}
}




var (
	token string
	whitelistedID string
)
