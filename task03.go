package tasks

import (
	"encoding/json"
	"errors"
	"os"
)

type AddressesDB struct {
	addresses []AddressInfo `json:"Addresses"`
}

type AddressInfo struct {
	FIO      string
	address  Address
	phoneNum string
}

type Address struct {
	city string
	index uint
	street string
	houseNum uint
	apartNum uint
	buildingType string
	storeysNum uint8
	inhabitantsNum uint
}

func (db *AddressesDB) AddAddress(addressInfo AddressInfo) {
	db.addresses = append(db.addresses, addressInfo)
}

func (db *AddressesDB) RemoveAddress(index int) {
	db.addresses = append(db.addresses[:index], db.addresses[index+1:]...)
}

func (db *AddressesDB) GetStringFromAddress(index int) (string, error) {
	resultBytes, err := json.Marshal(db.addresses[index])
	return string(resultBytes), err
}

func (db *AddressesDB) GetPhoneNumber(index int) (string, error) {
	if index >= len(db.addresses) {
		return "", errors.New("index out of range")
	}
	return db.addresses[index].phoneNum, nil
}

func (db *AddressesDB) GetCity(index int) (string, error) {
	if index >= len(db.addresses) {
		return "", errors.New("index out of range")
	}
	return db.addresses[index].address.city, nil
}

func (db *AddressesDB) GetIndex(index int) (uint, error) {
	if index >= len(db.addresses) {
		return 0, errors.New("index out of range")
	}
	return db.addresses[index].address.index, nil
}

func (db *AddressesDB) CountAddressesInCity(city string) (count uint) {
	for _, value := range db.addresses {
		if value.address.city == city {
			count++
		}
	}
	return
}

func (db *AddressesDB) Save() error {
	resultBytes, err := json.Marshal(db)
	if err != nil {
		return err
	}

	file, err := os.Create("hello.txt")
	if err != nil {
		return err
	}

	defer file.Close()
	file.WriteString(string(resultBytes))

	return nil
}

func (db *AddressesDB) Load() error {
	file, err := os.Open("hello.txt")
	if err != nil {
		return err
	}

	bytes := make([]byte, 10000)

	defer file.Close()
	file.Read(bytes)

	err = json.Unmarshal(bytes, db)
	if err != nil {
		return err
	}

	return nil
}