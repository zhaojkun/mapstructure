package mapstructure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type WeakBasic struct {
	Vbool bool
}

type WeakBasicResult struct {
	Bool2String1     string `json:"bool2string1"`
	Bool2String2     string `json:"bool2string2"`
	Int2String       string `json:"int2string"`
	Float2String     string `json:"float2string"`
	String2Bool1     bool   `json:"string2bool1"`
	String2Bool2     bool   `json:"string2bool2"`
	Int2Bool1        bool   `json:"int2bool1"`
	Int2Bool2        bool   `json:"int2bool2"`
	Bool2Int1        int    `json:"bool2int1"`
	Bool2Int2        int    `json:"bool2int2"`
	String2Int1      int    `json:"string2int1"`
	String2Int2      int    `json:"string2int2"`
	String2Int3      int    `json:"string2int3"`
	Map2String       string `json:"map2string"`
	Slice2String1    string `json:"slice2string1"`
	Slice2String2    string `json:"slice2string2"`
	Slice2String3    string `json:"slice2string3"`
	Interface2String string `json:"interface2string"`
}

func TestWeakInput(t *testing.T) {
	input := map[string]interface{}{
		"bool2string1": true,
		"bool2string2": false,
		"int2string":   0,
		"float2string": float64(0.0),
		"string2bool1": "true",
		"string2bool2": "false",
		"int2bool1":    0,
		"int2bool2":    100,
		"bool2int1":    false,
		"bool2int2":    true,
		"string2int1":  "0",
		"string2int2":  "1",
		"string2int3":  "",
		"map2string": map[string]int{
			"abc": 1,
		},
		"slice2string1": []byte("hello"),
		"slice2string2": []string{"hello", "world"},
		"slice2string3": []int{1, 2, 3},
		"interface2string": map[string]interface{}{
			"name": "zhaojkun",
			"age":  12,
		},
	}
	var res WeakBasicResult
	config := &DecoderConfig{
		WeaklyTypedInput: true,
		Result:           &res,
	}

	decoder, err := NewDecoder(config)
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	err = decoder.Decode(input)
	if err != nil {
		t.Fatalf("got an err: %s", err)
	}
	assert.Equal(t, "true", res.Bool2String1)
	assert.Equal(t, "false", res.Bool2String2)
	assert.Equal(t, "0", res.Int2String)
	assert.Equal(t, "0", res.Float2String)
	assert.Equal(t, true, res.String2Bool1)
	assert.Equal(t, false, res.String2Bool2)
	assert.Equal(t, false, res.Int2Bool1)
	assert.Equal(t, true, res.Int2Bool2)
	assert.Equal(t, 0, res.Bool2Int1)
	assert.Equal(t, 1, res.Bool2Int2)
	assert.Equal(t, 0, res.String2Int1)
	assert.Equal(t, 1, res.String2Int2)
	assert.Equal(t, 0, res.String2Int3)
	assert.Equal(t, `{"abc":1}`, res.Map2String)
	assert.Equal(t, "hello", res.Slice2String1)
	assert.Equal(t, `["hello","world"]`, res.Slice2String2)
	assert.Equal(t, `[1,2,3]`, res.Slice2String3)
	assert.Equal(t, `{"age":12,"name":"zhaojkun"}`, res.Interface2String)
}
