package llo

import (
	"context"
	"fmt"

	"github.com/lib/pq"
	ocr2types "github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"github.com/smartcontractkit/chainlink-common/pkg/sqlutil"
)

type RetirementReportCacheORM interface {
	StoreAttestedRetirementReport(ctx context.Context, cd ocr2types.ConfigDigest, attestedRetirementReport []byte) error
	LoadAttestedRetirementReports(ctx context.Context) (map[ocr2types.ConfigDigest][]byte, error)
	StoreConfig(ctx context.Context, cd ocr2types.ConfigDigest, signers [][]byte, f uint8) error
	LoadConfigs(ctx context.Context) ([]Config, error)
}

type retirementReportCacheORM struct {
	ds sqlutil.DataSource
}

// TODO: Test ORM
// TODO: Test whole thing
func (o *retirementReportCacheORM) StoreAttestedRetirementReport(ctx context.Context, cd ocr2types.ConfigDigest, attestedRetirementReport []byte) error {
	_, err := o.ds.ExecContext(ctx, `
INSERT INTO llo_retirement_report_cache (config_digest, attested_retirement_report, updated_at)
VALUES ($1, $2, NOW())
ON CONFLICT (config_digest) DO NOTHING
`, cd, attestedRetirementReport)
	if err != nil {
		return fmt.Errorf("StoreAttestedRetirementReport failed: %w", err)
	}
	return nil
}

func (o *retirementReportCacheORM) LoadAttestedRetirementReports(ctx context.Context) (map[ocr2types.ConfigDigest][]byte, error) {
	rows, err := o.ds.QueryContext(ctx, "SELECT config_digest, attested_retirement_report FROM llo_retirement_report_cache")
	if err != nil {
		return nil, fmt.Errorf("LoadAttestedRetirementReports failed: %w", err)
	}
	defer rows.Close()

	reports := make(map[ocr2types.ConfigDigest][]byte)
	for rows.Next() {
		var cd ocr2types.ConfigDigest
		var arr []byte
		if err := rows.Scan(&cd, &arr); err != nil {
			return nil, fmt.Errorf("LoadAttestedRetirementReports failed: %w", err)
		}
		reports[cd] = arr
	}

	return reports, nil
}

func (o *retirementReportCacheORM) StoreConfig(ctx context.Context, cd ocr2types.ConfigDigest, signers [][]byte, f uint8) error {
	_, err := o.ds.ExecContext(ctx, `INSERT INTO llo_retirement_report_cache_configs (config_digest, signers, f, updated_at) VALUES ($1, $2, $3, NOW())`, cd, signers, f)
	return err
}

type Config struct {
	Digest  [32]byte      `db:"config_digest"`
	Signers pq.ByteaArray `db:"signers"`
	F       uint8         `db:"f"`
}

// TODO: Load on start
func (o *retirementReportCacheORM) LoadConfigs(ctx context.Context) (configs []Config, err error) {
	err = o.ds.SelectContext(ctx, &configs, `SELECT config_digest, signers, f FROM llo_retirement_report_cache_configs`)
	return
}
