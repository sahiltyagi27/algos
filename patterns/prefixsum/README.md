# Prefix Sum

Use prefix sum when you need fast range sums or subarray sum counting.

Visual:

```text
nums:   [2, 4, 6]
prefix: [0, 2, 6, 12]

sum(1,2) = prefix[3] - prefix[1] = 12 - 2 = 10
```

Code:

- [prefix_sum.go](prefix_sum.go)
