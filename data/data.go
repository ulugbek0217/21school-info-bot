package data

import "github.com/go-telegram/bot/models"

// Info keeps answers of FAQ
var Info = map[string]map[string]string{
	"ru": {
		"item_1": "«SCHOOL 21» — образовательный проект Республики Узбекистан.\n\nМы готовим специалистов в сфере программирования по методологии «равный равному»: без менторов, без лекций, без оценок. Участники обучаются и общаются между собой, решают учебные задания, основанные на практике и проверяют друг друга.\n\t\n\t\n\t\n\t\n\t\n\tКак это работает? Каждый обучающийся получает техническое задание. Например, описание проекта. У вас есть вся необходимая информация о нем. Если вам чего-то не хватает или вы чего-то не понимаете, здесь включается метод peer-to-peer: вы идете к соседу справа или соседу слева. Если и они не знают, найдется человек или запрос в гугле, который поможет вам. Получается, что вы разбираетесь с задачей от самого первого пункта и до тех пор, пока окончательно не поймете, что от вас требуется и как это сделать. Очень крутой опыт.\n\t\n\tДальше коротко и по пунктам:  \n\t\n\t\n\t\n\t— Обучение бесплатное\n\t\n\t— Школа работает 24/7\n\t\n\t— В программе обучения – обязательная стажировка\n\t\n\t— Учебные кластеры оборудованы отличной техникой\n\t\n\t— Чтобы поступить, нужно пройти 3 отборочных этапа: 2 онлайн игры, видео-интервью и финальный отборочный интенсив, который проходит офлайн в кампусе SCHOOL 21.\n\t\n\t\n\t\n\tЖдем вас в Школе. Оставляйте заявку на сайте https://21-school.uz/\n\t\n\t\n\t\n\tПодписывайся на наш Телеграм-канал — @skd21school",

		"item_2": "Отборочный интенсив длится 26 дней и проводится на языке С. В это время усваивается большой объем информации всего за 4 недели. Участник получает повышенное количество заданий, ежедневные дедлайны и строгие критерии проверки заданий. Он занимает от 8 до 12 часов в день и совместить его с чем-то ещё будет практически невозможно.\n\n\n\n\tОтборочный интенсив – это финальный и последний шаг перед основным обучением.\n\t\n\tПроходя его, вы сможете для себя понять, насколько вам подходит эта система обучения и интересно ли вам программирование.\n\t\n\t\n\t\n\tЖдем вас в «SCHOOL 21». Переходите на сайт и оставляйте заявку: https://21-school.uz/",
		"item_3": "Ближайшие интенсивы:\n\n\tСамарканд: 15 апреля\n\t\n\tСамарканд: 27 мая\n\t\n\tСамарканд: 15 июня\n\t\n\t\n\t\n\tЧтобы попасть на интенсив, нужно оставить заявку на сайте https://21-school.uz/ , пройти онлайн-игры, записать интервью и выбрать одну из онлайн-встреч на сайте https://applicant.21-school.uz.\n\t\n\tЕсли у вас нет активных встреч, значит все свободные места разобрали.\n\t\n\tСледите за новостями в группе. Мы расскажем про открытие новых встреч.",

		"item_4": "После успешного завершения отборочного интенсива, вам предлагается подписать договор на основное обучение, которое длится от 1,5 до 3,5 лет. Все зависит от вашей мотивации и скорости усвоения материала.\n\n\n\n\tВы сможете сами выбрать траекторию обучения. Школа открыта 24/7, основная часть процесса обучения проходит в прямом взаимодействии с другими участниками. Все учебные материалы хранятся в электронном виде во внутренней сети школы. Чтобы получить их, вы должны приехать в школу.\n\t\n\t\n\t\n\tПо опыту работающих участников школы, на основное обучение приходится тратить несколько дней в неделю после работы и в выходные дни. То есть основное обучение вы сможете, при желании, совмещать с работой или учебой в вузе. Все зависит от того, сколько времени именно вам понадобится на выполнение проектов.\n\t\n\t\n\t\n\tЖдем вас в школе. Переходите на сайт и оставляйте заявку: https://21-school.uz/",

		"item_5": "Обучение на интенсиве устроено таким образом, что всю информацию вы получаете «по ходу» отборочного этапа. Подготовиться к нему невозможно и совершенно не нужно. 50% наших текущих участников не имели опыта программирования ранее. Поэтому ничего не бойтесь и оставляйте заявку на сайте\n\n\thttps://21-school.uz/",

		"item_6": "В школу может поступить любой желающий от 18 лет.\n\n\n\n\tЭтапы поступления:\n\t\n\t— Оставить заявку на нашем сайте:\n\t\n\thttps://21-school.uz/\n\t\n\t— Нужно создать свой профиль в нашей системе\n\t\n— Пройти 2 онлайн игры (на память и логику)\n\t\n\t— Записаться на онлайн-встречу, на ней вы узнаете всю подробную информацию и сможете задать вопросы.\n\t\n\t— Записаться на финальный отборочный интенсив и успешно пройти его",

		"item_7": "Вам нужно зайти на сайт\n\n\thttps://applicant.21-school.uz, выбрать вкладку «встреча» и выбрать подходящую вам дату. Если доступных встреч нет, значит их еще не открыли, как только они появятся — вам придет уведомление на почту. Встречи открываются перед интенсивами, которые чаще всего проходят два раза в год: летом и зимой.\n\t\n\tСейчас встречи открываются постепенно и проходят в онлайн режиме. Ехать в кампус не нужно.\n\t\n\tСледите за новостями, чтобы узнать, когда можно записаться на встречу.",
	},

	"uz": {
		"item_1": "Raqamli texnologiyalar maktabi “School 21” – O‘zbekiston Respublikasi ta’lim loyihasi.\n\nBiz dasturlash sohasida mutaxassislarni murabbiylarsiz, ma’ruzalarsiz, baholarsiz, \"peer to peer\" usulida tayyorlaymiz. Ishtirokchilar o‘zaro muloqot qiladilar va o‘rganadilar, amaliyotga asoslangan o‘quv topshiriqlarini muhokama qiladilar va bir-birlarini sinovdan o‘tkazadilar.\n\n\nBu qanday ishlaydi? Har bir o‘rganuvchi texnik topshiriq oladi. Masalan, loyiha tavsifi. Sizda u haqida barcha kerakli ma’lumotlar mavjud bo‘ladi. Agar sizga biror narsa yetishmayotgan bo‘lsa yoki tushunmasangiz, bu yerda peer-to-peer usuli ishlaydi: siz yon tomoningizdagi sheriklaringizga murojaat qilishingiz mumkin. Agar ular ham bilishmasa, sizga yordam beradigan odam topiladi yoki “Google”da so‘rov berish mumkin bo‘ladi. Demak, siz vazifangizni birinchi bosqichidan boshlab, sizdan nima talab qilinishini va buni qanday qilishni tushunmaguningizcha o‘rganasiz. Bu juda ajoyib tajriba.\n\nEndi bandlar bo‘yicha qisqacha:\n\n— Ta’lim bepul\n\n— Maktab 24/7 ishlaydi\n\n— O‘quv dasturida majburiy amaliyot mavjud\n\n— O‘quv klasterlari zamonaviy texnika bilan jihozlangan\n\n— Kirish uchun siz 3 ta saralash bosqichidan o‘tishingiz kerak: 2 ta onlayn o‘yin, video-suhbat va “School 21” kampusida oflayn rejimda bo‘lib o‘tadigan saralash Intensiv bosqichi.\n\n\nSizni raqamli texnologiyalar maktabi “School 21”da kutamiz.  https://21-school.uz/ saytida ariza qoldiring.\n\n\nBizning @skd21school telegram-kanalimizga obuna bo‘ling.",
		"item_2": "Saralash Intensivi 26 kun davom etadi va C tilida amalga oshiriladi. Bu vaqtda katta miqdordagi ma’lumotlar atigi 4 hafta ichida o‘zlashtiriladi. Ishtirokchi ko‘p sonli topshiriqlarni, har kunlik muddatlarni va topshiriqlarni tekshirish bo‘yicha qat’iy mezonlarni oladi. Bu kuniga 8 soatdan 12 soatgacha vaqtni egallaydi va uni boshqa biror-bir faoliyat bilan birgalikda olib borish imkonsiz bo‘ladi.\n\nSaralash Intensivi – bu asosiy ta’limdan oldingi yakuniy va oxirgi qadamdir.\n\nBu bosqich mazkur o‘quv tizimi sizga qanchalik mos kelishini va dasturlash siz uchun qiziq yoki yo‘qligini tushunishingizga imkon beradi.\n\nSizni raqamli texnologiyalar maktabi “School 21”da kutamiz. https://21-school.uz/ saytiga o‘ting va ariza qoldiring",
		"item_3": "Yaqin orada bo‘ladigan intensivlar:\n\n27-may Samarqand \n\n\nSaralash Intensiv bosqichiga ishtirok etish uchun siz https://21-school.uz/ saytida ariza qoldirishingiz, onlayn o‘yinlardan o‘tishingiz, https://applicant.21-school.uz saytidagi onlayn uchrashuvlardan birini tanlashingiz va ishtirok etishingiz lozim.\n\nAgar sizda faol uchrashuvlar bo‘lmasa, unda barcha bo‘sh joylar band qilingan.\n\nGuruhdagi yangiliklarni kuzatib boring.Yangi uchrashuvlarning ochilishi bo‘yicha ma’lumot beramiz.",
		"item_4": "Saralash Intensivini muvaffaqiyatli tugatgandan so‘ng, sizga 1,5 yildan 3 yilgacha davom etadigan asosiy ta’lim shartnomasini imzolash taklif etiladi. Bularning barchasi sizning shijoatingizga va ma’lumotlarni o‘zlashtirish tezligingizga bog’liq.\n\n\nSiz ta’lim yo‘nalishini o‘zingiz tanlashingiz mumkin. Maktab kuniga 24 soat va haftasiga 7 kun ochiq, o‘quv jarayonining asosiy qismi boshqa ishtirokchilar bilan bevosita hamkorlikda o‘tadi. Barcha o‘quv materiallari elektron shaklda maktabning ichki tarmog’ida saqlanadi. Ularni olish uchun maktabga kelishingiz lozim bo‘ladi.\nMaktabda o‘qish bilan birga ishlaydigan ishtirokchilarning tajribasiga ko‘ra, asosiy mashg’ulotlarga haftada bir necha kun ishdan keyingi vaqt va dam olish kunlarini sarflash kerak. Ya’ni, agar xohlasangiz, Asosiy ta’limni oliy ta’limda o‘qish yoki ishlash bilan birgalikda olib bora olasiz. Bularning barchasi loyihalarni amalga oshirish uchun qancha vaqt kerakligiga bog’liq.\n\nSizni raqamli texnologiyalar maktabi “School 21”da kutamiz. https://21-school.uz/ saytiga o‘ting va ariza qoldiring.",
		"item_5": "Saralash Intensivida ta’lim jarayoni shu tarzda tuzilganki, siz barcha ma’lumotlarga saralash bosqichi davomida ega bo‘lasiz. Unga tayyorgarlik ko‘rish mumkin ham, shart ham emas. Hozirgi a’zolarimizning 50 foizi ilgari dasturlash tajribasiga ega emas edi. Shuning uchun, hech narsadan qo‘rqmang va https://21-school.uz/ veb-saytimizda ariza qoldiring",
		"item_6": "Raqamli texnologiyalar maktabi “School 21”ga 18 yosh va undan yuqori boʻlgan har bir kishi o‘qishga kirishi mumkin.\n\nQabul bosqichlari:\n\n— https://21-school.uz/ veb-saytida ariza qoldirish;\n\n— Tizimda o‘z profilini yaratish;\n\n— 2 ta onlayn o‘yindan o‘tish (xotira va mantiq uchun)\n\n— Onlayn uchrashuvga yozilish. Unda siz barcha zarur ma’lumotlarga ega bo‘lasiz va savollaringizga javob olasiz.\n\n— Yakuniy saralash intensiv bosqichiga yozilish va undan muvaffaqiyatli o‘tish.",
		"item_7": "Siz saytga kirishingiz kerak\n\nhttps://applicant.21-school.uz saytida \"uchrashuv\" yorlig’ini va sizga mos sanani tanlang. Agar uchrashuvlar mavjud bo‘lmasa, demak ular hali ochilmagan. Ular paydo bo‘lishi bilanoq sizga pochta orqali xabarnoma keladi. Uchrashuvlar doimiy tarzda o‘tkaziladigan Saralash intensivlaridan oldin ochiladi. Uchrashuvlar onlayn tarzda  Microsoft Teams ilovasida o‘tkaziladi.  \n\n\nUchrashuvga qachon yozilish mumkinligini bilish uchun yangiliklarimizni kuzatib boring.",
	},
}

