package typemap

import (
	"context"
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

	var err error

	for _, o := range v.ObjectSources {
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
