package config

var Lang = map[string]interface{}{
	//Русский
	"ru" : map[string]interface{}{
		"btn" : map[string]string{
			"default" : "Примерить онлайн",
			"furniture" : "Примерить мебель",
			"view3d" : "Посмотреть в 3D",
		},
		"modal" : map[string]string{
			"title" : "Просмотр трехмерной модели товара",
			"ar_btn" : "Смотреть в пространстве",
			"qr_text" : "Наведите камеру телефона или планшета на QR-код, чтобы посмотреть на объект в интерьере с помощью дополненной реальности.",
			"qr_note" : "*Для смартфонов и планшетов с поддержкой дополненной реальности.",
		},
	},
	//Английский
	"en" : map[string]interface{}{
		"btn" : map[string]string{
			"default" : "Try Online",
			"furniture" : "Try Online",
			"view3d" : "View in 3D",
		},
		"modal" : map[string]string{
			"title" : "3D product model",
			"ar_btn" : "View in space",
			"qr_text" : "Point your phone or tablet camera at the QR code to see an object in the interior using Augmented Reality.",
			"qr_note" : "*For smartphones and tablets with Augmented Reality support.",
		},
	},
	//Сербский
	"sr" : map[string]interface{}{
		"btn" : map[string]string{
			"default" : "Pogledajte online",
			"furniture" : "Detaljan pregled proizvoda",
			"view3d" : "Pogledajte u 3D",
		},
		"modal" : map[string]string{
			"title" : "Pregled trodimenzionalnog modela proizvoda",
			"ar_btn" : "Isprobajte u vašem enterijeru",
			"qr_text" : "Skenirajte QR-kod uz pomoć telefona ili tableta, kako biste pogledali objekat u enterijeru uz pomoć proširene stvarnosti.",
			"qr_note" : "*Za pametne telefone i tablete sa podrškom za proširenu stvarnost.",
		},
	},
	//Немецкий
	"de" : map[string]interface{}{
		"btn" : map[string]string{
			"default" : "Online anprobieren",
			"furniture" : "Möbel anprobieren",
			"view3d" : "Ansicht in 3D",
		},
		"modal" : map[string]string{
			"title" : "Ansicht eines dreidimensionalen Modells des Produkts",
			"ar_btn" : "Beobachten im Weltraum",
			"qr_text" : "Bewegen Sie die Kamera Ihres Telefons oder Tablets über einen QR-Code, um ein Innenobjekt mit Hilfe von Augmented Reality zu betrachten.",
			"qr_note" : "*Für Augmented-Reality-fähige Smartphones und Tablets.",
		},
	},
};