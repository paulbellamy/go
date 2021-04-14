package main

import (
	"context"
	"fmt"

	backends "github.com/stellar/go/ingest/ledgerbackend"
)

func helloworld() {
	backend, err := backends.NewCaptive(config)
	panicIf(err)
	defer backend.Close()
	ctx := context.Background()

	// Prepare a single ledger to be ingested,
	err = backend.PrepareRange(ctx, backends.BoundedRange(123456, 123456))
	panicIf(err)

	// then retrieve it:
	ok, ledger, err := backend.GetLedger(ctx, 123456)
	if !ok {
		err = fmt.Errorf("The ledger doesn't exist on the backend.")
	}

	panicIf(err)

	// Now `ledger` is a raw `xdr.LedgerCloseMeta` object containing the
	// transactions contained within this ledger.
	fmt.Printf("\nHello, Sequence %d.\n", ledger.LedgerSequence())
}
