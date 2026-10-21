// Package ui provides the terminal user interface for arsenal-ng.
//
// This file contains the help view rendering function. It displays a
// comprehensive help screen with keyboard shortcuts, usage examples, and
// workflow instructions for the application.
package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"

	"github.com/halilkirazkaya/arsenal-ng/internal/config"
)

// =============================================================================
// Help View
// =============================================================================

func (m App) viewHelp() string {
	var b strings.Builder

	// Local styles for help screen
	sectionStyle := lipgloss.NewStyle().Foreground(primaryColor).Bold(true)
	lineStyle := lipgloss.NewStyle().Foreground(secondaryColor)
	keyStyle := lipgloss.NewStyle().Foreground(accentColor).Bold(true)
	helpDescStyle := lipgloss.NewStyle().Foreground(lipgloss.Color(config.ColorBright))
	cmdStyle := lipgloss.NewStyle().Foreground(secondaryColor)
	arrowStyle := lipgloss.NewStyle().Foreground(dimColor)
	exampleStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#98D8C8"))
	varStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#F7DC6F")).Bold(true)

	separator := "  ════════════════════════════════════════════════"

	b.WriteString(m.renderHeader())
	b.WriteString("\n\n")

	// Title
	b.WriteString(sectionStyle.Render("📖 帮助"))
	b.WriteString("\n\n")

	// Navigation section
	b.WriteString(sectionStyle.Render("  导航"))
	b.WriteString("\n")
	b.WriteString(lineStyle.Render(separator))
	b.WriteString("\n")
	b.WriteString(fmt.Sprintf("  %s  %s\n", keyStyle.Render("↑ / ↓        "), helpDescStyle.Render("在列表中上/下移动")))
	b.WriteString(fmt.Sprintf("  %s  %s\n", keyStyle.Render("PgUp / PgDown"), helpDescStyle.Render("上/下翻一页")))
	b.WriteString(fmt.Sprintf("  %s  %s\n", keyStyle.Render("Enter        "), helpDescStyle.Render("选中高亮的命令")))
	b.WriteString(fmt.Sprintf("  %s  %s\n", keyStyle.Render("ESC          "), helpDescStyle.Render("退出 arsenal 或返回")))
	b.WriteString(fmt.Sprintf("  %s  %s\n", keyStyle.Render("help         "), helpDescStyle.Render("显示本帮助界面")))
	b.WriteString("\n")

	// Arguments section
	b.WriteString(sectionStyle.Render("  参数输入"))
	b.WriteString("\n")
	b.WriteString(lineStyle.Render(separator))
	b.WriteString("\n")
	b.WriteString(fmt.Sprintf("  %s  %s\n", keyStyle.Render("Tab          "), helpDescStyle.Render("跳到下一个参数字段")))
	b.WriteString(fmt.Sprintf("  %s  %s\n", keyStyle.Render("Shift+Tab    "), helpDescStyle.Render("跳到上一个参数字段")))
	b.WriteString(fmt.Sprintf("  %s  %s\n", keyStyle.Render("Enter        "), helpDescStyle.Render("用填好的参数执行命令")))
	b.WriteString("\n")

	// Global Variables section
	b.WriteString(sectionStyle.Render("  🚀 全局变量"))
	b.WriteString("\n")
	b.WriteString(lineStyle.Render(separator))
	b.WriteString("\n")
	b.WriteString(helpDescStyle.Render("  变量设置一次，处处可用！\n"))
	b.WriteString("\n")
	// Align arrows by padding commands to fixed width
	cmdWidth := 20
	b.WriteString(fmt.Sprintf("  %s%s %s %s\n", cmdStyle.Render("set key=value"), strings.Repeat(" ", cmdWidth-len("set key=value")), arrowStyle.Render("→"), helpDescStyle.Render("设置一个变量")))
	b.WriteString(fmt.Sprintf("  %s%s %s %s\n", cmdStyle.Render("unset key"), strings.Repeat(" ", cmdWidth-len("unset key")), arrowStyle.Render("→"), helpDescStyle.Render("删除一个变量")))
	b.WriteString(fmt.Sprintf("  %s%s %s %s\n", cmdStyle.Render("variables"), strings.Repeat(" ", cmdWidth-len("variables")), arrowStyle.Render("→"), helpDescStyle.Render("列出你的所有变量")))
	b.WriteString(fmt.Sprintf("  %s%s %s %s\n", cmdStyle.Render("tools"), strings.Repeat(" ", cmdWidth-len("tools")), arrowStyle.Render("→"), helpDescStyle.Render("列出所有可用工具")))
	b.WriteString("\n")

	// Example workflow
	b.WriteString(sectionStyle.Render("  📋 使用示例"))
	b.WriteString("\n")
	b.WriteString(lineStyle.Render(separator))
	b.WriteString("\n")
	b.WriteString(fmt.Sprintf("  %s %s\n", exampleStyle.Render("1."), helpDescStyle.Render("输入：")+cmdStyle.Render("set key=value")+" "+keyStyle.Render("[Enter]")))
	b.WriteString(fmt.Sprintf("  %s %s\n", exampleStyle.Render("2."), helpDescStyle.Render("例如：")+cmdStyle.Render("set ip=10.10.10.10")+" "+keyStyle.Render("[Enter]")))
	b.WriteString(fmt.Sprintf("  %s %s\n", exampleStyle.Render("3."), helpDescStyle.Render("选择任意包含 ")+varStyle.Render("{{ip}}")+" "+helpDescStyle.Render("或 ")+varStyle.Render("{{domain}}")))
	b.WriteString(fmt.Sprintf("  %s %s\n", exampleStyle.Render("4."), successStyle.Render("参数会自动填充！✓")))
	b.WriteString("\n")

	// Before/After example
	b.WriteString(fmt.Sprintf("  %s nmap -sV %s %s %s\n",
		helpStyle.Render("之前："), varStyle.Render("{{ip}}"), arrowStyle.Render("→"), helpStyle.Render("每次都要手动输入 IP")))
	b.WriteString(fmt.Sprintf("  %s nmap -sV %s  %s %s\n",
		successStyle.Render("之后："), exampleStyle.Render("10.10.10.10"), arrowStyle.Render("→"), successStyle.Render("自动填充！✓")))

	b.WriteString("\n\n")
	b.WriteString(helpStyle.Render("  按 ESC 或 Enter 返回"))

	return b.String()
}
