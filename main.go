package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
	"os"
	"os/signal"
	"strings"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
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
	defer func(db *pgx.Conn, ctx context.Context) {
		err := db.Close(ctx)
		if err != nil {

		}
	}(db, context.Background())

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithMessageTextHandler("/start", bot.MatchTypeExact, start),
		bot.WithCallbackQueryDataHandler("item", bot.MatchTypePrefix, answer),
		bot.WithCallbackQueryDataHandler("lang", bot.MatchTypePrefix, setLang),
		bot.WithCallbackQueryDataHandler("questions", bot.MatchTypeExact, questionsMenu),
		bot.WithCallbackQueryDataHandler("main_menu", bot.MatchTypeExact, mainMenu),
		bot.WithCallbackQueryDataHandler("ask_group", bot.MatchTypeExact, askGroup),
		bot.WithCallbackQueryDataHandler("contacts", bot.MatchTypeExact, contacts),
		bot.WithMiddlewares(),
	}

	b, err := bot.New(os.Getenv("TOKEN"), opts...)

	if err != nil {
		panic(err)
	}

	b.Start(ctx)
}

// LoadEnv serves to load the .env files from the given directory
func LoadEnv(path string) {
	err := godotenv.Overload(path)
	if err != nil {
		log.Fatal("error loading .env: ", err)
		return
	}
}

// start serves as an entry point of the bot
func start(ctx context.Context, b *bot.Bot, update *models.Update) {
	if !users.UserExists(db, update.Message.From.ID) {
		users.RegMin(db, update.Message.From.ID)
	}
	lang := users.GetLang(db, update.Message.From.ID)

	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   data.Greeting[lang],
		ReplyMarkup: &models.InlineKeyboardMarkup{
			InlineKeyboard: data.MainMenu[lang],
		},
	})
	if err != nil {
		log.Fatal(err)
	}
}

// mainMenu returns the main menu buttons
func mainMenu(ctx context.Context, b *bot.Bot, update *models.Update) {
	_, err := b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: update.CallbackQuery.ID,
		ShowAlert:       false,
	})
	if err != nil {
		log.Println(err)
		return
	}
	_, err = b.DeleteMessage(ctx, &bot.DeleteMessageParams{
		ChatID:    update.CallbackQuery.Message.Message.Chat.ID,
		MessageID: update.CallbackQuery.Message.Message.ID,
	})
	if err != nil {
		log.Println(err)
		return
	}
	lang := users.GetLang(db, update.CallbackQuery.From.ID)
	_, err = b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.CallbackQuery.Message.Message.Chat.ID,
		Text:   data.HowCanHelp[lang],
		ReplyMarkup: &models.InlineKeyboardMarkup{
			InlineKeyboard: data.MainMenu[lang],
		},
	})
	if err != nil {
		log.Println(err)
		return
	}
}

// questionsMenu returns the faq buttons
func questionsMenu(ctx context.Context, b *bot.Bot, update *models.Update) {
	inlineKeyboard := &models.InlineKeyboardMarkup{
		InlineKeyboard: data.QuestKbd[users.GetLang(db, update.CallbackQuery.From.ID)],
	}
	_, err := b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: update.CallbackQuery.ID,
		ShowAlert:       false,
	})
	if err != nil {
		log.Println(err)
		return
	}
	_, err = b.DeleteMessage(ctx, &bot.DeleteMessageParams{
		ChatID:    update.CallbackQuery.Message.Message.Chat.ID,
		MessageID: update.CallbackQuery.Message.Message.ID,
	})
	if err != nil {
		log.Println(err)
		return
	}
	_, err = b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      update.CallbackQuery.Message.Message.Chat.ID,
		Text:        data.ShowMenu[users.GetLang(db, update.CallbackQuery.From.ID)], //
		ReplyMarkup: inlineKeyboard,
	})
	if err != nil {
		log.Println(err)
		return
	}
}

func contacts(ctx context.Context, b *bot.Bot, update *models.Update) {
	_, err := b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: update.CallbackQuery.ID,
		ShowAlert:       false,
	})
	if err != nil {
		log.Println(err)
		return
	}

	_, err = b.DeleteMessage(ctx, &bot.DeleteMessageParams{
		ChatID:    update.CallbackQuery.Message.Message.Chat.ID,
		MessageID: update.CallbackQuery.Message.Message.ID,
	})
	if err != nil {
		log.Println(err)
		return
	}
	_, err = b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.CallbackQuery.Message.Message.Chat.ID,
		Text:   data.ContactData[users.GetLang(db, update.CallbackQuery.From.ID)],
		ReplyMarkup: &models.InlineKeyboardMarkup{
			InlineKeyboard: [][]models.InlineKeyboardButton{
				{
					{
						Text: data.GoBack[users.GetLang(db, update.CallbackQuery.From.ID)], CallbackData: "main_menu",
					},
				},
			},
		},
		ParseMode: models.ParseModeHTML,
	})
}

// askGroup returns button with group link
func askGroup(ctx context.Context, b *bot.Bot, update *models.Update) {
	_, err := b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: update.CallbackQuery.ID,
		ShowAlert:       false,
	})
	if err != nil {
		log.Println(err)
		return
	}

	_, err = b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.CallbackQuery.Message.Message.Chat.ID,
		Text:   data.OtherQuestions[users.GetLang(db, update.CallbackQuery.From.ID)],
		ReplyMarkup: &models.InlineKeyboardMarkup{
			InlineKeyboard: [][]models.InlineKeyboardButton{
				{
					{
						Text: data.GroupButton[users.GetLang(db, update.CallbackQuery.From.ID)],
						URL:  os.Getenv("GROUP_LINK"),
					},
				},
			},
		},
	})
}

// answer serves to send answers for each faq
func answer(ctx context.Context, b *bot.Bot, update *models.Update) {
	userId := update.CallbackQuery.From.ID
	_, err := b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: update.CallbackQuery.ID,
		ShowAlert:       false,
	})
	if err != nil {
		log.Println(err)
		return
	}
	_, err = b.DeleteMessage(ctx, &bot.DeleteMessageParams{
		ChatID:    update.CallbackQuery.Message.Message.Chat.ID,
		MessageID: update.CallbackQuery.Message.Message.ID,
	})
	if err != nil {
		log.Println(err)
		return
	}
	_, err = b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.CallbackQuery.Message.Message.Chat.ID,
		Text:   data.Info[users.GetLang(db, userId)][update.CallbackQuery.Data],
		ReplyMarkup: &models.InlineKeyboardMarkup{
			InlineKeyboard: [][]models.InlineKeyboardButton{
				{
					{
						Text:         data.QuestionsList[users.GetLang(db, userId)],
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
	if err != nil {
		log.Println(err)
		return
	}
}

// setLang serves to set the chosen language for the given user
func setLang(ctx context.Context, b *bot.Bot, update *models.Update) {
	var lang = strings.Split(update.CallbackQuery.Data, "_")[1]
	tgId := update.CallbackQuery.From.ID
	users.SetLang(db, lang, tgId)

	_, err := b.EditMessageText(ctx, &bot.EditMessageTextParams{
		ChatID:    update.CallbackQuery.Message.Message.Chat.ID,
		MessageID: update.CallbackQuery.Message.Message.ID,
		Text:      fmt.Sprintf("%s\n%s", data.Greeting[lang], data.HowCanHelp[lang]),
		ReplyMarkup: &models.InlineKeyboardMarkup{
			InlineKeyboard: data.MainMenu[lang],
		},
	})
	if err != nil {
		log.Println(err)
		return
	}
}
