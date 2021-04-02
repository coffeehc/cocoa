package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/hsiafan/cocoa/generator/gen"
	"github.com/hsiafan/cocoa/generator/typing"
	"github.com/hsiafan/glow/httpx"
	"github.com/hsiafan/glow/iox/filex"
	"github.com/hsiafan/glow/stringx"
	"github.com/hsiafan/glow/timex"
)

// Fetcher fetch and parse one doc
type Fetcher struct {
	client *httpx.Client
}

func NewFetcher() *Fetcher {
	return &Fetcher{
		client: httpx.NewClient(
			httpx.DialTimeout(timex.SecondsDuration(30)),
		),
	}
}

func (f *Fetcher) fetchType(identifier string) (*gen.Class, error) {
	response, err := f.fetchResponse(identifier)
	if err != nil {
		return nil, err
	}
	metadata := response.Metadata
	if metadata.SymbolKind != "cl" {
		return nil, fmt.Errorf("%v is not a class type, but: %v", identifier, metadata.SymbolKind)
	}

	coder := &gen.Class{}
	module := f.moduleForIdentifier(identifier)
	coder.Type = &typing.ClassType{
		Name:   metadata.Title,
		GName:  metadata.Title,
		Module: &module,
	}

	for _, section := range response.TopicSections {
		switch section.Kind {
		case "taskGroup":
			for _, identifier := range section.Identifiers {
				reference := response.References[identifier]
				if reference.Kind != "symbol" || reference.Deprecated {
					continue
				}
				ir, err := f.fetchResponse(identifier)
				if err != nil {
					return nil, fmt.Errorf("fetch identifier response failed %v: %w ", identifier, err)
				}
				iMeta := ir.Metadata
				sections := ir.PrimaryContentSections
				switch iMeta.SymbolKind {
				case "instm":
					method := f.parseMethod(sections)
					coder.Methods = append(coder.Methods, method)
				case "instp":
					property := f.parseProperty(sections)
					coder.Properties = append(coder.Properties, property)
				case "clm":
					method := f.parseMethod(sections)
					coder.Methods = append(coder.Methods, method)
				case "econst":
				case "enum":
				case "data":
				case "intf":
				case "tdef":
				default:
					fmt.Println("unknown symbol kind:", ir.Metadata.SymbolKind, ir.Metadata.RoleHeading)
				}
			}
		default:
			fmt.Println("unknown topic section type:", section.Kind)
		}
	}
	return coder, nil
}

func (f *Fetcher) parseMethod(sections []*PrimaryContentSection) *gen.Method {
	method := &gen.Method{}
	for _, section := range sections {
		if section.Kind == "declarations" {
			signature, typeMap := f.joinTokenText(section.Declarations[0].Tokens)
			if strings.Contains(signature, "^") {
				//TODO: block type
				fmt.Println("skip block type:", signature)
				continue
			}
			signature = stringx.AppendIfMissing(signature, ";")
			t := &Tokenizer{strings.NewReader(signature)}
			prefix := t.NextTokenTill('(')
			if prefix == "+ " {
				method.Static = true
			}
			method.ReturnType = f.ParseType(t.NextParenthesesContent('(', ')'), typeMap)
			method.Name = t.NextTokenTill(':', ';')
			c := t.NextChar()
			if c == ';' {
				continue
			}
			t.PutBack()
			for {
				var param = &gen.Param{}
				param.FieldName = t.NextTokenTill(':')
				t.NextChar()
				param.Type = f.ParseType(t.NextParenthesesContent('(', ')'), typeMap)
				param.Name = t.NextTokenTill(' ', ';')
				method.Params = append(method.Params, param)
				c := t.NextChar()
				if c == ';' {
					break
				}
			}
			break
		}
	}
	return method
}

func (f *Fetcher) parseProperty(sections []*PrimaryContentSection) *gen.Property {
	property := &gen.Property{}
	for _, section := range sections {
		if section.Kind == "declarations" {
			signature, typeMap := f.joinTokenText(section.Declarations[0].Tokens)
			t := &Tokenizer{strings.NewReader(signature[len("@property"):])}
			c := t.NextChar()
			t.PutBack()
			if c == '(' {
				modifiers := t.NextParenthesesContent('(', ')')
				for _, item := range strings.Split(modifiers, ",") {
					s := strings.TrimSpace(item)
					if s == "readonly" {
						property.ReadOnly = true
					} else if s == "class" {
						property.Static = true
					} else if strings.HasPrefix(s, "getter=") {
						property.GetterName = s[len("getter="):]
					}
				}
			}
			t.SkipWhiteSpaces()
			typeStr := t.NextTokenTill(' ', '<')
			c = t.NextChar()
			t.PutBack()
			if c == '<' {
				typeStr = typeStr + "<" + t.NextParenthesesContent('<', '>') + ">"
			}
			t.SkipWhiteSpaces()
			c = t.NextChar()
			if c == '*' {
				typeStr = typeStr + "*"
			} else {
				t.PutBack()
			}
			property.Type = f.ParseType(typeStr, typeMap)
			property.Name = t.NextTokenTill(';')
			t.NextChar()
			break
		}
	}
	return property
}

func (f *Fetcher) parseAliasType(sections []*PrimaryContentSection) typing.Type {
	for _, section := range sections {
		if section.Kind == "declarations" {
			signature, typeMap := f.joinTokenText(section.Declarations[0].Tokens)
			typeStr := strings.TrimPrefix(signature, "typedef ")
			typeStr = stringx.SubstringBeforeLast(typeStr, " ")
			return f.ParseType(typeStr, typeMap)
		}
	}
	panic("parse alias type failed")
}

