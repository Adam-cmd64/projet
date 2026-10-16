package main

import "testing"

func TestProcess(t *testing.T) {
	cases := []struct{ in, want string }{
		{"1E (hex) files were added", "30 files were added"},
		{"It has been 10 (bin) years", "It has been 2 years"},
		{"Ready, set, go (up) !", "Ready, set, GO!"},
		{"I should stop SHOUTING (low)", "I should stop shouting"},
		{"Welcome to the Brooklyn bridge (cap)", "Welcome to the Brooklyn Bridge"},
		{"This is so exciting (up, 2)", "This is SO EXCITING"},
		{"over there ,and then BAMM !!", "over there, and then BAMM!!"},
		{"I was thinking ... You were right", "I was thinking... You were right"},
		{"they describe me: ' awesome '", "they describe me: 'awesome'"},
		{"As Elton John said: ' I am the most well-known homosexual in the world '", "As Elton John said: 'I am the most well-known homosexual in the world'"},
		{"There it was. A amazing rock!", "There it was. An amazing rock!"},

		{"If I make you BREAKFAST IN BED (low, 3) just say thank you instead of: how (cap) did you get in my house (up, 2) ?",
			"If I make you breakfast in bed just say thank you instead of: How did you get in MY HOUSE?"},
		{"I have to pack 101 (bin) outfits. Packed 1a (hex) just to be sure", "I have to pack 5 outfits. Packed 26 just to be sure"},
		{"Don not be sad ,because sad backwards is das . And das not good", "Don not be sad, because sad backwards is das. And das not good"},
		{"harold wilson (cap, 2) : ' I am a optimist ,but a optimist who carries a raincoat . '",
			"Harold Wilson: 'I am an optimist, but an optimist who carries a raincoat.'"},
		{"Simply add 42 (hex) and 10 (bin) and you will see the result is 68.", "Simply add 66 and 2 and you will see the result is 68."},
		{"Punctuation tests are ... kinda boring ,what do you think ?", "Punctuation tests are... kinda boring, what do you think?"},

		{"(up) hello", "hello"},
		{"go now (up, 9)", "GO NOW"},
		{"go (up) (low)", "go"},
		{"a, b (cap, 2)", "A, B"},
		{"don't stop", "don't stop"},
		{"élan (cap)", "Élan"},
		{"bRIDGE (cap)", "Bridge"},
		{"A hour", "An hour"},
		{"a 'amazing' day", "an 'amazing' day"},
		{"zz (hex)", "zz"},
	}
	for _, c := range cases {
		got := Process(c.in)
		if got != c.want {
			t.Error("entrée :", c.in, "| obtenu :", got, "| attendu :", c.want)
		}
	}
}
