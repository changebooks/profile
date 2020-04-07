# profile
读配置
==

<pre>
b := profile.Builder{}
b.SetDirectory("./github.com/changebooks/profile/profile").SetActive("dev").SetGlobal("database")

if err := b.AddName("database-w"); err != nil {
    fmt.Println(err)
    return
}

if err := b.AddName("database-r"); err != nil {
    fmt.Println(err)
    return
}

if err := b.AddName("database-b"); err != nil {
    fmt.Println(err)
    return
}

p, err := b.Build()
if err != nil {
    fmt.Println(err)
    return
}

fmt.Println(p.ToString())
</pre>