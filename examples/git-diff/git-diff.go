package main

import (
    "fmt"
    "os"
    "os/exec"

    "github.com/mauricelam/gocmdln/params"
)

type separatedParamSpec struct {
    params.ParamSpec
}

func (spec *separatedParamSpec) CaptureLength(argvSlice []string) (int, error) {
    for i, arg := range argvSlice {
        if arg == "--" {
            argvSlice = argvSlice[:i+1]
            break
        }
    }
    return spec.ParamSpec.CaptureLength(argvSlice)
}

func (spec *separatedParamSpec) Set(args []string) error {
    if l := len(args); l > 0 && args[l-1] == "--" {
        args = args[:l-1]
    }
    return spec.ParamSpec.Set(args)
}

/**
 git diff [options] [<commit>] [--] [<path>…​]
 git diff [options] --cached [<commit>] [--] [<path>…​]
 git diff [options] <commit> <commit> [--] [<path>…​]
 git diff [options] <blob> <blob>
 git diff [options] [--no-index] [--] <path> <path>
 */
func main() {
    var commitsOrPaths params.StringValueList
    params.Param(&separatedParamSpec{ params.NewListParamSpec(&commitsOrPaths, "commits", true, nil) })
    paths := params.StringList("paths", true, nil)

    err := params.Parse(os.Args[1:])
    if err != nil {
        println(err.Error())
        return
    }

    args := []string { "diff" }
    args = append(args, commitsOrPaths...)
    if *paths != nil {
        args = append(args, "--")
        args = append(args, *paths...)
    }

    output, err := exec.Command("git", args...).CombinedOutput()
    if err != nil {
      os.Stderr.WriteString(err.Error())
    }
    fmt.Println(string(output))
}
