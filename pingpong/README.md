Get two go routines to play tennis, - ping pong game-  challenge to understand the concept of concurrency and sharing data.


Hi everyone 🤗 , I hope you had an awesome weekend!


Be that as it may today I am going to share a piece of code though a part - it's easy to get the other part because it is almost a duplicate.


The code showcases the following 


1️⃣ How to get go routines running.


2️⃣ How to cancel go routines using context.


3️⃣ The importance of a closed channel.


4️⃣ How to check if the channel is closed


5️⃣ why channels are awesome as a means of communication for goroutines.


Anyways the code gets to play tennis for us 😎 !


Concurrency can be tricky business, coz a lot can go wrong like:


❌ closing the channel at the wrong place(or forgetting to close it all together)

❌ failing to check whether the channel is closed

❌ failing to close the go routine when done(forget to use wg.Done())

❌ forgetting wg.Wait()


these four might seem trivial, but make them and you have a deadlock!

#go #golang #golangtips #backend #concurrency #senior