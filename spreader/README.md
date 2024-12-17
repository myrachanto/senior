Some magic of `fmt package` -what's under the hood?

hello everyone 🤗 , wow Christmas is now just around the corner

so today I am just going to highlight a few points on how fmt works under the hood

If a stringer method is attached to a type, and you intend to print it?

```bash 
    type Stringer interface { 
    String() string
    }
```

 the `fmt package` will prioritize that string method- like in the example below.

 ```bash
 func (t total) String() string {
	if t == 0 {
		return "0"
	}
	return fmt.Sprintf("The fmt will prioritize this method by printing this line == %d \n", t)
}
```

the fmt package can also take 0 arguments to well lots and still work, the sum function tries to imitate that!

and don't get me started on how it converts various data types to strings by default.

There is a lot more of what the `fmt package` is capable of!

FYI `stringer` implies that the type has a String method that returns a string 

its an interface "stringer" 




Happy holidays everyone!

hashtag#go hashtag#golang hashtag#golangtips hashtag#golangdeveloper hashtag#backend hashtag#seniorgolang