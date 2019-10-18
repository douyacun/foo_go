package cache

type request struct {
	key      string
	response chan<- result
	done     chan struct{}
}

type result struct {
	value interface{}
	err   error
}

type Memo struct {
	requests chan request
}

type entry struct {
	ready chan struct{}
	res   result
}

type Func func(key string) (interface{}, error)

func New(f Func) *Memo {
	meme := &Memo{requests: make(chan request),}
	go meme.server(f)
	return meme
}

func (m *Memo) server(f Func) {
	cache := make(map[string]*entry)
	for req := range m.requests {
		e := cache[req.key]
		if e == nil {
			e = &entry{ready: make(chan struct{}),}
			cache[req.key] = e
			go e.call(f, req.key, req.done)
		}
		go e.deliver(req.response)
	}
}

func (e *entry) call(f Func, key string, done chan struct{}) {
	e.res.value, e.res.err = f(key)
	close(e.ready)
}

func (e *entry) deliver(response chan<- result) {
	<-e.ready
	response <- e.res
}

func (m *Memo) Get(key string, done chan struct{}) (interface{}, error) {
	response := make(chan result)
	m.requests <- request{key: key, response: response, done: done}
	res := <-response
	return res.value, res.err
}

func (m *Memo) Close() {
	close(m.requests)
}
