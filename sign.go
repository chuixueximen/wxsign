package wxsign

import (
	"crypto/sha1"
	"fmt"
	"github.com/chuixueximen/gutil"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// GetJsSign GetJsSign
func (wSign *WxSign) GetJsSign(url string) (*WxJsSign, error) {
	jsTicket, err := wSign.GetTicket()
	if err != nil {
		return nil, err
	}
	// splite url
	urlSlice := strings.Split(url, "#")

	// add /
	if urlSlice[0], err = addSlashToURL(urlSlice[0]); err != nil {
		return nil, err
	}

	jsSign := &WxJsSign{
		Appid:     wSign.Appid,
		Noncestr:  gutil.RandString(16),
		Timestamp: strconv.FormatInt(time.Now().UTC().Unix(), 10),
		Url:       urlSlice[0],
	}
	jsSign.Signature = Signature(jsTicket, jsSign.Noncestr, jsSign.Timestamp, jsSign.Url)
	return jsSign, nil
}

// Signature
func Signature(jsTicket, noncestr, timestamp, url string) string {
	h := sha1.New()
	h.Write([]byte(fmt.Sprintf("jsapi_ticket=%s&noncestr=%s&timestamp=%s&url=%s", jsTicket, noncestr, timestamp, url)))
	fmt.Printf("打印：\n\njsapi_ticket=%s&noncestr=%s&timestamp=%s&url=%s", jsTicket, noncestr, timestamp, url)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// addSlashToURL 只有域名的链接后面添加/
func addSlashToURL(urlStr string) (string, error) {
	u, err := url.Parse(urlStr)
	if err != nil {
		return "", err
	}

	if u.Path == "" {
		u.Path = "/"
	}

	return u.String(), nil
}
