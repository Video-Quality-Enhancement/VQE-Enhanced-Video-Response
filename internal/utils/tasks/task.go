package tasks

type HandlerFunc func(ctx *Context)

type Task interface {
	Perform(handlers ...HandlerFunc)
}

type task struct {
	ctx *Context
}

func NewTask() Task {
	return &task{
		ctx: &Context{
			index: -1,
		},
	}
}

func (t *task) Perform(handlers ...HandlerFunc) {
	t.ctx.handlers = handlers
	t.ctx.Next()
}

type Context struct {
	index    int8
	status   int
	err      error
	keys     map[string]any
	handlers []HandlerFunc
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
}

func (c *Context) Error(err error) {
	c.status = 500
	c.err = err
}

func (c *Context) AbortWithError(err error) {
	c.status = 500
	c.err = err
	c.index = int8(len(c.handlers))
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
