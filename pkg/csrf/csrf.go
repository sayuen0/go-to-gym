package csrf

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/sayuen0/go-to-gym/internal/infrastructure/logger"
	"io"
)

const (
	// CSRFHeader is the CSRF header
	CSRFHeader = "X-CSRF-Token"
	csrfSalt   = "pABhKHacAtMJR6fuUPXt"
)

// MakeToken creates CSRF token
func MakeToken(sid string, lg logger.Logger) string {
	hash := sha256.New()
	if _, err := io.WriteString(hash, fmt.Sprintf("%s%s", sid, csrfSalt)); err != nil {
		lg.Error("Make CSRF token", logger.Error(err))
	}
	token := base64.RawStdEncoding.EncodeToString(hash.Sum(nil))
	return token
}

func ValidToken(token, sid string, lg logger.Logger) bool {
	trueToken := MakeToken(sid, lg)
	return token == trueToken
}
