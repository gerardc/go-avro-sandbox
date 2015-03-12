package main

import (
	"os"
	"testing"
)

func TestReadAccount(t *testing.T) {
	file, _ := os.Open("account.avro")
	account, err := ReadAccount(file)
	if err != nil {
		t.Error(err)
	}
	if account.Id != 1 || account.Subdomain != "support" {
		t.Error("Account not parsed correctly", account)
	}
}

func TestReadAccountFile(t *testing.T) {
	account, err := ReadAccountFile("account.avro")
	if err != nil {
		t.Error(err)
	}
	if account.Id != 1 || account.Subdomain != "support" {
		t.Error("Account not parsed correctly", account)
	}
}

func TestReadPersonFile(t *testing.T) {
	person, err := ReadPersonFile("person.avro")
	if err != nil {
		t.Error(err)
	}

	if person.Name != "Johnny Logan" {
		t.Error("Did not parse name", person)
	}
	if person.Address.Street != "right here st." {
		t.Error("Did not parse Address.Street", person.Address.Street)
	}
}
