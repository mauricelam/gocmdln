package main

import (
    "flag"
    "fmt"
    "os"
    "os/exec"

    "github.com/mauricelam/gocmdln/params"
)

/**
    sed [OPTION]... <command> [input-file]...

    -n, --quiet, --silent
      suppress automatic printing of pattern space
    -e script, --expression=script
      add the script to the commands to be executed
    -f script-file, --file=script-file
      add the contents of script-file to the commands to be executed
    --follow-symlinks
      follow symlinks when processing in place; hard links will still be broken.
    -i[SUFFIX], --in-place[=SUFFIX]
      edit files in place (makes backup if extension supplied). The default operation mode is to break symbolic and hard links. This can be changed with --follow-symlinks and --copy.
    -c, --copy
      use copy instead of rename when shuffling files in -i mode. While this will avoid breaking links (symbolic or hard), the resulting editing operation is not atomic. This is rarely the desired mode; --follow-symlinks is usually enough, and it is both faster and more secure.
    -l N, --line-length=N
      specify the desired line-wrap length for the 'l' command
    --posix
      disable all GNU extensions.
    -r, --regexp-extended
      use extended regular expressions in the script.
    -s, --separate
      consider files as separate rather than as a single continuous long stream.
    -u, --unbuffered
      load minimal amounts of data from the input files and flush the output buffers more often
    --help
      display this help and exit

    --version
      output version information and exit
 */

func main() {
    command := params.String("command", true, nil)
    inputFiles := params.StringList("inputFiles", true, nil)

    quiet := flag.Bool("quiet", false, "suppress automatic printing of pattern space")
    script := flag.String("e", "", "add the script to the commands to be executed")
    scriptFile := flag.String("f", "", "add the contents of script-file to the commands to be executed")
    followSymlinks := flag.Bool("follow-symlinks", false, "follow symlinks when processing in place; hard links will still be broken.")
    inplace := flag.Bool("i", false, "edit files in place (makes backup if extension supplied). The default operation mode is to break symbolic and hard links. This can be changed with --follow-symlinks and --copy.")
    copyFlag := flag.Bool("copy", false, "use copy instead of rename when shuffling files in -i mode. While this will avoid breaking links (symbolic or hard), the resulting editing operation is not atomic. This is rarely the desired mode; --follow-symlinks is usually enough, and it is both faster and more secure.")
    lineLength := flag.Int("line-length", 0, "specify the desired line-wrap length for the 'l' command")
    posix := flag.Bool("posix", false, "disable all GNU extensions")
    regexpExtended := flag.Bool("regexp-extended", false, "use extended regular expressions in the script.")
    separate := flag.Bool("separate", false, "consider files as separate rather than as a single continuous long stream.")
    unbuffered := flag.Bool("unbuffered", false, "load minimal amounts of data from the input files and flush the output buffers more often")

    flag.Parse()
    params.Parse(flag.Args())

    // Reconstruct the command flags, just because
    args := []string{}
    if *quiet { args = append(args, "--quiet") }
    if *script != "" { args = append(args, "-e", *script) }
    if *scriptFile != "" { args = append(args, "-f", *scriptFile) }
    if *followSymlinks { args = append(args, "-follow-symlinks") }
    if *inplace { args = append(args, "-i") }
    if *copyFlag { args = append(args, "-copy") }
    if *lineLength != 0 { args = append(args, "-line-length", string(*lineLength)) }
    if *posix { args = append(args, "-posix") }
    if *regexpExtended { args = append(args, "-regexp-extended") }
    if *separate { args = append(args, "-separate") }
    if *unbuffered { args = append(args, "-unbuffered") }
    args = append(args, *command)
    args = append(args, *inputFiles...)

    fmt.Println(args)
    output, err := exec.Command("sed", args...).CombinedOutput()
    if err != nil {
      os.Stderr.WriteString(err.Error())
    }
    fmt.Println(string(output))
}
