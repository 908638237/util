package _string

import (
	"fmt"
	"testing"
)

func TestSubString(t *testing.T) {
	type args struct {
		str    string
		begin  int
		length int
	}
	tests := []struct {
		name       string
		args       args
		wantSubstr string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSubstr := SubString(tt.args.str, tt.args.begin, tt.args.length); gotSubstr != tt.wantSubstr {
				t.Errorf("SubString() = %v, want %v", gotSubstr, tt.wantSubstr)
			}
		})
	}
}

func TestToHump(t *testing.T) {
	for i := 0; i <= 20000; i++ {
		t.Run("role_info", func(t *testing.T) {
			fmt.Println(ToHump("role_info"))
		})
	}
}