// Questions FAQ
var Questions = map[string]map[string]string{
	"ru": {
		"item_1": "Расскажи, что такое \"School 21\"",
		"item_2": "Что такое отборочный интенсив? 🤔",
		"item_3": "Даты ближайших отборочных интенсивов 📆",
		"item_4": "Можно совмещать обучение с чем-то ещё? 🔄",
		"item_5": "Как подготовиться к интенсиву? 🏃",
		"item_6": "Как поступить к вам?",
		"item_7": "Как записаться на онлайн встречу? 📺",
	},
	"uz": {
		"item_1": "\"School 21\" haqida",
		"item_2": "Saralash intensivi nima? 🤔",
		"item_3": "Yaqin saralash intensiv sanalari 📆",
		"item_4": "O'qishni parallel olib borish 🔄",
		"item_5": "Intensivga qanday tayyorlanish kerak? 🏃",
		"item_6": "Ro'yxatdan o'tish bosqichlari",
		"item_7": "Onlayn uchrashuvga yozilish 📺",
	},
}

// QuestKbd is a keyboard of FAQ
var QuestKbd = map[string][][]models.InlineKeyboardButton{
	"ru": {
		{
			{
				Text: Questions["ru"]["item_1"], CallbackData: "item_1",
			},
		}, {
			{
				Text: Questions["ru"]["item_2"], CallbackData: "item_2",
			},
		}, {
			{
				Text: Questions["ru"]["item_3"], CallbackData: "item_3",
			},
		}, {
			{
				Text: Questions["ru"]["item_4"], CallbackData: "item_4",
			},
		}, {
			{
				Text: Questions["ru"]["item_5"], CallbackData: "item_5",
			},
		}, {
			{
				Text: Questions["ru"]["item_6"], CallbackData: "item_6",
			},
		}, {
			{
				Text: Questions["ru"]["item_7"], CallbackData: "item_7",
			},
		}, {
			{
				Text: "❓Хочу задать вопрос❓", CallbackData: "ask_group",
			},
		}, {
			{
				Text: "Каково быть программистом? 🧑‍💻", URL: "https://new-games.21-school.uz",
			},
		}, {
			{
				Text: "Главное меню 📔", CallbackData: "main_menu",
			},
		},
	},

	"uz": {{
		{
			Text: Questions["uz"]["item_1"], CallbackData: "item_1",
		},
	}, {
		{
			Text: Questions["uz"]["item_2"], CallbackData: "item_2",
		},
	}, {
		{
			Text: Questions["uz"]["item_3"], CallbackData: "item_3",
		},
	}, {
		{
			Text: Questions["uz"]["item_4"], CallbackData: "item_4",
		},
	}, {
		{
			Text: Questions["uz"]["item_5"], CallbackData: "item_5",
		},
	}, {
		{
			Text: Questions["uz"]["item_6"], CallbackData: "item_6",
		},
	}, {
		{
			Text: Questions["uz"]["item_7"], CallbackData: "item_7",
		},
	}, {
		{
			Text: "Dasturchi bo'lish qanaqa? 🧑‍💻", URL: "https://new-games.21-school.uz",
		},
	}, {
		{
			Text: "❓Savol yo'llash❓", CallbackData: "ask_group",
		},
	}, {
		{
			Text: "Asosiy menyu 📔", CallbackData: "main_menu",
		},
	},
	},
}

