package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/kenyako/chat-server/internal/api/chat"
	"github.com/kenyako/chat-server/internal/model"
	serviceMock "github.com/kenyako/chat-server/internal/service/mocks"
	desc "github.com/kenyako/chat-server/pkg/chat_v1"
	"github.com/stretchr/testify/require"
)

func TestImplementation_SuccessCreate(t *testing.T) {

	type mocker func() *serviceMock.ChatAPIService

	type args struct {
		ctx context.Context
		req *desc.CreateRequest
	}

	var (
		ctx = context.Background()

		chatId = gofakeit.Uint64()

		title   = gofakeit.BeerName()
		members = []string{"Alice", "Robert", "Lisa", "Barney"}

		req = &desc.CreateRequest{
			Title:     title,
			Usernames: members,
		}

		info = &model.CreateChat{
			Title:     title,
			Usernames: members,
		}

		res = &desc.CreateResponse{
			Id: int64(chatId),
		}
	)

	tests := []struct {
		name string
		args args
		want *desc.CreateResponse
		err  error
		mock mocker
	}{
		{
			name: "success chat create",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			mock: func() *serviceMock.ChatAPIService {

				serviceMock := serviceMock.NewChatAPIService(t)
				serviceMock.On("Create", ctx, info).Return(int64(chatId), nil)

				return serviceMock
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			serviceMock := tt.mock()

			api := chat.NewImplementation(serviceMock)

			result, err := api.Create(tt.args.ctx, tt.args.req)

			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, result)
		})
	}
}

func TestImplementation_FailCreate(t *testing.T) {

	type mocker func() *serviceMock.ChatAPIService

	type args struct {
		ctx context.Context
		req *desc.CreateRequest
	}

	var (
		ctx = context.Background()

		serviceErr = fmt.Errorf("failed to create chat")

		title   = gofakeit.BeerName()
		members = []string{"Alice", "Robert", "Lisa", "Barney"}

		req = &desc.CreateRequest{
			Title:     title,
			Usernames: members,
		}

		info = &model.CreateChat{
			Title:     title,
			Usernames: members,
		}
	)

	tests := []struct {
		name string
		args args
		want *desc.CreateResponse
		err  error
		mock mocker
	}{
		{
			name: "fail chat create",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  serviceErr,
			mock: func() *serviceMock.ChatAPIService {

				serviceMock := serviceMock.NewChatAPIService(t)
				serviceMock.On("Create", ctx, info).Return(int64(0), serviceErr)

				return serviceMock
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			serviceMock := tt.mock()

			api := chat.NewImplementation(serviceMock)

			result, err := api.Create(tt.args.ctx, tt.args.req)

			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, result)
		})
	}
}
