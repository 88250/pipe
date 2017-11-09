package turtle

// Version of the turtle library
const Version = "v0.1.0"

// Emojis maps a name to an Emoji
var Emojis = make(map[string]*Emoji)

func init() {
	for _, e := range emojis {
		Emojis[e.Name] = e
	}
}

// Search emojis by a name
func Search(s string) []*Emoji {
	return search(emojis, s)
}

// Keyword filters the emojis by a keyword
func Keyword(k string) []*Emoji {
	return keyword(emojis, k)
}

// Category filters the emojis by a category
func Category(c string) []*Emoji {
	return category(emojis, c)
}
