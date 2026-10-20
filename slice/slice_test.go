package slice

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCopyWithout(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		target int
		want   []int
	}{
		{
			name:   "nil срез",
			input:  nil,
			target: 5,
			want:   nil,
		},
		{
			name:   "пустой срез",
			input:  []int{},
			target: 5,
			want:   []int{},
		},
		{
			name:   "отсутствующее значение",
			input:  []int{1, 2, 3},
			target: 4,
			want:   []int{1, 2, 3},
		},
		{
			name:   "повторяющиеся элементы",
			input:  []int{1, 2, 2, 2, 3, 2, 4},
			target: 2,
			want:   []int{1, 3, 4},
		},
		{
			name:   "удаление всех элементов",
			input:  []int{2, 2, 2, 2},
			target: 2,
			want:   []int{},
		},
		{
			name:   "значение в начале среза",
			input:  []int{2, 1, 3},
			target: 2,
			want:   []int{1, 3},
		},
		{
			name:   "значение в конце среза",
			input:  []int{1, 3, 2},
			target: 2,
			want:   []int{1, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			origInputCopy := slices.Clone(tt.input)

			got := CopyWithout(tt.input, tt.target)

			assert.Equal(t, tt.want, got)

			if tt.input == nil {
				assert.Nil(t, got)
			} else {
				assert.NotNil(t, got)
				assert.Equal(t, origInputCopy, tt.input)
			}
		})
	}
}

func TestDeleteInPlace(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		target int
		want   []int
	}{
		{
			name:   "nil срез",
			input:  nil,
			target: 5,
			want:   nil,
		},
		{
			name:   "пустой срез",
			input:  []int{},
			target: 5,
			want:   []int{},
		},
		{
			name:   "отсутствующее значение",
			input:  []int{1, 2, 3},
			target: 4,
			want:   []int{1, 2, 3},
		},
		{
			name:   "повторяющиеся элементы",
			input:  []int{1, 2, 2, 2, 3, 2, 4},
			target: 2,
			want:   []int{1, 3, 4},
		},
		{
			name:   "удаление всех элементов",
			input:  []int{2, 2, 2, 2},
			target: 2,
			want:   []int{},
		},
		{
			name:   "значение в начале среза",
			input:  []int{2, 1, 3},
			target: 2,
			want:   []int{1, 3},
		},
		{
			name:   "значение в конце среза",
			input:  []int{1, 3, 2},
			target: 2,
			want:   []int{1, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputCopy := slices.Clone(tt.input)
			got := DeleteInPlace(inputCopy, tt.target)

			assert.Equal(t, tt.want, got)

			if tt.input == nil {
				assert.Nil(t, got)
			} else {
				assert.NotNil(t, got)
			}
		})
	}
}

func TestDeleteInPlace_ZeroesTail(t *testing.T) {
	underlying := []int{10, 20, 20, 30, 20, 40}
	initialCap := cap(underlying)

	res := DeleteInPlace(underlying, 20)

	assert.Equal(t, []int{10, 30, 40}, res)
	assert.Equal(t, 3, len(res))

	fullBackingArray := underlying[:initialCap]
	assert.Equal(t, []int{10, 30, 40, 0, 0, 0}, fullBackingArray)
}

func TestGenericsSupport(t *testing.T) {
	t.Run("CopyWithout со строками", func(t *testing.T) {
		input := []string{"apple", "banana", "apple", "cherry"}
		got := CopyWithout(input, "apple")
		assert.Equal(t, []string{"banana", "cherry"}, got)
	})

	t.Run("DeleteInPlace со строками и занулением хвоста", func(t *testing.T) {
		input := []string{"go", "rust", "go", "python"}
		got := DeleteInPlace(input, "go")
		assert.Equal(t, []string{"rust", "python"}, got)

		assert.Equal(t, []string{"rust", "python", "", ""}, input[:4])
	})
}
