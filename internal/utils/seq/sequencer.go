package seq

type HandlerFunc func(ctx *Context)

type Sequencer interface {
	Sequence(handlers ...HandlerFunc)
}

type sequencer struct {
	ctx *Context
}

func NewSequencer() Sequencer {
	return &sequencer{
		ctx: &Context{
			index: -1,
		},
	}
}

func (seq *sequencer) Sequence(handlers ...HandlerFunc) {
	seq.ctx.handlers = handlers
	seq.ctx.Next()
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
