package qrcode

import (
	"testing"
)

// func TestInit() {
// 	load(defaultPathToCfg)
// }

func TestEncodeNum(t *testing.T) {
	enc := Encoder{
		ecLv:    Low,
		mode:    EncModeNumeric,
		version: loadVersion(1, Low),
	}

	b, err := enc.Encode([]byte("12312312"))
	if err != nil {
		t.Errorf("could not encode: %v", err)
		t.Fail()
	}
	t.Log(b, b.Len())
}

func TestEncodeAlphanum(t *testing.T) {
	enc := Encoder{
		ecLv:    Low,
		mode:    EncModeAlphanumeric,
		version: loadVersion(1, Low),
	}

	b, err := enc.Encode([]byte("AKJA*:/"))
	if err != nil {
		t.Errorf("could not encode: %v", err)
		t.Fail()
	}
	t.Log(b, b.Len())
}

func TestEncodeByte(t *testing.T) {
	enc := Encoder{
		ecLv:    Quart,
		mode:    EncModeByte,
		version: loadVersion(1, Quart),
	}

	b, err := enc.Encode([]byte("helloworld"))
	if err != nil {
		t.Errorf("could not encode: %v", err)
		t.Fail()
	}
	t.Log(b, b.Len())
}
