package myhttp

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"
	"time"
)

type RequestsMaker struct {
	addresses     []string
	parallelCount int

	client *http.Client
}

func NewRequestsMaker(addresses []string, parallelCount int) *RequestsMaker {
	return &RequestsMaker{
		addresses:     addresses,
		parallelCount: parallelCount,

		client: &http.Client{
			Timeout: time.Second * 5,
		},
	}
}

func (t *RequestsMaker) Run() {
	addrChan := make(chan string, t.parallelCount)
	defer close(addrChan)

	wt := &sync.WaitGroup{}

	wt.Add(len(t.addresses))

	for _, addr := range t.addresses {
		go func(addr string) {
			addrChan <- addr

			defer func() {
				wt.Done()
				<-addrChan
			}()

			res := getHashResponse(t.client, addr)
			printResult(addr, res)
		}(addr)
	}

	wt.Wait()
}

func getHashResponse(client *http.Client, addr string) []byte {
	uri, err := getUrl(addr)
	if err != nil {
		return nil
	}

	resp, err := client.Get(uri.String())
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil
	}

	return hashBytes(bytes)
}

const httpScheme = "http"

func getUrl(addr string) (*url.URL, error) {
	uri, err := url.Parse(addr)
	if err != nil {
		return nil, err
	}

	if uri.Scheme == "" {
		uri.Scheme = httpScheme
	}

	return uri, nil
}

func hashBytes(bytes []byte) []byte {
	hashMd5 := md5.New()

	_, err := hashMd5.Write(bytes)
	if err != nil {
		return nil
	}

	return hashMd5.Sum(nil)
}

func printResult(addr string, hashData []byte) {
	fmt.Printf("%v %x\n", addr, hashData)
}
