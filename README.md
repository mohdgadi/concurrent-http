# concurrent-http
A golang micro-libary to make concurrent http request to your server

# Usage
1) Simply make a net/http Request.
2) Determine the number of request to be made.
3) Specify the concurrency i.e parallelism of the requests.

```golang
	url := "http://localhost:8080/"
	req, _ := http.NewRequest("GET", url, nil)
    numberOfRequests := 100000
    concurrency := 2500

    <!-- Construct the request -->
	concurrentRequest := concurrent.NewRequest(req, numberOfRequests, concurrency)

    <!-- Start sending out the request. This is blocking -->
    resChan :=  concurrentRequest.MakeSync()

    <!-- Get the status i.e percentage of request completed -->
    percentage := concurrentRequest.Status()

    <!-- Print all the http.Responses and error -->
    for res := range resChan {
        httpResponse = res.HttpRespone()
        error = res.Error()
    }
```