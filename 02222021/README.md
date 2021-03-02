# Go Proverbs

1. Don't communicate by sharing memory, share memory by communicating.
- Share memory by communicating: passing to a channel the addresses of objects. If you pass to a channel, keep the pointer otherwise you lose access. We are sharing memory not to other programs, but to corroutines. When we communicate we share a space of memory, which is a channel.

2. Concurrency is not parellelism.
- Concurrency is a structure of your program to make it scalable: to support multiple clients for example.
- Parallelism is the execution of multiple goroutines in paralel.

3. Channels orchestrate; mutexes serialize.
- Channels are the mechanism to organize the communication.
- Mutexes locking and unlocking resources. They serialize the execution: they either give access or deny it, simply.

## Kill Process in Linux

In linux, run this so that you get the pid of the process that is occupying that port (9090). <br />

```bash
ss -tunlp | grep 9090
```