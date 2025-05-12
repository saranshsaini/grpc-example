load("@rules_go//go:def.bzl", "go_binary")
# load("@io_bazel_rules_docker//contrib:push-all.bzl", "docker_push")
# load("@io_bazel_rules_docker//container:container.bzl", "container_push")
# load("@io_bazel_rules_docker//go:image.bzl", "go_image")

go_binary(
    name = "server",
    srcs = ["server.go"],
    deps = [
        "//proto:arithmetic_go_proto",
        "//remote:grpc_go",
        "//arithmetic_server:arithmetic_server",
        "//remote:reflection",
    ],
    importpath = "saranshsaini/arithmetic_server/cmd/arithmetic_server",
    visibility = ["//visibility:public"],
)

# go_image(
#     name = "server_image",
#     binary = ":server",
#     tags = ["manual"]
# )

# container_push(
#     name = "push_server_image",
#     format = "Docker",
#     image = ":server_image",
#     registry = "gcr.io",
#     repository = "saranshsaini",
#     tag = "latest",
#     visibility = ["//visibility:public"],
#     stamp = "@io_bazel_rules_docker//stamp:always",
#     tags = ["manual"],
# )