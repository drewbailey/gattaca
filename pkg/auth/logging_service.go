package auth

import (
	"context"

	"github.com/go-kit/kit/log"
)

type Middleware func(Service) Service

type loggingMiddleware struct {
	logger log.Logger
	next   Service
}

func NewLoggingMiddleware(logger log.Logger) Middleware {
	return func(next Service) Service {
		return loggingMiddleware{logger, next}
	}
}

func (mw loggingMiddleware) Signup(ctx context.Context, user, pass string) error {
	defer func() {
		mw.logger.Log("method", "Signup", "user", user, "pass", pass)
	}()
	return mw.next.Signup(ctx, user, pass)
}

func (mw loggingMiddleware) Login(ctx context.Context, user, pass string) (token string, err error) {
	defer func() {
		mw.logger.Log("method", "Login", "user", user, "pass", pass)
	}()
	return mw.next.Login(ctx, user, pass)
}

func (mw loggingMiddleware) Logout(ctx context.Context, user, token string) error {
	defer func() {
		mw.logger.Log("method", "Logout", "user", user)
	}()
	return mw.next.Logout(ctx, user, token)
}

func (mw loggingMiddleware) Validate(ctx context.Context, user, token string) error {
	defer func() {
		mw.logger.Log("method", "Validate", "user", user)
	}()
	return mw.next.Validate(ctx, user, token)
}
