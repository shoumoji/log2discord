# log2discord
## 概要
ファイルをDiscordの特定のチャンネルに送信するためのツールです。

## 使い方
.env以下に

```.env
TOKEN=fasdkglake9faefasedasfasf
logChannelID=34141343151
```

といった形でDiscord Botのtokenと、ファイルを送るDiscord ChannelIDを保存します。

あとは

```go:log2discord
./log2discord filename
```

とすれば、対象のチャンネルにファイルが送信されます。
