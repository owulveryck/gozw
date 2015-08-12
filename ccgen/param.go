package ccgen

import "errors"

type Param struct {
	Key            string `xml:"key,attr"`
	Name           string `xml:"name,attr"`
	Type           string `xml:"type,attr"`
	TypeHashCode   string `xml:"typehashcode,attr"`
	Comment        string `xml:"comment,attr"`
	EncapType      string `xml:"encaptype,attr"`
	OptionalOffset string `xml:"optionaloffs,attr"`
	OptionalMask   string `xml:"optionalmask,attr"`
	Encapsulated   bool   `xml:"encapsulated,attr"`
	CommandMask    string `xml:"cmd_mask,attr"`
	Affix          bool   `xml:"affix,attr"`

	ArrayAttrib []ArrayAttrib `xml:"arrayattrib"`
	Bit24       []Bit24       `xml:"bit_24"`
	BitField    []BitField    `xml:"bitfield"`
	BitFlag     []BitFlag     `xml:"bitflag"`
	Const       []Const       `xml:"const"`
	DWord       []DWord       `xml:"dword"`
	FieldEnum   []FieldEnum   `xml:"fieldenum"`
	ValueAttrib []ValueAttrib `xml:"valueattrib"`
	Variant     []Variant     `xml:"variant"`
	Bitmask     []Bitmask     `xml:"bitmask"`
	Word        []Word        `xml:"word"`
}

func (p Param) GetGoType() (string, error) {
	switch p.Type {

	case "ARRAY":
		if p.ArrayAttrib != nil && len(p.ArrayAttrib) == 1 {
			if p.ArrayAttrib[0].IsAscii {
				return "string", nil
			} else {
				return "[]byte", nil
			}
		} else if p.ArrayAttrib != nil {
			return "", errors.New("Weird number of <ArrayAttrib> elements")
		}
		return "[]byte", nil

	case "BIT_24":
		return "uint32", nil

	case "BITMASK":
		return "byte", nil

	case "BYTE":
		return "byte", nil

	case "CONST":
		return "byte", nil

	case "DWORD":
		return "uint32", nil

	case "ENUM":
		return "", errors.New("Unimplemented param type: ENUM")

	case "ENUM_ARRAY":
		return "", errors.New("Unimplemented param type: ENUM_ARRAY")

	case "MARKER":
		return "", errors.New("Unimplemented param type: MARKER")

	case "MULTI_ARRAY":
		return "", errors.New("Unimplemented param type: MULTI_ARRAY")

	case "STRUCT_BYTE":
		return "", errors.New("Unimplemented param type: STRUCT_BYTE")

	case "VARIANT":
		return "[]byte", nil

	case "WORD":
		return "uint16", nil

	default:
		return "", errors.New("Unknown param type: " + p.Type)

	}
}