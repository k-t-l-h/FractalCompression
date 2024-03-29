// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package config

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonC80ae7adDecodeFractalCompressionInternalConfig(in *jlexer.Lexer, out *TableConfig) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "k":
			out.K = uint64(in.Uint64())
		case "name":
			out.Name = string(in.String())
		case "strategy":
			out.Strategy = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC80ae7adEncodeFractalCompressionInternalConfig(out *jwriter.Writer, in TableConfig) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"k\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.K))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"strategy\":"
		out.RawString(prefix)
		out.String(string(in.Strategy))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v TableConfig) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC80ae7adEncodeFractalCompressionInternalConfig(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v TableConfig) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC80ae7adEncodeFractalCompressionInternalConfig(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *TableConfig) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC80ae7adDecodeFractalCompressionInternalConfig(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *TableConfig) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC80ae7adDecodeFractalCompressionInternalConfig(l, v)
}
func easyjsonC80ae7adDecodeFractalCompressionInternalConfig1(in *jlexer.Lexer, out *KeyConfig) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "name":
			out.Name = string(in.String())
		case "users":
			out.Users = bool(in.Bool())
		case "len":
			out.Len = uint64(in.Uint64())
		case "script":
			out.Script = string(in.String())
		case "type":
			out.Type = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC80ae7adEncodeFractalCompressionInternalConfig1(out *jwriter.Writer, in KeyConfig) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix[1:])
		out.String(string(in.Name))
	}
	if in.Users {
		const prefix string = ",\"users\":"
		out.RawString(prefix)
		out.Bool(bool(in.Users))
	}
	if in.Len != 0 {
		const prefix string = ",\"len\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.Len))
	}
	if in.Script != "" {
		const prefix string = ",\"script\":"
		out.RawString(prefix)
		out.String(string(in.Script))
	}
	{
		const prefix string = ",\"type\":"
		out.RawString(prefix)
		out.String(string(in.Type))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v KeyConfig) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC80ae7adEncodeFractalCompressionInternalConfig1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v KeyConfig) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC80ae7adEncodeFractalCompressionInternalConfig1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *KeyConfig) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC80ae7adDecodeFractalCompressionInternalConfig1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *KeyConfig) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC80ae7adDecodeFractalCompressionInternalConfig1(l, v)
}
func easyjsonC80ae7adDecodeFractalCompressionInternalConfig2(in *jlexer.Lexer, out *DatabaseConfig) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "host":
			out.Host = string(in.String())
		case "user":
			out.User = string(in.String())
		case "password":
			out.Password = string(in.String())
		case "database":
			out.Database = string(in.String())
		case "port":
			out.Port = uint(in.Uint())
		case "max_conn":
			out.MaxConn = uint(in.Uint())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC80ae7adEncodeFractalCompressionInternalConfig2(out *jwriter.Writer, in DatabaseConfig) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"host\":"
		out.RawString(prefix[1:])
		out.String(string(in.Host))
	}
	{
		const prefix string = ",\"user\":"
		out.RawString(prefix)
		out.String(string(in.User))
	}
	{
		const prefix string = ",\"password\":"
		out.RawString(prefix)
		out.String(string(in.Password))
	}
	{
		const prefix string = ",\"database\":"
		out.RawString(prefix)
		out.String(string(in.Database))
	}
	{
		const prefix string = ",\"port\":"
		out.RawString(prefix)
		out.Uint(uint(in.Port))
	}
	{
		const prefix string = ",\"max_conn\":"
		out.RawString(prefix)
		out.Uint(uint(in.MaxConn))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DatabaseConfig) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC80ae7adEncodeFractalCompressionInternalConfig2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DatabaseConfig) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC80ae7adEncodeFractalCompressionInternalConfig2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DatabaseConfig) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC80ae7adDecodeFractalCompressionInternalConfig2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DatabaseConfig) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC80ae7adDecodeFractalCompressionInternalConfig2(l, v)
}
func easyjsonC80ae7adDecodeFractalCompressionInternalConfig3(in *jlexer.Lexer, out *CompressionConfig) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "database":
			(out.DC).UnmarshalEasyJSON(in)
		case "table":
			(out.TC).UnmarshalEasyJSON(in)
		case "key":
			(out.KC).UnmarshalEasyJSON(in)
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC80ae7adEncodeFractalCompressionInternalConfig3(out *jwriter.Writer, in CompressionConfig) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"database\":"
		out.RawString(prefix[1:])
		(in.DC).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"table\":"
		out.RawString(prefix)
		(in.TC).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"key\":"
		out.RawString(prefix)
		(in.KC).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CompressionConfig) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC80ae7adEncodeFractalCompressionInternalConfig3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CompressionConfig) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC80ae7adEncodeFractalCompressionInternalConfig3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CompressionConfig) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC80ae7adDecodeFractalCompressionInternalConfig3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CompressionConfig) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC80ae7adDecodeFractalCompressionInternalConfig3(l, v)
}
