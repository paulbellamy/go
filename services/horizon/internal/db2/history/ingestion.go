package history

// TruncateIngestStateTables clears out ingestion state tables.
// Ingestion state tables are horizon database tables populated by
// the ingestion system using history archive snapshots.
// Any horizon database tables which cannot be populated using
// history archive snapshots will not be truncated.
func (q *Q) TruncateIngestStateTables() error {
	return q.TruncateTables([]string{
		"accounts",
		"accounts_data",
		"accounts_signers",
		"claimable_balances",
		"exp_asset_stats",
		"offers",
		"trust_lines",
	})
}
