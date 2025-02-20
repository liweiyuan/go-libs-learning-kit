package mock

import (
	"io"
	"net/http"
	"testing"

	"github.com/h2non/gock"
	"github.com/nbio/st"
)

// TestSimpleGoMock 测试使用 gock 库进行 HTTP 请求的模拟。
// 该测试创建一个模拟的 HTTP 服务器，响应特定的 GET 请求，并验证返回的状态码和响应体内容。
// 最后，确保没有未完成的模拟请求。
func TestSimpleGoMock(t *testing.T) {
	defer gock.Off()

	gock.New("http://example.com").
		Get("/bar").
		Reply(200).
		JSON(map[string]string{"foo": "bar"})

	// Perform the HTTP request
	resp, err := http.Get("http://example.com/bar")
	st.Expect(t, err, nil)
	st.Expect(t, resp.StatusCode, 200)

	body, _ := io.ReadAll(resp.Body)
	t.Log(string(body))
	st.Expect(t, string(body)[:13], `{"foo":"bar"}`)

	// Ensure that we don't have pending mocks
	// This will throw an error if there are pending mocks
	st.Expect(t, gock.IsDone(), true)
}

func TestMatchHeaders(t *testing.T) {
	defer gock.Off()

	gock.New("http://foo.com").
		MatchHeader("Authorization", "^foo bar$").
		MatchHeader("API", "1.[0-9]+").
		HeaderPresent("Accept").
		Reply(200).
		BodyString("foo foo")

	req, err := http.NewRequest("GET", "http://foo.com", nil)
	st.Expect(t, err, nil)
	req.Header.Set("Authorization", "foo bar")
	req.Header.Set("API", "1.0")
	req.Header.Set("Accept", "text/plain")

	res, err := (&http.Client{}).Do(req)
	st.Expect(t, err, nil)
	st.Expect(t, res.StatusCode, 200)
	body, _ := io.ReadAll(res.Body)
	st.Expect(t, string(body), "foo foo")

	// Verify that we don't have pending mocks
	st.Expect(t, gock.IsDone(), true)
}
