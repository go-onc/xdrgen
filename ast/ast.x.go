// Code generated by xdrgen-go - DO NOT EDIT.

// XDR Abstract Syntax Tree definition - a binary interchange format for XDR specifications
package ast

import (
	"encoding"
	"errors"
	"fmt"
	"strconv"

	xdr "go.e43.eu/xdr/interfaces"
)

// Binary magic: the `magic` field of the `specification` should be set to this value
const XDR_BIN_MAGIC = 0x895844520D0A1A0A

// Root object of a specification
type Specification struct {
	// Magic number: set to XDR_BIN_MAGIC
	Magic uint64 `json:"magic"`
	// Spec attributes (set using pragma directives)
	Attributes Attributes `json:"attributes"`
	// List of all definitions
	Definitions []*Definition `json:"definitions"`
}

// A set of attributes
type Attributes map[string]*Constant

// Definition_Body is union definition.body
type Definition_Body struct {
	Kind DefinitionKind `xdr:"union:switch" json:"kind"`
	// Body, for type definitions
	Type *Type `xdr:"union:0" json:"type,omitempty"`
	// Body, for constant definitions
	Constant *Constant `xdr:"union:1" json:"constant,omitempty"`
}

func (u *Definition_Body) UnionDiscriminant() interface{} {
	return u.Kind
}

func (u *Definition_Body) UnionValue() (interface{}, error) {
	switch u.Kind {
	case DEFINITION_KIND_CONSTANT:
		return u.Constant, nil
	case DEFINITION_KIND_TYPE:
		return u.Type, nil
	default:
		return nil, errors.New("Invalid discriminant")
	}
}

// A top-level definition
type Definition struct {
	// The name of the definition
	Name string `json:"name"`
	// The attributes of the definition
	Attributes Attributes       `json:"attributes"`
	Body       *Definition_Body `json:"body"`
}

// An attribute of an object
type Attribute struct {
	Name  string    `json:"name"`
	Value *Constant `json:"value"`
}

// Constant is union constant
type Constant struct {
	Type    ConstantKind `xdr:"union:switch" json:"type"`
	_       struct{}     `xdr:"union:6" json:",omitempty"`
	VBool   bool         `xdr:"union:0" json:"v_bool,omitempty"`
	VPosInt uint64       `xdr:"union:1" json:"v_pos_int,omitempty"`
	VNegInt uint64       `xdr:"union:2" json:"v_neg_int,omitempty"`
	VFloat  float64      `xdr:"union:3" json:"v_float,omitempty"`
	VString string       `xdr:"union:4" json:"v_string,omitempty"`
	VEnum   uint32       `xdr:"union:5" json:"v_enum,omitempty"`
}

func (u *Constant) UnionDiscriminant() interface{} {
	return u.Type
}

func (u *Constant) UnionValue() (interface{}, error) {
	switch u.Type {
	case CONST_BOOL:
		return u.VBool, nil
	case CONST_ENUM:
		return u.VEnum, nil
	case CONST_FLOAT:
		return u.VFloat, nil
	case CONST_NEG_INT:
		return u.VNegInt, nil
	case CONST_POS_INT:
		return u.VPosInt, nil
	case CONST_STRING:
		return u.VString, nil
	case CONST_VOID:
		return nil, nil
	default:
		return nil, errors.New("Invalid discriminant")
	}
}

// The kind of a definition
type DefinitionKind uint32

const (
	DEFINITION_KIND_TYPE     DefinitionKind = 0
	DEFINITION_KIND_CONSTANT DefinitionKind = 1
)

var xDefinitionKindValToStr = map[DefinitionKind]string{
	DEFINITION_KIND_TYPE:     "DEFINITION_KIND_TYPE",     // 0
	DEFINITION_KIND_CONSTANT: "DEFINITION_KIND_CONSTANT", // 1
}

