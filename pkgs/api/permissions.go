package api

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// PermissionsIdentity represents the Identity of the object.
var PermissionsIdentity = elemental.Identity{
	Name:     "permissions",
	Category: "permissions",
	Package:  "a3s",
	Private:  false,
}

// PermissionsList represents a list of Permissions
type PermissionsList []*Permissions

// Identity returns the identity of the objects in the list.
func (o PermissionsList) Identity() elemental.Identity {

	return PermissionsIdentity
}

// Copy returns a pointer to a copy the PermissionsList.
func (o PermissionsList) Copy() elemental.Identifiables {

	copy := append(PermissionsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the PermissionsList.
func (o PermissionsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(PermissionsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*Permissions))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o PermissionsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o PermissionsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the PermissionsList converted to SparsePermissionsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o PermissionsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparsePermissionsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparsePermissions)
	}

	return out
}

// Version returns the version of the content.
func (o PermissionsList) Version() int {

	return 1
}

// Permissions represents the model of a permissions
type Permissions struct {
	// The optional ID of the object to check permission for.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// IP of the client.
	IP string `json:"IP" msgpack:"IP" bson:"-" mapstructure:"IP,omitempty"`

	// The list of claims.
	Claims []string `json:"claims" msgpack:"claims" bson:"-" mapstructure:"claims,omitempty"`

	// Return an eventual error.
	Error string `json:"error,omitempty" msgpack:"error,omitempty" bson:"-" mapstructure:"error,omitempty"`

	// The namespace where to check permission from.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"-" mapstructure:"namespace,omitempty"`

	// The computed permissions.
	Permissions map[string]map[string]bool `json:"permissions,omitempty" msgpack:"permissions,omitempty" bson:"-" mapstructure:"permissions,omitempty"`

	// Sets the namespace restrictions that should apply.
	RestrictedNamespace string `json:"restrictedNamespace" msgpack:"restrictedNamespace" bson:"-" mapstructure:"restrictedNamespace,omitempty"`

	// Sets the networks restrictions that should apply.
	RestrictedNetworks []string `json:"restrictedNetworks" msgpack:"restrictedNetworks" bson:"-" mapstructure:"restrictedNetworks,omitempty"`

	// Sets the permissions restrictions that should apply.
	RestrictedPermissions []string `json:"restrictedPermissions" msgpack:"restrictedPermissions" bson:"-" mapstructure:"restrictedPermissions,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewPermissions returns a new *Permissions
func NewPermissions() *Permissions {

	return &Permissions{
		ModelVersion:          1,
		Claims:                []string{},
		Permissions:           map[string]map[string]bool{},
		RestrictedNetworks:    []string{},
		RestrictedPermissions: []string{},
	}
}

// Identity returns the Identity of the object.
func (o *Permissions) Identity() elemental.Identity {

	return PermissionsIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Permissions) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Permissions) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Permissions) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesPermissions{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Permissions) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesPermissions{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *Permissions) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *Permissions) BleveType() string {

	return "permissions"
}

// DefaultOrder returns the list of default ordering fields.
func (o *Permissions) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *Permissions) Doc() string {

	return `API to retrieve the permissions from a user identity.`
}

func (o *Permissions) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *Permissions) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparsePermissions{
			ID:                    &o.ID,
			IP:                    &o.IP,
			Claims:                &o.Claims,
			Error:                 &o.Error,
			Namespace:             &o.Namespace,
			Permissions:           &o.Permissions,
			RestrictedNamespace:   &o.RestrictedNamespace,
			RestrictedNetworks:    &o.RestrictedNetworks,
			RestrictedPermissions: &o.RestrictedPermissions,
		}
	}

	sp := &SparsePermissions{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "IP":
			sp.IP = &(o.IP)
		case "claims":
			sp.Claims = &(o.Claims)
		case "error":
			sp.Error = &(o.Error)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "permissions":
			sp.Permissions = &(o.Permissions)
		case "restrictedNamespace":
			sp.RestrictedNamespace = &(o.RestrictedNamespace)
		case "restrictedNetworks":
			sp.RestrictedNetworks = &(o.RestrictedNetworks)
		case "restrictedPermissions":
			sp.RestrictedPermissions = &(o.RestrictedPermissions)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparsePermissions to the object.
func (o *Permissions) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparsePermissions)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.IP != nil {
		o.IP = *so.IP
	}
	if so.Claims != nil {
		o.Claims = *so.Claims
	}
	if so.Error != nil {
		o.Error = *so.Error
	}
	if so.Namespace != nil {
		o.Namespace = *so.Namespace
	}
	if so.Permissions != nil {
		o.Permissions = *so.Permissions
	}
	if so.RestrictedNamespace != nil {
		o.RestrictedNamespace = *so.RestrictedNamespace
	}
	if so.RestrictedNetworks != nil {
		o.RestrictedNetworks = *so.RestrictedNetworks
	}
	if so.RestrictedPermissions != nil {
		o.RestrictedPermissions = *so.RestrictedPermissions
	}
}

// DeepCopy returns a deep copy if the Permissions.
func (o *Permissions) DeepCopy() *Permissions {

	if o == nil {
		return nil
	}

	out := &Permissions{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *Permissions.
func (o *Permissions) DeepCopyInto(out *Permissions) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy Permissions: %s", err))
	}

	*out = *target.(*Permissions)
}

// Validate valides the current information stored into the structure.
func (o *Permissions) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredExternal("claims", o.Claims); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("namespace", o.Namespace); err != nil {
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
func (*Permissions) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := PermissionsAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return PermissionsLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Permissions) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return PermissionsAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *Permissions) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "IP":
		return o.IP
	case "claims":
		return o.Claims
	case "error":
		return o.Error
	case "namespace":
		return o.Namespace
	case "permissions":
		return o.Permissions
	case "restrictedNamespace":
		return o.RestrictedNamespace
	case "restrictedNetworks":
		return o.RestrictedNetworks
	case "restrictedPermissions":
		return o.RestrictedPermissions
	}

	return nil
}

// PermissionsAttributesMap represents the map of attribute for Permissions.
var PermissionsAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": {
		AllowedChoices: []string{},
		ConvertedName:  "ID",
		Description:    `The optional ID of the object to check permission for.`,
		Exposed:        true,
		Name:           "ID",
		Type:           "string",
	},
	"IP": {
		AllowedChoices: []string{},
		ConvertedName:  "IP",
		Description:    `IP of the client.`,
		Exposed:        true,
		Name:           "IP",
		Type:           "string",
	},
	"Claims": {
		AllowedChoices: []string{},
		ConvertedName:  "Claims",
		Description:    `The list of claims.`,
		Exposed:        true,
		Name:           "claims",
		Required:       true,
		SubType:        "string",
		Type:           "list",
	},
	"Error": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Error",
		Description:    `Return an eventual error.`,
		Exposed:        true,
		Name:           "error",
		ReadOnly:       true,
		Type:           "string",
	},
	"Namespace": {
		AllowedChoices: []string{},
		ConvertedName:  "Namespace",
		Description:    `The namespace where to check permission from.`,
		Exposed:        true,
		Name:           "namespace",
		Required:       true,
		Type:           "string",
	},
	"Permissions": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Permissions",
		Description:    `The computed permissions.`,
		Exposed:        true,
		Name:           "permissions",
		ReadOnly:       true,
		SubType:        "map[string]map[string]bool",
		Type:           "external",
	},
	"RestrictedNamespace": {
		AllowedChoices: []string{},
		ConvertedName:  "RestrictedNamespace",
		Description:    `Sets the namespace restrictions that should apply.`,
		Exposed:        true,
		Name:           "restrictedNamespace",
		Type:           "string",
	},
	"RestrictedNetworks": {
		AllowedChoices: []string{},
		ConvertedName:  "RestrictedNetworks",
		Description:    `Sets the networks restrictions that should apply.`,
		Exposed:        true,
		Name:           "restrictedNetworks",
		SubType:        "string",
		Type:           "list",
	},
	"RestrictedPermissions": {
		AllowedChoices: []string{},
		ConvertedName:  "RestrictedPermissions",
		Description:    `Sets the permissions restrictions that should apply.`,
		Exposed:        true,
		Name:           "restrictedPermissions",
		SubType:        "string",
		Type:           "list",
	},
}

// PermissionsLowerCaseAttributesMap represents the map of attribute for Permissions.
var PermissionsLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"id": {
		AllowedChoices: []string{},
		ConvertedName:  "ID",
		Description:    `The optional ID of the object to check permission for.`,
		Exposed:        true,
		Name:           "ID",
		Type:           "string",
	},
	"ip": {
		AllowedChoices: []string{},
		ConvertedName:  "IP",
		Description:    `IP of the client.`,
		Exposed:        true,
		Name:           "IP",
		Type:           "string",
	},
	"claims": {
		AllowedChoices: []string{},
		ConvertedName:  "Claims",
		Description:    `The list of claims.`,
		Exposed:        true,
		Name:           "claims",
		Required:       true,
		SubType:        "string",
		Type:           "list",
	},
	"error": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Error",
		Description:    `Return an eventual error.`,
		Exposed:        true,
		Name:           "error",
		ReadOnly:       true,
		Type:           "string",
	},
	"namespace": {
		AllowedChoices: []string{},
		ConvertedName:  "Namespace",
		Description:    `The namespace where to check permission from.`,
		Exposed:        true,
		Name:           "namespace",
		Required:       true,
		Type:           "string",
	},
	"permissions": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Permissions",
		Description:    `The computed permissions.`,
		Exposed:        true,
		Name:           "permissions",
		ReadOnly:       true,
		SubType:        "map[string]map[string]bool",
		Type:           "external",
	},
	"restrictednamespace": {
		AllowedChoices: []string{},
		ConvertedName:  "RestrictedNamespace",
		Description:    `Sets the namespace restrictions that should apply.`,
		Exposed:        true,
		Name:           "restrictedNamespace",
		Type:           "string",
	},
	"restrictednetworks": {
		AllowedChoices: []string{},
		ConvertedName:  "RestrictedNetworks",
		Description:    `Sets the networks restrictions that should apply.`,
		Exposed:        true,
		Name:           "restrictedNetworks",
		SubType:        "string",
		Type:           "list",
	},
	"restrictedpermissions": {
		AllowedChoices: []string{},
		ConvertedName:  "RestrictedPermissions",
		Description:    `Sets the permissions restrictions that should apply.`,
		Exposed:        true,
		Name:           "restrictedPermissions",
		SubType:        "string",
		Type:           "list",
	},
}

// SparsePermissionsList represents a list of SparsePermissions
type SparsePermissionsList []*SparsePermissions

// Identity returns the identity of the objects in the list.
func (o SparsePermissionsList) Identity() elemental.Identity {

	return PermissionsIdentity
}

// Copy returns a pointer to a copy the SparsePermissionsList.
func (o SparsePermissionsList) Copy() elemental.Identifiables {

	copy := append(SparsePermissionsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparsePermissionsList.
func (o SparsePermissionsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparsePermissionsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparsePermissions))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparsePermissionsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparsePermissionsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparsePermissionsList converted to PermissionsList.
func (o SparsePermissionsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparsePermissionsList) Version() int {

	return 1
}

// SparsePermissions represents the sparse version of a permissions.
type SparsePermissions struct {
	// The optional ID of the object to check permission for.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// IP of the client.
	IP *string `json:"IP,omitempty" msgpack:"IP,omitempty" bson:"-" mapstructure:"IP,omitempty"`

	// The list of claims.
	Claims *[]string `json:"claims,omitempty" msgpack:"claims,omitempty" bson:"-" mapstructure:"claims,omitempty"`

	// Return an eventual error.
	Error *string `json:"error,omitempty" msgpack:"error,omitempty" bson:"-" mapstructure:"error,omitempty"`

	// The namespace where to check permission from.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"-" mapstructure:"namespace,omitempty"`

	// The computed permissions.
	Permissions *map[string]map[string]bool `json:"permissions,omitempty" msgpack:"permissions,omitempty" bson:"-" mapstructure:"permissions,omitempty"`

	// Sets the namespace restrictions that should apply.
	RestrictedNamespace *string `json:"restrictedNamespace,omitempty" msgpack:"restrictedNamespace,omitempty" bson:"-" mapstructure:"restrictedNamespace,omitempty"`

	// Sets the networks restrictions that should apply.
	RestrictedNetworks *[]string `json:"restrictedNetworks,omitempty" msgpack:"restrictedNetworks,omitempty" bson:"-" mapstructure:"restrictedNetworks,omitempty"`

	// Sets the permissions restrictions that should apply.
	RestrictedPermissions *[]string `json:"restrictedPermissions,omitempty" msgpack:"restrictedPermissions,omitempty" bson:"-" mapstructure:"restrictedPermissions,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparsePermissions returns a new  SparsePermissions.
func NewSparsePermissions() *SparsePermissions {
	return &SparsePermissions{}
}

// Identity returns the Identity of the sparse object.
func (o *SparsePermissions) Identity() elemental.Identity {

	return PermissionsIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparsePermissions) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparsePermissions) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparsePermissions) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparsePermissions{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparsePermissions) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparsePermissions{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparsePermissions) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparsePermissions) ToPlain() elemental.PlainIdentifiable {

	out := NewPermissions()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.IP != nil {
		out.IP = *o.IP
	}
	if o.Claims != nil {
		out.Claims = *o.Claims
	}
	if o.Error != nil {
		out.Error = *o.Error
	}
	if o.Namespace != nil {
		out.Namespace = *o.Namespace
	}
	if o.Permissions != nil {
		out.Permissions = *o.Permissions
	}
	if o.RestrictedNamespace != nil {
		out.RestrictedNamespace = *o.RestrictedNamespace
	}
	if o.RestrictedNetworks != nil {
		out.RestrictedNetworks = *o.RestrictedNetworks
	}
	if o.RestrictedPermissions != nil {
		out.RestrictedPermissions = *o.RestrictedPermissions
	}

	return out
}

// DeepCopy returns a deep copy if the SparsePermissions.
func (o *SparsePermissions) DeepCopy() *SparsePermissions {

	if o == nil {
		return nil
	}

	out := &SparsePermissions{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparsePermissions.
func (o *SparsePermissions) DeepCopyInto(out *SparsePermissions) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparsePermissions: %s", err))
	}

	*out = *target.(*SparsePermissions)
}

type mongoAttributesPermissions struct {
}
type mongoAttributesSparsePermissions struct {
}
