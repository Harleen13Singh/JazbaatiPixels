package vendorgateway

import (
	"context"

	"github.com/google/uuid"

	"github.com/personal/jazbaatipixels/api/rpc"
	vgPb "github.com/personal/jazbaatipixels/api/vendorgateway"
	vgEnumsPb "github.com/personal/jazbaatipixels/api/vendorgateway/enums"
)

type Service struct {
	// UnimplementedVendorGatewayServer must be embedded to have forward compatible implementations.
	// this is useful to run time panics as this struct implements all the methods in the gRPC service
	vgPb.UnimplementedVendorGatewayServer
}

func NewService() *Service {
	return &Service{}
}

// UploadImage is useful to uplaad images with the respective vendors
func (s *Service) UploadImage(ctx context.Context, req *vgPb.UploadImageRequest) (*vgPb.UploadImageResponse, error) {
	// as of now, only claude is supported
	if req.GetVendor() != vgEnumsPb.Vendor_VENDOR_CLAUDE {
		return &vgPb.UploadImageResponse{
			Status: rpc.StatusInvalidArgumentWithDebugMsg("unknown vendor"),
		}, nil
	}

	client, clientErr := NewClient(&Config{
		ApiKey: ClaudeApiKeys,
	})
	if clientErr != nil {
		return &vgPb.UploadImageResponse{
			Status: rpc.StatusInternalWithDebugMsg(clientErr.Error()),
		}, nil
	}
	_, uploadErr := client.Send(&VendorRequest{
		Message: &Message{
			User: string(req.GetImage()),
		},
		SessionId: uuid.NewString(),
		Url:       UploadImageApiUrl,
	})
	if uploadErr != nil {
		return &vgPb.UploadImageResponse{
			Status: rpc.StatusInternalWithDebugMsg(uploadErr.Error()),
		}, nil
	}
	return &vgPb.UploadImageResponse{
		Status: rpc.StatusOk(),
	}, nil
}

func (s *Service) Chat(ctx context.Context, req *vgPb.ChatRequest) (*vgPb.ChatResponse, error) {
	if req.GetVendor() != vgEnumsPb.Vendor_VENDOR_CLAUDE {
		return &vgPb.ChatResponse{
			Status: rpc.StatusInvalidArgumentWithDebugMsg("unknown vendor"),
		}, nil
	}
	client, clientErr := NewClient(&Config{
		ApiKey: ClaudeApiKeys,
	})
	if clientErr != nil {
		return &vgPb.ChatResponse{
			Status: rpc.StatusInternalWithDebugMsg(clientErr.Error()),
		}, nil
	}
	resp, respErr := client.Send(&VendorRequest{
		Message: &Message{
			User: req.GetUserQuery(),
		},
		SessionId: uuid.NewString(),
		Url:       ChatApiUrl,
	})
	if respErr != nil {
		return &vgPb.ChatResponse{
			Status: rpc.StatusInternalWithDebugMsg(respErr.Error()),
		}, nil
	}
	return &vgPb.ChatResponse{
		Status:  rpc.StatusOk(),
		Respone: resp.ToString(),
	}, nil
}
