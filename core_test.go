package main

import "testing"

func TestCreateUserAccount(t *testing.T) {
	login := "testLogin"
	email := "test@mail.ru"
	password := "testPassword"
	testCore := Core{users: make(map[string]User)}
	testRequest := SignupRequest{
		Login:    login,
		Password: password,
		Email:    email,
	}

	testCore.CreateUserAccount(testRequest)

	_, foundAccount := testCore.FindUserAccount(login)
	if !foundAccount {
		t.Errorf("user not found")
	}
}

func TestCreateAndKillSession(t *testing.T) {
	login := "testLogin"
	testCore := Core{sessions: make(map[string]Session)}

	sid, _ := testCore.CreateSession(login)
	isFound := testCore.FindActiveSession(sid)
	if !isFound {
		t.Errorf("session not found")
	}

	testCore.KillSession(sid)

	isFound = testCore.FindActiveSession(sid)
	if isFound {
		t.Errorf("found killed session")
	}
}
