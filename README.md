# Protoc Go Hello World Plugin

> A plugin is just a program that reads a CodeGeneratorRequest from stdin and writes a CodeGeneratorResponse to stdout.

## Generate Go code from a proto file with existing plugin `protoc-gen-go`

1. Prepare `example/example.proto`
1. Run `protoc` with `--<plugin>_out`

    ```
    protoc --go_out=. --go_opt=paths=source_relative -I . example.proto
    ```

    this commands generates Go codes in `example/example.pb.go` file.

## Create own plugin `protoc-gen-go-hello-world`

1. Init module

    ```
    go mod init protoc-sample-plugin
    go mod tidy
    ```

1. Write `cmd/protoc-gen-go-hello-world/main.go`

    ```go
    package main

    import (
        "google.golang.org/protobuf/compiler/protogen"
    )

    func main() {

        protogen.Options{}.Run(func(gen *protogen.Plugin) error {
            for _, f := range gen.Files {
                if !f.Generate {
                    continue
                }
                generateFile(gen, f)
            }
            return nil
        })
    }

    // generateFile generates a _ascii.pb.go file containing gRPC service definitions.
    func generateFile(gen *protogen.Plugin, file *protogen.File) *protogen.GeneratedFile {
        filename := file.GeneratedFilenamePrefix + "_hello_world.pb.go"
        g := gen.NewGeneratedFile(filename, file.GoImportPath)
        g.P("// Code generated by protoc-gen-go-hello-world. DO NOT EDIT.")
        g.P()
        g.P("package ", file.GoPackageName)
        g.P()
        for _, msg := range file.Messages {
            g.P("func (x *", msg.GoIdent, ") Hello() string {")
            g.P("return \"Hello, World!\"")
            g.P("}")
        }

        return g
    }
    ```

1. Install the package.

    ```
    go install ./cmd/protoc-gen-go-hello-world
    ```

1. Run `protoc`.

    ```
    protoc --go_out=. --go_opt=paths=source_relative \
        --go-hello-world_out=. --go-hello-world_opt=paths=source_relative \
        example/example.proto
    ```

    `example/example_hello_world.pb.go` will be generated.

1. You can use the generated code in `main.go`

    ```
    go run main.go
    Hello, World!
    ```

## References
1. https://rotemtam.com/2021/03/22/creating-a-protoc-plugin-to-gen-go-code/
1. https://pkg.go.dev/google.golang.org/protobuf/compiler/protogen
