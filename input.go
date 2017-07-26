package main

import (
        "fmt"
        "github.com/adammck/venv"
        "github.com/blang/vfs"
        "io/ioutil"
)

func GetInputPath(fs vfs.Filesystem, env venv.Env) string {

        var fn string

        fn = env.Getenv("TF_STATE")
        if fn != "" {
                return fn
        }

        fn = env.Getenv("TI_TFSTATE")
        if fn != "" {
                return fn
        }

        contents, err := ioutil.ReadFile(".terraform/environment")
        if err == nil {
                fn = fmt.Sprintf("terraform.tfstate.d/%s/terraform.tfstate", contents)
                _, err := fs.Stat(fn)
                if err == nil {
                        return fn
                }
        }

        fn = "terraform.tfstate"
        _, err = fs.Stat(fn)
        if err == nil {
                return fn
        }

        return "."
}
