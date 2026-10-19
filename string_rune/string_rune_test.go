package string_rune

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input string
		want  bool
	}{
		// 1. Граничные случаи длины
		{
			name:  "пустая строка",
			input: "",
			want:  true,
		},
		{
			name:  "один символ ASCII",
			input: "h",
			want:  true,
		},
		{
			name:  "один символ кириллицы",
			input: "я",
			want:  true,
		},

		// 2. ASCII палиндромы
		{
			name:  "ASCII палиндром нечётной длины",
			input: "racecar",
			want:  true,
		},
		{
			name:  "ASCII палиндром чётной длины",
			input: "noon",
			want:  true,
		},
		{
			name:  "ASCII палиндром с разным регистром",
			input: "Madam",
			want:  true,
		},
		{
			name:  "ASCII фраза с пробелами, запятыми и двоеточием",
			input: "A man, a plan, a canal: Panama",
			want:  true,
		},

		// 3. Кириллица
		{
			name:  "кириллица нечётной длины",
			input: "радар",
			want:  true,
		},
		{
			name:  "кириллица чётной длины",
			input: "Алла",
			want:  true,
		},
		{
			name:  "кириллица с разным регистром",
			input: "Шалаш",
			want:  true,
		},
		{
			name:  "кириллица фраза с пробелами и знаками препинания",
			input: "А роза упала на лапу Азора!",
			want:  true,
		},
		{
			name:  "кириллица фраза с буквой ё",
			input: "Лёша на полке клопа нашёл",
			want:  true,
		},

		// 4. Цифры и небуквенные строки
		{
			name:  "палиндром из цифр",
			input: "12321",
			want:  true,
		},
		{
			name:  "строка только из пробелов и знаков препинания",
			input: " ,.-!? ",
			want:  true,
		},

		// 5. Не-палиндромы
		{
			name:  "не-палиндром ASCII",
			input: "golang",
			want:  false,
		},
		{
			name:  "не-палиндром кириллица",
			input: "привет",
			want:  false,
		},
		{
			name:  "почти палиндром с одним несовпадающим символом",
			input: "abcbz",
			want:  false,
		},
		{
			name:  "не-палиндром из цифр",
			input: "12345",
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := IsPalindrome(tt.input); got != tt.want {
				t.Errorf("IsPalindrome(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
