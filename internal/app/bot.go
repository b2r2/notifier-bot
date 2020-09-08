package app

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
)

// BotAPI ...
type BotAPI struct {
	config *Config
	logger *logrus.Logger
	bot    *tgbotapi.BotAPI
}

// New ...
func New(config *Config) *BotAPI {
	return &BotAPI{
		config: config,
		logger: logrus.New(),
	}
}

// Run ...
func (b *BotAPI) Run() error {
	if err := b.configureLogger(); err != nil {
		return err
	}

	if err := b.configureBot(); err != nil {
		return err
	}

	// others configure ...

	logrus.Infof("Authorized on account %s, debuging mode %t", b.bot.Self.UserName, b.config.BotLogLevel)
	return nil
}

func (b *BotAPI) configureBot() error {
	bot, err := tgbotapi.NewBotAPI(b.config.Token)
	if err != nil {
		return err
	}

	bot.Debug = b.config.BotLogLevel

	b.bot = bot
	return nil
}

func (b *BotAPI) configureLogger() error {
	lvl, err := logrus.ParseLevel(b.config.LogLevel)
	if err != nil {
		return err
	}
	b.logger.SetLevel(lvl)
	return nil
}
