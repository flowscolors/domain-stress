package statistics

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPrintMap(t *testing.T) {

	tt := map[string]struct {
		input  map[int]int
		result string
	}{
		"test1": {input: map[int]int{
			200: 50,
			500: 10,
			100: 20,
		}, result: "100:20;200:50;500:10"},
	}

	for _, value := range tt {
		str := printMap(value.input)
		fmt.Println(str)
		if !reflect.DeepEqual(value.result, str) {
			t.Errorf("数据不一致 预期:%v 实际:%v", value.result, str)
		}
	}
}
