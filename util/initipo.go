package util

import (
	"fmt"
	"os"

	"github.com/bingoohuang/statiq/fs"
)

func Ipo(ipo bool) {
	if !ipo {
		return
	}

	if err := ipoInit(); err != nil {
		fmt.Println(err)
	}
	os.Exit(0)
}

func ipoInit() error {
	sfs, err := fs.New()
	if err != nil {
		return err
	}

	if err = InitCtl(sfs, "/ctl.tpl.sh", "./ctl"); err != nil {
		return err
	}

	if err = InitConfigFile(sfs, "/config.tpl.toml", "./config.toml"); err != nil {
		return err
	}

	return nil
}
