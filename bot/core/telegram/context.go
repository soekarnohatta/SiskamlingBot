package telegram

import (
	"strings"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2"
)

// TgContext contains the context information used to invoke a command.
type TgContext struct {
	Bot     *gotgbot.Bot
	Context *ext.Context

	// Convenience fields
	Message  *gotgbot.Message
	Chat     *gotgbot.Chat
	User     *gotgbot.User
	Callback *gotgbot.CallbackQuery

	// Miscellaneous
	CmdSegment string
	TimeInit   time.Duration
	TimeProc   time.Duration

	// Lazy values
	args        []string
	haveRawArgs bool
	rawArgs     string
}

func newContext(bot *gotgbot.Bot, ctx *ext.Context, cmdSeg string) *TgContext {
	newTgContext := &TgContext{
		Bot:        bot,
		Context:    ctx,
		CmdSegment: cmdSeg,
	}

	// use EffectiveMessage as it already handles all possible updates
	newTgContext.Message = ctx.EffectiveMessage
	newTgContext.User = newTgContext.Message.From
	newTgContext.Chat = &newTgContext.Message.Chat
	newTgContext.TimeInit = time.Now().UTC().Sub(time.Unix(newTgContext.Message.Date, 0).UTC())

	if ctx.Update.CallbackQuery != nil || ctx.CallbackQuery != nil {
		newTgContext.Callback = ctx.Update.CallbackQuery
		return newTgContext
	}

	return newTgContext
}

// Args returns a slice of whitespace-separated arguments from the command message.
func (c *TgContext) Args() []string {
	if c.args == nil {
		c.args = strings.Fields(c.Message.Text)[1:]
	}

	return c.args
}

// RawArgs returns a string with everything in the command message except the command invocation segment.
func (c *TgContext) RawArgs() string {
	if !c.haveRawArgs {
		c.rawArgs = strings.TrimSpace(c.Message.Text[len(c.CmdSegment):])
		c.haveRawArgs = true
	}

	return c.rawArgs
}
