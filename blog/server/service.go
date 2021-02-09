package main

import (
	"github.com/zemags/gRPSstudy/blog/pb"
)

// Service implement empty service with backward compatibility
type Service struct {
	pb.UnimplementedBlogServiceServer
}

// NewService make new empty Service
func NewService() *Service {
	return &Service{}
}
