# Timeline

Track events in yaml, output them as a timeline in your terminal.

```yaml
# critters.yaml

- date: 2024-06-03
  type: bird
  string: robin

- date: 2024-07-01
  type: snails
  number: 6

- date: 2024-08-11
  type: bird
  string: cardinal

- date: 2024-09-22
  type: snails
  number: 20

- date: 2025-04-20
  type: bird
  string: goldfinch
```

```
▸ go run . critters.yaml

bird

┌ 2024-06-03; 2 months
│ ┌ 2024-08-11; 8 months
│ │       ┌ 2025-04-20; 4 months
│ │       │   ┌ 2025-08-30
│-│-------│---│
│ │       │   └ now
│ │       └ goldfinch
│ └ cardinal
└ robin

snails

┌ 2024-07-01; 2 months
│ ┌ 2024-09-22; 11 months; 14.00; 233.3%
│ │          ┌ 2025-08-30
│-│----------│
│ │          └ now
│ └ 20
└ 6
```

We saw a robin, 2 months later a cardinal, 8 months later a goldfinch. 4 months
have passed since we last saw a bird.

We saw 6 snails, 2 months later we saw 20 snails. That's 14 more snails than
last time. A 233.3% increase in snails! That was 11 months ago.
