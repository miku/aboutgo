# Context

> Package context defines the Context type, which carries deadlines,
> cancellation signals, and other request-scoped values across API boundaries
> and between processes.

By convention, if context is used, it should be the first parameter.

> Do not store Contexts inside a struct type; instead, pass a Context explicitly
> to each function that needs it. The Context should be the first parameter,
> typically named ctx.


```go
func DoSomething(ctx context.Context, arg Arg) error {
	// ... use ctx ...
}
```

There are four ways to create a context:

* WithValue
* WithDeadline
* WithTimeout
* WithCancel

All function return a pair of `ctx, cancelFunc`.

A [`CancelFunc`](https://golang.org/pkg/context/#CancelFunc) tells an operation to abandon its work.

The CancelFunc should stay where it was created.

## Top level context

Provided via `context.Background`, can be used to derive other contexts.

```go
ctx, cancel := context.Background()
```

## Placeholder

Use `context.TODO` as a placeholder.

```go
ctx, cancel := context.TODO()
```

## Code Review

* [Example 1](example1/main.go)
* [Example 2](example2/main.go)
* [Example 3](example3/main.go)
