# Mastering Go Concurrency 🐹

A hands-on workshop that takes you from a naive, broken implementation all the way to a clean worker pool — one concurrency primitive at a time.

## What you'll learn

- The difference between parallelism and concurrency
- Goroutines, Mutex, WaitGroups, Channels, Select
- How to avoid race conditions, deadlocks, and goroutine leaks
- The worker pool pattern

## Prerequisites

- Latest version of Go
- git
- A terminal and your favourite editor/IDE

## Repository structure

Each directory contains a self-contained, runnable example. The numbering follows the order of the workshop.

```
code/
  bar/                  # shared domain package (coffee bar)
  001-race-condition/   # the starting point — things go wrong
  002-mutex/            # protecting shared memory
  003-waitgroup/        # waiting for goroutines to finish
  004-channels/         # communicating between goroutines
  005-select/           # handling multiple channels
  006-worker-pool/      # putting it all together
```

## Running an example

```bash
cd code/001-race-condition
go run main.go

# to detect race conditions
go run -race main.go
```

## ⚠️ Disclaimer

The code in this repository is intentionally written for **didactic purposes** and is not meant to be used in production.

In particular:
- Some examples are **deliberately broken** — they contain race conditions, deadlocks, or goroutine leaks. This is by design: they serve as a starting point to understand the problem before introducing the solution.
- The code is **not optimised** and may not follow Go best practices outside of the concurrency topics covered in the workshop.

If you find yourself thinking *"I would never write this in production"* — you're probably right, and that's the point.
