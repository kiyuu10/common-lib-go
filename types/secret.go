package types

import (
	"encoding"
	"encoding/base64"
	"encoding/hex"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"

	"gitea.alchemymagic.app/snap/go-common/erroy"
)

type (
	Secret interface {
		Get() ([]byte, error)
	}
	SecretValue interface {
		encoding.TextUnmarshaler
		Bytes() []byte
	}
)

type PlaceholderSecret struct {
	secret Secret
}

func (s *PlaceholderSecret) UnmarshalText(input []byte) error {
	inputText := string(input)
	inputDSN, err := url.Parse(inputText)
	if err != nil {
		return erroy.WrapStack(err, "secret: parse input DSN")
	}
	switch strings.ToLower(inputDSN.Scheme) {
	case "const":
		s.secret = NewConstantSecret(input)
	case "const+hex":
		data, err := hex.DecodeString(inputDSN.Host)
		if err != nil {
			return erroy.WrapStack(err, "secret: decode hex")
		}
		s.secret = NewConstantSecret(data)
	case "const+base64":
		data, err := base64.StdEncoding.DecodeString(inputDSN.Host)
		if err != nil {
			return erroy.WrapStack(err, "secret: decode hex")
		}
		s.secret = NewConstantSecret(data)
	case "file":
		s.secret = NewFileSecret(inputDSN.Path, new(RawBytes))
	case "file+hex":
		s.secret = NewFileSecret(inputDSN.Path, new(HexBytes))
	case "file+base64":
		s.secret = NewFileSecret(inputDSN.Path, new(Base64Bytes))
	default:
		return erroy.NewWithStack("secret: unsupported scheme").
			WithField("scheme", inputDSN.Scheme)
	}
	return nil
}

func (s *PlaceholderSecret) Get() ([]byte, error) {
	if s.secret == nil {
		return nil, erroy.New("secret: placeholder hasn't been loaded")
	}
	return s.secret.Get()
}

type CachedSecret struct {
	secret      Secret
	timeout     time.Duration
	cachedValue []byte
	cachedMux   sync.Mutex
	cachedTime  time.Time
}

var _ Secret = (*CachedSecret)(nil)

func NewCachedSecret(secret Secret, timeout time.Duration) *CachedSecret {
	return &CachedSecret{
		secret:  secret,
		timeout: timeout,
	}
}

func (s *CachedSecret) isExpired() bool {
	return s.cachedValue == nil || time.Now().Sub(s.cachedTime) > s.timeout
}

func (s *CachedSecret) load() error {
	s.cachedMux.Lock()
	defer s.cachedMux.Unlock()
	if !s.isExpired() {
		return nil
	}
	value, err := s.secret.Get()
	if err != nil {
		return err
	}
	s.cachedValue = value
	s.cachedTime = time.Now()
	return nil
}

func (s *CachedSecret) Get() ([]byte, error) {
	if s.isExpired() {
		if err := s.load(); err != nil {
			return nil, err
		}
	}
	return s.cachedValue, nil
}

type ConstantSecret struct {
	value []byte
}

var _ Secret = (*ConstantSecret)(nil)

func (s *ConstantSecret) Get() ([]byte, error) {
	return s.value, nil
}

func NewConstantSecret(value []byte) *ConstantSecret {
	return &ConstantSecret{
		value: value,
	}
}

type FileSecret struct {
	path       string
	value      SecretValue
	loadedFlag bool
	loadedMux  sync.Mutex
}

var _ Secret = (*FileSecret)(nil)

func NewFileSecret(path string, value SecretValue) *FileSecret {
	return &FileSecret{
		path:  path,
		value: value,
	}
}

func (s *FileSecret) load() (err error) {
	s.loadedMux.Lock()
	defer s.loadedMux.Unlock()
	if s.loadedFlag {
		return nil
	}
	data, err := os.ReadFile(s.path)
	if err != nil {
		return erroy.WrapStack(err, "file secret: reader run file")
	}
	if err = s.value.UnmarshalText(data); err != nil {
		return erroy.WrapStack(err, "file secret: parse data")
	}
	s.loadedFlag = true
	return nil
}

func (s *FileSecret) Get() ([]byte, error) {
	if !s.loadedFlag {
		if err := s.load(); err != nil {
			return nil, err
		}
	}
	return s.value.Bytes(), nil
}

func NewDockerSecret(name string, value SecretValue) *FileSecret {
	return NewFileSecret("/run/secrets/"+name, value)
}

func NewCachedFileSecret(timeout time.Duration, path string, value SecretValue) *CachedSecret {
	return NewCachedSecret(NewFileSecret(path, value), timeout)
}

func NewCachedDockerSecret(timeout time.Duration, name string, value SecretValue) *CachedSecret {
	return NewCachedSecret(NewDockerSecret(name, value), timeout)
}
