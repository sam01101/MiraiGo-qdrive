package tlv

import (
	"crypto/md5"

	"github.com/sam01101/MiraiGo-qdrive/binary"
)

func T188(androidId []byte) []byte {
	return binary.NewWriterF(func(w *binary.Writer) {
		w.WriteUInt16(0x188)
		w.WriteTlv(binary.NewWriterF(func(w *binary.Writer) {
			h := md5.Sum(androidId)
			w.Write(h[:])
		}))
	})
}
