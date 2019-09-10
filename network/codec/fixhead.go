package codec

import (
	"encoding/binary"
	"errors"
)

//ErrFixHeadConnIsNull self define error
var ErrFixHeadConnIsNull = errors.New("FixHeadProtocol:conn is nil")

//ErrFixHeadTooLargePacket self define error
var ErrFixHeadTooLargePacket = errors.New("FixHeadProtocol:too large packet")

//ErrFixHeadUnknownMsgInterface self define error
var ErrFixHeadUnknownMsgInterface = errors.New("FixHeadProtocol:unknown msg interface")

// ErrFixHeadSizeIsNotEnough self define error
var ErrFixHeadSizeIsNotEnough = errors.New("FixHead read message less than package head size ")

// DefaultFixHeadWord page const
var DefaultFixHeadWord = 4

// FixHeadDecoder fix head decoder
func FixHeadDecoder(isLittleEndian bool, headWord int, buf []byte) int {
	switch headWord {
	case 1:
		return int(buf[0])
	case 2:
		if isLittleEndian {
			return int(binary.LittleEndian.Uint16(buf))
		}
		return int(binary.BigEndian.Uint16(buf))
	case 4:
		if isLittleEndian {
			return int(binary.LittleEndian.Uint32(buf))
		}
		return int(binary.BigEndian.Uint32(buf))
	default:
		panic("fixHeadDecoder: unsupported head size")
	}
}

// FixHeadEncoder fix headEncoder
func FixHeadEncoder(isLittleEndian bool, headWord int, buf []byte, len int) {
	switch headWord {
	case 1:
		buf[0] = byte(len)
	case 2:
		if isLittleEndian {
			binary.LittleEndian.PutUint16(buf, uint16(len))
		} else {
			binary.BigEndian.PutUint16(buf, uint16(len))
		}

	case 4:
		if isLittleEndian {
			binary.LittleEndian.PutUint32(buf, uint32(len))
		} else {
			binary.BigEndian.PutUint32(buf, uint32(len))
		}
	default:
		panic("fixHeadEncoder: unsupported head size")
	}
}
