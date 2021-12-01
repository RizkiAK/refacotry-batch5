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
	PurchasedAt int64     `json:"purchased_at"`
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

	fmt.Println("1. Find items in the Meeting Room.")
	GetMeetingRoom(items)

	fmt.Println("=============================")
	fmt.Println("2. Find all electronic devices.")
	ElectronicDevice(items)

	fmt.Println("=============================")
	fmt.Println("3. Find all the furniture.")
	GetFurniture(items)

	fmt.Println("=============================")
	fmt.Println("4. Find all items were purchased on 16 Januari 2020.")
	GetDate(items)

	fmt.Println("=============================")
	fmt.Println("5. Find all items with brown color.")
	BrownColor(items)

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

func GetFurniture(items []Items) {
	for _, v := range items {
		if v.Type == "furniture" {
			fmt.Println(v)
		}
	}
}

func GetDate(items []Items) {
	layout := "2006-01-02"
	parse, _ := time.Parse(layout, "2020-01-16")
	p := parse.String()
	q := p[5:10]
	for _, v := range items {
		time := time.Unix(v.PurchasedAt, 0)
		timeString := time.String()
		if timeString[5:10] == q {
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
