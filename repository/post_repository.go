package repository

import (
	"context"
	"database/sql"
	"errors"
	"fajar7xx/pzn-golang-restful-api/helper"
	"fajar7xx/pzn-golang-restful-api/model/domain"
)

type PostRepository interface{
	Create(ctx context.Context, tx *sql.Tx, post domain.Post) domain.Post
	Update(ctx context.Context, tx *sql.Tx, post domain.Post) domain.Post
	Delete(ctx context.Context, tx *sql.Tx, post domain.Post)
	FindyById(ctx context.Context, tx *sql.Tx, postId int) (domain.Post, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Post
}

type PostRepositoryImpl struct{
	DB *sql.DB
}

func (p *PostRepositoryImpl)Create(ctx context.Context, tx *sql.Tx, post domain.Post)domain.Post{
	query := "INSERT INTO posts(name, post)VALUES(?,?)"
	result, err := tx.ExecContext(ctx, query, post.Name, post.Post)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	post.Id = int(id)
	return post

}

func (p *PostRepositoryImpl)Update(ctx context.Context, tx *sql.Tx, post domain.Post)domain.Post{
	query := "UPDATE post SET name=?, post=? WHERE id=?"
	_, err := tx.ExecContext(ctx, query, post.Name, post.Post, post.Id)
	helper.PanicIfError(err)

	return post
}

func (p *PostRepositoryImpl)Delete(ctx context.Context, tx *sql.Tx, post domain.Post){
	query := "DELETE FROM post where id=?"
	_, err := tx.ExecContext(ctx, query, post.Id)
	helper.PanicIfError(err)
}

func (p *PostRepositoryImpl)FindyById(ctx context.Context, tx *sql.Tx, postId int)(domain.Post, error){
	query := "SELECT id, name, post FROM posts where id=?"
	row, err := tx.QueryContext(ctx, query, postId)
	helper.PanicIfError(err)

	post := domain.Post{} 
	if row.Next(){
		err := row.Scan(
			&post.Id,
			&post.Name,
			&post.Post,
		)
		helper.PanicIfError(err)
		return post, nil
	}else{
		return post, errors.New("post is not found")
	}

}

func (p *PostRepositoryImpl)FindAll(ctx context.Context, tx *sql.Tx)[]domain.Post{
	query := "SELECT id, name, post FROM posts"
	rows, err := tx.QueryContext(ctx, query)
	helper.PanicIfError(err)
	 
	var posts []domain.Post
	for rows.Next(){
		post := domain.Post{}
		err := rows.Scan(
			&post.Id,
			&post.Name,
			&post.Post,
		)
		helper.PanicIfError(err)
		posts = append(posts, post)
	}

	return posts
}