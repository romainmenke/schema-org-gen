package genphp

import (
	"context"

	"github.com/tyler-sommer/stick"
	"github.com/tyler-sommer/stick/twig"
)

type templatesKeyType string

const templatesKey = templatesKeyType("schema-org-gen/templates")

func contextWithTemplates(ctx context.Context, tmpls *stick.Env) context.Context {
	return context.WithValue(ctx, templatesKey, tmpls)
}

func templatesFromContext(ctx context.Context) *stick.Env {
	v := ctx.Value(templatesKey)
	if v == nil {
		return nil
	}

	vv, ok := v.(*stick.Env)
	if !ok {
		return nil
	}

	return vv
}

func loadTemplates() (*stick.Env, error) {
	structTypesTmpl, err := FSString(false, "/templates/structtypes.twig")
	if err != nil {
		return nil, err
	}

	utilTmpl, err := FSString(false, "/templates/util.twig")
	if err != nil {
		return nil, err
	}

	loadTmpl, err := FSString(false, "/templates/load.twig")
	if err != nil {
		return nil, err
	}

	templates := &stick.MemoryLoader{
		Templates: map[string]string{
			"/templates/structtypes.twig": structTypesTmpl,
			"/templates/util.twig":        utilTmpl,
			"/templates/load.twig":        loadTmpl,
		},
	}

	return twig.New(templates), nil
}
