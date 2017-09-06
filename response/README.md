# response intercepter

## using

```go
	q := make(chan []byte)
	w := &response.Intercepter{
		Listener: q,
	}
	go SomeResponseWriterWantToIntercept(w)
	body := <-q
	fmt.Println(body)
```
