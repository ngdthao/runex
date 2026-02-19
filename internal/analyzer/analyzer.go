package analyzer

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/runex/runex/internal/detector"
)

type Analyzer struct {
	noColor bool
}

func New(noColor bool) *Analyzer {
	return &Analyzer{
		noColor: noColor,
	}
}

func (a *Analyzer) FormatError(err *detector.DetectedError) string {
	if a.noColor {
		return a.formatPlain(err)
	}
	return a.formatColored(err)
}

func (a *Analyzer) formatColored(err *detector.DetectedError) string {
	var b strings.Builder

	errorColor := color.New(color.FgRed, color.Bold)
	typeColor := color.New(color.FgYellow)
	langColor := color.New(color.FgCyan)
	reset := color.New(color.Reset)

	b.WriteString("\n")
	errorColor.Print("  ╔══════════════════════════════════════════╗\n")
	b.WriteString(errorColor.Sprint("  ║           RUNEX ERROR DETECTED            ║\n"))
	errorColor.Print("  ╚══════════════════════════════════════════╝\n\n")

	typeColor.Print("  Type: ")
	reset.Println(string(err.Type))

	langColor.Print("  Language: ")
	reset.Println(string(err.Language))

	typeColor.Print("  Error: ")
	reset.Println(err.Line)

	if len(err.StackTrace) > 0 {
		typeColor.Print("\n  Stack Trace:\n")
		reset.Println("  " + strings.Join(err.StackTrace[:min(5, len(err.StackTrace))], "\n  "))
	}

	return b.String()
}

func (a *Analyzer) formatPlain(err *detector.DetectedError) string {
	var b strings.Builder

	b.WriteString("\n  === RUNEX ERROR DETECTED ===\n\n")
	fmt.Fprintf(&b, "  Type: %s\n", err.Type)
	fmt.Fprintf(&b, "  Language: %s\n", err.Language)
	fmt.Fprintf(&b, "  Error: %s\n", err.Line)

	if len(err.StackTrace) > 0 {
		fmt.Fprintf(&b, "\n  Stack Trace:\n")
		b.WriteString("  " + strings.Join(err.StackTrace[:min(5, len(err.StackTrace))], "\n  "))
	}

	return b.String()
}

func (a *Analyzer) Summarize(err *detector.DetectedError) string {
	if err == nil {
		return ""
	}

	summary := fmt.Sprintf("%s error in %s", err.Type, err.Language)

	if a.noColor {
		return summary
	}

	errorColor := color.New(color.FgRed)
	return errorColor.Sprint(summary)
}
