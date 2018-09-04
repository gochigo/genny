package attrs

import "github.com/gobuffalo/flect/name"

type Attr struct {
	Original   string
	Name       name.Ident
	commonType string
	goType     string
}

func (a Attr) String() string {
	return a.Original
}

func (a Attr) GoType() string {
	if a.goType != "" {
		return a.goType
	}
	switch a.commonType {
	case "text":
		return "string"
	case "timestamp", "datetime", "date", "time":
		return "time.Time"
	}
	return a.commonType
}

func (a Attr) CommonType() string {
	if a.commonType != "" {
		return a.commonType
	}
	return a.commonType
}

type Attrs []Attr
