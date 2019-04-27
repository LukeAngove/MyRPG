def _gen_cards_impl(ctx):
    # Generate a datafile of concatenated fortunes.
    action = ctx.actions.run(
        outputs = [ctx.outputs.out],
        inputs = ctx.files.srcs + ctx.files.template,
        executable = ctx.executable.exe,
        arguments = [
            "--html_template",
            "{}".format([f.path for f in ctx.files.template][0]),
            "--card_outlines",
            "{}".format([f.path for f in ctx.files.srcs][0]),
            "--output",
            "{}".format(ctx.outputs.out.path),
            ]
        )

gen_cards = rule(
    implementation = _gen_cards_impl,
    attrs = {
        "srcs": attr.label(
            allow_files = [".yml"],
            doc = "Input card YAML files.",
            mandatory = True,
        ),
        "template": attr.label(
            allow_files = [".html"],
            doc = "Input card templates files.",
            mandatory = True,
        ),
        "out": attr.output(mandatory = True),
        "exe": attr.label(
            executable = True,
            cfg = "host",
            default = Label("//src/card_maker/cmd:card_maker")
        )
    }
)
