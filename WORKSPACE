# load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
# http_archive(
#     name = "io_bazel_rules_docker",
#     sha256 = "b1e80761a8a8243d03ebca8845e9cc1ba6c82ce7c5179ce2b295cd36f7e394bf",
#     urls = ["https://github.com/bazelbuild/rules_docker/releases/download/v0.25.0/rules_docker-v0.25.0.tar.gz"],
# )

# load(
#     "@io_bazel_rules_docker//repositories:repositories.bzl",
#     container_repositories = "repositories",
# )
# container_repositories()

# load("@io_bazel_rules_docker//repositories:deps.bzl", container_deps = "deps")

# container_deps()

# load("@io_bazel_rules_docker//docker/package_managers:download_pkgs.bzl", "download_pkgs")
# load("@io_bazel_rules_docker//docker/package_managers:install_pkgs.bzl", "install_pkgs")
# load("@io_bazel_rules_docker//docker/util:run.bzl", "container_run_and_commit")

# load("@io_bazel_rules_docker//go:image.bzl", _go_image_repos = "repositories")
# _go_image_repos()

# load(
#     "@io_bazel_rules_docker//container:container.bzl",
#     "container_pull",
# )

# container_pull(
#   name = "java_base",
#   registry = "gcr.io",
#   repository = "distroless/java",
#   # 'tag' is also supported, but digest is encouraged for reproducibility.
#   digest = "sha256:deadbeef",
# )