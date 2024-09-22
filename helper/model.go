package helper

import (
	"fajar7xx/pzn-golang-restful-api/model/domain"
	"fajar7xx/pzn-golang-restful-api/model/web"
)

func ToPostResponse(post domain.Post)web.PostResponse{
	return web.PostResponse{
		Id: post.Id,
		Name: post.Name,
		Post: post.Post,
	}
}

func ToPostResponses(posts []domain.Post)[]web.PostResponse{
	var postResponses []web.PostResponse
	for _, post := range posts{
		postResponses = append(postResponses, ToPostResponse(post))
	}

	return postResponses
}