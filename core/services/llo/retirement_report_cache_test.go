package llo

import (
	"testing"

	"github.com/stretchr/testify/assert"

	ocrtypes "github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"github.com/smartcontractkit/chainlink/v2/core/logger"
)

func Test_RetirementReportCache(t *testing.T) {
	lggr := logger.TestLogger(t)
	// db := pgtest.NewSqlxDB(t)
	// orm := &retirementReportCacheORM{ds: db}
	rrc := newRetirementReportCache(lggr, nil)
	exampleAttestedRetirementReport := []byte{1, 2, 3, 4, 5}
	exampleDigest := ocrtypes.ConfigDigest{1}

	t.Run("AttestedRetirementReport", func(t *testing.T) {
		attestedRetirementReport, exists := rrc.AttestedRetirementReport(exampleDigest)
		assert.False(t, exists)
		assert.Nil(t, attestedRetirementReport)

		rrc.arrs[exampleDigest] = exampleAttestedRetirementReport

		attestedRetirementReport, exists = rrc.AttestedRetirementReport(exampleDigest)
		assert.True(t, exists)
		assert.Equal(t, exampleAttestedRetirementReport, attestedRetirementReport)
	})
	t.Run("StoreConfig", func(t *testing.T) {
	})
	t.Run("Config", func(t *testing.T) {})
	t.Fatal("TODO")
}

func Test_RetirementReportCache_ORM(t *testing.T) {
	t.Fatal("TODO")
}
