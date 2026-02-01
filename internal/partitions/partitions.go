package partitions

import (
	"strconv"
	"sync"
	"sync/atomic"
)

type Stats struct {
	KeysCount   atomic.Uint64
	OpsCount    atomic.Uint64
	WritesCount atomic.Uint64
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
	var existed bool

	p.Lock.Lock()
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
	p.Lock.Unlock()

	if existed {
		p.Stats.KeysCount.Add(-1)
		p.Stats.WritesCount.Add(1)
	}
	p.Stats.OpsCount.Add(1)

	return existed
}

func (p *Partition) Get(key string) (any, bool) {
	p.Stats.OpsCount.Add(1)

	p.Lock.Lock()
	defer p.Lock.Unlock()
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

func (p *Partition) Incr(key string) (int64, error) {
	// check schema to support Int values
	if p.Schema != INT {
		return 0, ErrInvalidSchema
	}

	m := p.IntData

	p.Lock.Lock()
	v, ok := m[key]

	// if not create the key create with value 1
	if !ok {
		m[key] = 1
		p.Lock.Unlock()
		p.Stats.KeysCount.Add(1)
		p.Stats.WritesCount.Add(1)
		p.Stats.OpsCount.Add(1)
		return 1, nil
	}

	// if exist increment it and return response
	v++
	m[key] = v
	p.Lock.Unlock()
	p.Stats.WritesCount.Add(1)
	p.Stats.OpsCount.Add(1)

	return v, nil
}

func (p *Partition) Describe() []string {
	p.Lock.RLock()
	defer p.Lock.RUnlock()

	p.Stats.OpsCount.Add(1)

	cfg := p.cfg

	s := []string{
		"name", cfg.Name,
		"schema", cfg.Schema.String(),
		"persisMode", cfg.PersistMode.String(),
	}

	return s
}

func (p *Partition) Stat() []string {
	p.Stats.OpsCount.Add(1)
	return []string{
		"keyCount", strconv.FormatUint(p.Stats.KeysCount.Load(), 10),
		"opsCount", strconv.FormatUint(p.Stats.OpsCount.Load(), 10),
		"writesCount", strconv.FormatUint(p.Stats.WritesCount.Load(), 10),
	}
}

func (p *Partition) setInt(key string, rawValue []byte) error {
	intValue, err := strconv.ParseInt(string(rawValue), 10, 64)
	if err != nil {
		return ErrInvalidValue
	}

	p.Lock.Lock()
	_, existed := p.IntData[key]
	p.IntData[key] = intValue
	p.Lock.Unlock()

	if !existed {
		p.Stats.KeysCount.Add(1)
	}
	p.Stats.WritesCount.Add(1)
	p.Stats.OpsCount.Add(1)

	return nil
}

func (p *Partition) setString(key string, rawValue []byte) error {
	strValue := string(rawValue)

	p.Lock.Lock()
	_, existed := p.StringData[key]
	p.StringData[key] = strValue
	p.Lock.Unlock()

	if !existed {
		p.Stats.KeysCount.Add(1)
	}
	p.Stats.WritesCount.Add(1)
	p.Stats.OpsCount.Add(1)

	return nil
}
