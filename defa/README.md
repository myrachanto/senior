Be very careful about how you use deferred functions with anonymous functions!


Hi everyone 🤗 , I hope you had a fabulous weekend.


Well, I was very unlucky to learn this lesson via interview which contributed to my failure.


Be that as it may, I choose to either win or learn better yet both. oooh losing sucks big time!


But, what defines winners is "they are losers who never quit"

```bash
func checker() {
	var num int
	defer fmt.Println(num) //prints second output 0 (takes num as it is at this time 0)
	num++
	defer func() {
		num++
		fmt.Println(num) //prints second output 3(num has been incremented thrice)
	}()
	defer func() {
		num++
		fmt.Println(num) //prints first output 2(num has been incremented twice)
	}()
}

```

So, write in the comments what you think the results will be!

#go #golang #golangips #backend #softwareengineer