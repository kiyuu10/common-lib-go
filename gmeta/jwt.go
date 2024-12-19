package gmeta

import (
	"crypto/ecdsa"
	"crypto/rsa"
	"crypto/x509"

	"github.com/golang-jwt/jwt/v4"

	"gitea.alchemymagic.app/snap/go-common/erroy"
)

const (
	JwtEcKeyTextPrefix = "der:"
)

type JwtPrivateKeyECDSA struct {
	Key *ecdsa.PrivateKey
}

func (k *JwtPrivateKeyECDSA) PublicKey() *ecdsa.PublicKey {
	return &k.Key.PublicKey
}

func (k *JwtPrivateKeyECDSA) UnmarshalText(data []byte) (err error) {
	var (
		key    *ecdsa.PrivateKey
		prefix = string(data[:len(JwtEcKeyTextPrefix)])
	)
	if prefix == JwtEcKeyTextPrefix {
		if key, err = x509.ParseECPrivateKey(data[4:]); err != nil {
			return erroy.WrapStack(err, "jwt: parse ECDSA public key from DER")
		}
	} else {
		if key, err = jwt.ParseECPrivateKeyFromPEM(data); err != nil {
			return erroy.WrapStack(err, "jwt: parse ECDSA private key from PEM")
		}
	}
	k.Key = key
	return nil
}

func (k *JwtPrivateKeyECDSA) MarshalText() ([]byte, error) {
	keyBytes, err := x509.MarshalECPrivateKey(k.Key)
	if err != nil {
		return nil, erroy.WrapStack(err, "jwt: dump ECDSA private key to DER")
	}
	var textBytes = []byte(JwtEcKeyTextPrefix)
	textBytes = append(textBytes, keyBytes...)
	return textBytes, nil
}

type JwtPublicKeyECDSA struct {
	Key *ecdsa.PublicKey
}

func (k *JwtPublicKeyECDSA) UnmarshalText(data []byte) (err error) {
	var (
		key    any
		prefix = string(data[:len(JwtEcKeyTextPrefix)])
	)
	if prefix == JwtEcKeyTextPrefix {
		if key, err = x509.ParsePKIXPublicKey(data[4:]); err != nil {
			return erroy.WrapStack(err, "jwt: parse ECDSA public key from DER")
		}
	} else {
		if key, err = jwt.ParseECPublicKeyFromPEM(data); err != nil {
			return erroy.WrapStack(err, "jwt: parse ECDSA public key from PEM")
		}
	}
	k.Key = key.(*ecdsa.PublicKey)
	return nil
}

func (k *JwtPublicKeyECDSA) MarshalText() ([]byte, error) {
	keyBytes, err := x509.MarshalPKIXPublicKey(k.Key)
	if err != nil {
		return nil, erroy.WrapStack(err, "jwt: dump ECDSA public key to DER")
	}
	var textBytes = []byte(JwtEcKeyTextPrefix)
	textBytes = append(textBytes, keyBytes...)
	return textBytes, nil
}

type JwtPublicKeyRSA struct {
	Key *rsa.PublicKey
}

func (k *JwtPublicKeyRSA) UnmarshalText(data []byte) error {
	pubKey, err := jwt.ParseRSAPublicKeyFromPEM(data)
	if err != nil {
		return erroy.WrapStack(err, "jwt: parse RSA public key")
	}
	k.Key = pubKey
	return nil
}
