package main

import "testing"

func Test_getSizeAndModTime(t *testing.T) {
	type args struct {
		pathAndFileName string
	}
	tests := []struct {
		name  string
		args  args
		want  int64
		want1 string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := getSizeAndModTime(tt.args.pathAndFileName)
			if got != tt.want {
				t.Errorf("getSizeAndModTime() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("getSizeAndModTime() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
