load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["config.go"],
    importmap = "k8s.io/cloud-provider-gcp/vendor/k8s.io/cloud-provider/app/config",
    importpath = "k8s.io/cloud-provider/app/config",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/k8s.io/apiserver/pkg/server:go_default_library",
        "//vendor/k8s.io/client-go/informers:go_default_library",
        "//vendor/k8s.io/client-go/kubernetes:go_default_library",
        "//vendor/k8s.io/client-go/rest:go_default_library",
        "//vendor/k8s.io/client-go/tools/record:go_default_library",
        "//vendor/k8s.io/cloud-provider:go_default_library",
        "//vendor/k8s.io/cloud-provider/config:go_default_library",
    ],
)
