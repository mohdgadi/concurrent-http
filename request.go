package concurrent

import (
	"net/http"
	"sync"
	"sync/atomic"
)

type Request struct {
	count       int64
	concurrency int
	httpRequest *http.Request
	status      int64
}

type Response struct {
	httpResponse *http.Response
	err          error
}

func (res Response) HttpResponse() *http.Response {
	return res.httpResponse
}

func (res Response) Error() error {
	return res.err
}

func NewRequest(httpRequest *http.Request, count int64, concurrency int) (req Request) {
	return Request{
		count:       count,
		concurrency: concurrency,
		httpRequest: httpRequest,
		status:      0,
	}
}

func (req Request) MakeSync() (res chan Response) {
	res = make(chan Response, req.count)
	defer close(res)

	wg := sync.WaitGroup{}
	mutex := sync.Mutex{}

	for i := 0; i < req.concurrency; i++ {
		wg.Add(1)
		go func() {
			for {
				mutex.Lock()
				if req.status >= req.count {
					mutex.Unlock()
					break
				}

				req.status++
				mutex.Unlock()
				newRes := Response{}
				newRes.httpResponse, newRes.err = http.DefaultClient.Do(req.httpRequest)
				res <- newRes
			}

			wg.Done()
			return
		}()
	}

	wg.Wait()

	return
}

func (req Request) Status() (completed float32) {
	return (float32(atomic.LoadInt64(&req.status)) / float32(req.count)) * 100
}
