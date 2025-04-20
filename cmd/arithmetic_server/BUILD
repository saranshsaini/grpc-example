load("@rules_go//go:def.bzl", "go_binary")

go_binary(
    name = "server",
    srcs = ["server.go"],
    deps = [
        "//proto:arithmetic_go_proto",
        "//remote:grpc_go",
        "//arithmetic_server:arithmetic_server",
    ],
    importpath = "saranshsaini/arithmetic_server/cmd/arithmetic_server",  # Optional, but recommended for clarity
    visibility = ["//visibility:public"],
)