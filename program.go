package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gerardc/go-avro"
)

type Account struct {
	Id        int64
	Subdomain string
}

type Person struct {
	Name    string
	Age     int64
	Address *Address
}

type Address struct {
	Street string
}

func FetchAccount(id int64) (io.Reader, error) {
	res, err := http.Get(fmt.Sprintf("http://127.0.0.1:9292/accounts/%d", id))
	if err != nil {
		return nil, err
	}
	return res.Body, err
}

func ReadPerson(filename string) (Person, error) {
	datumReader := avro.NewGenericDatumReader()
	reader, err := avro.NewDataFileReader(filename, datumReader)

	if err != nil {
		return Person{}, err
		log.Panic("Unexpected error reading file", err)
	}
	person := Person{}
	reader.Next(&person)

	return person, nil
}

func ReadAccount(data io.Reader) (Account, error) {
	datumReader := avro.NewGenericDatumReader()
	reader, err := avro.NewDataReader(data, datumReader)

	if err != nil {
		log.Panic("Unexpected error reading file", err)
	}
	account := Account{}
	reader.Next(&account)

	return account, nil
}

func ReadAccountFile(filename string) (Account, error) {
	datumReader := avro.NewGenericDatumReader()
	reader, err := avro.NewDataFileReader(filename, datumReader)

	if err != nil {
		log.Panic("Unexpected error reading file", err)
	}
	account := Account{}
	reader.Next(&account)

	return account, nil
}

func main() {
	data, err := FetchAccount(1)
	if err != nil {
		log.Fatal(err)
	}

	account, err := ReadAccount(data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account)
}
