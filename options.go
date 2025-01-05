package idphotosdk

import (
	"log/slog"
)

type Options struct {
	baseUrl string

	secretToken string

	// log
	logger     *slog.Logger
	logEnabled bool
}

type Option func(*Client)

func WithBaseUrl(s string) Option {
	return func(c *Client) {
		c.baseUrl = s
	}
}

func WithSecretToken(s string) Option {
	return func(c *Client) {
		c.secretToken = s
	}
}

func WithLogger(logger *slog.Logger) Option {
	return func(c *Client) {
		c.logger = logger
	}
}

func WithLogEnabled(logEnabled bool) Option {
	return func(c *Client) {
		c.logEnabled = logEnabled
	}
}
