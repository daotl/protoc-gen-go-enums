//go:generate protoc -I . -I ../../include --go-enums_opt=include_nested=true --go-enums_out=. --go-enums_opt=paths=source_relative --go_out=. --go_opt=paths=source_relative example.proto

package uint32
