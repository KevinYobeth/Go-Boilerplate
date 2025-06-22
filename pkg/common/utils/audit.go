package utils

import "time"

func NewAuditTrail(issuer *string) (time.Time, string) {
	now := time.Now().UTC()
	by := newIssuer(issuer)

	return now, by
}

func newIssuer(issuer *string) string {
	by := "system"
	if issuer != nil {
		by = *issuer
	}

	return by
}
