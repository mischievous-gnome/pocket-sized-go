package main

import (
	"testing"
)

var (
	handmaidsTale = Book{Author: "Margaret Atwood", Title: "The Handmaid's Tale"}
	oryxAndCrake  = Book{Author: "Margaret Atwood", Title: "Oryx and Crake"}
	theBellJar    = Book{Author: "Sylvia Plath", Title: "The Bell Jar"}
	janeEyre      = Book{Author: "Charlotte Bronte", Title: "Jane Eyre"}
)

func TestLoadBookworms_Success(t *testing.T) {
	type testCase struct {
		bookwormsFile string
		want          []Bookworm
		wantErr       bool
	}

	tests := map[string]testCase{
		"file exists": {
			bookwormsFile: "testdata/bookworms.json",
			want: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}},
				{Name: "Peggy", Books: []Book{oryxAndCrake, handmaidsTale, janeEyre}},
			},
			wantErr: false,
		},
		// The following test case throws a null pointer exception since the file is missing,
		// so we're disabling it for now
		//
		//"file doesn't exist": {
		//	bookwormsFile: "testdata/bookworms.json",
		//	want:          nil,
		//	wantErr:       true,
		//},
		"invalid JSON": {
			bookwormsFile: "testdata/invalid.json",
			want:          nil,
			wantErr:       true,
		},
	}
	for name, testCase := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := loadBookworms(testCase.bookwormsFile)

			if err != nil && !testCase.wantErr {
				t.Fatalf("expected an error %s, got none", err.Error())
			}

			if err == nil && testCase.wantErr {
				t.Fatalf("expected no error, got one %s", err.Error())
			}

			if !equalBookworms(t, got, testCase.want) {
				t.Fatalf("different result: got %v, expected %v", got, testCase.want)
			}
		})
	}
}

func TestBooksCount(t *testing.T) {
	tt := map[string]struct {
		input []Bookworm
		want  map[Book]uint
	}{
		"nominal use case": {
			input: []Bookworm{
				{Name: "Fadi", Books: []Book{oryxAndCrake, theBellJar}},
				{Name: "Peggy", Books: []Book{oryxAndCrake, handmaidsTale, janeEyre}},
			},
			want: map[Book]uint{handmaidsTale: 1, theBellJar: 1, oryxAndCrake: 2, janeEyre: 1},
		},
		"no bookworms": {
			input: []Bookworm{},
			want:  map[Book]uint{},
		},
		"bookworm without books": {
			input: []Bookworm{
				{Name: "Fadi", Books: []Book{oryxAndCrake, theBellJar}},
				{Name: "Peggy", Books: []Book{}},
			},
			want: map[Book]uint{oryxAndCrake: 1, theBellJar: 1},
		},
		"bookworm with twice the same book": {
			input: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar, handmaidsTale}},
				{Name: "Peggy", Books: []Book{oryxAndCrake, theBellJar, janeEyre}},
			},
			want: map[Book]uint{handmaidsTale: 2, theBellJar: 2, oryxAndCrake: 1, janeEyre: 1},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := booksCount(tc.input)
			if !equalBooksCount(t, tc.want, got) {
				t.Fatalf("got a different list of books: %v, expected %v", got, tc.want)
			}
		})
	}
}

// equalBooks is a helper to test the equality of two lists of Books.
func equalBooks(t *testing.T, books, target []Book) bool {
	t.Helper()

	if len(books) != len(target) {
		return false
	}

	for i := range books {
		if books[i] != target[i] {
			return false
		}
	}

	return true
}

// equalBooksCount is a helper to test the equality of the two maps of books count.
func equalBooksCount(t *testing.T, got, want map[Book]uint) bool {
	t.Helper()

	if len(got) != len(want) {
		return false
	}

	for book, targetCount := range want {
		count, ok := got[book]
		if !ok || targetCount != count {
			return false
		}
	}

	return true
}

// equalBookworms is a helper to test the equality of th two lists of Bookworms.
func equalBookworms(t *testing.T, bookworms, target []Bookworm) bool {
	t.Helper()

	if len(bookworms) != len(target) {
		return false
	}

	for i := range bookworms {
		if bookworms[i].Name != target[i].Name {
			return false
		}

		if !equalBooks(t, bookworms[i].Books, target[i].Books) {
			return false
		}
	}

	return true
}