var xDefinitionKindStrToVal = map[string]DefinitionKind{
	"DEFINITION_KIND_TYPE":     DEFINITION_KIND_TYPE,
	"DEFINITION_KIND_CONSTANT": DEFINITION_KIND_CONSTANT,
}

// String satisfies fmt.Stringer
func (v DefinitionKind) String() string {
	if s, ok := xDefinitionKindValToStr[v]; ok {
		return s
	}
	return strconv.Itoa(int(v))
}

// MarshalText satisfies encoding.TextMarshaler
func (v DefinitionKind) MarshalText() ([]byte, error) {
	if s, ok := xDefinitionKindValToStr[v]; ok {
		return []byte(s), nil
	}
	return nil, errors.New("Invalid enum value")
}

// UnmarshalText satisfies encoding.TextUnmarshaler
func (v *DefinitionKind) UnmarshalText(buf []byte) error {
	if nv, ok := xDefinitionKindStrToVal[string(buf)]; ok {
		*v = nv
		return nil
	}
	return errors.New("Invalid enum value")
}

func (v DefinitionKind) IsKnown() bool {
	_, ok := xDefinitionKindValToStr[v]
	return ok
}

var (
	_ fmt.Stringer             = DefinitionKind(0)
	_ encoding.TextMarshaler   = DefinitionKind(0)
	_ encoding.TextUnmarshaler = new(DefinitionKind)
)

// Definition of a type
type Type struct {
	Kind       TypeKind     `xdr:"union:switch" json:"kind"`
	_          struct{}     `xdr:"union:5,9,3,4,7,2,1,6,0,8" json:",omitempty"`
	EnumSpec   *EnumSpec    `xdr:"union:10" json:"enum_spec,omitempty"`
	StructSpec *StructSpec  `xdr:"union:11" json:"struct_spec,omitempty"`
	UnionSpec  *UnionSpec   `xdr:"union:12" json:"union_spec,omitempty"`
	Ref        uint32       `xdr:"union:13" json:"ref,omitempty"`
	TypeDef    *Declaration `xdr:"union:14" json:"type_def,omitempty"`
}

func (u *Type) UnionDiscriminant() interface{} {
	return u.Kind
}

func (u *Type) UnionValue() (interface{}, error) {
	switch u.Kind {
	case TYPE_BOOL:
		return nil, nil
	case TYPE_DOUBLE:
		return nil, nil
	case TYPE_ENUM:
		return u.EnumSpec, nil
	case TYPE_FLOAT:
		return nil, nil
	case TYPE_HYPER:
		return nil, nil
	case TYPE_INT:
		return nil, nil
	case TYPE_OPAQUE:
		return nil, nil
	case TYPE_REF:
		return u.Ref, nil
	case TYPE_STRING:
		return nil, nil
	case TYPE_STRUCT:
		return u.StructSpec, nil
	case TYPE_TYPEDEF:
		return u.TypeDef, nil
	case TYPE_UNION:
		return u.UnionSpec, nil
	case TYPE_UNSIGNED_HYPER:
		return nil, nil
	case TYPE_UNSIGNED_INT:
		return nil, nil
	case TYPE_VOID:
		return nil, nil
	default:
		return nil, errors.New("Invalid discriminant")
	}
}

// The kind of the type
type TypeKind uint32

const (
	TYPE_VOID           TypeKind = 0
	TYPE_BOOL           TypeKind = 1
	TYPE_INT            TypeKind = 2
	TYPE_UNSIGNED_INT   TypeKind = 3
	TYPE_HYPER          TypeKind = 4
	TYPE_UNSIGNED_HYPER TypeKind = 5
	TYPE_FLOAT          TypeKind = 6
	TYPE_DOUBLE         TypeKind = 7
	TYPE_STRING         TypeKind = 8
	TYPE_OPAQUE         TypeKind = 9
	TYPE_ENUM           TypeKind = 10
	TYPE_STRUCT         TypeKind = 11
	TYPE_UNION          TypeKind = 12
	TYPE_REF            TypeKind = 13
	TYPE_TYPEDEF        TypeKind = 14
)

