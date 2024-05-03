package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/ulugbek0217/s21-info-bot/data"
	"github.com/ulugbek0217/s21-info-bot/users"
)

var db *pgx.Conn

func main() {
	LoadEnv("config/.env")
	var err error
	db, err = pgx.Connect(context.Background(), os.Getenv("DB_PATH"))
	if err != nil {
		panic(err)
	}
	defer db.Close(context.Background())

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithCallbackQueryDataHandler("item", bot.MatchTypePrefix, answer),
	}

	b, err := bot.New(os.Getenv("TOKEN"), opts...)

	if err != nil {
		panic(err)
	}

	b.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact, start)
	b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "lang", bot.MatchTypePrefix, setLang)
	b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "questions", bot.MatchTypeExact, questions_menu)
	b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "main_menu", bot.MatchTypeExact, main_menu)

	b.Start(ctx)
}

func LoadEnv(path string) {
	err := godotenv.Overload(path)
	if err != nil {
		log.Fatal("error loading .env: ", err)
	}
}

func start(ctx context.Context, b *bot.Bot, update *models.Update) {
	if !users.UserExists(db, update.Message.From.ID) {
		users.RegMin(db, update.Message.From.ID)
	}
	lang := users.GetLang(db, update.Message.From.ID)

	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   data.Greeting[lang],
		ReplyMarkup: &models.InlineKeyboardMarkup{
			InlineKeyboard: data.StartMenu[lang],
		},
	})
	if err != nil {
		log.Fatal(err)
	}
}

func main_menu(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: update.CallbackQuery.ID,
		ShowAlert:       false,
	})
	b.DeleteMessage(ctx, &bot.DeleteMessageParams{
		ChatID:    update.CallbackQuery.Message.Message.Chat.ID,
		MessageID: update.CallbackQuery.Message.Message.ID,
	})

	lang := users.GetLang(db, update.CallbackQuery.From.ID)
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.CallbackQuery.Message.Message.Chat.ID,
		Text:   data.HowCanHelp[lang],
		ReplyMarkup: &models.InlineKeyboardMarkup{
			InlineKeyboard: data.StartMenu[lang],
		},
	})
	if err != nil {
		log.Fatal(err)
	}
}

func questions_menu(ctx context.Context, b *bot.Bot, update *models.Update) {
	inlineKeyboard := &models.InlineKeyboardMarkup{
		InlineKeyboard: data.Quest_kbd[users.GetLang(db, update.CallbackQuery.From.ID)],
	}
	b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: update.CallbackQuery.ID,
		ShowAlert:       false,
	})
	b.DeleteMessage(ctx, &bot.DeleteMessageParams{
		ChatID:    update.CallbackQuery.Message.Message.Chat.ID,
		MessageID: update.CallbackQuery.Message.Message.ID,
	})
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      update.CallbackQuery.Message.Message.Chat.ID,
		Text:        data.ShowMenu[users.GetLang(db, update.CallbackQuery.From.ID)], //
		ReplyMarkup: inlineKeyboard,
	})
}

func answer(ctx context.Context, b *bot.Bot, update *models.Update) {
	user_id := update.CallbackQuery.From.ID
	b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: update.CallbackQuery.ID,
		ShowAlert:       false,
	})
	b.DeleteMessage(ctx, &bot.DeleteMessageParams{
		ChatID:    update.CallbackQuery.Message.Message.Chat.ID,
		MessageID: update.CallbackQuery.Message.Message.ID,
	})
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.CallbackQuery.Message.Message.Chat.ID,
		Text:   data.Info[users.GetLang(db, user_id)][update.CallbackQuery.Data],
		ReplyMarkup: &models.InlineKeyboardMarkup{
			InlineKeyboard: [][]models.InlineKeyboardButton{
				{
					{
						Text:         data.QuestionsList[users.GetLang(db, user_id)],
						CallbackData: "questions",
					},
				},
			},
		},
		LinkPreviewOptions: &models.LinkPreviewOptions{
			IsDisabled: bot.True(),
		},
		ParseMode: models.ParseModeHTML,
	})
}

func setLang(ctx context.Context, b *bot.Bot, update *models.Update) {
	var lang string = strings.Split(update.CallbackQuery.Data, "_")[1]
	tg_id := update.CallbackQuery.From.ID
	users.SetLang(db, lang, tg_id)

	b.EditMessageText(ctx, &bot.EditMessageTextParams{
		ChatID:    update.CallbackQuery.Message.Message.Chat.ID,
		MessageID: update.CallbackQuery.Message.Message.ID,
		Text:      fmt.Sprintf("%s\n%s", data.Greeting[lang], data.HowCanHelp[lang]),
		ReplyMarkup: &models.InlineKeyboardMarkup{
			InlineKeyboard: data.StartMenu[lang],
		},
	})
}
