package models

const (
	All Crawler = "*"

	Google      Crawler = "Googlebot"
	GoogleImage Crawler = "Googlebot-Image"
	GoogleNews  Crawler = "Googlebot-News"
	GoogleAds   Crawler = "Mediapartners-Google"
	GoogleOther Crawler = "GoogleOther"

	Bing            Crawler = "Bingbot"
	BingAds         Crawler = "AdIdxBot"
	BingPreview     Crawler = "MicrosoftPreview"
	MicrosoftSearch Crawler = "msnbot"

	Facebook Crawler = "FacebookBot"
	Twitter  Crawler = "Twitterbot"
	Discord  Crawler = "Discordbot"
	Yahoo    Crawler = "Slurp"

	ChatGPTUser Crawler = "ChatGPT-User"
)

type Crawler string