var xTypeKindValToStr = map[TypeKind]string{
	TYPE_VOID:           "TYPE_VOID",           // 0
	TYPE_BOOL:           "TYPE_BOOL",           // 1
	TYPE_INT:            "TYPE_INT",            // 2
	TYPE_UNSIGNED_INT:   "TYPE_UNSIGNED_INT",   // 3
	TYPE_HYPER:          "TYPE_HYPER",          // 4
	TYPE_UNSIGNED_HYPER: "TYPE_UNSIGNED_HYPER", // 5
	TYPE_FLOAT:          "TYPE_FLOAT",          // 6
	TYPE_DOUBLE:         "TYPE_DOUBLE",         // 7
	TYPE_STRING:         "TYPE_STRING",         // 8
	TYPE_OPAQUE:         "TYPE_OPAQUE",         // 9
	TYPE_ENUM:           "TYPE_ENUM",           // 10
	TYPE_STRUCT:         "TYPE_STRUCT",         // 11
	TYPE_UNION:          "TYPE_UNION",          // 12
	TYPE_REF:            "TYPE_REF",            // 13
	TYPE_TYPEDEF:        "TYPE_TYPEDEF",        // 14
}

var xTypeKindStrToVal = map[string]TypeKind{
	"TYPE_VOID":           TYPE_VOID,
	"TYPE_BOOL":           TYPE_BOOL,
	"TYPE_INT":            TYPE_INT,
	"TYPE_UNSIGNED_INT":   TYPE_UNSIGNED_INT,
	"TYPE_HYPER":          TYPE_HYPER,
	"TYPE_UNSIGNED_HYPER": TYPE_UNSIGNED_HYPER,
	"TYPE_FLOAT":          TYPE_FLOAT,
	"TYPE_DOUBLE":         TYPE_DOUBLE,
	"TYPE_STRING":         TYPE_STRING,
	"TYPE_OPAQUE":         TYPE_OPAQUE,
	"TYPE_ENUM":           TYPE_ENUM,
	"TYPE_STRUCT":         TYPE_STRUCT,
	"TYPE_UNION":          TYPE_UNION,
	"TYPE_REF":            TYPE_REF,
	"TYPE_TYPEDEF":        TYPE_TYPEDEF,
}

// String satisfies fmt.Stringer
func (v TypeKind) String() string {
	if s, ok := xTypeKindValToStr[v]; ok {
		return s
	}
	return strconv.Itoa(int(v))
}

// MarshalText satisfies encoding.TextMarshaler
func (v TypeKind) MarshalText() ([]byte, error) {
	if s, ok := xTypeKindValToStr[v]; ok {
		return []byte(s), nil
	}
	return nil, errors.New("Invalid enum value")
}

// UnmarshalText satisfies encoding.TextUnmarshaler
func (v *TypeKind) UnmarshalText(buf []byte) error {
	if nv, ok := xTypeKindStrToVal[string(buf)]; ok {
		*v = nv
		return nil
	}
	return errors.New("Invalid enum value")
}

func (v TypeKind) IsKnown() bool {
	_, ok := xTypeKindValToStr[v]
	return ok
}

var (
	_ fmt.Stringer             = TypeKind(0)
	_ encoding.TextMarshaler   = TypeKind(0)
	_ encoding.TextUnmarshaler = new(TypeKind)
)

// Definition of an enum. Represented by referencing the assoicaed global constants
type EnumSpec struct {
	// First constant that is a part of this enumeration
	Base uint32 `json:"base"`
	// Number of constants (which must consecutively follow base) that define components of this enumeration
	Count uint32 `json:"count"`
}

// Definition of an enum
type StructSpec struct {
	// Set of struct members
	Members []*Declaration `json:"members"`
}

