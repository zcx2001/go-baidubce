package bceClient

import (
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

type BceClient struct {
	accessKey                 string
	secretKey                 string
	expirationPeriodInSeconds uint //默认1800
	httpClient                *http.Client
}

func NewBceClient(accessKey, secretKey string) *BceClient {
	return &BceClient{
		accessKey:                 accessKey,
		secretKey:                 secretKey,
		expirationPeriodInSeconds: 1800,
		httpClient:                http.DefaultClient,
	}
}

func (c *BceClient) GetAccessKey() string {
	return c.accessKey
}

func (c *BceClient) GetSecretKey() string {
	return c.secretKey
}

func (c *BceClient) SetExpirationPeriodInSeconds(s uint) {
	c.expirationPeriodInSeconds = s
}

func (c *BceClient) GetExpirationPeriodInSeconds() uint {
	return c.expirationPeriodInSeconds
}

func (c *BceClient) SetHttpClient(httpClient *http.Client) {
	c.httpClient = httpClient
}

func (c *BceClient) Do(req *http.Request) (*http.Response, error) {
	timestamp := time.Now().UTC().Format("2006-01-02T15:04:05Z")
	requestId := uuid.New().String()

	//计算authStringPrefix
	authStringPrefix := fmt.Sprintf("bce-auth-v1/%s/%s/%d",
		c.accessKey, timestamp, c.expirationPeriodInSeconds)

	//计算signingKey
	signingKey := hmacSha256Hex(c.secretKey, authStringPrefix)

	req.Header.Add("Host", req.Host)
	req.Header.Add("x-bce-request-id", requestId) //请求id 36位uuid
	req.Header.Add("x-bce-date", timestamp)       //UTC时间格式%YYYY-%mm-%ddT%HH:%MM:%SSZ

	var canonicalRequest string
	//添加HTTP Method
	canonicalRequest += strings.ToUpper(req.Method) + "\n"

	//添加CanonicalURI
	canonicalRequest += req.URL.EscapedPath() + "\n"

	//添加CanonicalQueryString
	canonicalRequest += req.URL.Query().Encode() + "\n"

	//添加CanonicalHeaders
	headers := make([]string, 0)
	headerKeys := make([]string, 0)
	for k := range req.Header {
		headers = append(headers, fmt.Sprintf("%s:%s", strings.ToLower(k),
			url.QueryEscape(req.Header.Get(k))))
		headerKeys = append(headerKeys, strings.ToLower(k))
	}
	sort.Strings(headers)
	canonicalRequest += strings.Join(headers, "\n")

	//计算最终签名
	signature := hmacSha256Hex(signingKey, canonicalRequest)

	//生成最终使用的headerKey字符串
	sort.Strings(headerKeys)
	headerKey := strings.Join(headerKeys, ";")

	//添加最终认证信息
	req.Header.Add("Authorization",
		fmt.Sprintf("bce-auth-v1/%s/%s/%d/%s/%s",
			c.accessKey, timestamp, c.expirationPeriodInSeconds, headerKey, signature))

	return http.DefaultClient.Do(req)
}
