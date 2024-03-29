// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package model

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

func easyjson9e1087fdDecodeProjectInternalModel(in *jlexer.Lexer, out *UserContact) {
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
		case "id_user":
			out.IdUser = uint64(in.Uint64())
		case "id_contact":
			out.IdContact = uint64(in.Uint64())
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
func easyjson9e1087fdEncodeProjectInternalModel(out *jwriter.Writer, in UserContact) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id_user\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.IdUser))
	}
	{
		const prefix string = ",\"id_contact\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.IdContact))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UserContact) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9e1087fdEncodeProjectInternalModel(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserContact) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9e1087fdEncodeProjectInternalModel(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserContact) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9e1087fdDecodeProjectInternalModel(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserContact) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9e1087fdDecodeProjectInternalModel(l, v)
}
func easyjson9e1087fdDecodeProjectInternalModel1(in *jlexer.Lexer, out *User) {
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
		case "id":
			out.Id = uint64(in.Uint64())
		case "username":
			out.Username = string(in.String())
		case "nickname":
			out.Nickname = string(in.String())
		case "email":
			out.Email = string(in.String())
		case "status":
			out.Status = string(in.String())
		case "avatar":
			out.Avatar = string(in.String())
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
func easyjson9e1087fdEncodeProjectInternalModel1(out *jwriter.Writer, in User) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.Id))
	}
	{
		const prefix string = ",\"username\":"
		out.RawString(prefix)
		out.String(string(in.Username))
	}
	{
		const prefix string = ",\"nickname\":"
		out.RawString(prefix)
		out.String(string(in.Nickname))
	}
	{
		const prefix string = ",\"email\":"
		out.RawString(prefix)
		out.String(string(in.Email))
	}
	{
		const prefix string = ",\"status\":"
		out.RawString(prefix)
		out.String(string(in.Status))
	}
	{
		const prefix string = ",\"avatar\":"
		out.RawString(prefix)
		out.String(string(in.Avatar))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v User) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9e1087fdEncodeProjectInternalModel1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v User) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9e1087fdEncodeProjectInternalModel1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *User) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9e1087fdDecodeProjectInternalModel1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *User) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9e1087fdDecodeProjectInternalModel1(l, v)
}
func easyjson9e1087fdDecodeProjectInternalModel2(in *jlexer.Lexer, out *UpdateUser) {
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
		case "email":
			out.Email = string(in.String())
		case "new_avatar_url":
			out.NewAvatarUrl = string(in.String())
		case "nickname":
			out.Nickname = string(in.String())
		case "status":
			out.Status = string(in.String())
		case "current_password":
			out.CurrentPassword = string(in.String())
		case "new_password":
			out.NewPassword = string(in.String())
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
func easyjson9e1087fdEncodeProjectInternalModel2(out *jwriter.Writer, in UpdateUser) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"email\":"
		out.RawString(prefix[1:])
		out.String(string(in.Email))
	}
	{
		const prefix string = ",\"new_avatar_url\":"
		out.RawString(prefix)
		out.String(string(in.NewAvatarUrl))
	}
	{
		const prefix string = ",\"nickname\":"
		out.RawString(prefix)
		out.String(string(in.Nickname))
	}
	{
		const prefix string = ",\"status\":"
		out.RawString(prefix)
		out.String(string(in.Status))
	}
	{
		const prefix string = ",\"current_password\":"
		out.RawString(prefix)
		out.String(string(in.CurrentPassword))
	}
	{
		const prefix string = ",\"new_password\":"
		out.RawString(prefix)
		out.String(string(in.NewPassword))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UpdateUser) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9e1087fdEncodeProjectInternalModel2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UpdateUser) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9e1087fdEncodeProjectInternalModel2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UpdateUser) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9e1087fdDecodeProjectInternalModel2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UpdateUser) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9e1087fdDecodeProjectInternalModel2(l, v)
}
func easyjson9e1087fdDecodeProjectInternalModel3(in *jlexer.Lexer, out *RegistrationUser) {
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
		case "nickname":
			out.Nickname = string(in.String())
		case "email":
			out.Email = string(in.String())
		case "password":
			out.Password = string(in.String())
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
func easyjson9e1087fdEncodeProjectInternalModel3(out *jwriter.Writer, in RegistrationUser) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"nickname\":"
		out.RawString(prefix[1:])
		out.String(string(in.Nickname))
	}
	{
		const prefix string = ",\"email\":"
		out.RawString(prefix)
		out.String(string(in.Email))
	}
	{
		const prefix string = ",\"password\":"
		out.RawString(prefix)
		out.String(string(in.Password))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v RegistrationUser) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9e1087fdEncodeProjectInternalModel3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RegistrationUser) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9e1087fdEncodeProjectInternalModel3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RegistrationUser) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9e1087fdDecodeProjectInternalModel3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RegistrationUser) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9e1087fdDecodeProjectInternalModel3(l, v)
}
func easyjson9e1087fdDecodeProjectInternalModel4(in *jlexer.Lexer, out *LoginUser) {
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
		case "email":
			out.Email = string(in.String())
		case "password":
			out.Password = string(in.String())
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
func easyjson9e1087fdEncodeProjectInternalModel4(out *jwriter.Writer, in LoginUser) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"email\":"
		out.RawString(prefix[1:])
		out.String(string(in.Email))
	}
	{
		const prefix string = ",\"password\":"
		out.RawString(prefix)
		out.String(string(in.Password))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v LoginUser) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9e1087fdEncodeProjectInternalModel4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v LoginUser) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9e1087fdEncodeProjectInternalModel4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *LoginUser) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9e1087fdDecodeProjectInternalModel4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *LoginUser) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9e1087fdDecodeProjectInternalModel4(l, v)
}
func easyjson9e1087fdDecodeProjectInternalModel5(in *jlexer.Lexer, out *Contacts) {
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
		case "contacts":
			if in.IsNull() {
				in.Skip()
				out.Contacts = nil
			} else {
				in.Delim('[')
				if out.Contacts == nil {
					if !in.IsDelim(']') {
						out.Contacts = make([]User, 0, 0)
					} else {
						out.Contacts = []User{}
					}
				} else {
					out.Contacts = (out.Contacts)[:0]
				}
				for !in.IsDelim(']') {
					var v1 User
					(v1).UnmarshalEasyJSON(in)
					out.Contacts = append(out.Contacts, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
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
func easyjson9e1087fdEncodeProjectInternalModel5(out *jwriter.Writer, in Contacts) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"contacts\":"
		out.RawString(prefix[1:])
		if in.Contacts == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Contacts {
				if v2 > 0 {
					out.RawByte(',')
				}
				(v3).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Contacts) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9e1087fdEncodeProjectInternalModel5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Contacts) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9e1087fdEncodeProjectInternalModel5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Contacts) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9e1087fdDecodeProjectInternalModel5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Contacts) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9e1087fdDecodeProjectInternalModel5(l, v)
}
func easyjson9e1087fdDecodeProjectInternalModel6(in *jlexer.Lexer, out *AuthorizedUser) {
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
		case "id":
			out.Id = uint64(in.Uint64())
		case "avatar":
			out.Avatar = string(in.String())
		case "username":
			out.Username = string(in.String())
		case "nickname":
			out.Nickname = string(in.String())
		case "email":
			out.Email = string(in.String())
		case "status":
			out.Status = string(in.String())
		case "password":
			out.Password = string(in.String())
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
func easyjson9e1087fdEncodeProjectInternalModel6(out *jwriter.Writer, in AuthorizedUser) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.Id))
	}
	{
		const prefix string = ",\"avatar\":"
		out.RawString(prefix)
		out.String(string(in.Avatar))
	}
	{
		const prefix string = ",\"username\":"
		out.RawString(prefix)
		out.String(string(in.Username))
	}
	{
		const prefix string = ",\"nickname\":"
		out.RawString(prefix)
		out.String(string(in.Nickname))
	}
	{
		const prefix string = ",\"email\":"
		out.RawString(prefix)
		out.String(string(in.Email))
	}
	{
		const prefix string = ",\"status\":"
		out.RawString(prefix)
		out.String(string(in.Status))
	}
	{
		const prefix string = ",\"password\":"
		out.RawString(prefix)
		out.String(string(in.Password))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v AuthorizedUser) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9e1087fdEncodeProjectInternalModel6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v AuthorizedUser) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9e1087fdEncodeProjectInternalModel6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *AuthorizedUser) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9e1087fdDecodeProjectInternalModel6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *AuthorizedUser) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9e1087fdDecodeProjectInternalModel6(l, v)
}