// Mapping from values to union member. `member` is the index of the member in `members`
type UnionSpec_Options struct {
	Value  uint32 `json:"value"`
	Member uint32 `json:"member"`
}

// Definition of a union
type UnionSpec struct {
	// Discriminant field
	Discriminant *Declaration `json:"discriminant"`
	// Set of union member fields
	Members []*Declaration `json:"members"`
	// Mapping from values to union member. `member` is the index of the member in `members`
	Options map[uint32]uint32 `json:"options"`
	// If a default member is present, defines it
	DefaultMember *uint32 `xdr:"opt" json:"default_member,omitempty"`
}

// Modifier of the type
type Declaration_Modifier struct {
	Kind DeclarationModifier `xdr:"union:switch" json:"kind"`
	_    struct{}            `xdr:"union:0,1,5" json:",omitempty"`
	Size uint32              `xdr:"union:2,3" json:"size,omitempty"`
}

func (u *Declaration_Modifier) UnionDiscriminant() interface{} {
	return u.Kind
}

func (u *Declaration_Modifier) UnionValue() (interface{}, error) {
	switch u.Kind {
	case DECLARATION_MODIFIER_FIXED:
		return u.Size, nil
	case DECLARATION_MODIFIER_FLEXIBLE:
		return u.Size, nil
	case DECLARATION_MODIFIER_NONE:
		return nil, nil
	case DECLARATION_MODIFIER_OPTIONAL:
		return nil, nil
	case DECLARATION_MODIFIER_UNBOUNDED:
		return nil, nil
	default:
		return nil, errors.New("Invalid discriminant")
	}
}

// Field declaration
type Declaration struct {
	// Type of the field
	Type *Type `json:"type"`
	// Name of the field
	Name string `json:"name"`
	// Modifier of the type
	Modifier *Declaration_Modifier `json:"modifier"`
	// Field attributes
	Attributes Attributes `json:"attributes"`
}

// How a declaration modifies its type
type DeclarationModifier uint32

const (
	DECLARATION_MODIFIER_NONE      DeclarationModifier = 0
	DECLARATION_MODIFIER_OPTIONAL  DeclarationModifier = 1
	DECLARATION_MODIFIER_FIXED     DeclarationModifier = 2
	DECLARATION_MODIFIER_FLEXIBLE  DeclarationModifier = 3
	DECLARATION_MODIFIER_UNBOUNDED DeclarationModifier = 5
)

var xDeclarationModifierValToStr = map[DeclarationModifier]string{
	DECLARATION_MODIFIER_NONE:      "DECLARATION_MODIFIER_NONE",      // 0
	DECLARATION_MODIFIER_OPTIONAL:  "DECLARATION_MODIFIER_OPTIONAL",  // 1
	DECLARATION_MODIFIER_FIXED:     "DECLARATION_MODIFIER_FIXED",     // 2
	DECLARATION_MODIFIER_FLEXIBLE:  "DECLARATION_MODIFIER_FLEXIBLE",  // 3
	DECLARATION_MODIFIER_UNBOUNDED: "DECLARATION_MODIFIER_UNBOUNDED", // 5
}

var xDeclarationModifierStrToVal = map[string]DeclarationModifier{
	"DECLARATION_MODIFIER_NONE":      DECLARATION_MODIFIER_NONE,
	"DECLARATION_MODIFIER_OPTIONAL":  DECLARATION_MODIFIER_OPTIONAL,
	"DECLARATION_MODIFIER_FIXED":     DECLARATION_MODIFIER_FIXED,
	"DECLARATION_MODIFIER_FLEXIBLE":  DECLARATION_MODIFIER_FLEXIBLE,
	"DECLARATION_MODIFIER_UNBOUNDED": DECLARATION_MODIFIER_UNBOUNDED,
}

