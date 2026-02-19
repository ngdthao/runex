package detector

import (
	"regexp"
	"strings"
)

type Language string

const (
	LangGo     Language = "go"
	LangPython Language = "python"
	LangNodeJS Language = "nodejs"
	LangJava   Language = "java"
	LangRuby   Language = "ruby"
	LangRust   Language = "rust"
	LangUnknown Language = "unknown"
)

type ErrorType string

const (
	ErrPanic        ErrorType = "panic"
	ErrException    ErrorType = "exception"
	ErrRuntimeError ErrorType = "runtime_error"
	ErrCompileError ErrorType = "compile_error"
	ErrCrash        ErrorType = "crash"
)

type DetectedError struct {
	Type      ErrorType
	Language  Language
	Message   string
	StackTrace []string
	Line      string
}

type Detector struct {
	rules    []Rule
	language Language
}

type Rule struct {
	Pattern   *regexp.Regexp
	ErrorType ErrorType
	Language  Language
}

func New(language Language) *Detector {
	d := &Detector{
		language: language,
		rules:    make([]Rule, 0),
	}
	d.initRules()
	return d
}

func (d *Detector) initRules() {
	d.rules = []Rule{
		{Language: LangGo, ErrorType: ErrPanic, Pattern: regexp.MustCompile(`(?i)^panic:`)},
		{Language: LangGo, ErrorType: ErrRuntimeError, Pattern: regexp.MustCompile(`(?i)^runtime error:`)},
		{Language: LangPython, ErrorType: ErrException, Pattern: regexp.MustCompile(`(?i)^Traceback \(most recent call last\):`)},
		{Language: LangPython, ErrorType: ErrException, Pattern: regexp.MustCompile(`(?i)^.*Error:`)},
		{Language: LangPython, ErrorType: ErrException, Pattern: regexp.MustCompile(`(?i)^.*Exception:`)},
		{Language: LangNodeJS, ErrorType: ErrException, Pattern: regexp.MustCompile(`(?i)^.*Error:`)},
		{Language: LangNodeJS, ErrorType: ErrException, Pattern: regexp.MustCompile(`(?i)^TypeError:`)},
		{Language: LangNodeJS, ErrorType: ErrException, Pattern: regexp.MustCompile(`(?i)^ReferenceError:`)},
		{Language: LangNodeJS, ErrorType: ErrException, Pattern: regexp.MustCompile(`(?i)^UnhandledPromiseRejection`)},
		{Language: LangJava, ErrorType: ErrException, Pattern: regexp.MustCompile(`(?i)^Exception in thread`)},
		{Language: LangJava, ErrorType: ErrException, Pattern: regexp.MustCompile(`(?i)^java\.lang\.\w+Exception`)},
		{Language: LangRuby, ErrorType: ErrException, Pattern: regexp.MustCompile(`(?i)^\w+Error:`)},
		{Language: LangRust, ErrorType: ErrCrash, Pattern: regexp.MustCompile(`(?i)^thread '.*' panicked at`)},
		{Language: LangRust, ErrorType: ErrCompileError, Pattern: regexp.MustCompile(`(?i)^error\[`)},
	}
}

func (d *Detector) Detect(input string) *DetectedError {
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		for _, rule := range d.rules {
			if rule.Pattern.MatchString(line) {
				lang := rule.Language
				if d.language != LangUnknown {
					lang = d.language
				}

				return &DetectedError{
					Type:      rule.ErrorType,
					Language:  lang,
					Message:   line,
					StackTrace: d.extractStackTrace(lines),
					Line:      line,
				}
			}
		}
	}

	return nil
}

func (d *Detector) extractStackTrace(lines []string) []string {
	var stack []string
	inStackTrace := false

	for _, line := range lines {
		if strings.Contains(line, "at ") || strings.HasPrefix(line, "\t") ||
		   strings.HasPrefix(line, "    ") || strings.Contains(line, ".go:") {
			inStackTrace = true
			stack = append(stack, line)
		} else if inStackTrace && strings.TrimSpace(line) == "" {
			break
		}
	}

	return stack
}

func DetectLanguage(cmd string) Language {
	switch cmd {
	case "go", "python", "python3", "node", "npm", "java", "ruby", "cargo", "rustc":
		switch cmd {
		case "go":
			return LangGo
		case "python", "python3":
			return LangPython
		case "node", "npm":
			return LangNodeJS
		case "java":
			return LangJava
		case "ruby":
			return LangRuby
		case "cargo", "rustc":
			return LangRust
		}
	}
	return LangUnknown
}

func (d *Detector) SetLanguage(lang Language) {
	d.language = lang
}
