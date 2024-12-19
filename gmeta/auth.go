package gmeta

type (
	JwtKeyPair struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}

	Password interface {
		String() string
		GetSalt() []byte
		IsValidPassword(offerPassword []byte) bool
	}

	AuthMethodActionType string
)

const (
	JwtRefreshLifeTimeMinRate = 0.7
)
