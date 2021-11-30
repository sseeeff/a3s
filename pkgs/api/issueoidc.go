package api

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// IssueOIDC represents the model of a issueoidc
type IssueOIDC struct {
	// Contains the auth URL is noAuthRedirect is set to true.
	AuthURL string `json:"authURL,omitempty" msgpack:"authURL,omitempty" bson:"-" mapstructure:"authURL,omitempty"`

	// OIDC ceremony code.
	Code string `json:"code" msgpack:"code" bson:"-" mapstructure:"code,omitempty"`

	// If set, instruct the server to return the OIDC auth url in authURL instead of
	// performing an HTTP redirection.
	NoAuthRedirect bool `json:"noAuthRedirect" msgpack:"noAuthRedirect" bson:"-" mapstructure:"noAuthRedirect,omitempty"`

	// OIDC redirect url in case of error.
	RedirectErrorURL string `json:"redirectErrorURL" msgpack:"redirectErrorURL" bson:"-" mapstructure:"redirectErrorURL,omitempty"`

	// OIDC redirect url.
	RedirectURL string `json:"redirectURL" msgpack:"redirectURL" bson:"-" mapstructure:"redirectURL,omitempty"`

	// OIDC ceremony state.
	State string `json:"state" msgpack:"state" bson:"-" mapstructure:"state,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewIssueOIDC returns a new *IssueOIDC
func NewIssueOIDC() *IssueOIDC {

	return &IssueOIDC{
		ModelVersion: 1,
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *IssueOIDC) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesIssueOIDC{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *IssueOIDC) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesIssueOIDC{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// BleveType implements the bleve.Classifier Interface.
func (o *IssueOIDC) BleveType() string {

	return "issueoidc"
}

// DeepCopy returns a deep copy if the IssueOIDC.
func (o *IssueOIDC) DeepCopy() *IssueOIDC {

	if o == nil {
		return nil
	}

	out := &IssueOIDC{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *IssueOIDC.
func (o *IssueOIDC) DeepCopyInto(out *IssueOIDC) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy IssueOIDC: %s", err))
	}

	*out = *target.(*IssueOIDC)
}

// Validate valides the current information stored into the structure.
func (o *IssueOIDC) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if len(requiredErrors) > 0 {
		return requiredErrors
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (*IssueOIDC) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := IssueOIDCAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return IssueOIDCLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*IssueOIDC) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return IssueOIDCAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *IssueOIDC) ValueForAttribute(name string) interface{} {

	switch name {
	case "authURL":
		return o.AuthURL
	case "code":
		return o.Code
	case "noAuthRedirect":
		return o.NoAuthRedirect
	case "redirectErrorURL":
		return o.RedirectErrorURL
	case "redirectURL":
		return o.RedirectURL
	case "state":
		return o.State
	}

	return nil
}

// IssueOIDCAttributesMap represents the map of attribute for IssueOIDC.
var IssueOIDCAttributesMap = map[string]elemental.AttributeSpecification{
	"AuthURL": {
		AllowedChoices: []string{},
		ConvertedName:  "AuthURL",
		Description:    `Contains the auth URL is noAuthRedirect is set to true.`,
		Exposed:        true,
		Name:           "authURL",
		ReadOnly:       true,
		Type:           "string",
	},
	"Code": {
		AllowedChoices: []string{},
		ConvertedName:  "Code",
		Description:    `OIDC ceremony code.`,
		Exposed:        true,
		Name:           "code",
		Type:           "string",
	},
	"NoAuthRedirect": {
		AllowedChoices: []string{},
		ConvertedName:  "NoAuthRedirect",
		Description: `If set, instruct the server to return the OIDC auth url in authURL instead of
performing an HTTP redirection.`,
		Exposed: true,
		Name:    "noAuthRedirect",
		Type:    "boolean",
	},
	"RedirectErrorURL": {
		AllowedChoices: []string{},
		ConvertedName:  "RedirectErrorURL",
		Description:    `OIDC redirect url in case of error.`,
		Exposed:        true,
		Name:           "redirectErrorURL",
		Type:           "string",
	},
	"RedirectURL": {
		AllowedChoices: []string{},
		ConvertedName:  "RedirectURL",
		Description:    `OIDC redirect url.`,
		Exposed:        true,
		Name:           "redirectURL",
		Type:           "string",
	},
	"State": {
		AllowedChoices: []string{},
		ConvertedName:  "State",
		Description:    `OIDC ceremony state.`,
		Exposed:        true,
		Name:           "state",
		Type:           "string",
	},
}

// IssueOIDCLowerCaseAttributesMap represents the map of attribute for IssueOIDC.
var IssueOIDCLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"authurl": {
		AllowedChoices: []string{},
		ConvertedName:  "AuthURL",
		Description:    `Contains the auth URL is noAuthRedirect is set to true.`,
		Exposed:        true,
		Name:           "authURL",
		ReadOnly:       true,
		Type:           "string",
	},
	"code": {
		AllowedChoices: []string{},
		ConvertedName:  "Code",
		Description:    `OIDC ceremony code.`,
		Exposed:        true,
		Name:           "code",
		Type:           "string",
	},
	"noauthredirect": {
		AllowedChoices: []string{},
		ConvertedName:  "NoAuthRedirect",
		Description: `If set, instruct the server to return the OIDC auth url in authURL instead of
performing an HTTP redirection.`,
		Exposed: true,
		Name:    "noAuthRedirect",
		Type:    "boolean",
	},
	"redirecterrorurl": {
		AllowedChoices: []string{},
		ConvertedName:  "RedirectErrorURL",
		Description:    `OIDC redirect url in case of error.`,
		Exposed:        true,
		Name:           "redirectErrorURL",
		Type:           "string",
	},
	"redirecturl": {
		AllowedChoices: []string{},
		ConvertedName:  "RedirectURL",
		Description:    `OIDC redirect url.`,
		Exposed:        true,
		Name:           "redirectURL",
		Type:           "string",
	},
	"state": {
		AllowedChoices: []string{},
		ConvertedName:  "State",
		Description:    `OIDC ceremony state.`,
		Exposed:        true,
		Name:           "state",
		Type:           "string",
	},
}

type mongoAttributesIssueOIDC struct {
}
