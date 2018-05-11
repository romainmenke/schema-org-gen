package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"golang.org/x/net/html"
)

type walkFunc func(context.Context, *ObjectSource, error) error

type ObjectSource struct {
	URL      string
	Object   *Object
	Children []*ObjectSource
	Parent   *ObjectSource `json:"-"`
}

func (v *ObjectSource) walk(ctx context.Context, walker walkFunc) error {
	if v == nil {
		return nil
	}

	var err error

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		//
	}

	err = walker(ctx, v, err)
	if err != nil {
		return err
	}

	if len(v.Children) == 0 {
		return nil
	}

	for _, o := range v.Children {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			//
		}

		err = o.walk(ctx, walker)
		if err != nil {
			return err
		}
	}

	return err
}

type TypeMap struct {
	UpdatedAt     time.Time
	ObjectSources []*ObjectSource
}

func (v *TypeMap) walk(ctx context.Context, walker walkFunc) error {
	if v == nil {
		return nil
	}

	if len(v.ObjectSources) == 0 {
		return nil
	}

	var err error

	for _, o := range v.ObjectSources {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			//
		}

		err = o.walk(ctx, walker)
		if err != nil {
			return err
		}
	}

	return err
}

func getTypeMap(ctx context.Context) (*TypeMap, error) {
	req, err := http.NewRequest("GET", "https://schema.org/docs/full.html", nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.Body != nil {
		defer resp.Body.Close()
	}

	if resp.StatusCode/100 != 2 {
		return nil, fmt.Errorf("unexpected status code : %d", resp.StatusCode)
	}
	if resp.Body == nil {
		return nil, errors.New("nil resp body")
	}

	z := html.NewTokenizer(resp.Body)

	inTree := false

	tMap := &TypeMap{
		UpdatedAt: time.Now(),
	}
	var currentObjectSource *ObjectSource
	var parentObjectSource *ObjectSource

TOKENIZER_LOOP:
	for {
		tt := z.Next()
		if z.Err() != nil {
			break TOKENIZER_LOOP
		}

		switch {
		case tt == html.ErrorToken:
			break TOKENIZER_LOOP
		case tt == html.StartTagToken:
			t := z.Token()

			if t.Data == "div" {
				for _, attr := range t.Attr {
					if attr.Key == "id" && attr.Val == "thing_tree" {
						inTree = true
					}
				}
			}

			if inTree {

				if t.Data == "ul" {
					if currentObjectSource != nil {
						parentObjectSource = currentObjectSource
					}
				}

				if t.Data == "a" {
					currentObjectSource = &ObjectSource{}

					for _, attr := range t.Attr {
						if attr.Key == "href" {
							currentObjectSource.URL = attr.Val
						}
					}

					if parentObjectSource != nil {
						currentObjectSource.Parent = parentObjectSource
						parentObjectSource.Children = append(parentObjectSource.Children, currentObjectSource)
					}
				}

			}

		case tt == html.EndTagToken:
			t := z.Token()

			if t.Data == "div" && inTree {
				inTree = false
				break TOKENIZER_LOOP
			}

			if inTree {

				if t.Data == "ul" {

					if parentObjectSource == nil {
						tMap.ObjectSources = append(tMap.ObjectSources, currentObjectSource)
					}

					if currentObjectSource != nil {
						currentObjectSource = currentObjectSource.Parent
					}

					if parentObjectSource != nil {
						parentObjectSource = parentObjectSource.Parent
					}

				}

			}

		}
	}
	err = z.Err()
	if err == io.EOF {
		err = nil
	}
	if err != nil {
		return nil, err
	}

	return tMap, nil
}

func getObject(ctx context.Context, o *ObjectSource, err error) error {
	if err != nil {
		return err
	}
	if o == nil {
		return nil
	}

	req, err := http.NewRequest("GET", "https://schema.org"+o.URL, nil)
	if err != nil {
		return err
	}

	req = req.WithContext(ctx)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if resp.Body != nil {
		defer resp.Body.Close()
	}

	if resp.StatusCode/100 != 2 {
		return fmt.Errorf("unexpected status code : %d", resp.StatusCode)
	}
	if resp.Body == nil {
		return errors.New("nil resp body")
	}

	z := html.NewTokenizer(resp.Body)
	err = objectParser(z, o)
	if err != nil {
		return err
	}

	return nil
}

func printObject(ctx context.Context, o *ObjectSource, err error) error {
	if err != nil {
		return err
	}
	if o == nil || o.Object == nil {
		return nil
	}

	ob := *(o.Object)

	fmt.Println(ob.Name)
	fmt.Println(ob.URL)
	fmt.Println(ob.Fields)

	return nil
}

func newTypeMapCacheFile() (*os.File, error) {
	err := os.RemoveAll("./tm-cache")
	if err != nil {
		return nil, err
	}

	err = os.MkdirAll("./tm-cache", 0700)
	if err != nil {
		return nil, err
	}

	f, err := os.Create("./tm-cache/tm-cache.json")
	if err != nil {
		return nil, err
	}

	return f, nil
}

func newOrExistingTypeMapCacheFile() (*os.File, error) {
	err := os.MkdirAll("./tm-cache", 0700)
	if err != nil {
		return nil, err
	}

	f, err := os.Open("./tm-cache/tm-cache.json")
	if os.IsNotExist(err) {
		f, err = os.Create("./tm-cache/tm-cache.json")
		if err != nil {
			return nil, err
		}
	}
	if err != nil {
		return nil, err
	}

	return f, nil
}
