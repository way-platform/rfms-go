package rfms

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"math"
	"math/rand"
	"net"
	"net/http"
	"strconv"
	"time"
)

type retryTransport struct {
	maxRetries int
	next       http.RoundTripper
}

func (t *retryTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// if body is present, it must be buffered if there is any chance of a retry
	// since it can only be consumed once.
	var br *bytes.Reader
	if req.Body != nil && req.Body != http.NoBody {
		var buf bytes.Buffer
		if _, err := io.Copy(&buf, req.Body); err != nil {
			req.Body.Close()
			return nil, fmt.Errorf("error buffering body before retry: %w", err)
		}
		req.Body.Close()
		br = bytes.NewReader(buf.Bytes())
		req.Body = io.NopCloser(br)
	}
	var attemptCount int
	for {
		res, err := t.next.RoundTrip(req)
		attemptCount++
		if attemptCount-1 >= t.maxRetries {
			return res, err
		}
		shouldRetry := shouldRetry(err, req, res)
		if !shouldRetry {
			return res, err
		}
		delay := retryDelay(attemptCount, res)
		if br != nil {
			if _, serr := br.Seek(0, 0); serr != nil {
				return res, fmt.Errorf("error seeking body buffer back to beginning after attempt: %w", serr)
			}
			req.Body = io.NopCloser(br)
		}
		if res != nil {
			_, _ = io.Copy(io.Discard, res.Body)
			res.Body.Close()
		}
		if err := sleepWithContext(req.Context(), delay); err != nil {
			return nil, err
		}
	}
}

func isDNSErr(err error) bool {
	var dnse *net.DNSError
	return errors.As(err, &dnse)
}

func isTimeoutErr(err error) bool {
	var netErr net.Error
	return errors.As(err, &netErr) && netErr.Timeout()
}

func shouldRetry(err error, request *http.Request, response *http.Response) bool {
	if err != nil {
		return isDNSErr(err) || (isIdempotent(request) && isTimeoutErr(err))
	}
	if response.Header.Get("Retry-After") != "" {
		return true
	}
	switch response.StatusCode {
	case http.StatusTooManyRequests:
		return true
	case http.StatusBadGateway, http.StatusServiceUnavailable, http.StatusGatewayTimeout, http.StatusInternalServerError:
		return isIdempotent(request)
	default:
		return false
	}
}

func retryDelay(attempt int, response *http.Response) time.Duration {
	if response != nil {
		if retryAfter := response.Header.Get("Retry-After"); retryAfter != "" {
			if i, err := strconv.Atoi(retryAfter); err == nil {
				return addJitter(time.Duration(i) * time.Second)
			}
			if t, err := time.Parse(http.TimeFormat, retryAfter); err == nil {
				return addJitter(time.Until(t))
			}
		}
	}
	return expBackoff(attempt)
}

func expBackoff(attempt int) time.Duration {
	// based on "full jitter": https://aws.amazon.com/blogs/architecture/exponential-backoff-and-jitter/
	const base = time.Millisecond * 250
	const cap = time.Second * 10
	exp := math.Pow(2, float64(attempt-1))
	v := float64(base) * exp
	return time.Duration(
		rand.Int63n(int64(math.Min(float64(cap), v))),
	)
}

func addJitter(d time.Duration) time.Duration {
	const magnitude = 0.333
	f := float64(d)
	mj := f * magnitude
	j := rand.Float64() * mj
	coin := rand.Float64()
	if coin < 0.5 {
		return time.Duration(f + j)
	}
	return time.Duration(f - j)
}

func isIdempotent(req *http.Request) bool {
	if req.Header.Get("Idempotency-Key") != "" || req.Header.Get("X-Idempotency-Key") != "" {
		return true
	}
	switch req.Method {
	case http.MethodGet, http.MethodHead, http.MethodOptions, http.MethodTrace, http.MethodPut, http.MethodDelete:
		return true
	}
	return false
}

func sleepWithContext(ctx context.Context, duration time.Duration) error {
	timer := time.NewTimer(duration)
	defer timer.Stop()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
		return nil
	}
}
