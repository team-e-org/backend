package mocks

import (
	"app/repository"
	"reflect"
	"testing"
)

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

func TestBoardPinMock_DeleteBoardPin(t *testing.T) {
	b := BoardPinMock{}

	err := b.CreateBoardPin(1, 1)
	if err != nil {
		t.Fatalf("an error occurred")
	}

	tests := []struct {
		name    string
		boardID int
		pinID   int
		wantErr bool
	}{
		{
			name:    "delete board pin",
			boardID: 1,
			pinID:   1,
			wantErr: false,
		},
		{
			name:    "wrong board id",
			boardID: 99,
			pinID:   1,
			wantErr: true,
		},
		{
			name:    "wrong pin id",
			boardID: 1,
			pinID:   99,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := b.DeleteBoardPin(tt.boardID, tt.pinID); (err != nil) != tt.wantErr {
				t.Errorf("DeleteBoardPin() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
