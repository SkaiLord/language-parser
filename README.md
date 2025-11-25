# language-parser

A tiny **language parser written in Go**, built as a learning project to understand how tokenization, parsing, and simple interpretation work under the hood.

The project is structured to be easy to extend with your own grammar rules and examples, so you can experiment with building small languages, DSLs, or expression evaluators.

---

## ✨ Goals

- Learn how lexers/tokenizers work
- Implement a basic hand-written parser in Go
- Represent input as an Abstract Syntax Tree (AST)
- Optionally evaluate / pretty-print parsed expressions
- Provide small, focused examples under `examples/`

---

## 📁 Project Structure

```text
language-parser/
├── src/          # Core parser / lexer / AST code
├── examples/     # Example inputs / small demos for the parser
└── .gitignore
```
