// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package events

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

func easyjsonCce3d1beDecodeGithubComMailgunMailgunGoEvents(in *jlexer.Lexer, out *Storage) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "key":
			out.Key = string(in.String())
		case "url":
			out.URL = string(in.String())
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
func easyjsonCce3d1beEncodeGithubComMailgunMailgunGoEvents(out *jwriter.Writer, in Storage) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"key\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Key))
	}
	{
		const prefix string = ",\"url\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.URL))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Storage) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCce3d1beEncodeGithubComMailgunMailgunGoEvents(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Storage) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCce3d1beEncodeGithubComMailgunMailgunGoEvents(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Storage) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCce3d1beDecodeGithubComMailgunMailgunGoEvents(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Storage) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCce3d1beDecodeGithubComMailgunMailgunGoEvents(l, v)
}
func easyjsonCce3d1beDecodeGithubComMailgunMailgunGoEvents1(in *jlexer.Lexer, out *MessageHeaders) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "to":
			out.To = string(in.String())
		case "message-id":
			out.MessageID = string(in.String())
		case "from":
			out.From = string(in.String())
		case "subject":
			out.Subject = string(in.String())
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
func easyjsonCce3d1beEncodeGithubComMailgunMailgunGoEvents1(out *jwriter.Writer, in MessageHeaders) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"to\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.To))
	}
	{
		const prefix string = ",\"message-id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.MessageID))
	}
	{
		const prefix string = ",\"from\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.From))
	}
	{
		const prefix string = ",\"subject\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Subject))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v MessageHeaders) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCce3d1beEncodeGithubComMailgunMailgunGoEvents1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v MessageHeaders) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCce3d1beEncodeGithubComMailgunMailgunGoEvents1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *MessageHeaders) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCce3d1beDecodeGithubComMailgunMailgunGoEvents1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *MessageHeaders) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCce3d1beDecodeGithubComMailgunMailgunGoEvents1(l, v)
}
func easyjsonCce3d1beDecodeGithubComMailgunMailgunGoEvents2(in *jlexer.Lexer, out *Message) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "headers":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Headers).UnmarshalJSON(data))
			}
		case "attachments":
			if in.IsNull() {
				in.Skip()
				out.Attachments = nil
			} else {
				in.Delim('[')
				if out.Attachments == nil {
					if !in.IsDelim(']') {
						out.Attachments = make([]Attachment, 0, 1)
					} else {
						out.Attachments = []Attachment{}
					}
				} else {
					out.Attachments = (out.Attachments)[:0]
				}
				for !in.IsDelim(']') {
					var v1 Attachment
					if data := in.Raw(); in.Ok() {
						in.AddError((v1).UnmarshalJSON(data))
					}
					out.Attachments = append(out.Attachments, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "recipients":
			if in.IsNull() {
				in.Skip()
				out.Recipients = nil
			} else {
				in.Delim('[')
				if out.Recipients == nil {
					if !in.IsDelim(']') {
						out.Recipients = make([]string, 0, 4)
					} else {
						out.Recipients = []string{}
					}
				} else {
					out.Recipients = (out.Recipients)[:0]
				}
				for !in.IsDelim(']') {
					var v2 string
					v2 = string(in.String())
					out.Recipients = append(out.Recipients, v2)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "size":
			out.Size = int(in.Int())
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
func easyjsonCce3d1beEncodeGithubComMailgunMailgunGoEvents2(out *jwriter.Writer, in Message) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"headers\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.Headers).MarshalJSON())
	}
	{
		const prefix string = ",\"attachments\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Attachments == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v3, v4 := range in.Attachments {
				if v3 > 0 {
					out.RawByte(',')
				}
				out.Raw((v4).MarshalJSON())
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"recipients\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Recipients == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Recipients {
				if v5 > 0 {
					out.RawByte(',')
				}
				out.String(string(v6))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"size\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Size))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Message) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCce3d1beEncodeGithubComMailgunMailgunGoEvents2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Message) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCce3d1beEncodeGithubComMailgunMailgunGoEvents2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Message) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCce3d1beDecodeGithubComMailgunMailgunGoEvents2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Message) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCce3d1beDecodeGithubComMailgunMailgunGoEvents2(l, v)
}
func easyjsonCce3d1beDecodeGithubComMailgunMailgunGoEvents3(in *jlexer.Lexer, out *MailingList) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "address":
			out.Address = string(in.String())
		case "list-id":
			out.ListID = string(in.String())
		case "sid":
			out.SID = string(in.String())
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
func easyjsonCce3d1beEncodeGithubComMailgunMailgunGoEvents3(out *jwriter.Writer, in MailingList) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"address\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Address))
	}
	{
		const prefix string = ",\"list-id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ListID))
	}
	{
		const prefix string = ",\"sid\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.SID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v MailingList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCce3d1beEncodeGithubComMailgunMailgunGoEvents3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v MailingList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCce3d1beEncodeGithubComMailgunMailgunGoEvents3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *MailingList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCce3d1beDecodeGithubComMailgunMailgunGoEvents3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *MailingList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCce3d1beDecodeGithubComMailgunMailgunGoEvents3(l, v)
}
func easyjsonCce3d1beDecodeGithubComMailgunMailgunGoEvents4(in *jlexer.Lexer, out *GeoLocation) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "city":
			out.City = string(in.String())
		case "country":
			out.Country = string(in.String())
		case "region":
			out.Region = string(in.String())
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
func easyjsonCce3d1beEncodeGithubComMailgunMailgunGoEvents4(out *jwriter.Writer, in GeoLocation) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"city\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.City))
	}
	{
		const prefix string = ",\"country\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Country))
	}
	{
		const prefix string = ",\"region\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Region))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GeoLocation) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCce3d1beEncodeGithubComMailgunMailgunGoEvents4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GeoLocation) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCce3d1beEncodeGithubComMailgunMailgunGoEvents4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GeoLocation) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCce3d1beDecodeGithubComMailgunMailgunGoEvents4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GeoLocation) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCce3d1beDecodeGithubComMailgunMailgunGoEvents4(l, v)
}
func easyjsonCce3d1beDecodeGithubComMailgunMailgunGoEvents5(in *jlexer.Lexer, out *Flags) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "is-authenticated":
			out.IsAuthenticated = bool(in.Bool())
		case "is-big":
			out.IsBig = bool(in.Bool())
		case "is-system-test":
			out.IsSystemTest = bool(in.Bool())
		case "is-test-mode":
			out.IsTestMode = bool(in.Bool())
		case "is-delayed-bounce":
			out.IsDelayedBounce = bool(in.Bool())
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
func easyjsonCce3d1beEncodeGithubComMailgunMailgunGoEvents5(out *jwriter.Writer, in Flags) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"is-authenticated\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.IsAuthenticated))
	}
	{
		const prefix string = ",\"is-big\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.IsBig))
	}
	{
		const prefix string = ",\"is-system-test\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.IsSystemTest))
	}
	{
		const prefix string = ",\"is-test-mode\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.IsTestMode))
	}
	{
		const prefix string = ",\"is-delayed-bounce\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.IsDelayedBounce))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Flags) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCce3d1beEncodeGithubComMailgunMailgunGoEvents5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Flags) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCce3d1beEncodeGithubComMailgunMailgunGoEvents5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Flags) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCce3d1beDecodeGithubComMailgunMailgunGoEvents5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Flags) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCce3d1beDecodeGithubComMailgunMailgunGoEvents5(l, v)
}
func easyjsonCce3d1beDecodeGithubComMailgunMailgunGoEvents6(in *jlexer.Lexer, out *Envelope) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "mail-from":
			out.MailFrom = string(in.String())
		case "sender":
			out.Sender = string(in.String())
		case "transport":
			out.Transport = string(in.String())
		case "targets":
			out.Targets = string(in.String())
		case "sending-host":
			out.SendingHost = string(in.String())
		case "sending-ip":
			out.SendingIP = string(in.String())
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
func easyjsonCce3d1beEncodeGithubComMailgunMailgunGoEvents6(out *jwriter.Writer, in Envelope) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"mail-from\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.MailFrom))
	}
	{
		const prefix string = ",\"sender\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Sender))
	}
	{
		const prefix string = ",\"transport\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Transport))
	}
	{
		const prefix string = ",\"targets\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Targets))
	}
	{
		const prefix string = ",\"sending-host\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.SendingHost))
	}
	{
		const prefix string = ",\"sending-ip\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.SendingIP))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Envelope) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCce3d1beEncodeGithubComMailgunMailgunGoEvents6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Envelope) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCce3d1beEncodeGithubComMailgunMailgunGoEvents6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Envelope) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCce3d1beDecodeGithubComMailgunMailgunGoEvents6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Envelope) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCce3d1beDecodeGithubComMailgunMailgunGoEvents6(l, v)
}
func easyjsonCce3d1beDecodeGithubComMailgunMailgunGoEvents7(in *jlexer.Lexer, out *DeliveryStatus) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "code":
			out.Code = int(in.Int())
		case "attempt-no":
			out.AttemptNo = int(in.Int())
		case "description":
			out.Description = string(in.String())
		case "message":
			out.Message = string(in.String())
		case "session-seconds":
			out.SessionSeconds = float64(in.Float64())
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
func easyjsonCce3d1beEncodeGithubComMailgunMailgunGoEvents7(out *jwriter.Writer, in DeliveryStatus) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"code\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Code))
	}
	{
		const prefix string = ",\"attempt-no\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.AttemptNo))
	}
	{
		const prefix string = ",\"description\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Description))
	}
	{
		const prefix string = ",\"message\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Message))
	}
	{
		const prefix string = ",\"session-seconds\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.SessionSeconds))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DeliveryStatus) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCce3d1beEncodeGithubComMailgunMailgunGoEvents7(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DeliveryStatus) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCce3d1beEncodeGithubComMailgunMailgunGoEvents7(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DeliveryStatus) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCce3d1beDecodeGithubComMailgunMailgunGoEvents7(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DeliveryStatus) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCce3d1beDecodeGithubComMailgunMailgunGoEvents7(l, v)
}
func easyjsonCce3d1beDecodeGithubComMailgunMailgunGoEvents8(in *jlexer.Lexer, out *ClientInfo) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "accept-language":
			out.AcceptLanguage = string(in.String())
		case "client-name":
			out.ClientName = string(in.String())
		case "client-os":
			out.ClientOS = string(in.String())
		case "client-type":
			out.ClientType = string(in.String())
		case "device-type":
			out.DeviceType = string(in.String())
		case "ip":
			out.IP = string(in.String())
		case "user-agent":
			out.UserAgent = string(in.String())
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
func easyjsonCce3d1beEncodeGithubComMailgunMailgunGoEvents8(out *jwriter.Writer, in ClientInfo) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"accept-language\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.AcceptLanguage))
	}
	{
		const prefix string = ",\"client-name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ClientName))
	}
	{
		const prefix string = ",\"client-os\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ClientOS))
	}
	{
		const prefix string = ",\"client-type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ClientType))
	}
	{
		const prefix string = ",\"device-type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.DeviceType))
	}
	{
		const prefix string = ",\"ip\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.IP))
	}
	{
		const prefix string = ",\"user-agent\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.UserAgent))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ClientInfo) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCce3d1beEncodeGithubComMailgunMailgunGoEvents8(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ClientInfo) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCce3d1beEncodeGithubComMailgunMailgunGoEvents8(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ClientInfo) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCce3d1beDecodeGithubComMailgunMailgunGoEvents8(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ClientInfo) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCce3d1beDecodeGithubComMailgunMailgunGoEvents8(l, v)
}
func easyjsonCce3d1beDecodeGithubComMailgunMailgunGoEvents9(in *jlexer.Lexer, out *Campaign) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = string(in.String())
		case "name":
			out.Name = string(in.String())
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
func easyjsonCce3d1beEncodeGithubComMailgunMailgunGoEvents9(out *jwriter.Writer, in Campaign) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ID))
	}
	{
		const prefix string = ",\"name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Name))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Campaign) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCce3d1beEncodeGithubComMailgunMailgunGoEvents9(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Campaign) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCce3d1beEncodeGithubComMailgunMailgunGoEvents9(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Campaign) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCce3d1beDecodeGithubComMailgunMailgunGoEvents9(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Campaign) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCce3d1beDecodeGithubComMailgunMailgunGoEvents9(l, v)
}
func easyjsonCce3d1beDecodeGithubComMailgunMailgunGoEvents10(in *jlexer.Lexer, out *Attachment) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "filename":
			out.FileName = string(in.String())
		case "content-type":
			out.ContentType = string(in.String())
		case "size":
			out.Size = int(in.Int())
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
func easyjsonCce3d1beEncodeGithubComMailgunMailgunGoEvents10(out *jwriter.Writer, in Attachment) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"filename\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.FileName))
	}
	{
		const prefix string = ",\"content-type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ContentType))
	}
	{
		const prefix string = ",\"size\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Size))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Attachment) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCce3d1beEncodeGithubComMailgunMailgunGoEvents10(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Attachment) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCce3d1beEncodeGithubComMailgunMailgunGoEvents10(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Attachment) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCce3d1beDecodeGithubComMailgunMailgunGoEvents10(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Attachment) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCce3d1beDecodeGithubComMailgunMailgunGoEvents10(l, v)
}
