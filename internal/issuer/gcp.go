package issuer

import (
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"go.aporeto.io/a3s/pkgs/token"
	"go.aporeto.io/tg/tglib"
)

const (
	gcpClaimsCertURL        = "https://www.googleapis.com/oauth2/v1/certs"
	gcpClaimsRequiredIssuer = "https://accounts.google.com"
)

// ErrGCP represents an error that happened
// during operations related to GCP.
type ErrGCP struct {
	Err error
}

func (e ErrGCP) Error() string {
	return fmt.Sprintf("gcp error: %s", e.Err)
}

func (e ErrGCP) Unwrap() error {
	return e.Err
}

type gcpJWT struct {
	Google struct {
		ComputeEngine struct {
			ProjectID     string `json:"project_id"`
			ProjectNumber int    `json:"project_number"`
			Zone          string `json:"zone"`
			InstanceID    string `json:"instance_id"`
			InstanceName  string `json:"instance_name"`
		} `json:"compute_engine"`
	} `json:"google"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	jwt.RegisteredClaims
}

// A GCPIssuer can issue IdentityToken from a valid
// GCP token.
type GCPIssuer struct {
	token *token.IdentityToken
}

// NewGCPIssuer returns a new GCPIssuer.
func NewGCPIssuer() *GCPIssuer {
	return &GCPIssuer{
		token: token.NewIdentityToken(token.Source{
			Type: "gcp",
		}),
	}
}

// FromToken computes the identity claims from the given gcp token.
func (c *GCPIssuer) FromToken(tokenString string, audience string) (err error) {

	resp, err := http.Get(gcpClaimsCertURL)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid status code returned: %d", resp.StatusCode)
	}

	certsMap := map[string]string{}
	if err = json.NewDecoder(resp.Body).Decode(&certsMap); err != nil {
		return err
	}

	gcpToken := gcpJWT{}
	for _, v := range certsMap {
		cert, err := tglib.ParseCertificate([]byte(v))
		if err != nil {
			return err
		}
		if _, err := jwt.ParseWithClaims(tokenString, &gcpToken, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodRSA); ok {
				return cert.PublicKey.(*rsa.PublicKey), nil
			}
			return nil, fmt.Errorf("unexpected signing method: %s", t.Header["alg"])
		}); err == nil {
			break
		}
	}

	if gcpToken.Issuer != gcpClaimsRequiredIssuer {
		return ErrGCP{Err: fmt.Errorf("Invalid issuer from gcp token '%s'. Want '%s'", gcpToken.Issuer, gcpClaimsRequiredIssuer)}
	}

	if audience != "" {
		if !gcpToken.VerifyAudience(audience, true) {
			return ErrGCP{Err: fmt.Errorf("invalid audience '%s' want '%s'", gcpToken.Audience, audience)}
		}
	}

	c.token.Identity = computeGCPClaims(gcpToken)

	return nil
}

// Issue returns the IdentityToken.
func (c *GCPIssuer) Issue() *token.IdentityToken {
	return c.token
}

func computeGCPClaims(token gcpJWT) []string {

	var out []string

	out = append(out, fmt.Sprintf("subject=%s", token.Subject))

	if token.Google.ComputeEngine.ProjectID != "" {
		out = append(out, fmt.Sprintf("projectid=%s", token.Google.ComputeEngine.ProjectID))
	}

	if token.Google.ComputeEngine.ProjectNumber != 0 {
		out = append(out, fmt.Sprintf("projectnumber=%d", token.Google.ComputeEngine.ProjectNumber))
	}

	if token.Google.ComputeEngine.Zone != "" {
		out = append(out, fmt.Sprintf("zone=%s", token.Google.ComputeEngine.Zone))
	}

	if token.Google.ComputeEngine.InstanceID != "" {
		out = append(out, fmt.Sprintf("instanceid=%s", token.Google.ComputeEngine.InstanceID))
	}

	if token.Google.ComputeEngine.InstanceName != "" {
		out = append(out, fmt.Sprintf("instancename=%s", token.Google.ComputeEngine.InstanceName))
	}

	if token.Email != "" {
		out = append(out, fmt.Sprintf("email=%s", token.Email))
	}

	if token.EmailVerified {
		out = append(out, fmt.Sprintf("emailverified=%t", token.EmailVerified))
	}

	return out
}