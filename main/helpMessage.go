package main

import (
	"github.com/bwmarrin/discordgo"
)

var helpMessageEmbeds = []*discordgo.MessageEmbedField{
	{
		Name: "カスタム絵文字作成",
		Value: "アカネチャンがカスタム絵文字を作ってくれます．\n" +
			"`TEXT`:   絵文字にしたい文字列\n" +
			"`ALIAS`:  絵文字のエイリアス\n" +
			"`URL`:    画像のURL\n" +
			"`COLOR`:  文字列の色  CSSで使える色名もしくは16進カラーコード(省略可 3桁/6桁 デフォルト: 黒)\n" +
			"`TRANSP`: 背景を透明にするかどうか(省略可 true/false デフォルト: false)\n" +
			"`@Akane emoji        ALIAS TEXT COLOR TRANSP`: 新しいカスタム絵文字を作ります\n" +
			"`@Akane emoji image  ALIAS`:              画像からカスタム絵文字を作ります．画像投稿時のコメントにコマンドを入力してください\n" +
			"`@Akane emoji url    ALIAS URL`:          画像のURLからカスタム絵文字を作ります\n" +
			"`@Akane emoji delete ALIAS`:              ALIASを指定してカスタム絵文字を削除します",
	},
	{
		Name: "Ping",
		Value: "`@Akane ping`\n" +
			"と打つと元気よく Pong! と返事をします",
	},
	{
		Name: "ヘルプ",
		Value: "使い方が分からない? そんなときは\n" +
			"`@Akane help`\n" +
			"と打ってみましょう! きっとすぐに使えるようになりますよ!",
	},
	{
		Name: "Good Job",
		Value: "`@Akane goodjob`\n" +
			"と打つと返事をします",
	},
}
