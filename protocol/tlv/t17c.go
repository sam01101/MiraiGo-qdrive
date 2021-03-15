package tlv

import "github.com/sam01101/MiraiGo-qdrive/binary"

func T17C(code string) []byte {
	return binary.NewWriterF(func(w *binary.Writer) {
		w.WriteUInt16(0x17c)
		w.WriteTlv(binary.NewWriterF(func(w *binary.Writer) {
			w.WriteStringShort(code)
		}))
	})
}
