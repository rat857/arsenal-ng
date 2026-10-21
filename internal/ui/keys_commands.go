// Package ui provides the terminal user interface for arsenal-ng.
//
// This file defines special command prefixes (set, unset, show) and common
// variable suggestions used for command hints in the search view.
package ui

// =============================================================================
// Special Command Prefixes
// =============================================================================

const (
	cmdSet   = "set"
	cmdUnset = "unset"
	cmdShow  = "show"
)

// =============================================================================
// Common Variable Suggestions
// =============================================================================

// variableHint represents a variable suggestion for the set command hints.
type variableHint struct {
	Name string
	Desc string
}

// commonVariables contains common variable suggestions for set command hints.
var commonVariables = []variableHint{
	{Name: "domain", Desc: "目标域名（例如 corp.local）"},
	{Name: "hash", Desc: "NTLM 哈希"},
	{Name: "ip", Desc: "目标 IP 地址（例如 10.10.10.10）"},
	{Name: "lhost", Desc: "本机地址（你的 IP）"},
	{Name: "lport", Desc: "本机端口（用于反弹 shell）"},
	{Name: "output", Desc: "输出文件名"},
	{Name: "pass", Desc: "密码"},
	{Name: "port", Desc: "端口号（例如 445）"},
	{Name: "url", Desc: "目标 URL"},
	{Name: "user", Desc: "用户名（例如 admin）"},
	{Name: "wordlist", Desc: "字典路径"},
}

