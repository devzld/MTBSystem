package handler

import (
	"context"
	"comment-srv/db"
	"go.uber.org/zap"
	"share/pb"
	"share/utils/log"
	errors "share/errors"
	"comment-srv/entity"
)

type CommentServiceExtHandler struct {
	logger *zap.Logger
}

func NewCommentServiceExtHandler() *CommentServiceExtHandler {
	return &CommentServiceExtHandler{
		logger: log.Instance(),
	}
}

// 获取评论
func (f *CommentServiceExtHandler) HotComment(ctx context.Context, req *pb.HotCommentReq, rsp *pb.HotCommentRsp) error {

	movieId := req.MovieId
	comments,err := db.SelectHotComment(movieId)
	if err != nil {
		return errors.ErrorCommentFailed
	}
	records := []*pb.CommentRecord{}
	for _,comment := range comments {
		record := pb.CommentRecord{
			Title: comment.Title,
			Content: comment.Content,
			HeadImg:comment.HeadImg,
			Nickname:comment.NickName,
		}
		records = append(records, &record)
	}

	plus := pb.CommentPlus{
		Total:3,
		List:records,
	}

	data := pb.CommentData{
		Plus: &plus,
	}
	rsp.Data = &data
	return nil
}

func (f *CommentServiceExtHandler) MakeComment(ctx context.Context, req *pb.MakeCommentReq, rsp *pb.MakeCommentRsp) error {

	comment := entity.Comment{
		Title:req.Title,
		Content:req.Content,
		HeadImg:req.HeadImg,
		FilmId:req.MovieId,
		NickName:req.Nickname,
	}
	err := db.InsertHotComment(&comment)
	if err != nil {
		return errors.ErrorCommentFailed
	}
	return nil
}