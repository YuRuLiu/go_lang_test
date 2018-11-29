package main

import (
	"testing"
)

//右鍵選擇Generate Unit Tests for Function就會自動產生了測試檔了
func TestGetMachineGame(t *testing.T) {
	type args struct {
		MID string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "DB取無資料",
			args: args{MID: "123"},
			want: 0,
		},
		{
			name: "開分1:1000,餘額10",
			args: args{MID: "456"},
			want: 1,
		},
		{
			name: "開分100:1,餘額100",
			args: args{MID: "789"},
			want: 1,
		},
		{
			name: "開分1:1,餘額50",
			args: args{MID: "987"},
			want: 1,
		},
		{
			name: "開分無資料",
			args: args{MID: "654"},
			want: 0,
		},
		{
			name: "開分1:1000,餘額0",
			args: args{MID: "321"},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMachineGame(tt.args.MID); got != tt.want {
				t.Errorf("GetMachineGame() = %v, want %v", got, tt.want)
			}
		})
	}
}
