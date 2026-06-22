# Monotonic Stack

Use a monotonic stack when you need the next greater/smaller element.

Daily Temperatures idea:

```text
Keep indexes whose warmer day has not been found.
When a warmer temperature appears, pop colder indexes.
```

Code:

- [daily_temperatures.go](daily_temperatures.go)
