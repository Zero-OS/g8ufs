package main

import (
	"flag"
	"fmt"
	"github.com/g8os/g8ufs"
	"github.com/g8os/g8ufs/meta"
	"github.com/g8os/g8ufs/storage"
	"net/url"
	"os"
)

type Cmd struct {
	MetaDB  string
	Backend string
	URL     string
	Reset   bool
}

func (c *Cmd) Validate() []error {
	var errors []error
	if c.MetaDB == "" {
		errors = append(errors,
			fmt.Errorf("meta is required"),
		)
	}

	return errors
}

func mount(cmd *Cmd, target string) error {
	u, err := url.Parse(cmd.URL)
	if err != nil {
		return err
	}

	store, err := meta.NewRocksMeta("", cmd.MetaDB)
	if err != nil {
		return fmt.Errorf("failed to initialize meta store: %s", err)
	}
	aydo, err := storage.NewAydoStorage(u)
	if err != nil {
		return err
	}

	fs, err := g8ufs.Mount(&g8ufs.Options{
		MetaStore: store,
		Backend:   cmd.Backend,
		Target:    target,
		Storage:   aydo,
		Reset:     cmd.Reset,
	})

	if err != nil {
		return err
	}

	fs.Wait()
	return nil
}

func main() {
	var cmd Cmd
	flag.BoolVar(&cmd.Reset, "reset", false, "Reset filesystem on mount")
	flag.StringVar(&cmd.MetaDB, "meta", "", "Path to metadata database (rocksdb)")
	flag.StringVar(&cmd.Backend, "backend", "/tmp/backend", "Working directory of the filesystem (cache and others)")
	flag.StringVar(&cmd.URL, "aydo-url", "https://stor.jumpscale.org/stor2/store/ubuntu-g8os-flist/", "Base aydo URL for the plist")

	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Fprintf(os.Stderr, "Missing mount point argument\n")
		os.Exit(1)
	}

	if errs := cmd.Validate(); errs != nil {
		for _, err := range errs {
			fmt.Fprintf(os.Stderr, "%s\n", err)
		}
		os.Exit(1)
	}
	fmt.Println("Reset:", cmd.Reset)
	if err := mount(&cmd, flag.Arg(0)); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
