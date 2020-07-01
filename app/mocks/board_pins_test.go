package mocks

import (
	"app/repository"
	"reflect"
	"testing"
)

func TestBoardPinMock_CreateBoardPin(t *testing.T) {
	tests := []struct {
		name    string
		boardID int
		pinID   int
		wantErr bool
	}{
		{
			name:    "create board pin",
			boardID: 0,
			pinID:   0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := BoardPinMock{}
			if err := b.CreateBoardPin(tt.boardID, tt.pinID); (err != nil) != tt.wantErr {
				t.Errorf("CreateBoardPin() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewBoardPinRepository(t *testing.T) {
	tests := []struct {
		name string
		want repository.BoardPinRepository
	}{
		{
			name: "new board pin repository",
			want: &BoardPinMock{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBoardPinRepository(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBoardPinRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}
