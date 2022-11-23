// @author cold bin
// @date 2022/10/22

package web

import (
	"errors"
	"io"
	"net/http"
	"net/url"
	"sync"
)

// 池化
var reqPool = sync.Pool{New: func() interface{} {
	return http.DefaultClient
}}

// HttpGet 封装http get方法请求数据（get方法默认没有请求体）
func HttpGet(url string) (data []byte, e error) {
	var c *http.Client
	var ok bool
	if c, ok = reqPool.Get().(*http.Client); !ok {
		return nil, errors.New("req pool is wrong")
	}

	response, err := c.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	reqPool.Put(c)

	return bytes, nil
}

func HttpPost(url string, contentType string, body io.Reader) (data []byte, e error) {
	var c *http.Client
	var ok bool
	if c, ok = reqPool.Get().(*http.Client); !ok {
		return nil, errors.New("req pool is wrong")
	}

	response, err := c.Post(url, contentType, body)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	reqPool.Put(c)

	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func HttpPostForm(url string, datas url.Values) (data []byte, e error) {
	var c *http.Client
	var ok bool
	if c, ok = reqPool.Get().(*http.Client); !ok {
		return nil, errors.New("req pool is wrong")
	}

	response, err := c.PostForm(url, datas)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	reqPool.Put(c)

	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}
