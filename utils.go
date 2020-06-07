package lfp_scrapper

import "net/url"

func URLParameters(parseUrl string) (url.Values, error) {
	u, err := url.Parse(parseUrl)
	if err != nil {

		return nil, err
	}

	params, err := url.ParseQuery(u.RawQuery)
	if err != nil {

		return nil, err
	}

	return params, nil
}
