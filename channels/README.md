When to use buffered and unbuffered channels in Golang


Hi everyone 🤗 , counting the days to Christmas and to prosperous ano nuevo!


🔆 unbuffered channel

An unbuffered channel has no capacity, requiring both the sender and receiver to be ready simultaneously for data to pass.

❗ Sending or receiving on an unbuffered channel blocks until the other side is ready.


vs


🔆 buffered channel:

A buffered channel has a capacity greater than zero, allowing it to store multiple values without requiring simultaneous sender and receiver.

❗ Sending to a buffered channel blocks only if the channel is full, and receiving blocks only if the channel is empty.


when to use unbuffered channels:


1️⃣ Communication requires strict synchronization

2️⃣ The timing of message passing is crucial


and buffered channels?


1️⃣ data needs a temporary storage

2️⃣ unbuffered channels can reduce performance

3️⃣ Avoid leaking goroutines (from abandoned channels)


special use case of buffered channels is - semaphores


below is a table prepared by a great Golang teacher Matt Holiday(YouTube).


✅ He showcases the various channel states and how they behave when you try to send or receive data from them or close them.

#go #golang #golangtips #backend #seniorgolang #sofwareengineer #channels