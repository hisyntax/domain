package tld

type TldObj struct {
	AlphaOrder string `json:"alpha_order"`
	Tld        []Tdl  `json:"tdl"`
}

type Tdl struct {
	Tld  string `json:"domain"`
	Type string `json:"type"`
}

var Tlds = []TldObj{
	{AlphaOrder: "a",
		Tld: []Tdl{
			{Tld: "0db", Type: "handshake"}, {Tld: "0z", Type: "handshake"},
			{Tld: "1", Type: "handshake"}, {Tld: "1d", Type: "handshake"},
			{Tld: "1q", Type: "handshake"}, {Tld: "247", Type: "handshake"},
			{Tld: "35", Type: "handshake"}, {Tld: "3dom", Type: "handshake"},
			{Tld: "49", Type: "handshake"}, {Tld: "4free", Type: "handshake"},
			{Tld: "4k", Type: "handshake"}, {Tld: "4sale", Type: "handshake"},
			{Tld: "4you", Type: "handshake"}, {Tld: "80proof", Type: "handshake"},
			{Tld: "8888", Type: "handshake"}, {Tld: "8s", Type: "handshake"},
			{Tld: "abo", Type: "handshake"}, {Tld: "aboutme", Type: "handshake"},
			{Tld: "aby", Type: "handshake"}, {Tld: "ac", Type: "ccTLD"},
			{Tld: "aca", Type: "handshake"}, {Tld: "academy", Type: "gTLD"},
			{Tld: "accountant ", Type: "gTLD"}, {Tld: "accountants", Type: "gTLD"},
			{Tld: "actor", Type: "gTLD"}, {Tld: "addme", Type: "handshake"},
			{Tld: "adlt", Type: "handshake"}, {Tld: "adult", Type: "gTLD"},
			{Tld: "advisor", Type: "handshake"}, {Tld: "ae.org", Type: "gTLD"},
			{Tld: "afam", Type: "handshake"}, {Tld: "africa", Type: "gTLD"},
			{Tld: "afz", Type: "handshake"}, {Tld: "agency", Type: "gTLD"},
			{Tld: "ags", Type: "handshake"}, {Tld: "agua", Type: "handshake"},
			{Tld: "ahoy", Type: "handshake"}, {Tld: "ai", Type: "ccTLD"},
			{Tld: "airforce", Type: "gTLD"}, {Tld: "aj", Type: "handshake"},
			{Tld: "albums", Type: "handshake"}, {Tld: "alto", Type: "handshake"},
			{Tld: "amg", Type: "handshake"}, {Tld: "amigo", Type: "handshake"},
			{Tld: "amor", Type: "handshake"}, {Tld: "ane", Type: "handshake"},
			{Tld: "annex", Type: "handshake"}, {Tld: "aotearoa", Type: "handshake"},
			{Tld: "apartment", Type: "handshake"}, {Tld: "apartments", Type: "gTLD"},
			{Tld: "api", Type: "handshake"}, {Tld: "app", Type: "gTLD"},
			{Tld: "arbitrator", Type: "handshake"}, {Tld: "army", Type: "gTLD"},
			{Tld: "art ", Type: "gTLD"}, {Tld: "artesanal", Type: "handshake"},
			{Tld: "articles", Type: "handshake"}, {Tld: "artificial", Type: "handshake"},
			{Tld: "asia", Type: "gTLD"}, {Tld: "associates", Type: "gTLD"},
			{Tld: "assurances", Type: "handshake"}, {Tld: "atc", Type: "handshake"},
			{Tld: "ath", Type: "handshake"}, {Tld: "attorney", Type: "gTLD"},
			{Tld: "atwork", Type: "handshake"}, {Tld: "auction", Type: "gTLD"},
			{Tld: "audio", Type: "gTLD"}, {Tld: "augmented", Type: "handshake"},
			{Tld: "auto", Type: "gTLD"}, {Tld: "autos", Type: "gTLD"},
		},
	},
	{AlphaOrder: "b",
		Tld: []Tdl{
			{Tld: "b2b", Type: "handshake"}, {Tld: "baas", Type: "handshake"},
			{Tld: "baby", Type: "gTLD"}, {Tld: "badly", Type: "handshake"},
			{Tld: "bakes", Type: "handshake"}, {Tld: "band", Type: "gTLD"},
			{Tld: "bar", Type: "gTLD"}, {Tld: "bargains", Type: "gTLD"},
			{Tld: "batch", Type: "handshake"}, {Tld: "beach", Type: "handshake"},
			{Tld: "beauty", Type: "gTLD"}, {Tld: "beef", Type: "handshake"},
			{Tld: "beer", Type: "gTLD"}, {Tld: "bem", Type: "handshake"},
			{Tld: "berlin", Type: "gTLD"}, {Tld: "best", Type: "gTLD"},
			{Tld: "bid", Type: "gTLD"}, {Tld: "bike", Type: "gTLD"},
			{Tld: "bingo", Type: "gTLD"}, {Tld: "bio", Type: "gTLD"},
			{Tld: "biometric", Type: "handshake"}, {Tld: "bitcoinfund", Type: "handshake"},
			{Tld: "biz", Type: "gTLD"}, {Tld: "bizdata", Type: "handshake"},
			{Tld: "bizdev", Type: "handshake"}, {Tld: "black", Type: "gTLD"},
			{Tld: "blackfriday", Type: "gTLD"}, {Tld: "blockchaindapps", Type: "handshake"},
			{Tld: "blog", Type: "gTLD"}, {Tld: "blogging", Type: "handshake"},
			{Tld: "blue", Type: "gTLD"}, {Tld: "bmp", Type: "handshake"},
			{Tld: "boats", Type: "gTLD"}, {Tld: "bob", Type: "handshake"},
			{Tld: "bog", Type: "handshake"}, {Tld: "bond", Type: "gTLD"},
			{Tld: "boo", Type: "gTLD"}, {Tld: "booked", Type: "handshake"},
			{Tld: "boredapes", Type: "handshake"}, {Tld: "boston", Type: "gTLD"},
			{Tld: "boutique", Type: "gTLD"}, {Tld: "bqw", Type: "handshake"},
			{Tld: "br.com", Type: "gTLD"}, {Tld: "brand", Type: "handshake"},
			{Tld: "brewery", Type: "handshake"}, {Tld: "brokers", Type: "handshake"},
			{Tld: "browsers", Type: "handshake"}, {Tld: "btk9", Type: "handshake"},
			{Tld: "btt", Type: "handshake"}, {Tld: "buddhist", Type: "handshake"},
			{Tld: "build", Type: "gTLD"}, {Tld: "builders", Type: "gTLD"},
			{Tld: "business", Type: "gTLD"}, {Tld: "buyaflat", Type: "handshake"},
			{Tld: "buzz", Type: "gTLD"}, {Tld: "byn", Type: "handshake"},
			{Tld: "bz", Type: "ccTLD"},
		},
	},
	{AlphaOrder: "c",
		Tld: []Tdl{
			{Tld: "c", Type: "handshake"}, {Tld: "c0m", Type: "handshake"},
			{Tld: "c4", Type: "handshake"}, {Tld: "ca", Type: "ccTLD"},
			{Tld: "cab", Type: "gTLD"}, {Tld: "cabins", Type: "handshake"},
			{Tld: "cafe", Type: "gTLD"}, {Tld: "cam", Type: "gTLD"},
			{Tld: "camera", Type: "gTLD"}, {Tld: "camp", Type: "gTLD"},
			{Tld: "capital", Type: "gTLD"}, {Tld: "car", Type: "gTLD"},
			{Tld: "cardcollector", Type: "handshake"}, {Tld: "cards", Type: "gTLD"},
			{Tld: "care", Type: "gTLD"}, {Tld: "careers", Type: "gTLD"},
			{Tld: "cares", Type: "handshake"}, {Tld: "cars", Type: "gTLD"},
			{Tld: "casa", Type: "gTLD"}, {Tld: "cash", Type: "gTLD"},
			{Tld: "casino", Type: "gTLD"}, {Tld: "catering", Type: "gTLD"},
			{Tld: "catgirl", Type: "handshake"}, {Tld: "causes", Type: "handshake"},
			{Tld: "cc", Type: "ccTLD"}, {Tld: "center", Type: "gTLD"},
			{Tld: "ceo", Type: "gTLD"}, {Tld: "cerrajero", Type: "handshake"},
			{Tld: "cfd", Type: "gTLD"}, {Tld: "ch", Type: "gTLD"},
			{Tld: "chat", Type: "gTLD"}, {Tld: "cheap", Type: "gTLD"},
			{Tld: "cheddar", Type: "handshake"}, {Tld: "christmas", Type: "gTLD"},
			{Tld: "cism", Type: "handshake"}, {Tld: "cita", Type: "handshake"},
			{Tld: "citizenship", Type: "handshake"}, {Tld: "city", Type: "gTLD"},
			{Tld: "ckq", Type: "handshake"}, {Tld: "claims", Type: "gTLD"},
			{Tld: "clc", Type: "handshake"}, {Tld: "cleaning", Type: "gTLD"},
			{Tld: "clic", Type: "handshake"}, {Tld: "click", Type: "gTLD"},
			{Tld: "client", Type: "handshake"}, {Tld: "clients", Type: "handshake"},
			{Tld: "clinic", Type: "gTLD"}, {Tld: "cliq", Type: "handshake"},
			{Tld: "clothing", Type: "gTLD"}, {Tld: "cloud", Type: "gTLD"},
			{Tld: "cloudbot", Type: "handshake"}, {Tld: "club", Type: "gTLD"},
			{Tld: "cm", Type: "ccTLD"}, {Tld: "cm3", Type: "handshake"},
			{Tld: "cn", Type: "ccTLD"}, {Tld: "cn.com", Type: "gTLD"},
			{Tld: "cn3", Type: "handshake"}, {Tld: "co", Type: "ccTLD"},
			{Tld: "co.bz", Type: "ccTLD"}, {Tld: "co.com", Type: "gTLD"},
			{Tld: "co.in", Type: "ccTLD"}, {Tld: "co.uk", Type: "ccTLD"},
			{Tld: "co3", Type: "handshake"}, {Tld: "coach", Type: "gTLD"},
			{Tld: "codes", Type: "gTLD"}, {Tld: "codeschool", Type: "handshake"},
			{Tld: "coffee", Type: "gTLD"}, {Tld: "coffees", Type: "handshake"},
			{Tld: "coinnet", Type: "handshake"}, {Tld: "college", Type: "gTLD"},
			{Tld: "com", Type: "gTLD"}, {Tld: "com.au", Type: "ccTLD"},
			{Tld: "com.cn", Type: "ccTLD"}, {Tld: "com.co", Type: "ccTLD"},
			{Tld: "com.de", Type: "gTLD"}, {Tld: "com.es", Type: "ccTLD"},
			{Tld: "com.mx", Type: "ccTLD"}, {Tld: "com.pe", Type: "ccTLD"},
			{Tld: "com.ph", Type: "ccTLD"}, {Tld: "com.se", Type: "gTLD"},
			{Tld: "com.sg", Type: "ccTLD"}, {Tld: "com.vc", Type: "ccTLD"},
			{Tld: "comic", Type: "handshake"}, {Tld: "comics", Type: "handshake"},
			{Tld: "commerce", Type: "handshake"}, {Tld: "communion", Type: "handshake"},
			{Tld: "community", Type: "gTLD"}, {Tld: "company", Type: "gTLD"},
			{Tld: "complete", Type: "handshake"}, {Tld: "computer", Type: "gTLD"},
			{Tld: "computers", Type: "handshake"}, {Tld: "concise", Type: "handshake"},
			{Tld: "condos", Type: "gTLD"}, {Tld: "conduct", Type: "handshake"},
			{Tld: "conductor", Type: "handshake"}, {Tld: "conf", Type: "handshake"},
			{Tld: "connectors", Type: "handshake"}, {Tld: "connoisseur", Type: "handshake"},
			{Tld: "construction", Type: "gTLD"}, {Tld: "consultancy", Type: "handshake"},
			{Tld: "consulting", Type: "gTLD"}, {Tld: "contact", Type: "gTLD"},
			{Tld: "contractors", Type: "gTLD"}, {Tld: "cooking", Type: "gTLD"},
			{Tld: "cool", Type: "gTLD"}, {Tld: "cork", Type: "handshake"},
			{Tld: "corporation", Type: "handshake"}, {Tld: "country", Type: "gTLD"},
			{Tld: "coupons", Type: "gTLD"}, {Tld: "courses", Type: "gTLD"},
			{Tld: "crazy", Type: "handshake"}, {Tld: "creations", Type: "handshake"},
			{Tld: "creativity", Type: "handshake"}, {Tld: "creator", Type: "handshake"},
			{Tld: "credit", Type: "gTLD"}, {Tld: "creditcard", Type: "gTLD"},
			{Tld: "crew", Type: "handshake"}, {Tld: "cricket", Type: "gTLD"},
			{Tld: "croatia", Type: "handshake"}, {Tld: "cruises", Type: "gTLD"},
			{Tld: "cryp", Type: "handshake"}, {Tld: "cryptobets", Type: "handshake"},
			{Tld: "cryptogamer", Type: "handshake"}, {Tld: "cryptoservice", Type: "handshake"},
			{Tld: "cuba", Type: "handshake"}, {Tld: "cx", Type: "ccTLD"},
			{Tld: "cycle", Type: "handshake"}, {Tld: "cymru", Type: "gTLD"},
			{Tld: "cyou", Type: "gTLD"}, {Tld: "cyprus", Type: "handshake"},

			{Tld: "d5", Type: "handshake"}, {Tld: "d8", Type: "handshake"},
		},
	},
}

func Validate() {

}