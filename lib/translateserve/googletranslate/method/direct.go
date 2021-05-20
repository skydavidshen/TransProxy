package method

import (
	"io/ioutil"
	"net/http"
	"time"
)

type Direct struct {

}

func (d *Direct) Content(resourceUrl string) ([]byte, error) {
	var r *http.Response
	tries := 3
	for tries > 0 {
		r, err := http.Get(resourceUrl)
		if err != nil {
			if err == http.ErrHandlerTimeout {
				return []byte{}, err
			}
			return []byte{}, err
		}

		if r.StatusCode == http.StatusOK {
			break
		}

		if r.StatusCode == http.StatusForbidden {
			tries--
			time.Sleep(time.Second)
		}
	}

	raw, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return []byte{}, err
	}
	return raw, nil
}
