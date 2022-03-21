package main

var (
	RunesRand  = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz01234567890123456789")
	RandDigits = []rune("1234567890")
	RandMail   = []string{"@mail.ru", "@gmail.com", "@yandex.ru"}
	RandCity   = map[string]string{
		"Moscow":           "Moscow region",
		"Podolsk":          "Moscow region",
		"Khimki":           "Moscow region",
		"Saint Petersburg": "Leningradskaya oblast",
		"Magnitogorsk":     "Chelyabinskay oblast",
	}
	RandAddress  = []string{"Ulica Pobedi 33", "Prospect Lenina 13", "Fevralskaya 24", "Matrosskaya 19"}
	RandName     = []string{"Peter", "Oleg", "Ivan", "Nikolay", "Serge", "Alex"}
	RandSurname  = []string{"Ivanov", "Petrov", "Vasechkin", "Sokolov", "Stepanov"}
	RandBank     = []string{"alpha", "vtb", "sber", "tinkoff"}
	ProductName  = []string{"Pullover", "Trousers", "Jeans", "T-shirt", "Sweater", "Coat", "Turtleneck"}
	ProductBrand = []string{"Annemore", "Pavlotti", "UZcotton", "Sentio", "Ticle", "Materia"}
	RandCurrency = []string{"USD", "EUR", "RUB", "GBP"}
)
