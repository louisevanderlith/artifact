package api

import (
	"encoding/json"
	"fmt"
	"github.com/louisevanderlith/artifact/core"
	"github.com/louisevanderlith/husk/hsk"
	"github.com/louisevanderlith/husk/records"
	"io/ioutil"
	"net/http"
)

func FetchUpload(web *http.Client, host string, k hsk.Key) (core.Upload, error) {
	url := fmt.Sprintf("%s/upload/%s", host, k.String())
	resp, err := web.Get(url)

	if err != nil {
		return core.Upload{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bdy, _ := ioutil.ReadAll(resp.Body)
		return core.Upload{}, fmt.Errorf("%v: %s", resp.StatusCode, string(bdy))
	}

	result := core.Upload{}
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&result)

	return result, err
}

func FetchAllUploads(web *http.Client, host, pagesize string) (records.Page, error) {
	url := fmt.Sprintf("%s/upload/%s", host, pagesize)
	resp, err := web.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bdy, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("%v: %s", resp.StatusCode, string(bdy))
	}

	result := records.NewResultPage(core.Upload{})
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&result)

	return result, err
}
