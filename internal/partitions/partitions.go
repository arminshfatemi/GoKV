package partitions

import (
	"strconv"
	"sync"
)

type Stats struct {
	KeysCount   uint64
	OpsCount    uint64
	WritesCount uint64
}

type Partition struct {
	cfg         PartitionConfig
	Name        string
	Schema      ValueType
	Persistence PersistMode
	Lock        sync.RWMutex
	Stats       Stats
	IntData     map[string]int64
	StringData  map[string]string
}

func (p *Partition) Set(key string, rawValue []byte) error {
	switch p.Schema {
	case INT:
		return p.setInt(key, rawValue)
	case STRING:
		return p.setString(key, rawValue)
	default:
		return ErrInvalidSchema
	}
}

func (p *Partition) Del(key string) bool {
	p.Lock.Lock()
	defer p.Lock.Unlock()

	var existed bool

	switch p.Schema {
	case INT:
		_, existed = p.IntData[key]
		if existed {
			delete(p.IntData, key)
		}
	case STRING:
		_, existed = p.StringData[key]
		if existed {
			delete(p.StringData, key)
		}
	}

	if existed {
		p.Stats.KeysCount--
		p.Stats.WritesCount++
	}
	p.Stats.OpsCount++

	return existed
}

func (p *Partition) Get(key string) (any, bool) {
	p.Lock.Lock()
	defer p.Lock.Unlock()

	p.Stats.OpsCount++

	switch p.Schema {
	case INT:
		v, ok := p.IntData[key]
		return v, ok
	case STRING:
		v, ok := p.StringData[key]
		return v, ok
	default:
		return nil, false
	}
}

func (p *Partition) setInt(key string, rawValue []byte) error {
	intValue, err := strconv.ParseInt(string(rawValue), 10, 64)
	if err != nil {
		return ErrInvalidValue
	}

	p.Lock.Lock()
	defer p.Lock.Unlock()

	_, existed := p.IntData[key]
	p.IntData[key] = intValue

	if !existed {
		p.Stats.KeysCount++
	}
	p.Stats.WritesCount++
	p.Stats.OpsCount++

	return nil
}

func (p *Partition) setString(key string, rawValue []byte) error {
	strValue := string(rawValue)

	p.Lock.Lock()
	defer p.Lock.Unlock()

	_, existed := p.StringData[key]
	p.StringData[key] = strValue

	if !existed {
		p.Stats.KeysCount++
	}
	p.Stats.WritesCount++
	p.Stats.OpsCount++

	return nil
}
