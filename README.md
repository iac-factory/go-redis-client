## Discussion 

### Handling Structured Data Unknown at Compilation

- Package `cty` (pronounced see-tie) provides some infrastructure for a type system that might be useful for applications that need to represent configuration values provided by the user whose types are not known at compile time, ***particularly if the calling application also allows such values to be used in expressions***.
  - Reading *"[...] particularly if the calling application also allows such values to be used in expressions"*, **as with
    `terraform`**, it's important to be able to evaluate and determine
    type-structure

## Reference

### Default `git` Branch Setup

```bash
git config --global init.defaultBranch Development
```