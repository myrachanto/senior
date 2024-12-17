Let's get crazy and figure out the magic box called channels in Golang, and how it manages to keep communication between goroutines (which never ever interact with each other)


Hi everyone 🤗 , yeah I know it's December and we do fun not deep stuff 😎 


Be that as it may, I am just going to highlight a few things!


✳️ you should know that goroutines don’t share memory by default.


The `hchan` in the screenshot below is the structure of a channel.


let us have a look at a few fields:


1️⃣ qcount: Tracks the number of items currently in the channel buffer.


2️⃣ dataqsiz: Specifies the channel's capacity (0 for unbuffered channels).


3️⃣ buf: Points to the circular buffer used to store elements in buffered channels.


4️⃣ elemsize: Indicates the size of a single element in the channel (in bytes).

closed: Tracks whether the channel is closed. A recv on a closed channel returns the zero value.


5️⃣ recvq : Represent waiting lists of goroutines blocked on receiving. (when a goroutine tries to receive from an empty channel, it's blocked and added to  recvq until there is something to receive)


6️⃣ sendq: Represent waiting lists of goroutines blocked on sending operations.(when a goroutine tries to send to a full channel, it's blocked and added to sendq until there is space to send stuff )


7️⃣ lock: Ensures that operations on the channel are thread-safe.


🔆 You should know that the Go Scheduler pauses Go Routines when the channel is full(sending) or empty(receiving), switches to other Go Routines, and unpauses them when ready!


sendq and recvq keep track of paused go-to routine for easier resumability.


To put it mildly, channels are indeed magic:


✅ they are simple and powerful and can work with almost any data type

✅ it's a safe way for goroutines to communicate.

✅ channels are the reason for `do not communicate by sharing memory; instead, share memory by communicating`

#go #golang #golangtips #deeperdive #senior #concurrency #channels #softwareengineer