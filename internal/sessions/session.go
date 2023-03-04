package sessions

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/edy4c7/go-discord-bot/internal/bots"
	"github.com/joho/godotenv"
	gogpt "github.com/sashabaranov/go-gpt3"
)

type sessionType int

const (
	Discord sessionType = iota + 1
)

func NewSession(sessionType sessionType) error {
	godotenv.Load()

	openaiToken := os.Getenv("OPENAI_TOKEN")
	gpt := gogpt.NewClient(openaiToken)
	bot := bots.NewChatGptBot(gpt)

	listening := make(chan any, 1)
	go func() {
		// 終了シグナルを待ち受ける
		sc := make(chan os.Signal, 1)
		signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
		<-sc
		close(listening)
	}()

	fmt.Println("Chat bot is listening...")
	discordToken := os.Getenv("DISCORD_TOKEN")
	NewDiscordSession(discordToken, bot, listening)
	fmt.Println("Chat bot is shutdown")

	return nil
}
