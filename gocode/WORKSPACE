load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_go",
    urls = ["https://github.com/bazelbuild/rules_go/releases/download/0.18.1/rules_go-0.18.1.tar.gz"],
    sha256 = "77dfd303492f2634de7a660445ee2d3de2960cbd52f97d8c0dffa9362d3ddef9",
)

http_archive(
    name = "bazel_gazelle",
    urls = ["https://github.com/bazelbuild/bazel-gazelle/releases/download/0.17.0/bazel-gazelle-0.17.0.tar.gz"],
    sha256 = "3c681998538231a2d24d0c07ed5a7658cb72bfb5fd4bf9911157c0e9ac6a2687",
)

load("@io_bazel_rules_go//go:deps.bzl", "go_rules_dependencies", "go_register_toolchains")

go_rules_dependencies()

go_register_toolchains()

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

gazelle_dependencies()

go_repository(
    name = "in_gopkg_yaml_v2",
    commit = "51d6538a90f86fe93ac480b35f37b2be17fef232",
    importpath = "gopkg.in/yaml.v2",
)

go_repository(
    name = "com_github_ajstarks_svgo",
    commit = "6ce6a3bcf6cde6c5007887677ebd148ec30f42a4",
    importpath = "github.com/ajstarks/svgo",
)

go_repository(
    name = "com_github_ajstarks_deck",
    commit = "6fe4637ccacd917c6de33c3c2bf4acb04e89bbb1",
    importpath = "github.com/ajstarks/deck",
)

go_repository(
    name = "in_gopkg_check_v1",
    commit = "788fd78401277ebd861206a03c884797c6ec5541",
    importpath = "gopkg.in/check.v1",
)

go_repository(
    name = "co_honnef_go_tools",
    commit = "d36bf90409063c0fe4fecdc07cc71e25b59a4050",
    importpath = "honnef.co/go/tools",
)

go_repository(
    name = "com_github_kr_pretty",
    commit = "73f6ac0b30a98e433b289500d779f50c1a6f0712",
    importpath = "github.com/kr/pretty",
)

go_repository(
    name = "com_github_kr_text",
    commit = "e2ffdb16a802fe2bb95e2e35ff34f0e53aeef34f",
    importpath = "github.com/kr/text",
)
