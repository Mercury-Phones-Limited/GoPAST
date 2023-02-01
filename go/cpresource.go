package cpresource

import (
        "github.com/sethvargo/go-password/password"
        "log"
        "os"
        "os/exec"
)

func NewPassword() string {
        // Customize the list of symbols.
        gen, err := password.NewGenerator(&password.GeneratorInput{
                Symbols: "!#$^ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789",
        })
        if err != nil {
                log.Fatal(err)
        }
        res, err := gen.Generate(32, 15, 15, false, true)
        if err != nil {
                log.Fatal(err)
        }
        return res
}

func ClearScreen() {
        c := exec.Command("clear")
        c.Stdout = os.Stdout
        c.Run()
}

func getFile(file string) string {
        var path string
        path = "/usr/local/etc/cpresource/"

        var1, err := os.ReadFile(path + file)
        if err != nil {
                log.Fatal(err)

        }
        return string(var1)
}

func StartHTML() string {
        var1 := getFile("start.html")
        return var1
}

func EndHTML() string {
        var1 := getFile("end.html")
        return var1
}

func DBReadDetails() string {
        var1 := getFile("dbReadDetails.txt")
        return var1
}

func DBWriteDetails() string {
       var1 := getFile("dbWriteDetails.txt")
       return var1
}
