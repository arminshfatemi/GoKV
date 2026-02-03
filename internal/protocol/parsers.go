package protocol

import "bytes"

var (
	partition  = []byte("PARTITION")
	partitions = []byte("PARTITIONS")
)

// CREATE PARTITION P0 INT WAL
func parseCreate(t [][]byte) (*Command, error) {
	if len(t) != 5 || !bytes.EqualFold(t[1], partition) {
		return nil, ErrWrongArgCount
	}

	return &Command{
		Type:      CmdCreatePartition,
		Partition: t[2],
		Args:      t[3:],
	}, nil
}

// DROP PARTITION P0
func parseDrop(t [][]byte) (*Command, error) {
	if len(t) != 3 || !bytes.EqualFold(t[1], partition) {
		return nil, ErrWrongArgCount
	}

	return &Command{
		Type:      CmdDropPartition,
		Partition: t[2],
	}, nil
}

// LIST PARTITIONS
func parseList(t [][]byte) (*Command, error) {
	if len(t) != 2 || !bytes.EqualFold(t[1], partitions) {
		return nil, ErrWrongArgCount
	}
	return &Command{Type: CmdListPartitions}, nil
}

// DESCRIBE PARTITION P0
func parseDescribe(t [][]byte) (*Command, error) {
	if len(t) != 3 || !bytes.EqualFold(t[1], partition) {
		return nil, ErrWrongArgCount
	}

	return &Command{
		Type:      CmdDescribePartition,
		Partition: t[2],
	}, nil
}

// SET P0 key value
func parseSet(t [][]byte) (*Command, error) {
	if len(t) != 4 {
		return nil, ErrWrongArgCount
	}

	return &Command{
		Type:      CmdSet,
		Partition: t[1],
		Args:      t[2:],
	}, nil
}

// GET P0 key
func parseGet(t [][]byte) (*Command, error) {
	if len(t) != 3 {
		return nil, ErrWrongArgCount
	}

	return &Command{
		Type:      CmdGet,
		Partition: t[1],
		Args:      t[2:],
	}, nil
}

// DEL P0 key
func parseDel(t [][]byte) (*Command, error) {
	if len(t) < 3 {
		return nil, ErrWrongArgCount
	}

	return &Command{
		Type:      CmdDel,
		Partition: t[1],
		Args:      t[2:],
	}, nil
}

// INCR P0 key
func parseIncr(t [][]byte) (*Command, error) {
	if len(t) != 3 {
		return nil, ErrWrongArgCount
	}

	return &Command{
		Type:      CmdIncr,
		Partition: t[1],
		Args:      t[2:],
	}, nil
}

// STATS PARTITION P0
func parseStats(t [][]byte) (*Command, error) {
	if len(t) != 3 || !bytes.EqualFold(t[1], partition) {
		return nil, ErrWrongArgCount
	}

	return &Command{
		Type:      CmdStatsPartition,
		Partition: t[2],
	}, nil
}

// EXISTS <partition> <key> [key...]
func parseExists(t [][]byte) (*Command, error) {
	if len(t) <= 2 {
		return nil, ErrWrongArgCount
	}

	return &Command{
		Type:      CmdExists,
		Partition: t[1],
		Args:      t[2:],
	}, nil
}
