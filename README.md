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
  bar/                            # shared domain package (coffee bar)
  000-start/                      # sequential baseline — no concurrency yet
  150-goroutines/                 # launching a single goroutine
  152-goroutines-multiple/        # launching multiple goroutines
  155-goroutines-race/            # race condition — things go wrong
  160-mutex/                      # protecting shared memory with Mutex
  170-wait-group/                 # waiting for goroutines with WaitGroup
  171-wait-group-deadlock/        # WaitGroup misuse leading to deadlock
  180-channels/                   # basic unbuffered channels
  182-buffered-channels/          # buffered channels
  183-buffered-channels-range/    # iterating over a channel with range
  184-channels-leak/              # goroutine leak via blocked channel
  186-channels-go-routine/        # channels combined with goroutines
  187-prefer-unbuffered/          # why unbuffered channels are preferred
  200-worker-pool-bar-skeleton/   # worker pool skeleton (exercise)
  201-worker-pool-bar-complete/   # worker pool — complete solution
  300-select-first-response-wins/ # select: first channel to respond wins
  302-select-timeout/             # select: implementing a timeout
  304-select-default/             # select: non-blocking default case
```

## Running an example

```bash
cd code/000-start
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
