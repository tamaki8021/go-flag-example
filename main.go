package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// サブコマンドの宣言
	fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	// 固有のフラグ
	fooEnable := fooCmd.Bool("enable", false, "enable")
	fooName := fooCmd.String("name", "", "name")

	barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
	barLevel := barCmd.Int("level", 0, "level")

	numbPtr := flag.Int("numb", 42, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")

	var showVersion bool
	var svar string

	// --version又は-versionでバージョン情報を表示
	flag.BoolVar(&showVersion, "version", false, "print version")
	flag.StringVar(&svar, "svar", "bar", "a string var")

	// flag.Parse()を呼ぶことでコマンドラインの引数のフラグが解析され、フラグが変数にバインドされる。
	flag.Parse()
	args := flag.Args()


	if showVersion {
		fmt.Println(getVersion())
		// version引数があった場合はバージョン情報を表示して終了
		os.Exit(0)
	}

	// サブコマンドはプログラムの最初の引数になる
	if len(os.Args) < 2 {
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "foo":
		fooCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'foo'")
		fmt.Println("  enable:", *fooEnable)
		fmt.Println("  name:", *fooName)
		fmt.Println("  tail:", fooCmd.Args())
	case "bar":
		barCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'bar'")
		fmt.Println("  level:", *barLevel)
		fmt.Println("  tail:", barCmd.Args())
	default:
		fmt.Println("expected 'foo' or 'bar' subcommands")
		fmt.Println("fork:", *boolPtr)
		fmt.Println("numb:", *numbPtr)

		fmt.Println("svar:", svar)

		fmt.Println(args)
		fmt.Println("This is CLI tool.")
		os.Exit(1)
	}

}

func getVersion() string {
	return "0.0.1"
}
