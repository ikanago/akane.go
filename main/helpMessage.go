package main

type HelpMessage struct {
	name  string
	value string
}

var helpMessages = map[string]HelpMessage{
	"emoji": {
		name: "カスタム絵文字作成",
		value: "アオイチャンがカスタム絵文字を作ってくれます．\n" +
			"`TEXT`:   絵文字にしたい文字列\n" +
			"`ALIAS`:  絵文字のエイリアス\n" +
			"`URL`:    画像のURL\n" +
			"`COLOR`:  文字列の色  CSSで使える色名もしくは16進カラーコード(3桁/6桁)(省略した場合は白になります)\n" +
			"`TRANSP`: 背景を透明にするかどうか省略可('true'を指定すると背景が透明になります)\n" +
			"`@Aoi emoji TEXT   ALIAS COLOR TRANSP`: 新しいカスタム絵文字を作ります\n" +
			"`@Aoi emoji image  ALIAS`:              画像からカスタム絵文字を作ります．画像投稿時のコメントにコマンドを入力してください\n" +
			"`@Aoi emoji url    ALIAS URL`:          画像のURLからカスタム絵文字を作ります\n" +
			"`@Aoi emoji delete ALIAS`:              ALIASを指定してカスタム絵文字を削除します．",
	},
	"help": {
		name: "ヘルプ",
		value: "使い方が分からない? そんなときは\n" +
			"`@Aoi help`\n" +
			"と打ってみましょう! きっとすぐに使えるようになりますよ!",
	},
}
