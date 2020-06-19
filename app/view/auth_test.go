package view

import (
	"io/ioutil"
	"reflect"
	"strings"
	"testing"
)

func TestNewSignInResponse(t *testing.T) {
	test := struct {
		token  string
		userID int
	}{
		"$2a$fe.hfefefefelplfepv6Y2DYZ/Dbwyr.AKtlAldFe",
		0,
	}

	want := &SignInResponse{Token: test.token, UserID: test.userID}
	got := NewLSignInResponse(test.token, test.userID)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("want: %v, got: %v", want, got)
	}
}

func TestNewSignUpResponse(t *testing.T) {
	test := struct {
		token  string
		userID int
	}{
		"$2a$fe.hfefefefelplfepv6Y2DYZ/Dbwyr.AKtlAldFe",
		0,
	}

	want := &SignUpResponse{Token: test.token, UserID: test.userID}
	got := NewLSignUpResponse(test.token, test.userID)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("want: %v, got: %v", want, got)
	}
}

func TestNewSignInReequest(t *testing.T) {
	tests := []struct {
		desc    string
		request string
		want    *SignInRequest
		success bool
	}{
		{
			"valid case1",
			`{"email": "email@example.com", "password": "KOK3884jffii233"}`,
			&SignInRequest{"email@example.com", "KOK3884jffii233"},
			true,
		},
		{
			"no email",
			`{"password": "KOK3884jffii233"}`,
			nil,
			false,
		},
		{
			"no password",
			`{"email": "email@example.com"`,
			nil,
			false,
		},
		{
			"invalid email",
			`{"email": "emailexample.com", "password": "KOK3884jffii233"}`,
			nil,
			false,
		},
	}

	for _, tt := range tests {
		got, err := NewSignInRequest(ioutil.NopCloser(strings.NewReader(tt.request)))

		if tt.success {
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("desc: %v, want: %v, got: %v", tt.desc, tt.want, got)
			}
		} else {
			if err == nil {
				t.Errorf("Error should occur: %v", tt.desc)
			}
		}
	}
}

func TestNewSignUpReequest(t *testing.T) {
	tests := []struct {
		desc    string
		request string
		want    *SignUpRequest
		success bool
	}{
		{
			"valid case1",
			`{"name": "ryoma", "email": "email@example.com", "password": "KOK3884jffii233"}`,
			&SignUpRequest{"ryoma", "email@example.com", "KOK3884jffii233"},
			true,
		},
		{
			"no email",
			`{"name": "ryoma", "password": "KOK3884jffii233"}`,
			nil,
			false,
		},
		{
			"no password",
			`{"name": "ryoma", "email": "email@example.com"`,
			nil,
			false,
		},
		{
			"invalid email",
			`{"name": "ryoma", "email": "emailexample.com", "password": "KOK3884jffii233"}`,
			nil,
			false,
		},
		{
			"password too short",
			`{"name": "ryoma", "email": "emailexample.com", "password": "Kdfd3"}`,
			nil,
			false,
		},
		{
			"password contains only lower",
			`{"name": "ryoma", "email": "emailexample.com", "password": "aaaaaaaaaaaaaaaaa"}`,
			nil,
			false,
		},
		{
			"password contains only capital",
			`{"name": "ryoma", "email": "emailexample.com", "password": "AAAAAAAAAAAAAAAAAA"}`,
			nil,
			false,
		},
		{
			"password contains only number",
			`{"name": "ryoma", "email": "emailexample.com", "password": "1111111111111111"}`,
			nil,
			false,
		},
	}

	for _, tt := range tests {
		got, err := NewSignUpRequest(ioutil.NopCloser(strings.NewReader(tt.request)))

		if tt.success {
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("desc: %v, want: %v, got: %v", tt.desc, tt.want, got)
			}
		} else {

			if err == nil {
				t.Errorf("Error should occur: %v", tt.desc)
			}
		}
	}
}
