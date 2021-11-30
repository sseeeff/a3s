package api

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// IssueAWS represents the model of a issueaws
type IssueAWS struct {
	// The ID of the AWS STS token.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// The secret associated to the AWS STS token.
	Secret string `json:"secret" msgpack:"secret" bson:"-" mapstructure:"secret,omitempty"`

	// The original token.
	Token string `json:"token" msgpack:"token" bson:"-" mapstructure:"token,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewIssueAWS returns a new *IssueAWS
func NewIssueAWS() *IssueAWS {

	return &IssueAWS{
		ModelVersion: 1,
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *IssueAWS) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesIssueAWS{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *IssueAWS) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesIssueAWS{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// BleveType implements the bleve.Classifier Interface.
func (o *IssueAWS) BleveType() string {

	return "issueaws"
}

// DeepCopy returns a deep copy if the IssueAWS.
func (o *IssueAWS) DeepCopy() *IssueAWS {

	if o == nil {
		return nil
	}

	out := &IssueAWS{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *IssueAWS.
func (o *IssueAWS) DeepCopyInto(out *IssueAWS) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy IssueAWS: %s", err))
	}

	*out = *target.(*IssueAWS)
}

// Validate valides the current information stored into the structure.
func (o *IssueAWS) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("ID", o.ID); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("secret", o.Secret); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("token", o.Token); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if len(requiredErrors) > 0 {
		return requiredErrors
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (*IssueAWS) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := IssueAWSAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return IssueAWSLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*IssueAWS) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return IssueAWSAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *IssueAWS) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "secret":
		return o.Secret
	case "token":
		return o.Token
	}

	return nil
}

// IssueAWSAttributesMap represents the map of attribute for IssueAWS.
var IssueAWSAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": {
		AllowedChoices: []string{},
		ConvertedName:  "ID",
		Description:    `The ID of the AWS STS token.`,
		Exposed:        true,
		Name:           "ID",
		Required:       true,
		Type:           "string",
	},
	"Secret": {
		AllowedChoices: []string{},
		ConvertedName:  "Secret",
		Description:    `The secret associated to the AWS STS token.`,
		Exposed:        true,
		Name:           "secret",
		Required:       true,
		Type:           "string",
	},
	"Token": {
		AllowedChoices: []string{},
		ConvertedName:  "Token",
		Description:    `The original token.`,
		Exposed:        true,
		Name:           "token",
		Required:       true,
		Type:           "string",
	},
}

// IssueAWSLowerCaseAttributesMap represents the map of attribute for IssueAWS.
var IssueAWSLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"id": {
		AllowedChoices: []string{},
		ConvertedName:  "ID",
		Description:    `The ID of the AWS STS token.`,
		Exposed:        true,
		Name:           "ID",
		Required:       true,
		Type:           "string",
	},
	"secret": {
		AllowedChoices: []string{},
		ConvertedName:  "Secret",
		Description:    `The secret associated to the AWS STS token.`,
		Exposed:        true,
		Name:           "secret",
		Required:       true,
		Type:           "string",
	},
	"token": {
		AllowedChoices: []string{},
		ConvertedName:  "Token",
		Description:    `The original token.`,
		Exposed:        true,
		Name:           "token",
		Required:       true,
		Type:           "string",
	},
}

type mongoAttributesIssueAWS struct {
}
