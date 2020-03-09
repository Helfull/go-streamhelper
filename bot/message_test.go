package bot

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

const (
	TestLine   = "@badges=broadcaster/1,premium/1;color=#FF0000;display-name=Helfull;emotes=;id=9ce8e195-cc7e-4e36-89d4-3f16f4dca2c9;mod=0;room-id=31870168;sent-ts=1515872523296;subscriber=0;tmi-sent-ts=1515872525337;turbo=0;user-id=31870168;user-type= :helfull!helfull@helfull.tmi.twitch.tv PRIVMSG #helfull :test message"
	TestMiddle = "helfull!helfull@helfull.tmi.twitch.tv PRIVMSG #helfull"
	TestTags   = "@badges=broadcaster/1,premium/1;color=#FF0000;display-name=Helfull;emotes=;id=9ce8e195-cc7e-4e36-89d4-3f16f4dca2c9;mod=0;room-id=31870168;sent-ts=1515872523296;subscriber=0;tmi-sent-ts=1515872525337;turbo=0;user-id=31870168;user-type="
)

func TestMessage(t *testing.T) {
	Convey("parseMiddle", t, func() {
		username, msgType, channel := parseMiddle(TestMiddle)
		Convey("It should extract the username", func() {
			So(username, ShouldEqual, "helfull")
		})
		Convey("It should extract the msgType", func() {
			So(msgType, ShouldEqual, MessagePrivate)
		})
		Convey("It should extract the channel", func() {
			So(channel, ShouldEqual, "#helfull")
		})
	})

	Convey("getMsgType", t, func() {
		So(getMsgType("PRIVMSG"), ShouldEqual, MessagePrivate)
		So(getMsgType("WHISPER"), ShouldEqual, MessageWhisper)
		So(getMsgType("CLEARCHAT"), ShouldEqual, MessageClearChat)
		So(getMsgType("SOME RANDOM STUFF"), ShouldEqual, MessageUnknown)
	})

	Convey("parseTags", t, func() {
		resultTags := parseTags(TestTags)

		Convey("It returns a map of tags equal to tags in the send message", func() {
			expectedType := make(map[string]string)
			So(resultTags, ShouldHaveSameTypeAs, expectedType)
			So(resultTags, ShouldHaveLength, 13)
		})

		Convey("It contains key value pairs", func() {
			So(resultTags, ShouldContainKey, "display-name")
			So(resultTags["display-name"], ShouldEqual, "Helfull")
		})
	})

	Convey("parseMessage", t, func() {
		result := parseMessage(TestLine)
		Convey("It should extract tags", func() {
			So(result.Tags, ShouldContainKey, "badges")
			So(result.Tags["badges"], ShouldEqual, "broadcaster/1,premium/1")
		})
		Convey("It should extract the message type", func() {
			So(result.Type, ShouldEqual, MessagePrivate)
		})
		Convey("It should have an instance of user", func() {
			user := &User{"helfull", "helfull"}
			So(result.User, ShouldHaveSameTypeAs, user)
			So(result.User.id, ShouldEqual, "helfull")
			So(result.User.Nickname, ShouldEqual, "Helfull")
		})
	})
}
