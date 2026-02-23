package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Client struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewClient(name, email string) (*Client, error) {
	client := &Client{
		ID:        uuid.New().String(),
		Name:      name,
		Email:     email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := client.ValidateClient(); err != nil {
		return nil, err
	}
	return client, nil
}

func (c *Client) ValidateClient() error {
	if c.Name == "" {
		return errors.New("name is required")
	}
	if c.Email == "" {
		return errors.New("email is required")
	}
	return nil
}

func (c *Client) Update(name, email string) error {
	c.Name = name
	c.Email = email
	c.UpdatedAt = time.Now()

	if err := c.ValidateClient(); err != nil {
		return err
	}
	return nil
}
