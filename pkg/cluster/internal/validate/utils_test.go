package validate

import (
	"fmt"
	"testing"
)

type TestStruct struct {
	Field1 int
	Field2 int
}

func TestValidateStruct(t *testing.T) {
	testCases := []struct {
		name string
		s    TestStruct
		want error
	}{
		{
			name: "All fields set",
			s:    TestStruct{Field1: 1, Field2: 1},
			want: nil,
		},
		{
			name: "Missing fields",
			s:    TestStruct{Field1: 1},
			want: fmt.Errorf("<nil>Field2 in not set; "),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := validateStruct(tc.s)
			if got == nil || tc.want == nil {
				if got != tc.want {
					t.Errorf("validateStruct() = %v, want %v", got, tc.want)
				}
			} else {
				gotMsg := got.Error()
				wantMsg := tc.want.Error()
				if gotMsg != wantMsg {
					t.Errorf("validateStruct() = %v (len %v), want %v (len %v)", gotMsg, len(gotMsg), wantMsg, len(wantMsg))
				}
			}
		})
	}
}
