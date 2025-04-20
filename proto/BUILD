load("@rules_proto//proto:defs.bzl", "proto_library")
load("@rules_go//proto:def.bzl", "go_grpc_library")

proto_library(
    name = "arithmetic_proto",
    srcs = ["arithmetic.proto"],
)

go_grpc_library(
    name = "arithmetic_go_proto",
    proto = ":arithmetic_proto",
    importpath = "saranshsaini/arithmetic_pb",
    visibility = [
        "//visibility:public",
    ],
)