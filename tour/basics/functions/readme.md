the type comes after the variable name.

"naked" return.

Variables with initializers
A var declaration can include initializers, one per variable.

If an initializer is present, the type can be omitted; the variable will take the type of the initializer.

---
Short variable declarations
Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.

Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.

---

Constants cannot be declared using the := syntax.