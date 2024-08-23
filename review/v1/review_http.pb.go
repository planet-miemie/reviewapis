// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.0
// - protoc             v3.20.1
// source: review/v1/review.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationReviewCreateReview = "/api.review.v1.Review/CreateReview"
const OperationReviewReplyReview = "/api.review.v1.Review/ReplyReview"

type ReviewHTTPServer interface {
	// CreateReview C端创建评价
	CreateReview(context.Context, *CreateReviewRequest) (*CreateReviewReply, error)
	// ReplyReview B端回复评价
	ReplyReview(context.Context, *ReplyReviewRequest) (*ReplyReviewReply, error)
}

func RegisterReviewHTTPServer(s *http.Server, srv ReviewHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/review", _Review_CreateReview0_HTTP_Handler(srv))
	r.POST("/v1/review/reply", _Review_ReplyReview0_HTTP_Handler(srv))
}

func _Review_CreateReview0_HTTP_Handler(srv ReviewHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateReviewRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationReviewCreateReview)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateReview(ctx, req.(*CreateReviewRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateReviewReply)
		return ctx.Result(200, reply)
	}
}

func _Review_ReplyReview0_HTTP_Handler(srv ReviewHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ReplyReviewRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationReviewReplyReview)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ReplyReview(ctx, req.(*ReplyReviewRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ReplyReviewReply)
		return ctx.Result(200, reply)
	}
}

type ReviewHTTPClient interface {
	CreateReview(ctx context.Context, req *CreateReviewRequest, opts ...http.CallOption) (rsp *CreateReviewReply, err error)
	ReplyReview(ctx context.Context, req *ReplyReviewRequest, opts ...http.CallOption) (rsp *ReplyReviewReply, err error)
}

type ReviewHTTPClientImpl struct {
	cc *http.Client
}

func NewReviewHTTPClient(client *http.Client) ReviewHTTPClient {
	return &ReviewHTTPClientImpl{client}
}

func (c *ReviewHTTPClientImpl) CreateReview(ctx context.Context, in *CreateReviewRequest, opts ...http.CallOption) (*CreateReviewReply, error) {
	var out CreateReviewReply
	pattern := "/v1/review"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationReviewCreateReview))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *ReviewHTTPClientImpl) ReplyReview(ctx context.Context, in *ReplyReviewRequest, opts ...http.CallOption) (*ReplyReviewReply, error) {
	var out ReplyReviewReply
	pattern := "/v1/review/reply"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationReviewReplyReview))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
