# GoKV

**GoKV** is a lightweight, Redis-inspired key-value database written in **Go**, designed for learning, experimentation, and building a clean storage engine from scratch.

The goal of GoKV is **clarity over magic**:

* Simple internal architecture
* Explicit data structures
* Clear separation between protocol, execution, storage, and networking

This is **not a Redis clone**, but it borrows proven ideas (RESP protocol, partitions, command execution pipeline).

---

## ✨ Features (Current)

* TCP server (Redis-like interface)
* Command parsing layer
* Execution layer with command dispatch table
* In-memory partitions
* Partition registry (Create / Drop / List)
* Typed partitions (INT, STRING)
* Thread-safe operations
* Redis-compatible RESP responses

---

## 🧱 Architecture Overview

```
┌──────────┐
│  Client  │
└────┬─────┘
     │ TCP
┌────▼─────┐
│  Server  │   (TCP, bufio)
└────┬─────┘
     │
┌────▼─────┐
│ Protocol │   (Parse commands)
└────┬─────┘
     │
┌────▼─────┐
│ Executor │   (Command → Handler)
└────┬─────┘
     │
┌────▼─────┐
│Partition │   (In-memory storage)
└────┬─────┘
     │
┌────▼─────┐
│  Writer  │   (RESP output)
└──────────┘
```

Each layer has **one job** and no hidden side effects.

---

## 🧩 Partitions

A **partition** is an isolated key-value namespace with:

* Name
* Schema (value type)
* Persistence mode (future: WAL, snapshots)
* Internal locking
* Statistics tracking

### Supported schemas

* `INT`
* `STRING`

Each partition enforces **one value type**, keeping storage fast and predictable.

---

## 📜 Supported Commands

### Partition Management

```text
CREATE PARTITION <name> <TYPE> <PERSIST>
DROP PARTITION <name>
LIST PARTITIONS
```

Examples:

```text
CREATE PARTITION users INT NONE
LIST PARTITIONS
```

RESP output (example):

```
*1
$5
users
```

---

## 🧪 Running the Server

```bash
go run cmd/server/main.go
```

Server listens on:

```
:6379
```

Test with `nc`:

```bash
nc localhost 6379
```

---

## 🧠 Design Philosophy

* Prefer **explicit types** over `any`
* Fail early on invalid commands
* Keep execution logic out of the network layer
* Make every component testable in isolation
* Avoid premature optimization — correctness first

---

## 🚧 Roadmap

Planned features:

* [ ] SET / GET / DEL commands
* [ ] Single-map storage model
* [ ] Write-Ahead Log (WAL)
* [ ] Snapshots
* [ ] TTL support
* [ ] Binary protocol parsing
* [ ] Benchmarks
* [ ] Persistence recovery

---

## ⚠️ Disclaimer

GoKV is a **learning project**, not production-ready software.
Expect breaking changes as the internals evolve.

---

## 🤝 Contributing

This project is currently experimental and evolving fast.
Ideas, discussions, and feedback are welcome.

---

## 📝 License

MIT


