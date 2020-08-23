package main

import (
	"github.com/bwmarrin/discordgo"
)

var HelpMessageEmbeds = []*discordgo.MessageEmbedField{
	{
		Name: "カスタム絵文字作成",
		Value: "アオイチャンがカスタム絵文字を作ってくれます．\n" +
			"`TEXT`:   絵文字にしたい文字列\n" +
			"`ALIAS`:  絵文字のエイリアス\n" +
			"`URL`:    画像のURL\n" +
			"`COLOR`:  文字列の色  CSSで使える色名もしくは16進カラーコード(3桁/6桁)(省略した場合は白になります)\n" +
			"`TRANSP`: 背景を透明にするかどうか省略可('true'を指定すると背景が透明になります)\n" +
			"`@Akane emoji TEXT   ALIAS COLOR TRANSP`: 新しいカスタム絵文字を作ります\n" +
			"`@Akane emoji image  ALIAS`:              画像からカスタム絵文字を作ります．画像投稿時のコメントにコマンドを入力してください\n" +
			"`@Akane emoji url    ALIAS URL`:          画像のURLからカスタム絵文字を作ります\n" +
			"`@Akane emoji delete ALIAS`:              ALIASを指定してカスタム絵文字を削除します．",
	},
	{
		Name: "ヘルプ",
		Value: "使い方が分からない? そんなときは\n" +
			"`@Akane help`\n" +
			"と打ってみましょう! きっとすぐに使えるようになりますよ!",
	},
}

func getHelpMessage() *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Type:   discordgo.EmbedTypeRich,
		Title:  "アカネチャンのコマンド",
		Fields: HelpMessageEmbeds,
	}
}

func CreateEmojiFromText() (message string, err error) {
	return "emoji", nil
}
