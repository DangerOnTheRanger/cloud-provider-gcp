load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["tokenfile.go"],
    importmap = "k8s.io/cloud-provider-gcp/vendor/k8s.io/apiserver/pkg/authentication/token/tokenfile",
    importpath = "k8s.io/apiserver/pkg/authentication/token/tokenfile",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/k8s.io/apiserver/pkg/authentication/authenticator:go_default_library",
        "//vendor/k8s.io/apiserver/pkg/authentication/user:go_default_library",
        "//vendor/k8s.io/klog/v2:go_default_library",
    ],
)
