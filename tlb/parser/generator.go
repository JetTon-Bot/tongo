package parser

import (
	"fmt"
	"github.com/startfellows/tongo/utils"
	"go/format"
	"strconv"
	"strings"
)

type DefaultType struct {
	Name          string
	IsPointerType bool
}

type Generator struct {
	knownTypes map[string]DefaultType
	newTlTypes []string
	typeName   string
}

var (
	defaultKnownTypes = map[string]DefaultType{
		"#":       {"uint32", false},
		"int32":   {"int32", false},
		"bits256": {"tongo.Hash", false},
	}
)

func NewGenerator(knownTypes map[string]DefaultType, typeName string) *Generator {
	if knownTypes == nil {
		knownTypes = defaultKnownTypes
	}
	return &Generator{
		knownTypes: knownTypes,
		typeName:   typeName,
	}
}

func (g *Generator) LoadTypes(declarations []CombinatorDeclaration) (string, error) {
	return generateGolangTypes(declarations)
}

func generateGolangTypes(t []CombinatorDeclaration) (string, error) {
	sumTypes := make(map[string][]CombinatorDeclaration)

	for _, c := range t {
		if len(c.Combinator.TypeExpressions) > 0 {
			return "", fmt.Errorf("combinators with paramaters '%v' are not supported", c.Combinator.Name)
		}
		sumTypes[c.Combinator.Name] = append(sumTypes[c.Combinator.Name], c)
	}
	s := ""
	for _, v := range sumTypes {
		t, err := generateGolangType(v)
		if err != nil {
			return "", err
		}
		s += "\n" + t
	}

	b, err := format.Source([]byte(s))
	if err != nil {
		return s, err
	}
	return string(b), err
}

func generateGolangStruct(declaration CombinatorDeclaration) (string, error) {
	builder := strings.Builder{}
	builder.WriteString("struct{")
	if len(declaration.FieldDefinitions) > 0 {
		builder.WriteRune('\n')
	}
	for i, field := range declaration.FieldDefinitions {
		if field.IsEmpty() {
			return "", fmt.Errorf("all types are nil in field %v in %v", i, declaration.Constructor.Name)
		}
		if field.Implicit != nil {
			continue
		}
		var name string
		var e TypeExpression
		if field.CellRef != nil {
			e = field.CellRef.TypeExpression
		} else if field.NamedField != nil {
			name = field.NamedField.Name
			e = field.NamedField.Expression
		}
		if name == "" || name == "_" {
			name = fmt.Sprintf("Field%v", i)
		}
		builder.WriteString(utils.ToCamelCase(name))
		builder.WriteRune('\t')
		t, err := e.ToGolangType()
		if err != nil {
			return "", err
		}
		builder.WriteString(t.String())
		_ = e
		builder.WriteRune('\n')
	}
	builder.WriteRune('}')
	return builder.String(), nil
}

func generateGolangSimpleType(declaration CombinatorDeclaration) (string, error) {
	s, err := generateGolangStruct(declaration)
	return fmt.Sprintf("type %v %v", declaration.Combinator.Name, s), err
}

func generateGolangSumType(declarations []CombinatorDeclaration) (string, error) {
	builder := strings.Builder{}
	builder.WriteString("type " + declarations[0].Combinator.Name + " struct{\ntlb.SumType\n")
	for _, d := range declarations {
		s, err := generateGolangStruct(d)
		if err != nil {
			return "", err
		}
		builder.WriteString(utils.ToCamelCase(d.Constructor.Name))
		builder.WriteRune(' ')
		builder.WriteString(s)
		builder.WriteString(fmt.Sprintf(" `tlbSumType:\"%v\"`", d.Constructor.Prefix))
		builder.WriteRune('\n')
	}
	builder.WriteRune('}')
	return builder.String(), nil

}

func generateGolangType(declarations []CombinatorDeclaration) (string, error) {
	if len(declarations) == 1 {
		return generateGolangSimpleType(declarations[0])
	} else {
		return generateGolangSumType(declarations)
	}
}

type golangType struct {
	name   string
	tag    string
	params []golangType
}

func (t TypeExpression) ToGolangType() (golangType, error) {
	if t.ParenExpression != nil {
		return t.ParenExpression.ToGolangType()
	}
	if t.NamedRef != nil {
		return mapToGoType(*t.NamedRef, false), nil
	}
	if t.BuiltIn != nil {
		return mapToGoType(*t.BuiltIn, false), nil
	}
	if t.Number != nil {
		return mapToGoType(fmt.Sprintf("%d", *t.Number), false), nil
	}
	if t.CellRef != nil {
		gt, err := t.CellRef.TypeExpression.ToGolangType()
		if err != nil {
			return golangType{}, err
		}
		gt.tag = "^"
		return mapToGoType(gt.String(), false), nil
	}

	return golangType{
		name: "UnknownType",
		tag:  "",
	}, nil
}

