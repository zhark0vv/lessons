package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProcessData(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		id        int
		mockFunc  func(*TestHelper)
		expected  string
		expectErr error
	}{
		{
			name: "successful data processing",
			id:   1,
			mockFunc: func(th *TestHelper) {
				th.DataProvider.On("GetData", context.Background(), 1).
					Return("mocked data", nil)
			},
			expected:  "Processed: mocked data",
			expectErr: nil,
		},
		{
			name: "provider returns error",
			id:   2,
			mockFunc: func(th *TestHelper) {
				th.DataProvider.On("GetData", context.Background(), 2).
					Return("", assert.AnError)
			},
			expected:  "",
			expectErr: assert.AnError,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel() // Запуск теста в параллельном режиме

			th := NewTestHelper(t)
			tt.mockFunc(th)

			gotResult, gotErr := th.Service.ProcessData(context.Background(), tt.id)

			require.ErrorIs(t, gotErr, tt.expectErr)
			require.Equal(t, tt.expected, gotResult)
		})
	}
}
