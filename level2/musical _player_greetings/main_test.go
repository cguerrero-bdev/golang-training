package main

import "testing"

func TestGreating(t *testing.T) {

	testCases := []struct {
		musicalPlayer MusicalPlayer
		want          string
	}{
		{&Trumpeter{"Mike"}, "Hi, my name is Mike and I'm a trumpeter\n"},
		{&Violinist{"John"}, "Hi, my name is John and I'm a violinist\n"},
		{&Violinist{"Lois"}, "Hi, my name is Lois and I'm a violinist\n"},
		{&Trumpeter{"Peter"}, "Hi, my name is Peter and I'm a trumpeter\n"},
	}

	for i, testCase := range testCases {
		if testCase.musicalPlayer.Greetings() != testCase.want {

			t.Errorf("Test case %d. Expected message \"%s\", returned message: \"%s\" ", i, testCase.want, testCase.musicalPlayer.Greetings())
		}
	}
}
