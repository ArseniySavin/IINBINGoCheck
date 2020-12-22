This solution may checking our individual tax number, 
using official government algorithm. 
If number was checked true. We will be getting true value.


### How to use
```go
	chek, _ := iinbincheker.NewValidator("iin or bin")
	
	// validation
	res := chek.DirectValidating()
	
	// get tax number type
	_, doc := chek.Identify()
	iin, _ := doc.(*iinbindata.Iin)
    
	// get tax number data
	iinData := iin.Data()
```

### Examples
- [Iin](examples/iin.go)
- [Bin](examples/bin.go)
