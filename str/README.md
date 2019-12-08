case
-----

### Reference

- https://github.com/iancoleman/strcase
- https://github.com/janos/casbab

### Example

```go
s := 'AnyKind of_string'
```

#### Snake style

| Function                                  | Result               |
|-------------------------------------------|----------------------|
| `Snake(s)`                                | `any_kind_of_string` |
| `CamelSnake(s)`                           | `Any_Kind_Of_String` |
| `ScreamingSnake(s)`                       | `ANY_KIND_OF_STRING` |

#### Kebab style

| Function                                  | Result               |
|-------------------------------------------|----------------------|
| `Kebab(s)`                                | `any-kind-of-string` |
| `CamelKebab(s)`                           | `Any-Kind-Of-String` |
| `ScreamingKebab(s)`                       | `ANY-KIND-OF-STRING` |

#### Camel style

| Function                                  | Result               |
|-------------------------------------------|----------------------|
| `Camel(s)` or `LowerCamel(s)`             | `anyKindOfString`    |
| `Pascal(s)` or `UpperCamel(s)`            | `AnyKindOfString`    |

#### Other style

| Function                                  | Result               |
|-------------------------------------------|----------------------|
| `Delimited(s, '.')`                       | `any.kind.of.string` |
| `Delimited(s, ':')`                       | `any:kind:of:string` |
| `Delimited(s, ',')`                       | `any,kind,of,string` |
