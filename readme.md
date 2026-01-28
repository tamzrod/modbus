# Truth-Preserving Modbus Client (Go)

## Purpose

This library provides a **minimal, truth-preserving Modbus client** for systems that operate on **fixed polling intervals** and **time-bounded observations**.

It was created to support architectures where:

- time is part of the signal
- failures are meaningful data
- silence is observable
- retries are managed by the poller, not the protocol
- protocol truth is separated from health, policy, and interpretation

This is **not** a convenience Modbus library.

---

## This Library Is Not Anti-goburrow

This project was **not** born because existing Modbus libraries (such as goburrow) are wrong.

Libraries like goburrow are well-engineered and optimized for:
- application-driven request/response flows
- developer convenience
- interpreted results
- Go-idiomatic error handling

They are a good fit for many systems.

This library exists because **poll-based, long-running industrial systems have different architectural constraints and goals**.

---

## The Core Difference: Polling vs Requests

Most Modbus libraries assume a mental model like this:

```
call → retry → retry → success or error
```

This fits:
- user-initiated actions
- APIs
- dashboards
- short-lived workflows

This library assumes a different model:

```
poll tick → one attempt → record outcome → wait for next tick
```

In poll-based systems:
- the next poll **is already the retry**
- time separation matters
- failure frequency matters
- the first failure matters

Adding retries inside the library collapses time and hides information.

---

## Why This Library Does Not Retry

Retries are a **policy decision**, not a protocol responsibility.

In a polling system:
- retrying inside the library duplicates the retry already provided by the poll rate
- retries increase load precisely when a device is struggling
- retries distort timing guarantees
- retries hide early failure signals

This library follows a simple rule:

> **Attempt once. Record what happened. Move on.**

If there is another poll, it will happen at the correct time.

---

## Late Responses Are Dropped (By Design)

This library treats Modbus polling as a **time-sliced system**, similar to real-time media (e.g. voice communication).

In such systems:
- data belongs to a specific time window
- late data is worse than missing data
- preserving causality matters more than completeness

If a Modbus response arrives too late to belong to its poll window, it is discarded.

The next poll provides a clean, time-correct observation.

---

## What This Library Guarantees

- **One request → one response or transport failure**
- **Modbus exception codes are preserved as data**
- **Transport failures are the only Go errors**
- **No retries**
- **No healing**
- **No interpretation**
- **No state across calls**

The library exposes **exact device truth**, nothing more.

---

## What This Library Does *Not* Do

This library does **not**:
- interpret register values
- normalize errors
- retry requests
- manage timing
- determine device health
- smooth failures
- optimize for convenience

Those responsibilities belong upstream.

---

## Intended Use Cases

This library is designed for:
- SCADA and OT systems
- solar farms, substations, BESS
- large fleets of Modbus devices
- long-running, always-on pollers
- systems that derive health from observed behavior over time

If you need:
- quick reads
- automatic retries
- friendly APIs
- interpreted values

You may be better served by other libraries.

---

## A Note on Architecture

This project follows a strict separation:

| Layer | Responsibility |
|-----|-----|
| Transport | Move bytes, detect transport failure |
| Protocol | Decode frames, expose responses |
| Poller | Owns timing and frequency |
| Health logic | Derived externally |
| Policy / retries | External and explicit |

> **Truth lives at the edge. Meaning lives upstream.**

---

## Final Thought

Different libraries exist because **different systems need different trade-offs**.

If an existing Modbus library solves your problem, you should use it.

If you need time-correct, truth-preserving Modbus behavior in a polling system —  
this library exists for you.

