package models

import (
	"fmt"

	"github.com/go-pg/pg/v10"
)

type Client struct {
	tableName struct{} `pg:"client"`
	ID        int      `json:"id" pg:"id,pk"`
	Name      string   `json:"name" pg:"name"`
}

var clients []Client

func (c *Client) InsertClient(db *pg.DB) error {
	_, err := db.Model(c).Insert()
	if err != nil {
		fmt.Println("cannot save client to db ", err)
	}

	return err
}

func (c *Client) SelectClients(db *pg.DB) ([]Client, error) {
	err := db.Model(&clients).Select()

	return clients, err
}

func (c *Client) SelectClient(db *pg.DB, id int) (*Client, error) {
	err := db.Model(c).Where("id=?", id).Select()

	return c, err
}

func (c *Client) UpdateClient(db *pg.DB, id int) error {
	_, err := db.Model(c).Where("id=?", id).Update()

	return err
}

func (c *Client) DeleteClient(db *pg.DB, id int) error {
	_, err := db.Model(c).Where("id=?", id).Delete()

	return err
}
