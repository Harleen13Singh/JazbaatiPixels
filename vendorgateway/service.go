package vendorgateway

import (
	"context"

	"github.com/personal/jazbaatipixels/api/rpc"
	vgPb "github.com/personal/jazbaatipixels/api/vendorgateway"
	vgEnumsPb "github.com/personal/jazbaatipixels/api/vendorgateway/enums"
)

type Service struct{}

// UploadImage is useful to uplaad images with the respective vendors
func (s *Service) UploadImage(ctx context.Context, req *vgPb.UploadImageRequest) (*vgPb.UploadImageResponse, error) {
	// as of now, only claude is supported
	if req.GetVendor() != vgEnumsPb.Vendor_VENDOR_CLAUDE {
		return &vgPb.UploadImageResponse{
			Status: rpc.StatusInvalidArgument(),
		}, nil
	}
}
