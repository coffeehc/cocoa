package main

type Response struct {
	Kind                   string                   `json:"kind"`
	Metadata               *Metadata                `json:"metadata"`
	Abstract               []*Item                  `json:"abstract"`
	PrimaryContentSections []*PrimaryContentSection `json:"primaryContentSections"`
	TopicSections          []*TopicSection          `json:"topicSections"`
	References             map[string]*Reference    `json:"references"`
}

type Metadata struct {
	Title       string    `json:"title"`
	Role        string    `json:"role"`
	RoleHeading string    `json:"roleHeading"` // Class|Instance Method|Instance Property|Enumeration Case|Enumeration|Global Variable|intf Protocol|Type Alias
	SymbolKind  string    `json:"symbolKind"`  // cl|instm|instp|econst|enum|data|intf|tdef
	Modules     []*Module `json:"modules"`
	ExternalID  string    `json:"externalID"`
}

type Module struct {
	Name string `json:"name"`
}

type PrimaryContentSection struct {
	Kind         string                              `json:"kind"`
	Declarations []*PrimaryContentSectionDeclaration `json:"declarations"`
	Content      []*PrimaryContentSectionContent     `json:"content"`
}

type PrimaryContentSectionDeclaration struct {
	Languages []string            `json:"languages"`
	Tokens    []*DeclarationToken `json:"tokens"`
	Platforms []string            `json:"platforms"`
}
type DeclarationToken struct {
	Kind              string `json:"kind"`
	Text              string `json:"text"`
	PreciseIdentifier string `json:"preciseIdentifier"`
	Identifier        string `json:"identifier"`
}
type PrimaryContentSectionContent struct {
	Type          string  `json:"type"`
	InlineContent []*Item `json:"inlineContent"`
}

type Item struct {
	Type       string `json:"type"`
	Text       string `json:"text"`
	IsActive   bool   `json:"isActive"`
	Identifier string `json:"identifier"`
}

type TopicSection struct {
	Kind        string   `json:"kind"`
	Title       string   `json:"title"`
	Identifiers []string `json:"identifiers"`
}

type Reference struct {
	Title      string `json:"title"`
	Identifier string `json:"identifier"`
	URL        string `json:"url"`
	Type       string `json:"type"`
	Topic      string `json:"topic"`
	Kind       string `json:"kind"`
	Role       string `json:"role"`
	Deprecated bool   `json:"deprecated"`
}