//func (t TypeExpression) String() string {
//	if t.NamedRef != nil {
//		return *t.NamedRef
//	}
//	if t.BuiltIn != nil {
//		return *t.BuiltIn
//	}
//	if t.Number != nil {
//		return fmt.Sprintf("%d", *t.Number)
//	}
//	if t.ParenExpression != nil {
//		s, _ := t.ParenExpression.ToGolangType()
//		return s.String()
//	}
//	return "Temp" //todo: implement
//}

func (t *ParenExpression) ToGolangType() (golangType, error) {
	var res golangType
	name, err := t.Name.ToGolangType()
	if err != nil {
		return golangType{}, err
	}
	res.name = name.String()
	switch res.name {
	case "HashmapE":
		if len(t.Parameter) != 2 {
			return golangType{}, fmt.Errorf("invalid parameters qty for HashmapE")
		}
		p, err := t.Parameter[0].ToGolangType()
		if err != nil {
			return golangType{}, err
		}
		res.tag = p.String()
		p, err = t.Parameter[1].ToGolangType()
		if err != nil {
			return golangType{}, err
		}
		res.params = []golangType{p}
		return res, nil
	case "Maybe":
		if len(t.Parameter) != 1 {
			return golangType{}, fmt.Errorf("invalid parameters qty for Maybe")
		}
		p, err := t.Parameter[0].ToGolangType()
		if err != nil {
			return golangType{}, err
		}
		res.params = []golangType{p}
		return res, nil
	case "VarUInteger":
		if len(t.Parameter) != 1 {
			return golangType{}, fmt.Errorf("invalid parameters qty for VarUInteger")
		}
		p, err := t.Parameter[0].ToGolangType()
		if err != nil {
			return golangType{}, err
		}
		res.tag = p.String()
		return res, nil
	case "##":
		if len(t.Parameter) != 1 {
			return golangType{}, fmt.Errorf("invalid parameters qty for ##")
		}
		p, err := t.Parameter[0].ToGolangType()
		if err != nil {
			return golangType{}, err
		}
		res.tag = p.String()
		return res, nil
	case "Bits":
		if len(t.Parameter) != 1 {
			return golangType{}, fmt.Errorf("invalid parameters qty for Bits")
		}
		if t.Parameter[0].Number != nil { // static type
			p, err := t.Parameter[0].ToGolangType()
			if err != nil {
				return golangType{}, err
			}
			res.tag = p.String()
		}
		return res, nil
	}

	for _, p := range t.Parameter {
		param, err := p.ToGolangType()
		if err != nil {
			return golangType{}, err
		}
		res.params = append(res.params, param)
	}
	return res, nil
}

func mapToGoType(name string, optional bool) golangType {
	goType, ok := defaultKnownTypes[name] // TODO: replace to generator field
	if ok {
		return golangType{
			name: goType.Name,
			//optional:    optional,
			//pointerType: goType.IsPointerType,
		}
	}
	t, ok := parseBuildInInt(name)
	if ok {
		return t
	}
	if name == "##" {
		return golangType{
			name: name,
			//optional:    optional,
			//pointerType: false,
		}
	}
	return golangType{
		name: utils.ToCamelCase(name),
		//optional:    optional,
		//pointerType: false,
	}
}

func parseBuildInInt(s string) (golangType, bool) {
	if strings.HasPrefix(s, "int") {
		last := strings.TrimPrefix(s, "int")
		bits, err := strconv.Atoi(last)
		if err != nil {
			return golangType{}, false
		}
		return golangType{
			name: fmt.Sprintf("uint64 `tlb:\"%dbits\"`", bits),
			//optional:    optional,
			//pointerType: false,
		}, true
	}
	return golangType{}, false
}

func (t golangType) String() string {
	switch t.name {
	case "HashmapE":
		if len(t.params) != 1 {
			return t.name
		}
		var pStr string
		if t.params[0].tag != "" {
			pStr = fmt.Sprintf("struct {Val %s}", t.params[0].String())
		} else {
			pStr = t.params[0].String()
		}
		tStr := fmt.Sprintf("tlb.%s[%s] `tlb:\"%sbits\"`", t.name, pStr, t.tag)
		return tStr
	case "Maybe":
		if len(t.params) != 1 {
			return t.name
		}
		var pStr string
		if t.params[0].tag != "" {
			pStr = fmt.Sprintf("struct {Val %s}", t.params[0].String())
		} else {
			pStr = t.params[0].String()
		}
		tStr := fmt.Sprintf("tlb.%s[%s]", t.name, pStr)
		return tStr
	case "VarUInteger":
		if t.tag == "" {
			return t.name
		}
		tStr := fmt.Sprintf("tlb.%s `tlb:\"%sbytes\"`", t.name, t.tag)
		return tStr
	case "##":
		if t.tag == "" {
			return t.name
		}
		tStr := fmt.Sprintf("uint64 `tlb:\"%sbits\"`", t.tag) // max 32 bits in block.tlb
		return tStr
	case "Bits":
		if t.tag == "" {
			return "tlb.BitString"
		}
		tStr := fmt.Sprintf("tlb.BitString `tlb:\"%sbits\"`", t.tag)
		return tStr
	default:
		return t.name
	}
}
