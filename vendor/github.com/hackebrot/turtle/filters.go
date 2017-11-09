package turtle

import "strings"

// filter a given slice of Emoji by f
func filter(emojis []*Emoji, f func(e *Emoji) bool) []*Emoji {
	var r []*Emoji
	for _, e := range emojis {
		if f(e) {
			r = append(r, e)
		}
	}
	return r
}

// category filters a slice of Emoji by Category
func category(emojis []*Emoji, c string) []*Emoji {
	return filter(emojis, func(e *Emoji) bool {
		return e.Category == c
	})
}

// keyword filters a slice of Emoji by Keywords
func keyword(emojis []*Emoji, k string) []*Emoji {
	return filter(emojis, func(e *Emoji) bool {
		for _, keyword := range e.Keywords {
			if keyword == k {
				return true
			}
		}
		return false
	})
}

// search Emoji in a slice by Name
func search(emojis []*Emoji, s string) []*Emoji {
	return filter(emojis, func(e *Emoji) bool {
		return strings.Contains(e.Name, s)
	})
}
