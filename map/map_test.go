package _map

import (
	"fmt"
	"testing"
)

func TestMGet_2levelsDeep(t *testing.T) {
	fmt.Println("test with map[string]string")
	tMap := map[string]interface{}{
		"key1": "value1",
		"key2": map[string]string{
			"key21": "value2",
		},
		"key3": map[string]interface{}{
			"key31": []string{
				"value31",
				"value32",
			},
		},
	}

	actual := MGet[[]string](tMap, "key3.key31")

	if actual == nil {
		fmt.Print("Element not found.")
		t.Fail()
	}
	if actual[0] != "value31" && actual[1] != "value32" {
		t.Fail()
	}
}

func TestMGet_1levelDeep(t *testing.T) {
	fmt.Println("test with map[string]string")
	tMap := map[string]interface{}{
		"key1": "value1",
		"key2": map[string]string{
			"key21": "value2",
		},
		"key3": map[string]interface{}{
			"key31": []string{
				"value31",
				"value32",
			},
		},
	}

	actual := MGet[string](tMap, "key1")

	if actual != "value1" {
		fmt.Print("Element not found.")
		t.Fail()
	}
}

func TestMGet_3levelDeep(t *testing.T) {
	fmt.Println("test with map[string]string")
	tMap := map[string]interface{}{
		"key1": "value1",
		"key2": map[string]interface{}{
			"key21": map[string]interface{}{
				"key211": 211,
				"key212": 212,
			},
			"key3": map[string]interface{}{
				"key31": []string{
					"value31",
					"value32",
				},
			},
		},
	}

	actual := MGet[int](tMap, "key2.key21.key212")

	if actual != 212 {
		fmt.Print("Element not found.")
		t.Fail()
	}
}
