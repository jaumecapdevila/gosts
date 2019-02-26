package logger

// Context store all the fields that will be logged with the message
type Context map[string]interface{}

// Add a new element to our current context
func (me *Context) Add(key string, value interface{}) *Context {
	(*me)[key] = value

	return me
}

// Append merges our current context with another context
func (me *Context) Append(context *Context) *Context {
	if len(*context) == 0 {
		return me
	}

	for key, value := range *context {
		me.Add(key, value)
	}

	return me
}
