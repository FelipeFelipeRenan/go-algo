package main

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "exemplo 1 - abcabcbb",
			input: "abcabcbb",
			want:  3, // "abc"
		},
		{
			name:  "exemplo 2 - bbbbb",
			input: "bbbbb",
			want:  1, // "b"
		},
		{
			name:  "exemplo 3 - pwwkew",
			input: "pwwkew",
			want:  3, // "wke"
		},
		{
			name:  "string vazia",
			input: "",
			want:  0,
		},
		{
			name:  "string sem repetição",
			input: "abcdefg",
			want:  7,
		},
		{
			name:  "string com um caractere",
			input: "a",
			want:  1,
		},
		{
			name:  "repetição no final",
			input: "dvdf",
			want:  3, // "vdf"
		},
		{
			name:  "repetição no início",
			input: "abba",
			want:  2, // "ab" ou "ba"
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Supondo que sua função se chame LengthOfLongestSubstring
			// Se nomear de forma diferente, apenas ajuste aqui.
			got := LengthOfLongestSubstring(tt.input)
			if got != tt.want {
				t.Errorf("LengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}