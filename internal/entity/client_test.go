package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewClient(t *testing.T) {
	client, err := NewClient("John Doe", "john@doe.com")
	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, "John Doe", client.Name)
	assert.Equal(t, "john@doe.com", client.Email)
}

func TestCreateNewClientWhenArgsInvalid(t *testing.T) {
	client, err := NewClient("", "")
	assert.NotNil(t, err)
	assert.Nil(t, client)
}

func TestUpdateClient(t *testing.T) {
	client, _ := NewClient("John Doe", "teste.gmail")
	err := client.Update("John Doe Update", "teste2@hotmail")
	assert.Nil(t, err)
	assert.Equal(t, "John Doe Update", client.Name)
	assert.Equal(t, "teste2@hotmail", client.Email)
}

func TestUpdateClientWhenArgsInvalid(t *testing.T) {
	client, _ := NewClient("John Doe", "teste.gmail")
	err := client.Update("", "teste.gmail")
	assert.NotNil(t, err)
	assert.Error(t, err, "name is required")
}

func TestAddAccountToClient(t *testing.T) {
	client, _ := NewClient("teste", "teste@hotmail.com")
	account := NewAccount(client)
	err := client.AddAccount(account)
	assert.Nil(t, err)
}
