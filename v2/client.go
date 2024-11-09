package v2

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/google/go-querystring/query"
)

const ContentType = "application/json"

type Client struct {
	client             *http.Client
	host               *url.URL
	username           string
	password           string
	onRequestCompleted RequestCompletionCallback

	Alias      AliasesAPI
	Cluster    ClusterAPI
	Collection CollectionAPI
	ConfigSet  ConfigSetAPI
	Core       CoreAPI
	FileStore  FileStoreAPI
	Node       NodeAPI
	Query      QueryAPI
	Replica    ReplicaAPI
	Schema     SchemaAPI
	Shard      ShardAPI
}

type RequestCompletionCallback func(*http.Request, *http.Response)

func NewClient(host string, username string, password string) (*Client, error) {
	httpClient := http.DefaultClient
	hostURL, err := url.Parse(host)
	if err != nil {
		return nil, err
	}

	client := Client{
		client:   httpClient,
		host:     hostURL,
		username: username,
		password: password,
	}
	client.InitAPI()

	return &client, nil
}

func (c *Client) InitAPI() {
	c.Alias = AliasesAPI{client: c}
	c.Cluster = ClusterAPI{client: c}
	c.Collection = CollectionAPI{client: c}
	c.ConfigSet = ConfigSetAPI{client: c}
	c.Core = CoreAPI{client: c}
	c.FileStore = FileStoreAPI{client: c}
	c.Node = NodeAPI{client: c}
	c.Query = QueryAPI{client: c}
	c.Replica = ReplicaAPI{client: c}
	c.Schema = SchemaAPI{client: c}
	c.Shard = ShardAPI{client: c}
}

func (c *Client) NewUpload(ctx context.Context, urlStr string, filepath string, queryStrings interface{}) (*Response, error) {
	u, err := c.host.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	params, _ := query.Values(queryStrings)
	u.RawQuery = params.Encode()

	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	resp, err := http.Post(u.String(), "application/octet-stream", file)
	if err != nil {
		return nil, err
	}

	b, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := Response{HttpResponse: resp}

	err = json.Unmarshal(b, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) NewRequest(ctx context.Context, method, urlStr string, body interface{}, queryStrings interface{}, headers *map[string]string) (*http.Request, error) {
	u, err := c.host.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	if body != nil {
		err = json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	params, _ := query.Values(queryStrings)
	u.RawQuery = params.Encode()

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", ContentType)
	if headers != nil {
		for key, value := range *headers {
			req.Header.Set(key, value)
		}
	}

	if c.username != "" && c.password != "" {
		req.SetBasicAuth(c.username, c.password)
	}

	return req, nil
}

func (c *Client) Do(ctx context.Context, req *http.Request) (*Response, error) {
	req = req.WithContext(ctx)
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	if c.onRequestCompleted != nil {
		c.onRequestCompleted(req, resp)
	}

	defer func() {
		if rerr := resp.Body.Close(); err == nil {
			err = rerr
		}
	}()

	b, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := Response{HttpResponse: resp}

	err = json.Unmarshal(b, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) OnRequestCompleted(rc RequestCompletionCallback) {
	c.onRequestCompleted = rc
}
