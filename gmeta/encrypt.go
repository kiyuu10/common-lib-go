package gmeta

import (
	"database/sql/driver"
	"strings"

	"gitea.alchemymagic.app/snap/go-common/erroy"
	"gitea.alchemymagic.app/snap/go-common/types"
	comutils "gitea.alchemymagic.app/snap/go-common/utils"
)

const (
	EncryptedSeparator    = ":"
	EncryptedPrefix       = "aes-256-gcm"
	EncryptedNonceLength  = 12
	EncryptedSecretLength = 32
)

type EncryptedValue struct {
	cipherText string
	secret     types.Secret
}

func NewEncryptedValue(raw []byte, secret types.Secret) (value EncryptedValue, err error) {
	if err = value.SetSecret(secret); err != nil {
		return
	}
	if value.cipherText, err = value.Encrypt(raw); err != nil {
		return
	}
	return
}

func LoadEncryptedValue(cipherText string, secret types.Secret) (value EncryptedValue, err error) {
	if err = value.SetSecret(secret); err != nil {
		return
	}
	value.cipherText = cipherText
	return
}

func (v *EncryptedValue) IsEmpty() bool {
	return v.cipherText == ""
}

func (v *EncryptedValue) Get() ([]byte, error) {
	if v.IsEmpty() {
		return nil, nil
	}
	originalData, err := v.Decrypt()
	if err != nil {
		return nil, err
	}
	return originalData, nil
}

func (v *EncryptedValue) String() string {
	return v.cipherText
}

func (v *EncryptedValue) Value() (driver.Value, error) {
	return v.cipherText, nil
}

func (v *EncryptedValue) Scan(input any) error {
	v.cipherText = string(input.([]byte))
	return nil
}

func (v *EncryptedValue) MarshalText() ([]byte, error) {
	return []byte(v.cipherText), nil
}

func (v *EncryptedValue) UnmarshalText(input []byte) error {
	v.cipherText = string(input)
	return nil
}

func (v *EncryptedValue) SetSecret(secret types.Secret) error {
	secretBytes, err := secret.Get()
	if err != nil {
		return err
	}
	if len(secretBytes) != EncryptedSecretLength {
		return erroy.NewWithStack(
			"encrypted value: requires %v-length secret to use AES-256",
			EncryptedSecretLength,
		)
	}
	v.secret = secret
	return nil
}

func (v *EncryptedValue) getSecret() ([]byte, error) {
	if v.secret == nil {
		return nil, erroy.NewWithStack("encrypted value: requires `secret` before execution")
	}
	return v.secret.Get()
}

func (v *EncryptedValue) Encrypt(data []byte) (_ string, err error) {
	secretBytes, err := v.getSecret()
	if err != nil {
		return
	}
	nonce, err := comutils.RandomBytes(EncryptedNonceLength)
	if err != nil {
		return
	}
	cipherBytes, err := comutils.AesGcmEncrypt(secretBytes, nonce, []byte(data))
	if err != nil {
		return
	}
	var (
		nonceBase64   = comutils.Base64Encode(nonce)
		cipherBase64  = comutils.Base64Encode(cipherBytes)
		encryptedText = EncryptedPrefix +
			EncryptedSeparator + nonceBase64 +
			EncryptedSeparator + cipherBase64
	)
	return encryptedText, nil
}

func (v *EncryptedValue) Decrypt() (_ []byte, err error) {
	secretBytes, err := v.getSecret()
	if err != nil {
		return
	}
	parts := strings.Split(v.cipherText, EncryptedSeparator)
	if len(parts) != 3 {
		return nil, erroy.NewWithStack(
			"encrypted value: requires 3 parts in encrypted value (not %v)",
			len(parts))
	}
	if parts[0] != EncryptedPrefix {
		return nil, erroy.NewWithStack("encrypted value: invalid header").
			WithField("header", parts[0])
	}
	nonceBytes, err := comutils.Base64Decode(parts[1])
	if err != nil {
		return
	}
	cipherBytes, err := comutils.Base64Decode(parts[2])
	if err != nil {
		return
	}
	originalData, err := comutils.AesGcmDecrypt(secretBytes, nonceBytes, cipherBytes)
	if err != nil {
		return
	}
	return originalData, nil
}

// EncryptedLiteValue doesn't hold the secret
type EncryptedLiteValue struct {
	cipherText string
}

func (v *EncryptedLiteValue) IsEmpty() bool {
	return v.cipherText == ""
}

func (v *EncryptedLiteValue) Get(secret types.Secret) ([]byte, error) {
	if v.IsEmpty() {
		return nil, nil
	}
	originalData, err := v.Decrypt(secret)
	if err != nil {
		return nil, err
	}
	return originalData, nil
}

func (v *EncryptedLiteValue) String() string {
	return v.cipherText
}

func (v *EncryptedLiteValue) Value() (driver.Value, error) {
	return v.cipherText, nil
}

func (v *EncryptedLiteValue) Scan(input any) error {
	v.cipherText = string(input.([]byte))
	return nil
}

func (v *EncryptedLiteValue) MarshalText() ([]byte, error) {
	return []byte(v.cipherText), nil
}

func (v *EncryptedLiteValue) UnmarshalText(input []byte) error {
	v.cipherText = string(input)
	return nil
}

func (v *EncryptedLiteValue) Encrypt(data []byte, secret types.Secret) (_ string, err error) {
	secretBytes, err := secret.Get()
	if err != nil {
		return
	}
	nonce, err := comutils.RandomBytes(EncryptedNonceLength)
	if err != nil {
		return
	}
	cipherBytes, err := comutils.AesGcmEncrypt(secretBytes, nonce, []byte(data))
	if err != nil {
		return
	}
	var (
		nonceBase64   = comutils.Base64Encode(nonce)
		cipherBase64  = comutils.Base64Encode(cipherBytes)
		encryptedText = EncryptedPrefix +
			EncryptedSeparator + nonceBase64 +
			EncryptedSeparator + cipherBase64
	)
	return encryptedText, nil
}

func (v *EncryptedLiteValue) Decrypt(secret types.Secret) (_ []byte, err error) {
	secretBytes, err := secret.Get()
	if err != nil {
		return
	}
	parts := strings.Split(v.cipherText, EncryptedSeparator)
	if len(parts) != 3 {
		return nil, erroy.NewWithStack(
			"encrypted value: requires 3 parts in encrypted value (not %v)",
			len(parts))
	}
	if parts[0] != EncryptedPrefix {
		return nil, erroy.NewWithStack("encrypted value: invalid header").
			WithField("header", parts[0])
	}
	nonceBytes, err := comutils.Base64Decode(parts[1])
	if err != nil {
		return
	}
	cipherBytes, err := comutils.Base64Decode(parts[2])
	if err != nil {
		return
	}
	originalData, err := comutils.AesGcmDecrypt(secretBytes, nonceBytes, cipherBytes)
	if err != nil {
		return
	}
	return originalData, nil
}
