package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Items struct {
	InventoryID int       `json:"inventory_id"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`
	Tags        []string  `json:"tags"`
	PurchasedAt int       `json:"purchased_at"`
	Placement   Placement `json:"placement"`
}

type Placement struct {
	RoomID int    `json:"room_id"`
	Name   string `json:"name"`
}

func main() {
	reader, _ := os.Open("data.json")
	decoder := json.NewDecoder(reader)

	items := []Items{}
	decoder.Decode(&items)

	// GetMeetingRoom(items)
	// ElectronicDevice(items)
	// GetAll(items)
	GetDate(items)
	// BrownColor(items)

}

func GetMeetingRoom(items []Items) {
	for _, v := range items {
		if v.Placement.Name == "Meeting Room" {
			fmt.Println(v)
		}
	}
}

func ElectronicDevice(items []Items) {
	for _, v := range items {
		if v.Type == "electronic" {
			fmt.Println(v)
		}
	}
}

func GetAll(items []Items) {
	for _, v := range items {
		fmt.Println(v)
	}
}

func GetDate(items []Items) {
	layout := "2006-01-02"
	parse, _ := time.Parse(layout, "2020-01-16")
	p := parse.String()
	q := p[5:10]
	for _, v := range items {
		purchaseAt := time.Date(0, 0, 0, 0, 0, v.PurchasedAt*1000, 0, time.UTC)
		purchaseAtString := purchaseAt.String()
		if purchaseAtString[6:11] == q {
			fmt.Println(v)
		}
	}
}

func BrownColor(items []Items) {
	for _, v := range items {
		for _, val := range v.Tags {
			if val == "brown" {
				fmt.Println(v)
			}
		}
	}
}
