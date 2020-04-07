package profile

import "testing"

func TestSanitiseName(t *testing.T) {
	got := sanitiseName("a'b\"c#d")
	want := "abc"
	if got != want {
		t.Errorf("got %q; want %q", got, want)
	}

	got2 := sanitiseName("a b\\c\td")
	want2 := "abc"
	if got2 != want2 {
		t.Errorf("got %q; want %q", got2, want2)
	}

	got3 := sanitiseName("  a bc   \nd ")
	want3 := "abc"
	if got3 != want3 {
		t.Errorf("got %q; want %q", got3, want3)
	}
}

func TestSanitiseValue(t *testing.T) {
	got := sanitiseValue("a'b\"c#d")
	want := "a'b\"c"
	if got != want {
		t.Errorf("got %q; want %q", got, want)
	}

	got2 := sanitiseValue("a b\\c\rd")
	want2 := "a b\\c"
	if got2 != want2 {
		t.Errorf("got %q; want %q", got2, want2)
	}

	got3 := sanitiseValue("  a bc   \nd ")
	want3 := "  a bc   "
	if got3 != want3 {
		t.Errorf("got %q; want %q", got3, want3)
	}
}
