package service

import (
	"github.com/inagacky/go_grpc_sample/pb/cat"
	"golang.org/x/net/context"
	"log"
)

type CatService struct{}

var catList = [...] *cat.Cat{
	{Name:"mike", Type:cat.Type_AMERICAN_SHORTHAIR},
	{Name:"tama", Type:cat.Type_AMERICAN_SHORTHAIR},
	{Name:"sora", Type:cat.Type_SCOTTISH_FOLD},
	{Name:"umi", Type:cat.Type_SCOTTISH_FOLD},
	{Name:"me", Type:cat.Type_PERSIAN},
	{Name:"chun", Type:cat.Type_PERSIAN},
}

func (s *CatService) Search(ctx context.Context, req *cat.CatSearchRequest) (*cat.CatSearchReply, error) {

	log.Println("call Cat Type:", req.Type)

	// Typeで検索する
	result := searchByType(req.Type)
	rsp := cat.CatSearchReply {
		Cat:result,
	}

	return &rsp, nil
}

func searchByType (catType cat.Type) []*cat.Cat {

	result := make([]*cat.Cat, 0)
	// Typeが一致するものを取得
	for _, cat := range catList {
		if cat.Type == catType {
			result = append(result, cat)
		}
	}

	return result
}
