package main

import (
	"context"
	"encoding/json"
	"net/http"
)

type Service interface {
	GetFatCatFact(context.Context) (*FatCatFact, error)
}

type FatCatFactService struct {
	url string
}

func NewFatCatFactService(url string) Service {
	return &FatCatFactService{
		url: url,
	}
}

func (s *FatCatFactService) GetFatCatFact(ctx context.Context) (*FatCatFact, error) {
	resp, err := http.Get(s.url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	fact := &FatCatFact{}
	if err := json.NewDecoder(resp.Body).Decode(&fact);err != nil {
		return nil, err
	}

	return fact, nil 

}