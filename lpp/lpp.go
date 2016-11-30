// lpp
package lpp

import (
	"bufio"

	"io"
	"os"
)

type Block_Reading struct {
	File     string
	Blocktag string
	Buffer   int
}
type IO struct {
	Io          *bufio.Reader
	BlockTag    []byte
	SplitTag    byte
	BlockString string
}

func (blockreading *Block_Reading) Read() (IO, error) {
	BlockIO := IO{}

	raw_file, err := os.Open(blockreading.File)
	if blockreading.Buffer == 0 {
		blockreading.Buffer = 99999999999
	}
	BlockIO.Io = bufio.NewReaderSize(raw_file, blockreading.Buffer)
	if blockreading.Blocktag == "" {
		BlockIO.BlockTag = []byte("\n")
	} else {
		BlockIO.BlockTag = []byte(blockreading.Blocktag)
	}
	BlockIO.SplitTag = byte([]byte(blockreading.Blocktag)[len(blockreading.Blocktag)-1])
	BlockIO.BlockString = blockreading.Blocktag
	return BlockIO, err

}

func (Reader IO) Next() ([]byte, error) {

	var out_tag []byte
	var status error

	for {
		line, err := Reader.Io.ReadSlice(Reader.SplitTag)
		if err == nil {
			if len(out_tag) > 1 {
				out_tag = append(out_tag, line...)
			} else {
				out_tag = line
			}

			if len(Reader.BlockTag) > 1 {
				if len(out_tag) >= len(Reader.BlockTag) && string(out_tag[(len(out_tag)-len(Reader.BlockTag)):]) == Reader.BlockString {

					break
				}

			} else {
				break
			}

		} else if err == bufio.ErrBufferFull || err == io.EOF {
			out_tag = append(out_tag, line...)
			if err == io.EOF {
				status = err
				break
			}

		}

	}

	return out_tag, status

}
