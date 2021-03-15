package tlv

import "github.com/sam01101/MiraiGo-qdrive/binary"

func T401(d []byte) []byte {
	return binary.NewWriterF(func(w *binary.Writer) {
		w.WriteUInt16(0x401)
		w.WriteTlv(d)
	})
}
