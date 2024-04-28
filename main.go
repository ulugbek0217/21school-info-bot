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

// type user struct {
// 	id      int
// 	tg_id   string
// 	name    string
// 	surname string
// 	phone   string
// 	role    string
// 	lang    string
// }

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
	b.RegisterHandler(bot.HandlerTypeMessageText, "Показать меню", bot.MatchTypeExact, menu)
	b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "lang", bot.MatchTypePrefix, setLang)

	b.Start(ctx)
}

func LoadEnv(path string) {
	err := godotenv.Overload(path)
	if err != nil {
		log.Fatal("error loading .env: ", err)
	}
}

func start(ctx context.Context, b *bot.Bot, update *models.Update) {
	if users.UserExists(db, update.Message.From.ID) {
		users.RegMin(db, update.Message.From.ID)
	}
	lang := users.GetLang(db, update.Message.From.ID)
	kb := &models.ReplyKeyboardMarkup{
		Keyboard: [][]models.KeyboardButton{
			{
				{
					Text: data.ShowMenu[lang],
				},
			},
		},
		ResizeKeyboard: true,
	}

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        data.Greeting[lang],
		ReplyMarkup: kb,
	})

	lang_kb := &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{
					Text: "O'zbek 🇺🇿", CallbackData: "lang_uz",
				},
			},
			{
				{
					Text: "Русский 🇷🇺", CallbackData: "lang_ru",
				},
			},
		},
	}

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        "Выберите язык",
		ReplyMarkup: lang_kb,
	})
}

func menu(ctx context.Context, b *bot.Bot, update *models.Update) {
	inlineKeyboard := &models.InlineKeyboardMarkup{
		InlineKeyboard: data.Quest_kbd,
	}

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        "Меню вопросов",
		ReplyMarkup: inlineKeyboard,
	})
}

func answer(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: update.CallbackQuery.ID,
		ShowAlert:       false,
	})

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.CallbackQuery.Message.Message.Chat.ID,
		Text:   data.Info[users.GetLang(db, update.CallbackQuery.From.ID)][update.CallbackQuery.Data],
		LinkPreviewOptions: &models.LinkPreviewOptions{
			IsDisabled: bot.True(),
		},
		ParseMode: models.ParseModeHTML,
	})
}

func setLang(ctx context.Context, b *bot.Bot, update *models.Update) {
	var lang string = strings.Split(update.CallbackQuery.Data, "_")[1]
	_, err := db.Exec(context.Background(), "update users set lang = $1 where tg_id = $2", lang,
		fmt.Sprintf("%d", update.CallbackQuery.From.ID))
	if err != nil {
		panic(err)
	}
	fmt.Println(lang)

	b.EditMessageText(ctx, &bot.EditMessageTextParams{
		ChatID:    update.CallbackQuery.Message.Message.Chat.ID,
		MessageID: update.CallbackQuery.Message.Message.ID,
		Text:      fmt.Sprintf("Set to %v", lang),
	})
}
