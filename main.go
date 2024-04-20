package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/ulugbek0217/s21-info-bot/data"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(start),
	}

	b, err := bot.New("5559681653:AAFbKDgbqpUHoycIhGEhoYs10uvCEDHtkZg", opts...)

	if err != nil {
		log.Fatal(err)
	}

	b.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact, start)
	b.RegisterHandler(bot.HandlerTypeMessageText, "Показать меню", bot.MatchTypeExact, menu)

	b.Start(ctx)
}

func start(ctx context.Context, b *bot.Bot, update *models.Update) {
	kb := &models.ReplyKeyboardMarkup{
		Keyboard: [][]models.KeyboardButton{
			{
				{
					Text: "Показать меню",
				},
			},
		},
		ResizeKeyboard: true,
	}

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        "Привет👋",
		ReplyMarkup: kb,
	})
}

func menu(ctx context.Context, b *bot.Bot, update *models.Update) {
	inlineKeyboard := &models.InlineKeyboardMarkup{
		InlineKeyboard: data.Questions,
	}

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        "Меню вопросов",
		ReplyMarkup: inlineKeyboard,
	})
}
