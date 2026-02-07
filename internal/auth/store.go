package auth

import (
	"GoKV/internal/config"
	"crypto/sha256"
	"crypto/subtle"
	"errors"
	"os"
	"sync"
)

var (
	invalidSaltError = errors.New("invalid salt file")
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

func NewStore(saltPath string) (*Store, error) {
	var salt []byte
	var err error

	if saltPath == "" {
		salt, err = generateSalt(32)
		if err != nil {
			return nil, err
		}
	} else {
		salt, err = loadSaltFromFile(saltPath)
		if err != nil {
			return nil, err
		}

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
	rec, ok := s.users[username]
	s.lock.RUnlock()

	if !ok {
		return nil, false
	}

	h := s.hashPassword(password)
	if subtle.ConstantTimeCompare(h[:], rec.hash[:]) != 1 {
		return nil, false
	}

	return &User{Username: username, Role: rec.role}, true
}

func (s *Store) hashPassword(password []byte) [32]byte {
	h := sha256.New()
	h.Write(s.salt)
	h.Write(password)
	var out [32]byte
	copy(out[:], h.Sum(out[:0]))
	return out
}

func (s *Store) AddUsersFromConfig(users []config.UserConfig) {
	for _, user := range users {
		role := ParseRoleStr(user.Role)
		s.AddUser(user.Username, []byte(user.Password), role)
	}
}

func loadSaltFromFile(filePath string) ([]byte, error) {
	b, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	if len(b) < 32 {
		return nil, invalidSaltError
	}

	return b, nil
}
