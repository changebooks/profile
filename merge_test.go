package profile

import "testing"

func TestMerge(t *testing.T) {
	got := Merge(nil, nil)
	if got != nil {
		t.Errorf("got %v; want map[]", got)
	}

	got2 := Merge(map[string]string{}, nil)
	if got2 != nil {
		t.Errorf("got %v; want map[]", got2)
	}

	got3 := Merge(nil, map[string]string{})
	if len(got3) != 0 {
		t.Errorf("got %v; want map[]", got3)
	}

	got4 := Merge(map[string]string{}, map[string]string{})
	if len(got4) != 0 {
		t.Errorf("got %v; want map[]", got4)
	}
}

func TestMerge2(t *testing.T) {
	got := Merge(nil, map[string]string{"a": "a1"})
	if len(got) != 1 || got["a"] != "a1" {
		t.Errorf("got %v; want map[a:a1]", got)
	}

	got2 := Merge(map[string]string{}, map[string]string{"b": "b2"})
	if len(got2) != 1 || got2["b"] != "b2" {
		t.Errorf("got %v; want map[b:b2]", got2)
	}

	got3 := Merge(map[string]string{"c": "c3"}, nil)
	if len(got3) != 1 || got3["c"] != "c3" {
		t.Errorf("got %v; want map[c:c3]", got3)
	}

	got5 := Merge(
		map[string]string{"a": "a1", "c": "c3", "d": "d4", "x": "x24"},
		map[string]string{"a": "a1-1", "b": "b2", "c": "c3-3", "y": "y25"})
	if len(got5) != 6 ||
		got5["a"] != "a1-1" || got5["c"] != "c3-3" || got5["d"] != "d4" ||
		got5["x"] != "x24" || got5["b"] != "b2" || got5["y"] != "y25" {
		t.Errorf("got %v; want [b:b2 y:y25 a:a1-1 c:c3-3 d:d4 x:x24]", got5)
	}
}
