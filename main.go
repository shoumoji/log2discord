package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"unicode/utf8"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	// DiscordのBotトークンとログチャンネルを.envから取得
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	Token := os.Getenv("TOKEN")
	logChannelID := os.Getenv("logChannelID")
	// どのディレクトリをDiscordに出力するか引数で選択
	logDir := os.Args[1]

	//discordと接続
	discord, err := discordgo.New("Bot " + Token)
	if err != nil {
		log.Fatal(err)
	}
	err = discord.Open()
	if err != nil {
		log.Fatal(err)
	}

	err = filepath.Walk(logDir, func(file string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("prevent panic by handling failure accessing a path %q: %v\n", file, err)
		}
		sendMessageDiscordLogChannel(discord, logChannelID, file)
		return nil
	})
}

func sendMessageDiscordLogChannel(discord *discordgo.Session, logChannelID string, path string) {
	if utf8.RuneCountInString(event.Name+"\n\n"+readData) <= 1999 {
		// logディレクトリに追加されたファイルをdiscordに送信
		discord.ChannelMessageSend(logChannelID, event.Name+"\n\n"+readData)
	} else {
		// 2000文字以上は一括で送れないので、分割して送信
		discord.ChannelMessageSend(logChannelID, "----------文字数が多いため分割します----------")
		discord.ChannelMessageSend(logChannelID, event.Name) //ファイルの絶対パス
		splitlen := 1999
		runes := []rune(readData)
		for i := 0; i < len(runes); i += splitlen {
			if i+splitlen < len(runes) {
				discord.ChannelMessageSend(logChannelID, string(runes[i:(i+splitlen)]))
			} else {
				discord.ChannelMessageSend(logChannelID, string(runes[i:]))
			}
		}
		discord.ChannelMessageSend(logChannelID, "-----------分割終了----------")
	}
}
