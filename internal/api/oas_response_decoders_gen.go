// Code generated by ogen, DO NOT EDIT.

package api

import (
	"fmt"
	"io"
	"mime"
	"net/http"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"

	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/validate"
)

func decodeCreateCatResponse(resp *http.Response) (res *Cat, _ error) {
	switch resp.StatusCode {
	case 201:
		// Code 201.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			buf, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}
			d := jx.DecodeBytes(buf)

			var response Cat
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				if err := d.Skip(); err != io.EOF {
					return errors.New("unexpected trailing data")
				}
				return nil
			}(); err != nil {
				err = &ogenerrors.DecodeBodyError{
					ContentType: ct,
					Body:        buf,
					Err:         err,
				}
				return res, err
			}
			// Validate response.
			if err := func() error {
				if err := response.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, errors.Wrap(err, "validate")
			}
			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	}
	return res, validate.UnexpectedStatusCode(resp.StatusCode)
}

func decodeDeleteCatResponse(resp *http.Response) (res DeleteCatRes, _ error) {
	switch resp.StatusCode {
	case 204:
		// Code 204.
		return &DeleteCatNoContent{}, nil
	case 404:
		// Code 404.
		return &DeleteCatNotFound{}, nil
	}
	return res, validate.UnexpectedStatusCode(resp.StatusCode)
}

func decodeGetCatResponse(resp *http.Response) (res GetCatRes, _ error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			buf, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}
			d := jx.DecodeBytes(buf)

			var response Cat
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				if err := d.Skip(); err != io.EOF {
					return errors.New("unexpected trailing data")
				}
				return nil
			}(); err != nil {
				err = &ogenerrors.DecodeBodyError{
					ContentType: ct,
					Body:        buf,
					Err:         err,
				}
				return res, err
			}
			// Validate response.
			if err := func() error {
				if err := response.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, errors.Wrap(err, "validate")
			}
			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	case 404:
		// Code 404.
		return &GetCatNotFound{}, nil
	}
	return res, validate.UnexpectedStatusCode(resp.StatusCode)
}

func decodeListCatsResponse(resp *http.Response) (res []Cat, _ error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			buf, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}
			d := jx.DecodeBytes(buf)

			var response []Cat
			if err := func() error {
				response = make([]Cat, 0)
				if err := d.Arr(func(d *jx.Decoder) error {
					var elem Cat
					if err := elem.Decode(d); err != nil {
						return err
					}
					response = append(response, elem)
					return nil
				}); err != nil {
					return err
				}
				if err := d.Skip(); err != io.EOF {
					return errors.New("unexpected trailing data")
				}
				return nil
			}(); err != nil {
				err = &ogenerrors.DecodeBodyError{
					ContentType: ct,
					Body:        buf,
					Err:         err,
				}
				return res, err
			}
			// Validate response.
			if err := func() error {
				if response == nil {
					return errors.New("nil is invalid value")
				}
				var failures []validate.FieldError
				for i, elem := range response {
					if err := func() error {
						if err := elem.Validate(); err != nil {
							return err
						}
						return nil
					}(); err != nil {
						failures = append(failures, validate.FieldError{
							Name:  fmt.Sprintf("[%d]", i),
							Error: err,
						})
					}
				}
				if len(failures) > 0 {
					return &validate.Error{Fields: failures}
				}
				return nil
			}(); err != nil {
				return res, errors.Wrap(err, "validate")
			}
			return response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	}
	return res, validate.UnexpectedStatusCode(resp.StatusCode)
}

func decodeUpdateCatResponse(resp *http.Response) (res UpdateCatRes, _ error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			buf, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}
			d := jx.DecodeBytes(buf)

			var response Cat
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				if err := d.Skip(); err != io.EOF {
					return errors.New("unexpected trailing data")
				}
				return nil
			}(); err != nil {
				err = &ogenerrors.DecodeBodyError{
					ContentType: ct,
					Body:        buf,
					Err:         err,
				}
				return res, err
			}
			// Validate response.
			if err := func() error {
				if err := response.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, errors.Wrap(err, "validate")
			}
			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	case 404:
		// Code 404.
		return &UpdateCatNotFound{}, nil
	}
	return res, validate.UnexpectedStatusCode(resp.StatusCode)
}
