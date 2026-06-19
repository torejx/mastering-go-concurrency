# Worker Pool — Bar Exercise

## Goal

Implement a worker pool that simulates a coffee bar with **3 machines** serving a queue of orders concurrently.

Each machine must process orders one at a time and emit a log line for every order it completes:

```
-- Machine #1    15:04:05.000 - 15:04:07.500    [Order #42 ready - espresso] --
```

The log stream should be printed to stdout so it can be captured to a file and loaded into the timeline visualizer.

## Expected output format

You don't need to build the log lines yourself — they are the return value of the `Brew()` method already provided. Just print whatever `Brew()` returns and the format will be correct:

```go
log := machine.Brew(order)
fmt.Println(log)
```

Each line looks like this:

```
-- Machine #N    HH:MM:SS.mmm - HH:MM:SS.mmm    [Order #ID ready - description] --
```

| Field | Description |
|-------|-------------|
| `N` | machine number (1–3) |
| `HH:MM:SS.mmm` | wall-clock start and end timestamps |
| `ID` | order number |
| `description` | drink name or any label |

## Saving the output

Run the program and redirect stdout to a file:

```bash
go run main.go > output.log
```
## Visualizing the timeline

Open `coffee-timeline.html` in a browser and drag-and-drop (or click to select) the `output.log` file.

The visualizer will render a Gantt-style timeline for each machine and highlight two classes of concurrency problems:

- **Concurrent executions** — a machine starts a new order before the previous one finishes (its bar segment is outlined in red).
- **Duplicate orders** — the same order ID is fulfilled by more than one machine (bar segments are highlighted in orange).

A correct implementation produces a green status: no overlaps, every order served exactly once.

NB: avoid debug/info logs or remove them from `output.log` file before uploading it.


## Constraints

- Exactly **3 coffee machines** (workers).
- Each machine handles **one order at a time**.
- All orders must be processed before the program exits.
