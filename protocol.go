package gomrjob

import (
	"bufio"
	"io"
	"log"
	"github.com/bitly/go-simplejson"
)

// returns a channel of simplejson.Json objects. This channel will be closed
// when the input stream closes. Errors will be logged
func JsonInputProtocol(input io.Reader) chan *simplejson.Json {
	out := make(chan *simplejson.Json, 100)
	go func() {
		var line []byte
		var err error
		r := bufio.NewReaderSize(input, 1024*1024*2)
		for {
			if err == io.EOF {
				break;
			}
			line, err = r.ReadBytes('\n');
			data, err := simplejson.NewJson(line)
			if err != nil {
				log.Printf("%s - failed parsing %s", err, line)
			} else {
				out <- data
			}
		}
		close(out)
	}()
	return out
}