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
	Comment    string
	URL        string
	Name       string
	Fields     []Field
	SubClassOf []Object
}

type Typemap []Object

func (t Typemap) Len() int {
	c := 0
	for _, o := range t {
		c += len(o.Comment)
		c += len(o.URL)
		c += len(o.Name)

		for _, s := range o.SubClassOf {
			c += len(s.Comment)
			c += len(s.URL)
			c += len(s.Name)
		}

		for _, f := range o.Fields {
			c += len(f.Comment)
			c += len(f.URL)
			c += len(f.Name)

			for _, ft := range f.Types {
				c += len(ft.URL)
				c += len(ft.Type)
			}
		}
	}

	return c
}

func AddField(o Object, f Field) Object {
	for _, oField := range o.Fields {
		if oField.Name == f.Name {
			return o
		}
	}

	o.Fields = append(o.Fields, f)
	return o
}

func AddSuperClass(o Object, p string) Object {
	for _, oParent := range o.SubClassOf {
		if oParent.Name == p {
			return o
		}
	}

	po := Object{
		Name: p,
	}

	o.SubClassOf = append(o.SubClassOf, po)

	return o
}

func SuperClasses(tm Typemap, o Object) []Object {
OUTER:
	for i, super := range o.SubClassOf {
		for _, fullSuper := range tm {
			if super.Name != fullSuper.Name {
				continue
			}

			next := SuperClasses(tm, fullSuper)
			if len(next) > 0 {
				o.SubClassOf = append(o.SubClassOf, next...)
			}

			o.SubClassOf[i] = fullSuper

			continue OUTER
		}
	}

	return o.SubClassOf
}
