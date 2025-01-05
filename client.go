package idphotosdk

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"net/url"
)

type Requester interface {
	Request(ctx context.Context) (*http.Request, error)
}

type Responder interface {
}

type baseResponse struct {
	Status bool `json:"status"`
}

type Client struct {
	Options

	respBodyRaw string
}

func NewClient(options ...Option) *Client {
	c := new(Client)
	c.logger = slog.Default()
	c.logEnabled = false
	for _, option := range options {
		option(c)
	}
	return c
}

func (c *Client) Do(ctx context.Context, request Requester, responder Responder) error {
	err := c.do(ctx, request, responder)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) do(ctx context.Context, request Requester, response Responder) error {
	req, err := request.Request(ctx)
	if err != nil {
		return err
	}

	baseUrl, err := url.Parse(c.baseUrl)
	if err != nil {
		if c.logEnabled {
			c.logger.InfoContext(ctx, "invalid baseUrl",
				slog.Any("error", err),
			)
		}
		return errors.New("invalid baseUrl")
	}
	req.URL = baseUrl.ResolveReference(req.URL)

	if c.secretToken != "" {
		req.Header.Set("Authorization", c.secretToken)
	}

	if c.logEnabled {
		c.logger.InfoContext(ctx, "starting request",
			slog.String("method", req.Method),
			slog.String("url", req.URL.String()),
			//slog.Any("headers", req.Header),
		)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var buf bytes.Buffer
	//io.TeeReader(resp.Body, &buf)
	_, err = io.Copy(&buf, resp.Body)
	if err != nil {
		return err
	}

	// response raw text
	c.respBodyRaw = buf.String()
	if c.logEnabled {
		c.logger.Info("response",
			slog.Int("status code", resp.StatusCode),
			slog.String("body raw", c.respBodyRaw),
		)
	}

	// response json to struct
	err = json.NewDecoder(&buf).Decode(response)
	if err != nil {
		return err
	}

	return nil
}
