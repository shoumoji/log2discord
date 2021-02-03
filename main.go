package main

import (
	"io/ioutil"
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

	err = filepath.Walk(logDir, func(path string, info os.FileInfo, err error) {

	})
	//discordと接続
	discord, err := discordgo.New("Bot " + Token)
	if err != nil {
		log.Fatal(err)
	}
	err = discord.Open()
	if err != nil {
		log.Fatal(err)
	}

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

func dirWalk(dir string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			paths = append(paths, dirWalk(filepath.Join(dir, file.Name()))...)
			continue
		}
		paths = append(paths, filepath.Join(dir, file.Name()))
	}
	return paths
}
