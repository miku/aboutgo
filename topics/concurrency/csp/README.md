# CSP

## Notes

* Communicating Sequential Processes (Hoare, 1978)
* Goroutines that communicate via channels
* Emphasis on communication, a goroutine does not address a target, but uses a channel
* *Do not communicate by sharing memory; instead, share memory by communicating.*

> One way to think about this model is to consider a typical single-threaded
> program running on one CPU. It has no need for synchronization primitives. Now
> run another such instance; it too needs no synchronization. Now let those two
> communicate; if the communication is the synchronizer, there's still no need
> for other synchronization. Unix pipelines, for example, fit this model
> perfectly. Although Go's approach to concurrency originates in Hoare's
> Communicating Sequential Processes (CSP), it can also be seen as a type-safe
> generalization of Unix pipes.

