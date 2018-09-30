package dal

import (
	"context"
	"database/sql"
)

const (
	ctxKeyPreparer = "preparer"
	//ctxKeyOperator = "operator"
)

type Preparer interface {
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
}

//func WithOperator(parent context.Context, op string) context.Context {
//	return context.WithValue(parent, ctxKeyOperator, op)
//}

func WithPreparer(parent context.Context, p Preparer) context.Context {
	return context.WithValue(parent, ctxKeyPreparer, p)
}

func prepare(ctx context.Context) Preparer {
	v := ctx.Value(ctxKeyPreparer)
	if v == nil {
		return nil
	}
	return v.(Preparer)
}
//
//func operator(ctx context.Context) string {
//	v := ctx.Value(ctxKeyOperator)
//	if v == nil {
//		return ""
//	}
//	return v.(string)
//}