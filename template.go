package main

// // this is taken and modified from:
// // https://raw.githubusercontent.com/drone-plugins/drone-slack/master/template.go
// import (
// 	"fmt"
// 	"net/url"
// 	"strings"
// 	"time"
// 	"unicode"
// 	"unicode/utf8"

// 	"github.com/aymerick/raymond"
// )

// const commentChars = "#"

// func init() {
// 	raymond.RegisterHelpers(funcs)
// }

// // RenderTrim parses and executes a template, returning the results in string
// // format. The result is trimmed to remove left and right padding and newlines
// // that may be added unintentially in the template markup.
// func RenderTrim(template string, playload interface{}) (string, error) {
// 	out, err := raymond.Render(template, playload)
// 	return strings.Trim(out, " \n"), err
// }

// var funcs = map[string]interface{}{
// 	"uppercasefirst": uppercaseFirst,
// 	"uppercase":      strings.ToUpper,
// 	"lowercase":      strings.ToLower,
// 	"duration":       toDuration,
// 	"datetime":       toDatetime,
// 	"success":        isSuccess,
// 	"failure":        isFailure,
// 	"truncate":       truncate,
// 	"urlencode":      urlencode,
// 	"since":          since,
// 	"trimleft":       trimLeft,
// 	"trimright":      trimRight,
// 	"trim":           trim,
// }

// func truncate(s string, len int) string {
// 	if utf8.RuneCountInString(s) <= len {
// 		return s
// 	}
// 	runes := []rune(s)
// 	return string(runes[:len])

// }

// func uppercaseFirst(s string) string {
// 	a := []rune(s)
// 	a[0] = unicode.ToUpper(a[0])
// 	s = string(a)
// 	return s
// }

// func toDuration(started, finished int64) string {
// 	return fmt.Sprintln(time.Duration(finished-started) * time.Second)
// }

// func toDatetime(timestamp int64, layout, zone string) string {
// 	if len(zone) == 0 {
// 		return time.Unix(int64(timestamp), 0).Format(layout)
// 	}
// 	loc, err := time.LoadLocation(zone)
// 	if err != nil {
// 		return time.Unix(int64(timestamp), 0).Local().Format(layout)
// 	}
// 	return time.Unix(int64(timestamp), 0).In(loc).Format(layout)
// }

// func isSuccess(conditional bool, options *raymond.Options) string {
// 	if !conditional {
// 		return options.Inverse()
// 	}

// 	switch options.ParamStr(0) {
// 	case "success":
// 		return options.Fn()
// 	default:
// 		return options.Inverse()
// 	}
// }

// func isFailure(conditional bool, options *raymond.Options) string {
// 	if !conditional {
// 		return options.Inverse()
// 	}

// 	switch options.ParamStr(0) {
// 	case "failure", "error", "killed":
// 		return options.Fn()
// 	default:
// 		return options.Inverse()
// 	}
// }

// func urlencode(options *raymond.Options) string {
// 	return url.QueryEscape(options.Fn())
// }

// func since(start int64) string {
// 	// NOTE: not using `time.Since()` because the fractional second component
// 	// will give us something like "40m12.917523438s" vs "40m12s". We lose
// 	// some precision, but the format is much more readable.
// 	now := time.Unix(time.Now().Unix(), 0)
// 	return fmt.Sprintln(now.Sub(time.Unix(start, 0)))
// }

// func trimLeft(s string, cutset string) string {
// 	return strings.TrimLeft(s, cutset)
// }

// func trimRight(s string, cutset string) string {
// 	return strings.TrimRight(s, cutset)
// }

// func trim(s string, cutset string) string {
// 	return strings.Trim(s, cutset)
// }

// func stripComment(source string) string {
// 	if cut := strings.IndexAny(source, commentChars); cut >= 0 {
// 		return strings.TrimRightFunc(source[:cut], unicode.IsSpace)
// 	}
// 	return source
// }
