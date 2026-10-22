package main

import (
	"fmt"
	"html"
	"html/template"
	"sort"
	"strings"
)

// The "fields" view (?view=fields on any page) renders the raw API response
// of that page as dense developer tables without a hand-written template:
// top-level scalars as a key/value strip, nested objects as key/value
// lists, and every array of objects as a table whose columns are inferred
// from the keys. Identity columns (evtlog id, block, tx hash, time, round,
// ids, addresses) sort first; each row carries a card with the full JSON.

type fieldKV struct {
	Key   string
	Value template.HTML
}

type fieldTable struct {
	Name string
	Cols []string
	Rows []fieldRow
}

type fieldRow struct {
	Cells []template.HTML
	Card  template.HTML
}

type fieldObject struct {
	Name string
	KV   []fieldKV
}

type fieldsView struct {
	Scalars []fieldKV
	Objects []fieldObject
	Tables  []fieldTable
}

// hiddenKeys are API envelope fields that carry no data.
var hiddenKeys = map[string]bool{"error": true, "status": true}

// colPriority ranks a (flattened) column name: lower sorts first.
func colPriority(name string) int {
	leaf := name
	if i := strings.LastIndexByte(name, '.'); i >= 0 {
		leaf = name[i+1:]
	}
	switch {
	case leaf == "EvtLogId", leaf == "Id", leaf == "RecordId":
		return 0
	case leaf == "BlockNum", leaf == "BlockNumber":
		return 1
	case leaf == "TxHash":
		return 2
	case leaf == "DateTime", leaf == "RealDate":
		return 3
	case leaf == "TimeStamp":
		return 4
	case leaf == "TxId":
		return 5
	case leaf == "RoundNum":
		return 6
	case strings.HasSuffix(leaf, "Id"), strings.HasSuffix(leaf, "Aid"):
		return 7
	case strings.HasSuffix(leaf, "Addr"), strings.HasSuffix(leaf, "Address"):
		return 8
	case strings.HasSuffix(leaf, "Eth"):
		return 9
	case strings.HasPrefix(leaf, "Num"), strings.HasPrefix(leaf, "Total"):
		return 10
	}
	return 20
}

// flatten copies m into out, expanding one level of nested objects as
// "Parent.Child" keys.
func flatten(m map[string]any, out map[string]any) {
	for k, v := range m {
		if hiddenKeys[k] {
			continue
		}
		if sub, ok := v.(map[string]any); ok {
			for sk, sv := range sub {
				out[k+"."+sk] = sv
			}
			continue
		}
		out[k] = v
	}
}

// fieldCell renders one value according to its key name and type.
func fieldCell(project, key string, v any) template.HTML {
	leaf := key
	if i := strings.LastIndexByte(key, '.'); i >= 0 {
		leaf = key[i+1:]
	}
	switch t := v.(type) {
	case nil:
		return `<span class="empty">null</span>`
	case bool:
		if t {
			return `<span class="ok">✓</span>`
		}
		return `<span class="empty">✗</span>`
	case jsonNum:
		switch {
		case leaf == "RoundNum" && project == "cosmicsignature" && t >= 0:
			return roundLink(t)
		case leaf == "TokenId" && t >= 0:
			if project == "randomwalk" {
				return rwTokenLink(t)
			}
			return tokenLink(t)
		case leaf == "TimeStamp" && t > 1e9:
			return template.HTML(fmt.Sprintf(`<span title="%s">%s</span>`, unixDate(t), num(t))) // #nosec G203 -- numeric
		}
		return template.HTML(html.EscapeString(num(t))) // #nosec G203 -- escaped
	case string:
		switch {
		case t == "":
			return `<span class="empty">—</span>`
		case isTxHash(t):
			return txHashCell(t)
		case isEthAddr(t):
			if project == "randomwalk" {
				return hexCell(t)
			}
			return ethAddrLink(t)
		case strings.HasPrefix(t, "0x") && len(t) > 14:
			return hexCell(t)
		case len(t) > 60:
			return template.HTML(fmt.Sprintf(`<span class="trunc" title="%s">%s…</span>`, // #nosec G203 -- escaped
				html.EscapeString(t), html.EscapeString(t[:57])))
		}
		return template.HTML(html.EscapeString(t)) // #nosec G203 -- escaped
	case []any, map[string]any:
		js := jsonPretty(t)
		short := js
		if len(short) > 40 {
			short = short[:37] + "…"
		}
		return template.HTML(fmt.Sprintf(`<span class="hx"><span class="trunc">%s</span><span class="hx-card"><pre>%s</pre></span></span>`, // #nosec G203 -- escaped
			html.EscapeString(short), html.EscapeString(js)))
	}
	return template.HTML(html.EscapeString(fmt.Sprint(v))) // #nosec G203 -- escaped
}

// buildFieldsView classifies the decoded API response.
func buildFieldsView(project string, data map[string]any) fieldsView {
	var fv fieldsView
	keys := make([]string, 0, len(data))
	for k := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		if hiddenKeys[k] || strings.HasPrefix(k, "_") {
			continue
		}
		switch v := data[k].(type) {
		case []any:
			if tbl, ok := buildTable(project, k, v); ok {
				fv.Tables = append(fv.Tables, tbl)
				continue
			}
			fv.Scalars = append(fv.Scalars, fieldKV{k, fieldCell(project, k, v)})
		case map[string]any:
			obj := fieldObject{Name: k}
			flat := map[string]any{}
			flatten(v, flat)
			obj.KV = sortedKV(project, flat)
			fv.Objects = append(fv.Objects, obj)
		default:
			fv.Scalars = append(fv.Scalars, fieldKV{k, fieldCell(project, k, v)})
		}
	}
	return fv
}

func sortedKV(project string, flat map[string]any) []fieldKV {
	keys := make([]string, 0, len(flat))
	for k := range flat {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		pi, pj := colPriority(keys[i]), colPriority(keys[j])
		if pi != pj {
			return pi < pj
		}
		return keys[i] < keys[j]
	})
	out := make([]fieldKV, 0, len(keys))
	for _, k := range keys {
		out = append(out, fieldKV{k, fieldCell(project, k, flat[k])})
	}
	return out
}

// buildTable turns an array of objects into a table; ok is false for arrays
// of scalars (rendered as a plain value instead).
func buildTable(project, name string, items []any) (fieldTable, bool) {
	if len(items) == 0 {
		return fieldTable{Name: name}, true
	}
	flatRows := make([]map[string]any, 0, len(items))
	colSet := map[string]bool{}
	for _, it := range items {
		m, ok := it.(map[string]any)
		if !ok {
			return fieldTable{}, false
		}
		flat := map[string]any{}
		flatten(m, flat)
		for k := range flat {
			colSet[k] = true
		}
		flatRows = append(flatRows, flat)
	}
	cols := make([]string, 0, len(colSet))
	for k := range colSet {
		cols = append(cols, k)
	}
	sort.Slice(cols, func(i, j int) bool {
		pi, pj := colPriority(cols[i]), colPriority(cols[j])
		if pi != pj {
			return pi < pj
		}
		return cols[i] < cols[j]
	})
	tbl := fieldTable{Name: name, Cols: cols}
	for i, flat := range flatRows {
		row := fieldRow{Cells: make([]template.HTML, 0, len(cols)), Card: recordCard(items[i])}
		for _, c := range cols {
			v, ok := flat[c]
			if !ok {
				row.Cells = append(row.Cells, `<span class="empty"></span>`)
				continue
			}
			row.Cells = append(row.Cells, fieldCell(project, c, v))
		}
		tbl.Rows = append(tbl.Rows, row)
	}
	return tbl, true
}
