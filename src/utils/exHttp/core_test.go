package exHttp

import (
	"context"
	"github.com/bytedance/mockey"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"
)

func TestSwitchVerbose(t *testing.T) {
	r := NewHttpRequest(context.Background(), "https://test.com", "", WithTimeout(time.Millisecond))
	r.SwitchVerbose(false)
	assert.False(t, r.verbose)
	r.SwitchVerbose(true)
	assert.True(t, r.verbose)
}

func TestFixHost1(t *testing.T) {
	host := fixHost("https://a.b.com/")
	assert.True(t, host == "https://a.b.com")
}

func TestFixHost2(t *testing.T) {
	host := fixHost("https://a.b.com")
	assert.True(t, host == "https://a.b.com")
}

func TestFixHost3(t *testing.T) {
	host := fixHost("")
	assert.True(t, host == "")
}

func TestFixRelativeUrl1(t *testing.T) {
	host := fixRelativeUrl("name/method/")
	assert.True(t, host == "/name/method/")
}

func TestFixRelativeUrl2(t *testing.T) {
	host := fixRelativeUrl("/name/method/")
	assert.True(t, host == "/name/method/")
}

func TestFixRelativeUrl3(t *testing.T) {
	host := fixRelativeUrl("")
	assert.True(t, host == "")
}

func TestWithTimeoutAndHeader(t *testing.T) {
	ctx := context.Background()
	timeoutCtx, _ := context.WithTimeout(ctx, time.Millisecond*300)
	r := NewHttpRequest(
		timeoutCtx,
		"https://test.com",
		"",
		WithHeaders(map[string]string{"Content-type": "application/json"}),
		WithRequestBody(`{"name":"test"}`),
		WithTimeout(time.Millisecond),
	)

	_, err := r.Post()
	assert.True(t, err != nil)
}

func TestGenHttpRequestFailed(t *testing.T) {
	r := NewHttpRequest(
		context.Background(),
		"https://test.com",
		"",
		WithTimeout(time.Millisecond),
	)

	mockey.PatchConvey(t.Name(), t, func() {
		mockey.Mock(r.genHttpRequest).Return(nil, errors.New("err")).Build()

		_, err := r.Post()
		assert.True(t, err.Error() != "")
	})
}

func TestDoRequest(t *testing.T) {
	ctx := context.Background()
	r := NewHttpRequest(
		ctx,
		"https://a.b.com",
		"",
		WithTimeout(time.Second),
	)

	content := `{"test":1}`
	httpClient := &http.Client{}

	mockey.PatchConvey(t.Name(), t, func() {
		mockey.Mock(mockey.GetMethod(httpClient, "do")).Return(&http.Response{
			Body: io.NopCloser(strings.NewReader(content)),
		}, nil).Build()

		resp, err := r.Get()
		assert.True(t, err == nil)
		assert.True(t, string(resp) == content)
	})
}