func (f *Fetcher) parseClassType(sections []*PrimaryContentSection, module *typing.Module) typing.Type {
	for _, section := range sections {
		if section.Kind == "declarations" {
			signature, _ := f.joinTokenText(section.Declarations[0].Tokens)
			typeStr := strings.TrimPrefix(signature, "@interface ")
			typeStr = stringx.SubstringBeforeLast(typeStr, " ")
			return &typing.ClassType{Name: typeStr, Module: module}
		}
	}
	panic("parse alias type failed")
}

func (f *Fetcher) joinTokenText(tokens []*DeclarationToken) (string, map[string]*DeclarationToken) {
	var sb strings.Builder
	typeMap := map[string]*DeclarationToken{}
	for _, token := range tokens {
		sb.WriteString(token.Text)
		if token.Kind == "typeIdentifier" {
			typeMap[token.Text] = token
		}
	}
	return sb.String(), typeMap
}

func (f *Fetcher) fetchResponse(identifier string) (*Response, error) {
	cachePathItems := []string{"response_cache"}
	id := f.idFromIdentifier(identifier)
	items := strings.Split(id, "_")
	cacheDir := path.Join(append(cachePathItems, items[:len(items)-1]...)...)
	_ = os.MkdirAll(cacheDir, os.ModePerm)
	cacheFile := path.Join(cacheDir, items[len(items)-1]+".json")
	body, err := filex.ReadAllToString(cacheFile)
	var response = &Response{}

	if err != nil || body == "" {
		url := f.urlForIdentifier(identifier)
		_, body, err = f.fetchUrl(url)
		if err != nil && body != "" {
			return nil, err
		}
		err = filex.WriteString(cacheFile, body)
		if err != nil {
			fmt.Println("write cache file error:", err)
		}
		err = json.Unmarshal([]byte(body), response)
		if err != nil {
			return nil, err
		}
	} else {
		err = json.Unmarshal([]byte(body), response)
		if err != nil {
			return nil, err
		}
	}

	return response, err
}

func (f *Fetcher) fetchUrl(url string) (*httpx.ResponseHeader, string, error) {
	return f.client.Get(url).ReadAllString()
}

func (f *Fetcher) idFromIdentifier(identifier string) string {
	id := strings.Replace(identifier, "doc://com.apple.documentation/documentation/", "", 1)
	id = strings.ReplaceAll(id, "/", "_")
	return id
}

func (f *Fetcher) identifierForType(module string, name string) string {
	return fmt.Sprintf("doc://com.apple.documentation/documentation/%v/%v", strings.ToLower(module), strings.ToLower(name))
}

func (f *Fetcher) urlForIdentifier(identifier string) string {
	return strings.Replace(identifier, "doc://com.apple.documentation/", "https://developer.apple.com/tutorials/data/", 1) + ".json?language=objc"
}

func (f *Fetcher) moduleForIdentifier(identifier string) typing.Module {
	s := stringx.SubstringBeforeLast(identifier, "/")
	s = stringx.SubstringAfterLast(s, "/")
	return typing.FindModule(s)
}

// ParseType parse objc type declaration to Type instance
func (f *Fetcher) ParseType(typeStr string, typeMap map[string]*DeclarationToken) typing.Type {
	switch typeStr {
	case "void":
		return &typing.VoidType{}
	case "BOOL":
		return typing.Bool
	case "NSInteger":
		return typing.Int
	case "NSUInteger":
		return typing.UInt
	case "NSString *":
		fallthrough
	case "NSString*":
		return &typing.StringType{}
	case "NSData *":
		fallthrough
	case "NSData*":
		return &typing.DataType{}
	case "id":
		return &typing.IDType{}
	case "instancetype":
		return &typing.InstanceType{}
	default:
		if strings.HasPrefix(typeStr, "NSArray") {
			elementTypeStr := stringx.SubstringAfter(typeStr, "<")
			elementTypeStr = stringx.SubstringBefore(elementTypeStr, ">")
			elementType := f.ParseType(elementTypeStr, typeMap)
			return &typing.ArrayType{Type: elementType}
		}
		if strings.HasPrefix(typeStr, "NSDictionary") {
			elementTypesStr := stringx.SubstringAfter(typeStr, "<")
			elementTypesStr = stringx.SubstringBefore(elementTypesStr, ">")
			items := strings.Split(elementTypesStr, ",")
			if len(items) != 2 {
				fmt.Println("Parse dict type error, type:", typeStr)
				return nil
			}
			keyTypeStr := strings.TrimSpace(items[0])
			keyType := f.ParseType(keyTypeStr, typeMap)
			valueTypeStr := strings.TrimSpace(items[1])
			valueType := f.ParseType(valueTypeStr, typeMap)
			return &typing.DictType{KeyType: keyType, ValueType: valueType}
		}
		if token, ok := typeMap[typeStr]; ok {
			module := f.moduleForIdentifier(token.Identifier)
			response, err := f.fetchResponse(token.Identifier)
			if err != nil {
				fmt.Println("parse type error:", typeStr)
				return nil
			}
			metadata := response.Metadata
			switch metadata.SymbolKind {
			case "tdef":
				// typedef
				_type := f.parseAliasType(response.PrimaryContentSections)
				return &typing.AliasType{Type: _type, Alias: metadata.Title, Module: module}
			case "cl":
				// class
				return f.parseClassType(response.PrimaryContentSections, &module)
			default:
				fmt.Println("parse type unknown kind:", metadata.SymbolKind)
				return nil
			}
		}

		fmt.Println("unknown type:", typeStr)
		return nil
	}
}

func main() {
	fetcher := NewFetcher()
	identifier := fetcher.identifierForType("Foundation", "NSURL")
	_, err := fetcher.fetchType(identifier)
	if err != nil {
		fmt.Println(err)
		return
	}
}
