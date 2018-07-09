package typemap

import (
	"context"
	"sort"
	"time"
)

type TypeMap struct {
	UpdatedAt     time.Time
	ObjectSources []*ObjectSource
}

func (v *TypeMap) Walk(ctx context.Context, walker WalkFunc) error {
	if v == nil {
		return nil
	}

	if len(v.ObjectSources) == 0 {
		return nil
	}

	for _, o := range v.ObjectSources {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			//
		}

		err := o.walk(ctx, walker)
		if err != nil {
			return err
		}
	}

	return nil
}

func (v *TypeMap) FlattenTypes(ctx context.Context) error {
	if v == nil {
		return nil
	}

	if len(v.ObjectSources) == 0 {
		return nil
	}

	typesList := []*ObjectSource{}

	for _, o := range v.ObjectSources {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			//
		}

		err := o.walk(ctx, gatherTypes(&typesList))
		if err != nil {
			return err
		}
	}

	sort.Slice(typesList, func(i int, j int) bool {
		if (typesList[i].Object == nil || typesList[i].Object.Name == "") && (typesList[j].Object == nil || typesList[i].Object.Name == "") {
			return false
		}

		if typesList[i].Object == nil || typesList[i].Object.Name == "" {
			return true
		}

		if typesList[j].Object == nil || typesList[i].Object.Name == "" {
			return true
		}

		return typesList[i].Object.Name < typesList[j].Object.Name
	})

	uniqueTypeList := make(map[string]*ObjectSource)
	for _, t := range typesList {
		if t.Object == nil || t.Object.Name == "" {
			continue
		}

		if u, ok := uniqueTypeList[t.Object.Name]; ok {
			if u.Parent != nil && u.Parent.Object != nil {

				currentParent := t.Parent
				for currentParent != nil && currentParent.Object != nil {
					u.Object.AddParent(currentParent.Object)
					currentParent = currentParent.Parent
				}
			}
		} else {
			if t.Parent != nil && t.Parent.Object != nil {

				currentParent := t.Parent
				for currentParent != nil && currentParent.Object != nil {
					t.Object.AddParent(currentParent.Object)
					currentParent = currentParent.Parent
				}
			}
			uniqueTypeList[t.Object.Name] = t
		}
	}

	v.ObjectSources = v.ObjectSources[:0]
	for _, u := range uniqueTypeList {
		v.ObjectSources = append(v.ObjectSources, u)
	}

	return nil
}

func gatherTypes(typesP *[]*ObjectSource) func(ctx context.Context, o *ObjectSource, err error) error {
	return func(ctx context.Context, o *ObjectSource, err error) error {
		if err != nil {
			return err
		}
		if o == nil || o.Object == nil || typesP == nil {
			return nil
		}

		types := *typesP
		*typesP = append(types, o)

		return nil
	}
}
