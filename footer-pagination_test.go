package main

import "testing"

func TestFooterPagination(t *testing.T) {
	cases := []struct {
		currentPage, totalPages, boundaries, around int
		want                                        string
	}{
		{4, 5, 1, 0, "1 ... 4 5"},
		{4, 10, 2, 2, "1 2 3 4 5 6 ... 9 10"},
		{13, 13, 0, 0, "... 13"},
		{13, 30, 3, 5, "1 2 3 ... 8 9 10 11 12 13 14 15 16 17 18 ... 28 29 30"},
		{3, 14, 0, 0, "... 3 ... "},
	}
	for _, c := range cases {
		got := FooterPagination(c.currentPage, c.totalPages, c.boundaries, c.around)
		if got != c.want {
			t.Errorf("FooterPagination(%q) == %q, want %q", c, got, c.want)
		}
	}
}
