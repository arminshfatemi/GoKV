package partitions

import (
	"bytes"
	"strings"
)

type ValueType uint8

const (
	INT ValueType = iota
	STRING
)

const (
	INTString    = "INT"
	STRINGString = "STRING"
)

var (
	INTByte    = []byte(INTString)
	StringByte = []byte(STRINGString)
)

func (vt ValueType) String() string {
	switch vt {
	case INT:
		return "INT"
	case STRING:
		return "STRING"
	default:
		return "UNKNOWN"
	}
}

func ParseValueType(s string) (ValueType, error) {
	switch strings.ToUpper(s) {
	case INTString:
		return INT, nil
	case STRINGString:
		return STRING, nil
	default:
		return 0, ErrInvalidValueType
	}
}

func ParseValueTypeBytes(b []byte) (ValueType, error) {
	switch {
	case bytes.EqualFold(b, INTByte):
		return INT, nil
	case bytes.EqualFold(b, StringByte):
		return STRING, nil
	default:
		return 0, ErrInvalidValueType
	}
}

type PersistMode uint8

const (
	NONE PersistMode = iota
	WAL
)

var (
	NoneByte = []byte("NONE")
	WALByte  = []byte("WAL")
)

func (pm PersistMode) String() string {
	switch pm {
	case NONE:
		return "NONE"
	case WAL:
		return "WAL"
	default:
		return "UNKNOWN"
	}
}

func ParsePersistMode(s string) (PersistMode, error) {
	switch strings.ToUpper(s) {
	case "NONE":
		return NONE, nil
	case "WAL":
		return WAL, nil
	default:
		return 0, ErrInvalidPersistMode
	}
}

func ParsePersistModeBytes(b []byte) (PersistMode, error) {
	switch {
	case bytes.EqualFold(b, NoneByte):
		return NONE, nil
	case bytes.EqualFold(b, WALByte):
		return WAL, nil
	default:
		return 0, ErrInvalidPersistMode
	}
}

type PartitionConfig struct {
	Name        string
	Schema      ValueType
	PersistMode PersistMode
}
