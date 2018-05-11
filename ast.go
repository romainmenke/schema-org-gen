package main

type FieldType struct {
	URL  string
	Type string
}

type Field struct {
	Comment string
	Types   []FieldType
	Name    string
	URL     string
}

type Object struct {
	URL          string
	Name         string
	ParentObject *Object
	Fields       []Field
}
