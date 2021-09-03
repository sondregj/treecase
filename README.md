# treecase

Reports whether a filetree contains files with same name, but different casing.

## Usage

```sh
$ treecase .

Conflicts found.

...
```

## Why

Having files with the same name and different casing may cause problems on case-insensitive file systems. They may also be considered ambiguous. This tool was built to detect such cases.
