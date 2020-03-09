package bot

import (
	"regexp"
	"strings"
)

type MsgTyp int

const (
	MessageDelimiter = " :"
	MessageUnknown   = 0
	MessagePrivate   = 1
	MessageWhisper   = 2
	MessageClearChat = 3
)

type Message struct {
	User    *User
	Tags    map[string]string
	raw     string
	Type    MsgTyp
	Channel string
	Text    string
}

func parseMessage(line string) *Message {
	msg := &Message{
		raw: line,
	}
	splitLine := strings.SplitN(line, MessageDelimiter, 3)

	msg.Tags = parseTags(splitLine[0])
	var username string
	username, msg.Type, msg.Channel = parseMiddle(splitLine[1])

	msg.User = &User{
		id:       username,
		Nickname: msg.Tags["display-name"],
	}

	msg.Text = splitLine[2]

	return msg
}

func parseMiddle(rawMiddle string) (username string, msgType MsgTyp, channel string) {

	re := regexp.MustCompile(`^(?P<username>[0-9A-Za-z][\w]{3,24})![0-9A-Za-z][\w]{3,24}@[0-9A-Za-z][\w]{3,24}.tmi.twitch.tv\s(?P<msgtyp>[A-Z]{1,})\s(?P<channel>#[0-9A-Za-z][\w]{3,24})$`)
	if re.MatchString(rawMiddle) {
		n1 := re.SubexpNames()
		r2 := re.FindAllStringSubmatch(rawMiddle, -1)[0]
		md := map[string]string{}
		for i, n := range r2 {
			md[n1[i]] = n
		}

		username, msgType, channel = md["username"], getMsgType(md["msgtyp"]), md["channel"]
	}
	return
}

func getMsgType(rawType string) MsgTyp {
	switch rawType {
	case "PRIVMSG":
		return MessagePrivate
	case "WHISPER":
		return MessageWhisper
	case "CLEARCHAT":
		return MessageClearChat
	default:
		return MessageUnknown
	}
}

func parseTags(rawTags string) map[string]string {

	mapTags := make(map[string]string)

	tags := strings.Split(strings.TrimLeft(rawTags, "@"), ";")

	for _, tag := range tags {
		splitTag := strings.Split(tag, "=")
		mapTags[splitTag[0]] = splitTag[1]
	}

	return mapTags
}
