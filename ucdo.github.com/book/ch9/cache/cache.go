package main

type entry struct {
	res   result
	ready chan struct{}
}

type Memo struct {
	requests chan request
}

type result struct {
	value interface{}
	err   error
}

type request struct {
	key      string
	response chan<- result
}

type Func func(key string) (interface{}, error)

func (m *Memo) New(f Func) *Memo {
	mono := &Memo{requests: make(chan request)}
	go m.server(f)
	return mono
}

func (m *Memo) Get(key string) (interface{}, error) {
	response := make(chan result)
	m.requests <- request{key, response}
	res := <-response
	return res.value, res.err
}

func (m *Memo) Close() {
	close(m.requests)
}

// 不要用同步的思维看并发的代码
func (m *Memo) server(f Func) {
	cache := make(map[string]*entry)

	for req := range m.requests {
		e := cache[req.key]
		if e == nil {
			e = &entry{ready: make(chan struct{})}
			cache[req.key] = e
			go e.call(f, req.key) // 获取数据，注意这里是goroutine
		}

		go e.deliver(req.response)
	}
}

// 因为这里要把数据写回到调用里面，所以是 *entry
func (e *entry) call(f Func, key string) {
	e.res.value, e.res.err = f(key)
	close(e.ready)
}

func (e *entry) deliver(resp chan<- result) {
	<-e.ready //  为什么这里会被阻塞？ 结合上面的代码来看，因为ready是个无buffer的channel。读空阻塞，call写入之后就可以了
	resp <- e.res
}

func main() {}
