package bot

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"net"
	"net/textproto"
	"strings"
	"time"
)

type MessageCallback func(msg *Message)

type Connection struct {
	conn net.Conn
	bot  *Bot

	onMessage map[MsgTyp][]MessageCallback
}

func NewConnection(owner *Bot) *Connection {
	conn := Connection{}
	conn.bot = owner
	conn.onMessage = make(map[MsgTyp][]MessageCallback, 4)
	return &conn
}

func (irc *Connection) Connect() *Connection {
	dialer := &net.Dialer{
		KeepAlive: time.Second * 10,
	}
	conf := &tls.Config{}
	conn, err := tls.DialWithDialer(dialer, "tcp", irc.bot.Settings.Server.Domain, conf)
	irc.conn = conn
	if err != nil {
		return irc
	}
	fmt.Printf("IRC Connected %s\n", irc.bot.Settings.Server.Domain)
	return irc
}

func (irc *Connection) Disconnect() *Connection {
	defer irc.conn.Close()
	fmt.Printf("IRC Disconnected %s\n", irc.bot.Settings.Server.Domain)
	return irc
}

func (irc *Connection) Join() *Connection {
	irc.write("PASS %s\r\n", irc.bot.Settings.Oauth)
	irc.write("NICK %s\r\n", irc.bot.Settings.Nickname)
	irc.write("JOIN %s\r\n", irc.bot.Settings.Channel)
	irc.write("CAP REQ :twitch.tv/tags\r\n")
	irc.write("CAP REQ :twitch.tv/commands\r\n")
	return irc
}

func (irc *Connection) On(msgType MsgTyp, callback MessageCallback) {
	irc.onMessage[msgType] = append(irc.onMessage[msgType], callback)
}

func (irc *Connection) readLoop() error {
	reader := bufio.NewReader(irc.conn)
	tp := textproto.NewReader(reader)
	for {
		fmt.Printf("Waiting -> ")
		line, err := tp.ReadLine()

		if err != nil {
			irc.handleError(err)
			return err
		}

		fmt.Printf("Read -> %s\n", line)

		switch {
		case irc.handlePing(line):
		case irc.handleUserstate(line):
		case irc.handleMessage(line):
		}
	}
}

func (irc *Connection) Loop() *Connection {

	irc.Join()

	defer func() {
		irc.Disconnect()
	}()

	for {
		err := irc.readLoop()
		if err != nil {
			break
		}
	}

	return irc
}

func (irc *Connection) handleError(err error) bool {
	fmt.Printf("ERROR -> %v\n", err)
	return true
}

func (irc *Connection) handleUserstate(line string) bool {
	if strings.Contains(line, "tmi.twitch.tv USERSTATE") {
		return true
	}

	return false
}

func (irc *Connection) handleMessage(line string) bool {
	if strings.HasPrefix(line, "@") {
		msg := parseMessage(line)
		handlers := irc.onMessage[msg.Type]
		for _, callback := range handlers {
			defer callback(msg)
			return true
		}
	}

	return false
}

func (irc *Connection) handlePing(line string) bool {

	if strings.Contains(line, "PING") {
		irc.write(strings.Replace(line, "PING", "PONG", 1))
		return true
	}

	return false
}

func (irc *Connection) SendChannel(message string, args ...interface{}) {
	irc.Send(irc.bot.Settings.Channel, message, args...)
}

func (irc *Connection) SendPrivate(receiver string, message string, args ...interface{}) {
	irc.Send(receiver, message, args...)
}

func (irc *Connection) Send(receiver string, message string, args ...interface{}) *Connection {
	compiledMessage := fmt.Sprintf(message, args...)
	preparedMessage := fmt.Sprintf("PRIVMSG %s :%s\r\n", receiver, compiledMessage)
	irc.write(preparedMessage)

	return irc
}

func (irc *Connection) write(message string, args ...interface{}) *Connection {
	bytesWritten, err := irc.conn.Write([]byte(fmt.Sprintf(message, args...)))
	if err != nil {
		irc.handleError(err)
	}
	fmt.Printf("Written [%v] -> %s", bytesWritten, fmt.Sprintf(message, args...))
	return irc
}
