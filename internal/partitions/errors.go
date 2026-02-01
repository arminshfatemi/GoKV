package partitions

import "errors"

var (
	ErrPartitionExists    = errors.New("partition already exists")
	ErrPartitionNotFound  = errors.New("partition not found")
	ErrInvalidSchema      = errors.New("invalid schema")
	ErrInvalidValue       = errors.New("invalid value")
	ErrInvalidValueType   = errors.New("invalid value type")
	ErrInvalidPersistMode = errors.New("invalid persist mode")
)
