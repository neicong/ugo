package common

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	exec2 "os/exec"
)

var (
	llPath  string
	outPath string
)

func init() {
	flag.StringVar(&llPath, "l", "ugo.out.ll", ".ll 路径")
	flag.StringVar(&outPath, "o", "ugo.out", ".out 路径")

}

func exec(opn []byte) error {
	flag.Parse()
	if llPath == "" {
		return errors.New(".ll 路径 不能为空")
	}
	if outPath == "" {
		return errors.New(".out 路径 不能为空")
	}
	err := os.WriteFile(llPath, opn, 0666)
	if err != nil {
		return err
	}
	err = exec2.Command("clang", "-Wno-override-module", "--target=x86_64-pc-mingw64", "-o", outPath, llPath).Run()
	if err != nil {
		close()
		return err
	}
	return nil
}

func Run(opn []byte) {
	err := exec(opn)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer close()
	if err = exec2.Command(fmt.Sprintf("./%s", outPath)).Run(); err != nil {
		log.Fatal(err.Error())
	}
}
func close() {
	if PathExists(llPath) {
		os.Remove(llPath)
	}
	if PathExists(outPath) {
		os.Remove(outPath)
	}
}

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