// String satisfies fmt.Stringer
func (v DeclarationModifier) String() string {
	if s, ok := xDeclarationModifierValToStr[v]; ok {
		return s
	}
	return strconv.Itoa(int(v))
}

// MarshalText satisfies encoding.TextMarshaler
func (v DeclarationModifier) MarshalText() ([]byte, error) {
	if s, ok := xDeclarationModifierValToStr[v]; ok {
		return []byte(s), nil
	}
	return nil, errors.New("Invalid enum value")
}

// UnmarshalText satisfies encoding.TextUnmarshaler
func (v *DeclarationModifier) UnmarshalText(buf []byte) error {
	if nv, ok := xDeclarationModifierStrToVal[string(buf)]; ok {
		*v = nv
		return nil
	}
	return errors.New("Invalid enum value")
}

func (v DeclarationModifier) IsKnown() bool {
	_, ok := xDeclarationModifierValToStr[v]
	return ok
}

var (
	_ fmt.Stringer             = DeclarationModifier(0)
	_ encoding.TextMarshaler   = DeclarationModifier(0)
	_ encoding.TextUnmarshaler = new(DeclarationModifier)
)

// Type of a constant. These are a subset of XDR types
type ConstantKind uint32

const (
	CONST_VOID    ConstantKind = 6
	CONST_BOOL    ConstantKind = 0
	CONST_POS_INT ConstantKind = 1
	CONST_NEG_INT ConstantKind = 2
	CONST_FLOAT   ConstantKind = 3
	CONST_STRING  ConstantKind = 4
	CONST_ENUM    ConstantKind = 5
)

var xConstantKindValToStr = map[ConstantKind]string{
	CONST_BOOL:    "CONST_BOOL",    // 0
	CONST_POS_INT: "CONST_POS_INT", // 1
	CONST_NEG_INT: "CONST_NEG_INT", // 2
	CONST_FLOAT:   "CONST_FLOAT",   // 3
	CONST_STRING:  "CONST_STRING",  // 4
	CONST_ENUM:    "CONST_ENUM",    // 5
	CONST_VOID:    "CONST_VOID",    // 6
}

var xConstantKindStrToVal = map[string]ConstantKind{
	"CONST_VOID":    CONST_VOID,
	"CONST_BOOL":    CONST_BOOL,
	"CONST_POS_INT": CONST_POS_INT,
	"CONST_NEG_INT": CONST_NEG_INT,
	"CONST_FLOAT":   CONST_FLOAT,
	"CONST_STRING":  CONST_STRING,
	"CONST_ENUM":    CONST_ENUM,
}

// String satisfies fmt.Stringer
func (v ConstantKind) String() string {
	if s, ok := xConstantKindValToStr[v]; ok {
		return s
	}
	return strconv.Itoa(int(v))
}

// MarshalText satisfies encoding.TextMarshaler
func (v ConstantKind) MarshalText() ([]byte, error) {
	if s, ok := xConstantKindValToStr[v]; ok {
		return []byte(s), nil
	}
	return nil, errors.New("Invalid enum value")
}

// UnmarshalText satisfies encoding.TextUnmarshaler
func (v *ConstantKind) UnmarshalText(buf []byte) error {
	if nv, ok := xConstantKindStrToVal[string(buf)]; ok {
		*v = nv
		return nil
	}
	return errors.New("Invalid enum value")
}

func (v ConstantKind) IsKnown() bool {
	_, ok := xConstantKindValToStr[v]
	return ok
}

var (
	_ fmt.Stringer             = ConstantKind(0)
	_ encoding.TextMarshaler   = ConstantKind(0)
	_ encoding.TextUnmarshaler = new(ConstantKind)
)

// Dummy type assertions - added to ensure that no errors are generated
// because we didn't use one of our imports
var (
	_ encoding.TextMarshaler = nil
	_ fmt.Stringer           = nil
	_ xdr.Marshaler          = nil
	_                        = strconv.ErrSyntax
	_                        = errors.New
)
