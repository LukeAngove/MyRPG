load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_go",
    urls = ["https://github.com/bazelbuild/rules_go/releases/download/0.18.1/rules_go-0.18.1.tar.gz"],
    sha256 = "77dfd303492f2634de7a660445ee2d3de2960cbd52f97d8c0dffa9362d3ddef9",
)

load("@io_bazel_rules_go//go:deps.bzl", "go_rules_dependencies", "go_register_toolchains")

go_rules_dependencies()

go_register_toolchains()

http_archive(
    name = "bazel_gazelle",
    urls = ["https://github.com/bazelbuild/bazel-gazelle/releases/download/0.17.0/bazel-gazelle-0.17.0.tar.gz"],
    sha256 = "3c681998538231a2d24d0c07ed5a7658cb72bfb5fd4bf9911157c0e9ac6a2687",
)

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

gazelle_dependencies()

http_archive(
    name = "bazel_latex",
    strip_prefix = "bazel-latex-luke.develop",
    url = "https://github.com/LukeAngove/bazel-latex/archive/luke.develop.tar.gz",
    sha256 = "e7203c34fd7091014e47b8d39db629a5a98c789f8761e2a30c12aaf17a49784d",
)

load("@bazel_latex//:repositories.bzl", "latex_repositories")

latex_repositories()

http_archive(
    name = "bazel_pandoc",
    strip_prefix = "bazel-pandoc-0.2",
    url = "https://github.com/ProdriveTechnologies/bazel-pandoc/archive/v0.2.tar.gz",
    sha256 = "47ad1f08db3e6c8cc104931c11e099fd0603c174400b9cc852e2481abe08db24",
)

load("@bazel_pandoc//:repositories.bzl", "pandoc_repositories")

pandoc_repositories()

http_archive(
    name = "rpg_style",
    strip_prefix = "DND-5e-LaTeX-Template-luke.develop",
    url = "https://github.com/LukeAngove/DND-5e-LaTeX-Template/archive/luke.develop.tar.gz",
    build_file = "//:BUILD.style",
    sha256 = "fb927ed9690d45e6f81826e70016db24d0e319fef4138eca923a21c53a4392ee",
)

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
