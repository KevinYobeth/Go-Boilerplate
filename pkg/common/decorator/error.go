package decorator

import "context"

type commandErrorDecorator[C any] struct {
	base CommandHandler[C]
}

func (d commandErrorDecorator[C]) Handle(ctx context.Context, cmd C) error {
	err := d.base.Handle(ctx, cmd)
	if err != nil {
		return err
	}

	return nil
}

type queryErrorDecorator[Q, R any] struct {
	base QueryHandler[Q, R]
}

func (d queryErrorDecorator[Q, R]) Handle(ctx context.Context, query Q) (R, error) {
	result, err := d.base.Handle(ctx, query)
	if err == nil {
		return result, nil
	}

	return result, err
}
