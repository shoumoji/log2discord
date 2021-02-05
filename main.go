package main

import (
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

func main() {
	// DiscordのBotトークンとログチャンネルを.envから取得
	Token := ""
	logChannelID := ""
	// どのディレクトリをDiscordに出力するか引数で選択
	sendFile := os.Args[1]

	//discordと接続
	discord, err := discordgo.New("Bot " + Token)
	if err != nil {
		log.Fatal(err)
	}
	err = discord.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer discord.Close()

	f, err := os.Open(sendFile)
	if err != nil {
		println(err)
	}
	defer f.Close()

	data := &discordgo.MessageSend{
		Files: []*discordgo.File{
			{
				Name:   sendFile,
				Reader: f,
			},
		},
	}

	discord.ChannelMessageSendComplex(logChannelID, data)
}
