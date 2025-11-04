# protoc-gen-go-enums

Protobuf plugin that generates renamed Go constants from enums.

Options:

```sh
protoc --go-enums_out=. \
--go-enums_opt=paths=source_relative \
--go-enums_opt=include_nested=true 
```

Enum options:

```proto
extend google.protobuf.EnumOptions {
  // Enable generation of renamed Go constants from enums. (Default: true)
  bool gen_go_enums = 91720;
  // Strip prefix. (Default: true)
  string gen_go_enums_strip_name_prefix = 91721;
  // Turn const name into pascal case
  // (apply before  `gen_go_enums_name_prefix` and `gen_go_enums_name_suffix`).
  bool gen_go_enums_name_pascal_case = 91722;
  // Turn const name into CAPS_CASE
  // (apply before `gen_go_enums_name_prefix` and `gen_go_enums_name_suffix`).
  // Can't be used together with gen_go_enums_name_pascal_case.
  bool gen_go_enums_name_caps_case = 91723;
  // Add string prefix to generated Go string constant names.
  string gen_go_enums_name_prefix = 91724;
  // Add string suffix to generated Go string constant names.
  string gen_go_enums_name_suffix = 91725;
}
```

See `example/` for examples.
