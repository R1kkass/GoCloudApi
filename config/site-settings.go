package config


var SiteSettingsConfig = map[string]interface{}{
	"versions":  map[int]interface{}{
        1: "Версия 1.0",
	},
    "tariffs":  map[int]interface{}{
        1: "Бесплатный",
        2: "Платный",
	},
    "styles" :  map[string]interface{}{
        "material" : "Прямоугольная", // Оставлять первой
        "fabposter" : "Прямоугольная с постером",
        "hidden" : "Скрытая",
	},
    "colors" :  map[string]interface{}{
        "red" : "Красный",
        "orange" : "Оранжевый",
        "lightorange" : "Светлооранжевый",
        "yellow" : "Желтый",
        "lightyellow": "Светложелтый",
        "green" : "Зеленый",
        "lightgreen" : "Светлозеленый",
        "blue" : "Синий",
        "turquoise": "Бирюзовый",
        "violet" : "Фиолетовый",
        "lilac": "Лиловый",
        "crimson": "Малиновый",
        "black" : "Черный",
        "lightblack": "Серый",
	},
    "contents" :  map[string]interface{}{
        "material" :  map[string]interface{}{
            "default" :  map[string]interface{}{
                "name" : "Примерь онлайн",
                "text" : "Примерь онлайн",
			},
            "furniture" :  map[string]interface{}{
                "name" : "Примерить мебель",
                "text" : "Примерить мебель",
			},
            "view3d" :  map[string]interface{}{
                "name" : "Посмотреть 3D",
                "text" : "Посмотреть 3D",
			},
		},
        "fabposter" :  map[string]interface{}{
            "default" :  map[string]interface{}{
                "name" : "Примерь онлайн",
                "text" : "Примерь онлайн",
                "image" : "poster.webp",
			},
            "furniture" :  map[string]interface{}{
                "name" : "Примерить мебель",
                "text" : "Примерить мебель",
                "image" : "poster.webp",
			},
		},
        "hidden" : map[string]interface{}{
            "default" :  map[string]interface{}{
                "name" : "По-умолчанию",
			},
		},
	},
    "hdrs" :  map[string]interface{}{
        "neutral" : "Нейтральное",
        "shadow" : "С тенями",
	},
    "positions" :  map[string]interface{}{
        "left-bottom" :  map[string]interface{}{
            "name" : "Левый нижний угол",
            "is_default" : false,
            "position" : map[string]interface{}{
                // "top" : 0,
                // "right" : 0,
                "bottom" : "10vh",
                "left" : "7.5vw",
			},
		},
        "default" :  map[string]interface{}{
            "name" : "Правый нижний угол",
            "is_default" : true,
            "position" :  map[string]interface{}{
                // "top" : 0,
                "right" : "7.5vw",
                "bottom" : "10vh",
                // "left" : 0
			},
		},
        "left-bottom-lower" :  map[string]interface{}{
            "name" : "Левый нижний угол (ниже)",
            "is_default" : false,
            "position" :  map[string]interface{}{
                // "top" : 0,
                // "right" : 0,
                "bottom" : "4vh",
                "left" : "2.5vw",
			},
		},
        "right-bottom-lower" :  map[string]interface{}{
            "name" : "Правый нижний угол (ниже)",
            "is_default" : false,
            "position" :  map[string]interface{}{
                // "top" : 0,
                "right" : "2.5vw",
                "bottom" : "4vh",
                // "left" : 0
			},
		},
	},
}