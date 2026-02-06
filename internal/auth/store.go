package auth

import (
	"crypto/sha256"
	"crypto/subtle"
	"sync"
)

type Store struct {
	users map[string]userRec
	lock  sync.RWMutex
	salt  []byte
}

type userRec struct {
	hash [32]byte
	role Role
}

func NewStore() (*Store, error) {
	salt, err := generateSalt(32)
	if err != nil {
		return nil, err
	}
	return &Store{
		users: make(map[string]userRec),
		salt:  salt,
	}, nil
}

func (s *Store) AddUser(username string, password []byte, role Role) {
	s.lock.Lock()
	defer s.lock.Unlock()

	h := s.hashPassword(password)
	s.users[username] = userRec{hash: h, role: role}
}

func (s *Store) Authenticate(username string, password []byte) (*User, bool) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	rec, ok := s.users[username]
	if !ok {
		return nil, false
	}

	h := sha256.Sum256(append(append([]byte{}, s.salt...), password...))
	if subtle.ConstantTimeCompare(h[:], rec.hash[:]) != 1 {
		return nil, false
	}

	return &User{Username: username, Role: rec.role}, true
}

func (s *Store) hashPassword(password []byte) [32]byte {
	// one allocation-free way is to reuse a scratch buffer via sync.Pool,
	// but that's probably overkill for AUTH.
	// For v1, keep it simple.
	h := sha256.New()
	h.Write(s.salt)
	h.Write(password)
	var out [32]byte
	copy(out[:], h.Sum(nil))
	return out
}
