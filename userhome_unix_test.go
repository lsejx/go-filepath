package fpath

import "testing"

func TestAbsHome(t *testing.T) {
	tests := []struct {
		arg string
		ret string
		err error
	}{
		{"/aiueo", "/aiueo", nil},
		{"~", "/root", nil},
		{"~~", "~~", nil},
		{"~/", "/root/", nil},
		{"~/aiueo", "/root/aiueo", nil},
		{"aiueo/~", "aiueo/~", nil},
	}

	for _, tt := range tests {
		s, err := AbsHome(tt.arg)
		if tt.err == nil {
			if err != nil {
				t.Fatalf("a:%v, err:%v", tt.arg, err)
			}
			if s != tt.ret {
				t.Fatalf("a:%v, s:%v, w:%v", tt.arg, s, tt.ret)
			}
		} else {
			if err == nil {
				t.Fatalf("a:%v, nilerr, s:%v", tt.arg, s)
			}
			if err.Error() != tt.err.Error() {
				t.Fatalf("a:%v, err:%v, w:%v", tt.arg, err, tt.err)
			}
		}
	}
}
