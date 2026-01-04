# Keystone

Keystone ORM is a modular ORM for Go, focused on **clarity of intent**, **expressive APIs**, and **full control over the generated SQL**.

It does not try to hide the database.
It provides a consistent, predictable, and safe query language on top of SQL.

---

## Philosophy

- SQL is a feature, not an implementation detail
- APIs should express intent, not mechanics
- Relationships are explicit
- No hidden magic
- The developer stays in control

Keystone is not a framework.
It is a library that can be used standalone or integrated into any stack.

---

## Installation

```bash
go get github.com/henrique998/keystone
