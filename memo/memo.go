package memo

/*
@Time : 2024/7/5 15:51
@Author : echo
@File : memo
@Software: GoLand
@Description:
*/
//type Memo struct {
//	f     Func
//	cache map[string]result
//}
type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

//	func New(f Func) *Memo {
//		return &Memo{f: f, cache: make(map[string]result)}
//	}
//
//	func (memo *Memo) Get(key string) (interface{}, error) {
//		res, ok := memo.cache[key]
//		if ok {
//			res.value, res.err = memo.f(key)
//			memo.cache[key] = res
//		}
//		return res.value, res.err
//	}
type entry struct {
	res   result
	ready chan struct{} // closed when res is ready
}

/*func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}
*/
//type Memo struct {
//	f     Func
//	mu    sync.Mutex // guards cache
//	cache map[string]*entry
//}

/*
	func (memo *Memo) Get(key string) (value interface{}, err error) {
		memo.mu.Lock()
		e := memo.cache[key]
		if e == nil {
			// This is the first request for this key.
			// This goroutine becomes responsible for computing
			// the value and broadcasting the ready condition.
			e = &entry{ready: make(chan struct{})}
			memo.cache[key] = e
			memo.mu.Unlock()

			e.res.value, e.res.err = memo.f(key)

			close(e.ready) // broadcast ready condition
		} else {
			// This is a repeat request for this key.
			memo.mu.Unlock()

			<-e.ready // wait for ready condition
		}
		return e.res.value, e.res.err
	}
*/
type request struct {
	key      string
	response chan<- result // the client wants a single result
}

type Memo struct{ requests chan request }

// New returns a memoization of f.  Clients must subsequently call Close.
func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request)}
	go memo.server(f)
	return memo
}

func (memo *Memo) Get(key string) (interface{}, error) {
	response := make(chan result)
	memo.requests <- request{key, response}
	res := <-response
	return res.value, res.err
}

func (memo *Memo) Close() { close(memo.requests) }
func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry)
	for req := range memo.requests {
		e := cache[req.key]
		if e == nil {
			// This is the first request for this key.
			e = &entry{ready: make(chan struct{})}
			cache[req.key] = e
			go e.call(f, req.key) // call f(key)
		}
		go e.deliver(req.response)
	}
}

func (e *entry) call(f Func, key string) {
	// Evaluate the function.
	e.res.value, e.res.err = f(key)
	// Broadcast the ready condition.
	close(e.ready)
}

func (e *entry) deliver(response chan<- result) {
	// Wait for the ready condition.
	<-e.ready
	// Send the result to the client.
	response <- e.res
}
