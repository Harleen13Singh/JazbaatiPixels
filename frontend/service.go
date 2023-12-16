package frontend

import (
	"context"
	"errors"

	fePb "github.com/personal/jazbaatipixels/api/frontend"
	feErrPb "github.com/personal/jazbaatipixels/api/frontend/errors"
	"github.com/personal/jazbaatipixels/api/frontend/header"
	"github.com/personal/jazbaatipixels/api/rpc"
	vgPb "github.com/personal/jazbaatipixels/api/vendorgateway"
	"github.com/personal/jazbaatipixels/api/vendorgateway/enums"
)

type Service struct {
	vgClient vgPb.VendorGatewayClient
}

func NewService(vgClient vgPb.VendorGatewayClient) *Service {
	return &Service{
		vgClient: vgClient,
	}
}

// frontend service will expose rpcs to client to interact with the gpt

func (s *Service) StartChatSession(ctx context.Context, req *fePb.StartChatSessionRequest) (*fePb.StartChatSessionResponse, error) {
	if req.GetUserId() == "" {
		return defaultInvalidArgumentResponse("empty user id"), nil
	}
	if req.GetRequestType() == nil {
		return defaultInvalidArgumentResponse("nil request type"), nil
	}

	switch req.GetRequestType().(type) {
	case *fePb.StartChatSessionRequest_Image:
		vgResp, vgRespErr := s.vgClient.UploadImage(ctx, &vgPb.UploadImageRequest{
			Vendor: enums.Vendor_VENDOR_CLAUDE,
			Image:  req.GetImage(),
		})
		if vgResp == nil || vgRespErr != nil || vgResp.GetStatus() != rpc.StatusOk() {
			return defaultErrorResponse(errors.New("unknown response")), nil
		}
		return defaultResponse(vgResp.GetResponse()), nil
	case *fePb.StartChatSessionRequest_Message:
		vgResp, vgRespErr := s.vgClient.Chat(ctx, &vgPb.ChatRequest{
			Vendor:    enums.Vendor_VENDOR_CLAUDE,
			UserQuery: req.GetMessage(),
		})
		if vgResp == nil || vgRespErr != nil || vgResp.GetStatus() != rpc.StatusOk() {
			return defaultErrorResponse(errors.New("unknown response")), nil
		}
		return defaultResponse(vgResp.GetRespone()), nil
	default:
		return defaultInvalidArgumentResponse("unknown request type"), nil
	}

}

func defaultResponse(vgResp string) *fePb.StartChatSessionResponse {
	return &fePb.StartChatSessionResponse{
		RespHeader: &header.ResponseHeader{
			Status: rpc.StatusOk(),
		},
		Response: vgResp,
	}
}

func defaultErrorResponse(err error) *fePb.StartChatSessionResponse {
	return &fePb.StartChatSessionResponse{
		RespHeader: &header.ResponseHeader{
			Status: rpc.StatusInternalWithDebugMsg(err.Error()),
			ErrorView: &feErrPb.ErrorView{
				Type: feErrPb.ErrorViewType_BOTTOM_SHEET,
				Options: &feErrPb.ErrorView_BottomSheetErrorView{
					BottomSheetErrorView: &feErrPb.BottomSheetErrorView{
						Title: "Something went wront, Please try again!!!",
					},
				},
			},
		},
	}
}

func defaultInvalidArgumentResponse(debugMsg string) *fePb.StartChatSessionResponse {
	return &fePb.StartChatSessionResponse{
		RespHeader: &header.ResponseHeader{
			Status: rpc.StatusInvalidArgumentWithDebugMsg(debugMsg),
			ErrorView: &feErrPb.ErrorView{
				Type: feErrPb.ErrorViewType_BOTTOM_SHEET,
				Options: &feErrPb.ErrorView_BottomSheetErrorView{
					BottomSheetErrorView: &feErrPb.BottomSheetErrorView{
						Title: "Invalid query!!!",
					},
				},
			},
		},
	}
}
