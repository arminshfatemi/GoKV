package protocol

import "bytes"

// CREATE PARTITION P0 INT WAL
func parseCreate(t [][]byte) (*Command, error) {
	if len(t) != 5 || !bytes.EqualFold(t[1], []byte("PARTITION")) {
		return nil, ErrWrongArgCount
	}

	return &Command{
		Type:        CmdCreatePartition,
		Partition:   string(t[2]),
		ValueType:   string(t[3]),
		PersistMode: string(t[4]),
	}, nil
}

// DROP PARTITION P0
func parseDrop(t [][]byte) (*Command, error) {
	if len(t) != 3 || !bytes.EqualFold(t[1], []byte("PARTITION")) {
		return nil, ErrWrongArgCount
	}

	return &Command{
		Type:      CmdDropPartition,
		Partition: string(t[2]),
	}, nil
}

// LIST PARTITIONS
func parseList(t [][]byte) (*Command, error) {
	if len(t) != 2 || !bytes.EqualFold(t[1], []byte("PARTITIONS")) {
		return nil, ErrWrongArgCount
	}

	return &Command{Type: CmdListPartitions}, nil
}

// // DESCRIBE PARTITION P0
//
//	func parseDescribe(t [][]byte) (*Command, error) {
//		if len(t) != 3 || !bytes.EqualFold(t[1], []byte("PARTITION")) {
//			return nil, ErrWrongArgCount
//		}
//
//		return &Command{
//			Schema:      CmdDescribePartition,
//			Partition: string(t[2]),
//		}, nil
//	}
//

// SET P0 key value
func parseSet(t [][]byte) (*Command, error) {
	if len(t) != 4 {
		return nil, ErrWrongArgCount
	}

	return &Command{
		Type:      CmdSet,
		Partition: string(t[1]),
		Key:       string(t[2]),
		Value:     string(t[3]),
	}, nil
}

// GET P0 key

func parseGet(t [][]byte) (*Command, error) {
	if len(t) != 3 {
		return nil, ErrWrongArgCount
	}

	return &Command{
		Type:      CmdGet,
		Partition: string(t[1]),
		Key:       string(t[2]),
	}, nil
}

// DEL P0 key
func parseDel(t [][]byte) (*Command, error) {
	if len(t) != 3 {
		return nil, ErrWrongArgCount
	}

	return &Command{
		Type:      CmdDel,
		Partition: string(t[1]),
		Key:       string(t[2]),
	}, nil
}

// INCR P0 key
func parseIncr(t [][]byte) (*Command, error) {
	if len(t) != 3 {
		return nil, ErrWrongArgCount
	}

	return &Command{
		Type:      CmdIncr,
		Partition: string(t[1]),
		Key:       string(t[2]),
	}, nil
}

//// STATS PARTITION P0
//func parseStats(t [][]byte) (*Command, error) {
//	if len(t) != 3 || !bytes.EqualFold(t[1], []byte("PARTITION")) {
//		return nil, ErrWrongArgCount
//	}
//
//	return &Command{
//		Schema:      CmdStatsPartition,
//		Partition: string(t[2]),
//	}, nil
//}