// MainMenu keeps main menu buttons
var MainMenu = map[string][][]models.InlineKeyboardButton{
	"ru": {
		{
			{
				Text: "Меню вопросов", CallbackData: "questions",
			},
		}, {
			{
				Text: "Tilni almashtirish 🇺🇿", CallbackData: "lang_uz",
			},
		},
	},
	"uz": {
		{
			{
				Text: "Savollar menyusi", CallbackData: "questions",
			},
		}, {
			{
				Text: "Сменить язык 🇷🇺", CallbackData: "lang_ru",
			},
		},
	},
}

var ShowMenu = map[string]string{
	"ru": "Меню вопросов",
	"uz": "Savollar menyusi",
}

var HowCanHelp = map[string]string{
	"ru": "Как я могу вам помочь?",
	"uz": "Sizga qanday yordam bera olaman?",
}

var Greeting = map[string]string{
	"ru": "Привет, добро пожаловать.",
	"uz": "Assalomu alaykum, xush kelibsiz.",
}

var QuestionsList = map[string]string{
	"ru": "Вернуться к списку вопросов ⬅️",
	"uz": "Savollar menyusiga qaytish ⬅️",
}

var OtherQuestions = map[string]string{
	"ru": "Остальные вопросы вы можете задать в нашу группу. Наши волонтеры ответят вам сразу 🙂\nНажмите кнопку ниже ⬇️",
	"uz": "Boshqa savollaringizni guruhimizga yo'llashingiz mumkin. Sizning savolingizga darhol javob berishadi 🙂\nQuyidagi tugmani bosing ⬇️",
}

var GroupButton = map[string]string{
	"ru": "Группа",
	"uz": "Guruh",
}
