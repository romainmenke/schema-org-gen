package typemap

import (
	"context"

	"github.com/romainmenke/schema-org-gen/internal/ast"
)

type ObjectSource struct {
	URL      string
	Object   *ast.Object
	Children []*ObjectSource
	Parent   *ObjectSource `json:"-"`
}

func (v *ObjectSource) walk(ctx context.Context, walker WalkFunc) error {
	if v == nil {
		return nil
	}

	var err error

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		//
	}

	err = walker(ctx, v, err)
	if err != nil {
		return err
	}

	if len(v.Children) == 0 {
		return nil
	}

	for _, o := range v.Children {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			//
		}

		err = o.walk(ctx, walker)
		if err != nil {
			return err
		}
	}

	return err
}
