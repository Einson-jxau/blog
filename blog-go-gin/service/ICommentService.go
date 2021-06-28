package service

import (
	pb "blog-go-gin/go_proto"
	"blog-go-gin/models/page"
)

type ICommentService interface {
	GetComments(articleId int, page *page.IPage) (*pb.CommentInfo, error)
	GetReplies(commentId int, page *page.IPage) ([]*pb.Reply, error)
	AddComment(comment *pb.CsComment) error
}
