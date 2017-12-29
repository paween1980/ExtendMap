package extend_map

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestExtendMap(t *testing.T) {
	a := map[string]interface{}{
		"ak1": "av1",
		"ak2": "av2",
		"dup1": map[string]interface{}{
			"ak4": "av4",
			"ak5": "av5",
			"ak6": map[string]interface{}{
				"ak7": "av7",
				"ak8": "av8",
			},
			"dup2": map[string]interface{}{
				"ak9":  "av9",
				"ak10": "av10",
			},
		},
		"dup3": map[string]interface{}{
			"ak11": "av11",
		},
		"dup4": "av12",
	}
	b := map[string]interface{}{
		"bk1": "bv1",
		"bk2": "bv2",
		"dup1": map[string]interface{}{
			"bk3": "bv3",
			"bk4": "bv4",
			"dup2": map[string]interface{}{
				"bk5": "bv5",
				"bk6": "bv6",
			},
		},
		"dup3": "bv7",
		"dup4": map[string]interface{}{
			"bk8": "bv8",
		},
	}
	x := map[string]interface{}{
		"ak1": "av1",
		"ak2": "av2",
		"dup1": map[string]interface{}{
			"ak4": "av4",
			"ak5": "av5",
			"ak6": map[string]interface{}{
				"ak7": "av7",
				"ak8": "av8",
			},
			"bk3": "bv3",
			"bk4": "bv4",
			"dup2": map[string]interface{}{
				"ak9":  "av9",
				"ak10": "av10",
				"bk5":  "bv5",
				"bk6":  "bv6",
			},
		},
		"bk1":  "bv1",
		"bk2":  "bv2",
		"dup3": "bv7",
		"dup4": map[string]interface{}{
			"bk8": "bv8",
		},
	}
	result := ExtendMap(a, b)
	if !reflect.DeepEqual(result, x) {
		t.Errorf("Expect ok but not\n%#v\n%#v\n", result, x)
	}
	if reflect.DeepEqual(a, result) {
		t.Error("Expect clone new map but it appends")
	}
	if reflect.DeepEqual(b, result) {
		t.Error("Expect clone new map but it appends")
	}
}

func TestExtendMap_inputNil(t *testing.T) {
	a := map[string]interface{}{
		"key1": "value1",
	}
	result := ExtendMap(a, nil)
	if !reflect.DeepEqual(result, a) {
		t.Errorf("Expect result should be %#v\n", a)
	}

	result = ExtendMap(nil, a)
	if !reflect.DeepEqual(result, a) {
		t.Errorf("Expect result should be %#v\n", a)
	}

	nilMap := map[string]interface{}{}
	result = ExtendMap(a, nilMap)
	if !reflect.DeepEqual(result, a) {
		t.Errorf("Expect result should be %#v\n", a)
	}

	result = ExtendMap(nilMap, a)
	if !reflect.DeepEqual(result, a) {
		t.Errorf("Expect result should be %#v\n", a)
	}

	result = ExtendMap(nil, nil)
	if !reflect.DeepEqual(result, nilMap) {
		t.Error("Expect result should be nilMap but", result)
	}

	b, _ := json.Marshal(result)
	if string(b) != "{}" {
		t.Error("Expect json.Marshal nilMap should be {} but", string(b))
	}
}
