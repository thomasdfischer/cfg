package rollback

import (
	"fmt"
	"os"
)

type Tx struct {
	DirsCreated  []string
	FilesCreated []string
}

func NewTx() *Tx {
	return &Tx{
		DirsCreated:  make([]string, 0),
		FilesCreated: make([]string, 0),
	}
}

func (tx *Tx) MustRollback() {
	fmt.Fprintf(os.Stderr, "rolling back changes\n")
	for _, f := range tx.FilesCreated {
		if err := os.Remove(f); err != nil {
			fmt.Fprintf(os.Stderr, "error: failed to remove %v\n", f)
			fmt.Fprintf(os.Stderr, "%v\n", err)
			fmt.Fprintf(os.Stderr, "failed to rollback changes\n")
			os.Exit(1)
		}
	}
	for _, d := range tx.DirsCreated {
		if err := os.RemoveAll(d); err != nil {
			fmt.Fprintf(os.Stderr, "error: failed to remove %v\n", d)
			fmt.Fprintf(os.Stderr, "%v\n", err)
			fmt.Fprintf(os.Stderr, "failed to rollback changes\n")
			os.Exit(1)
		}
	}
}
