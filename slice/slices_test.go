package slice

import (
	"fmt"
	"strconv"
	"testing"
)

func TestFind(t *testing.T) {
	fmt.Println("test with map[string]string")
	testArr := []map[string]string{
		{
			"key": "value1",
		},
		{
			"key": "value2",
		},
		{
			"key": "value3",
		},
	}

	k := "key"
	v := "value2"

	finder := func(el map[string]string) bool {
		return el[k] == v
	}
	actual := Find(testArr, finder)

	if actual == nil {
		fmt.Print("Element not found.")
		t.Fail()
	}
}

func TestMap(t *testing.T) {
	testArr := []map[string]string{
		{
			"key": "value1",
		},
		{
			"key": "value2",
		},
		{
			"key": "value3",
		},
	}

	mapper := func(el map[string]string, idx int) interface{} {
		el["key"] = "value" + strconv.Itoa(idx+1)
		return el
	}

	actual := Map[map[string]string](testArr, mapper)
	if len(actual) != len(testArr) {
		t.Fail()
	}
	if actual[0].(map[string]string)["key"] != "value1" &&
		actual[1].(map[string]string)["key"] != "value2" &&
		actual[2].(map[string]string)["key"] != "value3" {
		t.Fail()
	}
}

func TestPop(t *testing.T) {
	tArr := []string{
		"val1",
		"val2",
		"val3",
	}

	actualArr, actual := Pop(tArr)

	if len(actualArr) != 2 {
		fmt.Println("Invalid popped slice length")
		t.Fail()
	}

	if actual != "val3" {
		fmt.Println("Invalid popped element")
		t.Fail()
	}
}
