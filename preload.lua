function TestRequireModule(t)
    modules = {
        "base64",
        "cmd",
        "crypto",
        "filepath",
        "goos",
        "humanize",
        "inspect",
        "ioutil",
        "json",
        "log",
        "plugin",
        "regexp",
        "runtime",
        "shellescape",
        "stats",
        "storage",
        "strings",
        "template",
        "time",
        "xmlpath",
        "yaml",
    }
    for _, module in ipairs(modules) do
        t:Run(module, function(t)
            require(module)
        end)
    end
end

