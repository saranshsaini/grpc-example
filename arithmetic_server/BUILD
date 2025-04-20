load("@rules_go//go:def.bzl", "go_library", "go_binary")

go_library(
    name = "arithmetic_server",
    srcs = ["arithmetic_rpc.go"],
    importpath = "saranshsaini/arithmetic_server",
    visibility = [
        "//visibility:public",
    ],
    deps = [
        "//proto:arithmetic_go_proto",
    ],
)