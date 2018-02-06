package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/magik6k/git-remote-ipld/core"
)

const (
	IPLD_PREFIX = "ipld://"
	IPFS_PREFIX = "ipfs://"
)

func Main(args []string, reader io.Reader, writer io.Writer, logger *log.Logger) error {
	logger = log.New(os.Stderr, "", 0)
	logger.Printf("Args: %s", args)
	if len(args) < 3 {
		return fmt.Errorf("Usage: git-remote-ipld remote-name url")
	}

	hashArg := args[2]
	if strings.HasPrefix(hashArg, IPLD_PREFIX) || strings.HasPrefix(hashArg, IPFS_PREFIX) {
		hashArg = hashArg[len(IPLD_PREFIX):]
	}

	remote, err := core.NewRemote(&IpldHandler{remoteHash: hashArg, osArgs: args}, reader, writer, logger)
	if err != nil {
		return err
	}
	defer remote.Close()

	return remote.ProcessCommands()
}

func main() {
	if err := Main(os.Args, nil, nil, nil); err != nil {
		fmt.Fprintf(os.Stderr, "\x1b[K")
		log.Fatal(err)
	}
}
