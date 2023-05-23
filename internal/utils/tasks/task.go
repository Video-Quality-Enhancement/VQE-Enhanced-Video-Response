package tasks

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"golang.org/x/exp/slog"
)

type HandlerFunc func(ctx *Context)

type Task interface {
	Activities(handlers ...HandlerFunc)
	Perform(delivery amqp.Delivery)
}

type task struct {
	ctx *Context
}

func NewTask() Task {
	return &task{
		ctx: &Context{
			index:  -1,
			status: 0,
			err:    nil,
			keys:   make(map[string]any),
		},
	}
}

func (t *task) Activities(handlers ...HandlerFunc) {
	t.ctx.handlers = handlers
}

func (t *task) Perform(delivery amqp.Delivery) {
	t.ctx.Delivery = delivery
	t.ctx.reset()
	t.ctx.Next()
}

type Context struct {
	index    int8
	status   int
	err      error
	keys     map[string]any
	Delivery amqp.Delivery
	handlers []HandlerFunc
}

func (c *Context) reset() {
	c.index = -1
	c.status = 0
	c.err = nil
	c.keys = make(map[string]any)
}

func (c *Context) Next() {

	c.index++
	for c.index < int8(len(c.handlers)) {
		c.handlers[c.index](c)
		c.index++
	}

}

func (c *Context) Status() int {
	return c.status
}

func (c *Context) Success() {
	c.status = 200
	err := c.Delivery.Ack(false)
	if err != nil {
		slog.Error("Failed to ack message", "err", err)
	} else {
		slog.Debug("Acked message", "requestId", c.Get("X-Request-ID").(string), "userId", c.Get("X-User-ID").(string))
	}
}

func (c *Context) Failure(err error) {

	c.status = 500
	c.err = err
	c.index = int8(len(c.handlers))
	err = c.Delivery.Nack(false, true)
	if err != nil {
		slog.Error("Failed to nack message", "err", err)
	} else {
		slog.Error("Nacked message", "requestId", c.Get("X-Request-ID").(string), "userId", c.Get("X-User-ID").(string))
	}

}

func (c *Context) ErrorMessage() string {
	return c.err.Error()
}

func (c *Context) Set(key string, value any) {
	c.keys[key] = value
}

func (c *Context) Get(key string) any {
	if value, ok := c.keys[key]; ok {
		return value
	}
	return nil
}
