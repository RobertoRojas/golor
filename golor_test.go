package golor

import (
	"testing"
)

func TestPlain(t *testing.T) {
	s := "This is a test"
	m := New()
	m.Bold().Black().ColorBit256(255).ColorBitRGB(255, 255, 255).Add(s).Reset()
	sl := len(s)
	ml := m.Len()
	mrl := len(m.String())
	if got := sl == ml && ml == mrl; !got {
		t.Errorf("The plain mode doesn't send equals lengths; sl(%d), ml(%d), mrl(%d)", sl, ml, mrl)
	}
}

func TestBit8(t *testing.T) {
	s := "This is a test"
	Mode = Bit8
	m := New()
	m.Bold().Black().ColorBit256(255).ColorBitRGB(255, 255, 255).Add(s).Reset()
	sl := len(s)
	ml := m.Len()
	mrl := len(m.String())
	mrle := 27
	if got := sl == ml && mrl == mrle; !got {
		t.Errorf("The values are not the expected, sl==ml and mrl==mrle; sl(%d), ml(%d), mrl(%d), mrle(%d)",
			sl, ml, mrl, mrle)
	}
}

func TestBit256(t *testing.T) {
	s := "This is a test"
	Mode = Bit256
	m := New()
	m.Bold().Black().ColorBit256(255).ColorBitRGB(255, 255, 255).Add(s).Reset()
	sl := len(s)
	ml := m.Len()
	mrl := len(m.String())
	mrle := 33
	if got := sl == ml && mrl == mrle; !got {
		t.Errorf("The values are not the expected, sl==ml and mrl==mrle; sl(%d), ml(%d), mrl(%d), mrle(%d)",
			sl, ml, mrl, mrle)
	}
}

func TestBitRGB(t *testing.T) {
	s := "This is a test"
	Mode = BitRGB
	m := New()
	m.Bold().Black().ColorBit256(255).ColorBitRGB(255, 255, 255).Add(s).Reset()
	sl := len(s)
	ml := m.Len()
	mrl := len(m.String())
	mrle := 41
	if got := sl == ml && mrl == mrle; !got {
		t.Errorf("The values are not the expected, sl==ml and mrl==mrle; sl(%d), ml(%d), mrl(%d), mrle(%d)",
			sl, ml, mrl, mrle)
	}
}

func TestOnlyGraphics(t *testing.T) {
	s := "This is a test"
	Mode = BitRGB
	m := New()
	m.Bold().Black().ColorBit256(255).Add(s).Reset()
	sl := len(s)
	ml := m.Len()
	mrl := len(m.String())
	mrle := 22
	if got := sl == ml && mrl == mrle; !got {
		t.Errorf("The values are not the expected, sl==ml and mrl==mrle; sl(%d), ml(%d), mrl(%d), mrle(%d)",
			sl, ml, mrl, mrle)
	}
}
