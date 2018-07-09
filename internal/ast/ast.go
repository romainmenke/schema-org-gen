package ast

type FieldType struct {
	URL  string
	Type string
}

type Field struct {
	Comment   string
	Types     []FieldType
	Name      string
	URL       string
	Duplicate bool
}

type Object struct {
	URL           string
	Name          string
	ParentObjects []*Object
	Fields        []Field
}

func (o *Object) AddField(f Field) {
	for _, oField := range o.Fields {
		if oField.Name == f.Name {
			return
		}
	}

	o.Fields = append(o.Fields, f)
}

func (o *Object) AddParent(p *Object) {
	for _, oParent := range o.ParentObjects {
		if oParent.Name == p.Name {
			return
		}
	}

	cp := &Object{
		URL:    p.URL,
		Name:   p.Name,
		Fields: p.Fields,
	}

	o.ParentObjects = append(o.ParentObjects, cp)
}
