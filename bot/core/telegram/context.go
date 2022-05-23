package telegram

import (
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

type TgContext struct {
	haveRawArgs bool
	Date        int64
	CmdSegment  string
	TimeInit    string
	rawArgs     string
	args        []string

	Bot      *gotgbot.Bot
	Context  *ext.Context
	Message  *gotgbot.Message
	Chat     *gotgbot.Chat
	User     *gotgbot.User
	Callback *gotgbot.CallbackQuery

	sync.Mutex
}

func NewContext(bot *gotgbot.Bot, ctx *ext.Context, cmdSeg string) *TgContext {
	newTgContext := &TgContext{
		Bot:        bot,
		Context:    ctx,
		CmdSegment: cmdSeg,
	}

	// use EffectiveMessage as it already handles all possible updates
	newTgContext.Message = ctx.EffectiveMessage
	newTgContext.User = ctx.EffectiveUser
	newTgContext.Chat = ctx.EffectiveChat
	newTgContext.Callback = ctx.Update.CallbackQuery
	newTgContext.Date = ctx.EffectiveMessage.Date

	secs := time.Since(time.Unix(newTgContext.Date, 0)).Seconds()
	newTgContext.TimeInit = strconv.FormatFloat(secs, 'f', 3, 64)
	return newTgContext
}

func (c *TgContext) Args() []string {
	if c.args == nil {
		c.args = strings.Fields(c.Message.Text)[1:]
	}

	if c.args == nil {
		c.args[0] = "Not Specified"
	}

	return c.args
}

func (c *TgContext) RawArgs() string {
	if !c.haveRawArgs {
		c.rawArgs = strings.TrimSpace(c.Message.Text[len(c.CmdSegment):])
		c.haveRawArgs = true
	}

	return c.rawArgs
}
