package partitions

import (
	"sync"
)

var Registry = map[string]*Partition{}
var RegistryLock sync.RWMutex

func CreatePartition(cfg PartitionConfig) error {
	RegistryLock.Lock()
	defer RegistryLock.Unlock()

	// check if partition already exist
	if _, ok := Registry[cfg.Name]; ok {
		return ErrPartitionExists
	}
	p := &Partition{
		cfg:         cfg,
		Name:        cfg.Name,
		Schema:      cfg.Schema,
		Persistence: cfg.PersistMode,
		Stats:       Stats{},
	}

	switch cfg.Schema {
	case INT:
		p.IntData = make(map[string]int64)
	case STRING:
		p.StringData = make(map[string]string)
	}

	Registry[cfg.Name] = p

	return nil
}

func DropPartition(name string) error {
	RegistryLock.Lock()
	defer RegistryLock.Unlock()

	if _, ok := Registry[name]; !ok {
		return ErrPartitionNotFound
	}

	delete(Registry, name)

	return nil
}

func ListPartitions() []string {
	RegistryLock.RLock()
	defer RegistryLock.RUnlock()

	partitions := make([]string, 0, len(Registry))

	for name, _ := range Registry {
		partitions = append(partitions, name)
	}

	return partitions
}

func GetPartition(name string) (*Partition, bool) {
	RegistryLock.RLock()
	defer RegistryLock.RUnlock()

	p, ok := Registry[name]
	return p, ok

}
