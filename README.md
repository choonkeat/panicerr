``` go
func main() {
	defer func() {
		if r := recover(); r != nil {
			err := panicerr.New(r)
			log.Println(err)
		}
	}()
}
```
