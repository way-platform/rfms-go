package rfms

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"os"
)

type debugTransport struct {
	next http.RoundTripper
}

func (t *debugTransport) RoundTrip(request *http.Request) (*http.Response, error) {
	requestDump, err := httputil.DumpRequestOut(request, true)
	if err != nil {
		return nil, fmt.Errorf("failed to dump request for debug: %w", err)
	}
	prettyPrintDump(os.Stderr, requestDump, "> ")
	response, err := t.next.RoundTrip(request)
	if err != nil {
		return nil, err
	}
	responseDump, err := httputil.DumpResponse(response, true)
	if err != nil {
		return nil, fmt.Errorf("failed to dump response for debug: %w", err)
	}
	prettyPrintDump(os.Stderr, responseDump, "< ")
	return response, nil
}

func prettyPrintDump(w io.Writer, dump []byte, prefix string) {
	var output bytes.Buffer
	output.Grow(len(dump) * 2)
	for line := range bytes.Lines(dump) {
		output.WriteString(prefix)
		output.Write(line)
	}
	output.WriteByte('\n')
	_, _ = w.Write(output.Bytes())
}
