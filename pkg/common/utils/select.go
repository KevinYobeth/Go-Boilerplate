package utils

func SelectWithAuditTrail(columns ...string) []string {
	return append(columns, "created_at", "updated_at", "created_by", "updated_by")
}
