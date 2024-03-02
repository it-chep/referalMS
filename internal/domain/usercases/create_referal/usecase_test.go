package create_referal

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"referalMS/internal/domain/usercases/create_referal/mocks"
	"testing"
)

func TestCreateReferalUseCase_Create(t *testing.T) {
	t.Parallel()
	t.Run("test creating", func(t *testing.T) {
		t.Parallel()
		mockRepo := mocks.NewMockWriteRepo(gomock.NewController(t))
		uc := CreateRferalUseCase{repo: mockRepo}
		mockRepo.EXPECT().Create(gomock.Any()).Return(nil)

		require.NoError(t, uc.Create(context.Background()))
	})
}
