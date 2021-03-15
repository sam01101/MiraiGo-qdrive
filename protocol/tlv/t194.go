package tlv

import "github.com/sam01101/MiraiGo-qdrive/binary"

func T194(imsiMd5 []byte) []byte {
	return binary.NewWriterF(func(w *binary.Writer) {
		w.WriteUInt16(0x194)
		w.WriteTlv(binary.NewWriterF(func(w *binary.Writer) {
			w.Write(imsiMd5)
		}))
	})
}
