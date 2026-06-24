package repository

import (
	"database/sql"
	"strings"
	"time"
)

// parseTime 解析 ISO8601 UTC 字符串
func parseTime(ns sql.NullString) *time.Time {
	if !ns.Valid || ns.String == "" {
		return nil
	}
	t, err := time.Parse(time.RFC3339Nano, ns.String)
	if err != nil {
		// 兼容旧格式
		t, err = time.Parse("2006-01-02 15:04:05", ns.String)
		if err != nil {
			return nil
		}
	}
	return &t
}

// normalizeKeyword 关键词规范化（trim + lower）
func normalizeKeyword(s string) string {
	return strings.ToLower(strings.TrimSpace(s))
}

// normalizeTagName 标签名规范化
func normalizeTagName(s string) string {
	s = strings.TrimSpace(s)
	// 简单 lowercase，Unicode 折叠可后续扩展
	return strings.ToLower(s)
}

// nullStr 包装 nil 为空字符串
func nullStr(s string) sql.NullString {
	if s == "" {
		return sql.NullString{}
	}
	return sql.NullString{String: s, Valid: true}
}
