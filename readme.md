# modbus

Minimal, truth-preserving Modbus implementation.

Goals:
- Preserve device replies exactly as received
- Treat Modbus exceptions as data, not errors
- Separate protocol, transport, and semantics
- No retries, no interpretation, no hidden behavior

Non-goals:
- High-level Modbus abstractions
- Safety reinterpretation of device replies
- Feature completeness beyond real use cases

This library exists to support deterministic systems
that require full visibility into device behavior.
