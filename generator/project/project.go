package project

import (
	"context"
	"path"
)

type projectKey struct{}

type Project struct {
	Path string
}

func CtxWithProject(ctx context.Context, project *Project) context.Context {
	return context.WithValue(ctx, projectKey{}, project)
}

func FromCtx(ctx context.Context) *Project {
	return ctx.Value(projectKey{}).(*Project)
}

func AbsolutePath(ctx context.Context, relativePath string) string {
	return path.Join(FromCtx(ctx).Path, relativePath)
}
