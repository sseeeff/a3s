package processors

import (
	"crypto/x509"
	"net/http"

	"go.aporeto.io/a3s/pkgs/api"
	"go.aporeto.io/a3s/pkgs/authorizer"
	"go.aporeto.io/a3s/pkgs/permissions"
	"go.aporeto.io/a3s/pkgs/token"
	"go.aporeto.io/bahamut"
	"go.aporeto.io/elemental"
)

// A AuthzProcessor is a bahamut processor for Authzs.
type AuthzProcessor struct {
	authorizer authorizer.Authorizer
	jwtCert    *x509.Certificate
}

// NewAuthzProcessor returns a new AuthzProcessor.
func NewAuthzProcessor(authorizer authorizer.Authorizer, jwtCert *x509.Certificate) *AuthzProcessor {
	return &AuthzProcessor{
		authorizer: authorizer,
		jwtCert:    jwtCert,
	}
}

// ProcessCreate handles the creates requests for Authzs.
func (p *AuthzProcessor) ProcessCreate(bctx bahamut.Context) error {

	req := bctx.InputData().(*api.Authz)

	idt, err := token.Verify(req.Token, p.jwtCert)
	if err != nil {
		return elemental.NewError(
			"Bad Request",
			err.Error(),
			"a3s:authz",
			http.StatusBadRequest,
		)
	}

	var r permissions.Restrictions
	if idt.Restrictions != nil {
		r = *idt.Restrictions
	}

	ok, err := p.authorizer.CheckAuthorization(
		bctx.Context(),
		idt.Identity,
		req.Action,
		req.Namespace,
		req.Resource,
		authorizer.OptionCheckID(req.ID),
		authorizer.OptionCheckSourceIP(req.IP),
		authorizer.OptionCheckRestrictions(r),
	)
	if err != nil {
		return err
	}

	if ok {
		bctx.SetStatusCode(http.StatusOK)
	} else {
		bctx.SetStatusCode(http.StatusForbidden)
		bctx.SetOutputData("Forbidden")
	}

	return nil
}
