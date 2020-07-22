// Copyright 2019 The go-TITprotocol Authors
// This file is part of go-TITprotocol.
//
// go-TITprotocol is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// go-TITprotocol is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with go-TITprotocol. If not, see <http://www.gnu.org/licenses/>.

// Command bzzhash computes a swarm tree hash.
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/TITprotocol/go-TITprotocol/cmd/utils"
	"github.com/TITprotocol/go-TITprotocol/swarm/storage"
	"gopkg.in/urfave/cli.v1"
)

var hashesCommand = cli.Command{
	Action:             hashes,
	CustomHelpTemplate: helpTemplate,
	Name:               "hashes",
	Usage:              "print all hashes of a file to STDOUT",
	ArgsUsage:          "<file>",
	Description:        "Prints all hashes of a file to STDOUT",
}

func hashes(ctx *cli.Context) {
	args := ctx.Args()
	if len(args) < 1 {
		utils.Fatalf("Usage: swarm hashes <file name>")
	}
	f, err := os.Open(args[0])
	if err != nil {
		utils.Fatalf("Error opening file " + args[1])
	}
	defer f.Close()

	fileStore := storage.NewFileStore(&storage.FakeChunkStore{}, storage.NewFileStoreParams())
	refs, err := fileStore.GetAllReferences(context.TODO(), f, false)
	if err != nil {
		utils.Fatalf("%v\n", err)
	} else {
		for _, r := range refs {
			fmt.Println(r.String())
		}
	}
}
