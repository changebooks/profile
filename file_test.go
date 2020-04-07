package profile

import "testing"

func TestReadFile(t *testing.T) {
	_, got := ReadFile("undefined.ini")
	want := "file \"undefined.ini\" is not exist"
	if got == nil || got.Error() != want {
		t.Errorf("got %q; want %q", got, want)
	}
}

func TestReadDir(t *testing.T) {
	got, _ := ReadDir("profile")
	if len(got) != 6 ||
		got[0] != "app.ini" ||
		got[1] != "database.ini" ||
		got[2] != "keystore.ini" ||
		got[3] != "log.ini" ||
		got[4] != "redis.ini" ||
		got[5] != "sharding.ini" {
		t.Errorf("got %v; want [app.ini database.ini keystore.ini log.ini redis.ini sharding.ini]", got)
	}

	_, got2 := ReadDir("profile2")
	want2 := "directory \"profile2\" is not exist"
	if got2 == nil || got2.Error() != want2 {
		t.Errorf("got %q; want %q", got2, want2)
	}
}

func TestIsDirectory(t *testing.T) {
	got := IsDirectory("undefined")
	want := "directory \"undefined\" is not exist"
	if got == nil || got.Error() != want {
		t.Errorf("got %q; want %q", got, want)
	}

	got2 := IsDirectory("profile/app.ini")
	want2 := "directory \"profile/app.ini\" must be a directory"
	if got2 == nil || got2.Error() != want2 {
		t.Errorf("got %q; want %q", got2, want2)
	}

	got3 := IsDirectory("profile")
	if got3 != nil {
		t.Errorf("got %v; want <nil>", got3)
	}
}
