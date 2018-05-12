package typemap

import "context"

type WalkFunc func(context.Context, *ObjectSource, error) error
