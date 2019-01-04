//Package Amc comment
// This file war generated by tars2go 1.1
// Generated from Go.tars
package Amc

import (
	"fmt"
	"github.com/TarsCloud/TarsGo/tars/protocol/codec"
)

//GetTimeRsp strcut implement
type GetTimeRsp struct {
	Utc_timestamp   int64  `json:"utc_timestamp"`
	Local_timestamp int64  `json:"local_timestamp"`
	Local_time_str  string `json:"local_time_str"`
}

func (st *GetTimeRsp) resetDefault() {
}

//ReadFrom reads  from _is and put into struct.
func (st *GetTimeRsp) ReadFrom(_is *codec.Reader) error {
	var err error
	var length int32
	var have bool
	var ty byte
	st.resetDefault()

	err = _is.Read_int64(&st.Utc_timestamp, 0, true)
	if err != nil {
		return err
	}

	err = _is.Read_int64(&st.Local_timestamp, 1, true)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.Local_time_str, 2, true)
	if err != nil {
		return err
	}

	_ = length
	_ = have
	_ = ty
	return nil
}

//ReadBlock reads struct from the given tag , require or optional.
func (st *GetTimeRsp) ReadBlock(_is *codec.Reader, tag byte, require bool) error {
	var err error
	var have bool
	st.resetDefault()

	err, have = _is.SkipTo(codec.STRUCT_BEGIN, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require GetTimeRsp, but not exist. tag %d", tag)
		}
		return nil

	}

	st.ReadFrom(_is)

	err = _is.SkipToStructEnd()
	if err != nil {
		return err
	}
	_ = have
	return nil
}

//WriteTo encode struct to buffer
func (st *GetTimeRsp) WriteTo(_os *codec.Buffer) error {
	var err error

	err = _os.Write_int64(st.Utc_timestamp, 0)
	if err != nil {
		return err
	}

	err = _os.Write_int64(st.Local_timestamp, 1)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.Local_time_str, 2)
	if err != nil {
		return err
	}

	return nil
}

//WriteBlock encode struct
func (st *GetTimeRsp) WriteBlock(_os *codec.Buffer, tag byte) error {
	var err error
	err = _os.WriteHead(codec.STRUCT_BEGIN, tag)
	if err != nil {
		return err
	}

	st.WriteTo(_os)

	err = _os.WriteHead(codec.STRUCT_END, 0)
	if err != nil {
		return err
	}
	return nil
}