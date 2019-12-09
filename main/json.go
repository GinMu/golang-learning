package main

import "encoding/json"

import "fmt"

// Screen define screen struct
type Screen struct {
	Size       float32
	ResX, ResY int
}

// Battery define battery struct
type Battery struct {
	Capacity int
}

func getJSONData() []byte {
	raw := &struct {
		Screen
		Battery
		HasTouchID bool
	}{
		Screen: Screen{
			Size: 5.5,
			ResX: 1920,
			ResY: 1080,
		},
		Battery: Battery{
			Capacity: 2910,
		},
		HasTouchID: true,
	}
	jsonData, _ := json.Marshal(raw)
	return jsonData
}

func main() {
	jsonData := getJSONData()
	fmt.Println(string(jsonData))
	screenAndTouch := struct {
		Screen
		HasTouchID bool
	}{}
	json.Unmarshal(jsonData, &screenAndTouch)
	fmt.Printf("%+v\n", screenAndTouch)
	batteryAndTouch := struct {
		Battery
		HasTouchID bool
	}{}
	json.Unmarshal(jsonData, &batteryAndTouch)
	fmt.Printf("%+v\n", batteryAndTouch)
}
