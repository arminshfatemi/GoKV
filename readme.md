# GoKV

**GoKV** is an experimental, in-memory key-value database written in **Go**, designed to explore **storage engines**, **concurrency**, and **high-performance backend patterns**.
The project focuses on **partition-based isolation**, **low-latency operations**, and **production-oriented design tradeoffs**.

> ⚠️ GoKV is a learning and research project. Some features are still under active development.

---

## ✨ Key Concepts

### 🔹 Partitions

Data is organized into **partitions**, each with:

* Its own schema (`INT`, `STRING`)
* Independent configuration and lifecycle
* Isolated data and statistics

Partitions allow better separation of concerns, scalability, and future access control.

---

### 🔹 High-Performance Concurrency Model

* Fine-grained locking at the partition level
* **Atomic statistics** (Ops, Writes, Keys) to reduce lock contention
* Early lock release for write-heavy operations (intentional design choice)

This trades strict metric consistency for **higher throughput under load**, similar to real-world systems.

---

### 🔹 Supported Commands (Current)

#### Partition Management

```
CREATE PARTITION <name> <schema> <persist_mode>
DROP PARTITION <name>
LIST PARTITIONS
DESCRIBE PARTITION <name>
STATS PARTITION <name>
```

#### Data Operations

```
SET <partition> <key> <value>
GET <partition> <key>
DEL <partition> <key> [key...]
INCR <partition> <key>
EXISTS <partition> <key> [key...]
```

* `DEL` and `EXISTS` support **bulk operations**
* Operations are schema-aware and validated per partition

---

## 📊 Statistics

Each partition maintains atomic counters:

* Total operations
* Write operations
* Current key count

Stats are **eventually consistent by design** to minimize contention in hot paths.

---

## 🔐 Authentication & Authorization (Planned / In Progress)

GoKV is evolving toward a **production-like security model**.

### Planned features:

* User authentication (`LOGIN / AUTH`)
* Connection-bound sessions
* **Partition-level permissions**

    * `READ` (GET, EXISTS, STATS)
    * `WRITE` (SET, DEL, INCR)
    * `ADMIN` (CREATE/DROP partitions, ACL management)
* Admin-only partition creation
* Default-deny access model

This design aims to support **multi-tenant** and **sensitive-data** use cases.

---

## 🧪 Design Goals

* Understand database internals through implementation
* Explore real-world tradeoffs (consistency vs performance)
* Build systems that behave well under **high load and stress**
* Keep the codebase readable and intentional

---

## 🛠 Tech Stack

* Language: **Go**
* Concurrency: `sync.Mutex`, atomic counters
* Protocol: Custom text-based command protocol
* Architecture: Partition-based, modular executor & parser

---

## 🚧 Current Status

* Core command execution ✔
* Partition lifecycle ✔
* Bulk operations ✔
* Atomic stats ✔
* Authentication & ACL ❌ (in progress)

---

## 📌 Notes

This project is not intended to replace Redis or similar databases.
It is a hands-on exploration of **how such systems can be built** and the engineering decisions behind them.

