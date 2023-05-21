package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"project/internal/config"
	chatMock "project/internal/microservices/chat/repository/mocks"
	consumerMock "project/internal/microservices/consumer/usecase/mocks"
	messagesMock "project/internal/microservices/messages/repository/mocks"
	producerMock "project/internal/microservices/producer/usecase/mocks"
	"project/internal/model"
	myErrors "project/internal/pkg/errors"
	"testing"
	"time"
)

type testCase struct {
	name          string
	body          []byte
	dbError       error
	members       []model.ChatMembers
	membersError  error
	producerError error
	result        error
}

func Test_Messages_SendingChatMembers(t *testing.T) {
	tests := []testCase{
		{
			name: `send message; 
                      OK; 
                      zero members`,
			dbError:       nil,
			members:       []model.ChatMembers{},
			membersError:  nil,
			producerError: nil,
			result:        nil,
		},
		{
			name: `send message;
				   OK; 
                   two members`,
			dbError: nil,
			members: []model.ChatMembers{
				{
					ChatId:   1,
					MemberId: 1,
				},
				{
					ChatId:   1,
					MemberId: 2,
				},
			},
			membersError:  nil,
			producerError: nil,
			result:        nil,
		},
		{
			name: `send message;
				   unmarshal error; 
                   two members`,
			dbError: nil,
			members: []model.ChatMembers{
				{
					ChatId:   1,
					MemberId: 1,
				},
				{
					ChatId:   1,
					MemberId: 2,
				},
			},
			membersError:  nil,
			producerError: nil,
			result:        errors.New("invalid character 'd' in literal false (expecting 'a')"),
		},
		{
			name: `send message;
                      get chat members error;
                      zero members`,
			dbError:       nil,
			members:       []model.ChatMembers{},
			membersError:  myErrors.ErrMembersNotFound,
			producerError: nil,
			result:        myErrors.ErrMembersNotFound,
		},
		{
			name: `send message;
                      insert message in DB error;
                      zero members`,
			dbError: myErrors.ErrInternal,
			members: []model.ChatMembers{
				{
					ChatId:   1,
					MemberId: 1,
				},
				{
					ChatId:   1,
					MemberId: 2,
				},
			},
			membersError:  nil,
			producerError: nil,
			result:        nil,
		},
		{
			name: `send message;
				   producer error;
                   two members`,
			dbError: nil,
			members: []model.ChatMembers{
				{
					ChatId:   1,
					MemberId: 1,
				},
				{
					ChatId:   1,
					MemberId: 2,
				},
			},
			membersError:  nil,
			producerError: myErrors.ErrInternal,
			result:        myErrors.ErrInternal,
		},
	}

	messages := []model.WebSocketMessage{
		{
			Id:       "1",
			Type:     config.Create,
			Body:     "Hello world!",
			AuthorID: 1,
			ChatID:   1,
		},
		{
			Id:       "1",
			Type:     config.Create,
			Body:     "Hello world!",
			AuthorID: 1,
			ChatID:   1,
		},
		{},
		{
			Id:       "1",
			Type:     config.Create,
			Body:     "Hello world!",
			AuthorID: 1,
			ChatID:   1,
		},
		{
			Id:       "1",
			Type:     config.Create,
			Body:     "Hello world!",
			AuthorID: 1,
			ChatID:   1337,
		},
		{
			Id:       "1",
			Type:     config.Create,
			Body:     "Hello world!",
			AuthorID: 1,
			ChatID:   1,
		},
	}

	var err error
	jsonMessages := make([][]byte, len(messages))
	for idx, message := range messages {
		if idx == 2 {
			jsonMessages[idx] = []byte("fdafdafd")
			continue
		}
		jsonMessages[idx], err = json.Marshal(message)
		require.NoError(t, err)
	}

	for idx, _ := range messages {
		tests[idx].body = jsonMessages[idx]
	}

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	chatRepository := chatMock.NewMockRepository(ctl)
	consumerUsecase := consumerMock.NewMockUsecase(ctl)
	producerUsecase := producerMock.NewMockUsecase(ctl)
	messagesRepository := messagesMock.NewMockRepository(ctl)
	usecase := NewMessagesUsecase(chatRepository, consumerUsecase, producerUsecase, messagesRepository)

	for _, test := range tests {
		chatRepository.EXPECT().GetChatMembersByChatId(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, chatID uint64) ([]model.ChatMembers, error) {
			return test.members, test.membersError
		}).AnyTimes()

		messagesRepository.EXPECT().InsertMessageInDB(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, message model.Message) error {
			return test.dbError
		}).AnyTimes()

		producerUsecase.EXPECT().ProduceMessage(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, message model.ProducerMessage) error {
			return test.producerError
		}).AnyTimes()

		err := usecase.PutInProducer(context.TODO(), test.body)
		time.Sleep(100 * time.Millisecond)
		if test.result == nil {
			require.NoError(t, err)
			continue
		}

		require.Equal(t, test.result.Error(), err.Error())
	}
}

func Test_Messages_EditMessage(t *testing.T) {
	tests := []testCase{
		{
			name: `edit message; 
                      OK; 
                      zero members`,
			dbError:       nil,
			members:       []model.ChatMembers{},
			membersError:  nil,
			producerError: nil,
			result:        nil,
		},
		{
			name: `edit message;
				   OK; 
                   two members`,
			dbError: nil,
			members: []model.ChatMembers{
				{
					ChatId:   1,
					MemberId: 1,
				},
				{
					ChatId:   1,
					MemberId: 2,
				},
			},
			membersError:  nil,
			producerError: nil,
			result:        nil,
		},
		{
			name: `edit message;
				   unmarshal error; 
                   two members`,
			dbError: nil,
			members: []model.ChatMembers{
				{
					ChatId:   1,
					MemberId: 1,
				},
				{
					ChatId:   1,
					MemberId: 2,
				},
			},
			membersError:  nil,
			producerError: nil,
			result:        errors.New("invalid character 'd' in literal false (expecting 'a')"),
		},
		{
			name: `edit message;
                      get chat members error;
                      zero members`,
			dbError:       nil,
			members:       []model.ChatMembers{},
			membersError:  myErrors.ErrMembersNotFound,
			producerError: nil,
			result:        myErrors.ErrMembersNotFound,
		},
		{
			name: `edit message;
                      edit message in DB error;
                      zero members`,
			dbError: myErrors.ErrInternal,
			members: []model.ChatMembers{
				{
					ChatId:   1,
					MemberId: 1,
				},
				{
					ChatId:   1,
					MemberId: 2,
				},
			},
			membersError:  nil,
			producerError: nil,
			result:        nil,
		},
		{
			name: `edit message;
				   producer error;
                   two members`,
			dbError: nil,
			members: []model.ChatMembers{
				{
					ChatId:   1,
					MemberId: 1,
				},
				{
					ChatId:   1,
					MemberId: 2,
				},
			},
			membersError:  nil,
			producerError: myErrors.ErrInternal,
			result:        myErrors.ErrInternal,
		},
	}

	messages := []model.WebSocketMessage{
		{
			Id:       "1",
			Action:   config.Edit,
			Type:     config.NotSticker,
			Body:     "Hello world!",
			AuthorID: 1,
			ChatID:   1,
		},
		{
			Id:       "1",
			Action:   config.Edit,
			Type:     config.NotSticker,
			Body:     "Hello world!",
			AuthorID: 1,
			ChatID:   1,
		},
		{},
		{
			Id:       "1",
			Action:   config.Edit,
			Type:     config.NotSticker,
			Body:     "Hello world!",
			AuthorID: 1,
			ChatID:   1,
		},
		{
			Id:       "1",
			Action:   config.Edit,
			Type:     config.NotSticker,
			Body:     "Hello world!",
			AuthorID: 1,
			ChatID:   1337,
		},
		{
			Id:       "1",
			Action:   config.Edit,
			Type:     config.NotSticker,
			Body:     "Hello world!",
			AuthorID: 1,
			ChatID:   1,
		},
	}

	var err error
	jsonMessages := make([][]byte, len(messages))
	for idx, message := range messages {
		if idx == 2 {
			jsonMessages[idx] = []byte("fdafdafd")
			continue
		}
		jsonMessages[idx], err = json.Marshal(message)
		require.NoError(t, err)
	}

	for idx, _ := range messages {
		tests[idx].body = jsonMessages[idx]
	}

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	chatRepository := chatMock.NewMockRepository(ctl)
	consumerUsecase := consumerMock.NewMockUsecase(ctl)
	producerUsecase := producerMock.NewMockUsecase(ctl)
	messagesRepository := messagesMock.NewMockRepository(ctl)
	usecase := NewMessagesUsecase(chatRepository, consumerUsecase, producerUsecase, messagesRepository)

	for _, test := range tests {
		chatRepository.EXPECT().GetChatMembersByChatId(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, chatID uint64) ([]model.ChatMembers, error) {
			return test.members, test.membersError
		}).AnyTimes()

		messagesRepository.EXPECT().EditMessageById(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, message model.ProducerMessage) (model.Message, error) {
			return model.Message{}, test.dbError
		}).AnyTimes()

		producerUsecase.EXPECT().ProduceMessage(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, message model.ProducerMessage) error {
			return test.producerError
		}).AnyTimes()

		err := usecase.PutInProducer(context.TODO(), test.body)
		time.Sleep(100 * time.Millisecond)
		if test.result == nil {
			require.NoError(t, err)
			continue
		}

		require.Equal(t, test.result.Error(), err.Error())
	}
}

func Test_Messages_DeleteMessage(t *testing.T) {
	tests := []testCase{
		{
			name: `delete message; 
                      OK; 
                      zero members`,
			dbError:       nil,
			members:       []model.ChatMembers{},
			membersError:  nil,
			producerError: nil,
			result:        nil,
		},
		{
			name: `delete message;
				   OK; 
                   two members`,
			dbError: nil,
			members: []model.ChatMembers{
				{
					ChatId:   1,
					MemberId: 1,
				},
				{
					ChatId:   1,
					MemberId: 2,
				},
			},
			membersError:  nil,
			producerError: nil,
			result:        nil,
		},
		{
			name: `delete message;
				   unmarshal error; 
                   two members`,
			dbError: nil,
			members: []model.ChatMembers{
				{
					ChatId:   1,
					MemberId: 1,
				},
				{
					ChatId:   1,
					MemberId: 2,
				},
			},
			membersError:  nil,
			producerError: nil,
			result:        errors.New("invalid character 'd' in literal false (expecting 'a')"),
		},
		{
			name: `delete message;
                      get chat members error;
                      zero members`,
			dbError:       nil,
			members:       []model.ChatMembers{},
			membersError:  myErrors.ErrMembersNotFound,
			producerError: nil,
			result:        myErrors.ErrMembersNotFound,
		},
		{
			name: `delete message;
                      delete from DB error;
                      zero members`,
			dbError: myErrors.ErrInternal,
			members: []model.ChatMembers{
				{
					ChatId:   1,
					MemberId: 1,
				},
				{
					ChatId:   1,
					MemberId: 2,
				},
			},
			membersError:  nil,
			producerError: nil,
			result:        nil,
		},
		{
			name: `delete message;
				   producer error;
                   two members`,
			dbError: nil,
			members: []model.ChatMembers{
				{
					ChatId:   1,
					MemberId: 1,
				},
				{
					ChatId:   1,
					MemberId: 2,
				},
			},
			membersError:  nil,
			producerError: myErrors.ErrInternal,
			result:        myErrors.ErrInternal,
		},
	}

	messages := []model.WebSocketMessage{
		{
			Id:       "1",
			Action:   config.Delete,
			Type:     config.NotSticker,
			Body:     "Hello world!",
			AuthorID: 1,
			ChatID:   1,
		},
		{
			Id:       "1",
			Action:   config.Delete,
			Type:     config.NotSticker,
			Body:     "Hello world!",
			AuthorID: 1,
			ChatID:   1,
		},
		{},
		{
			Id:       "1",
			Action:   config.Delete,
			Type:     config.NotSticker,
			Body:     "Hello world!",
			AuthorID: 1,
			ChatID:   1,
		},
		{
			Id:       "1",
			Action:   config.Delete,
			Type:     config.NotSticker,
			Body:     "Hello world!",
			AuthorID: 1,
			ChatID:   1337,
		},
		{
			Id:       "1",
			Action:   config.Delete,
			Type:     config.NotSticker,
			Body:     "Hello world!",
			AuthorID: 1,
			ChatID:   1,
		},
	}

	var err error
	jsonMessages := make([][]byte, len(messages))
	for idx, message := range messages {
		if idx == 2 {
			jsonMessages[idx] = []byte("fdafdafd")
			continue
		}
		jsonMessages[idx], err = json.Marshal(message)
		require.NoError(t, err)
	}

	for idx, _ := range messages {
		tests[idx].body = jsonMessages[idx]
	}

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	chatRepository := chatMock.NewMockRepository(ctl)
	consumerUsecase := consumerMock.NewMockUsecase(ctl)
	producerUsecase := producerMock.NewMockUsecase(ctl)
	messagesRepository := messagesMock.NewMockRepository(ctl)
	usecase := NewMessagesUsecase(chatRepository, consumerUsecase, producerUsecase, messagesRepository)

	for _, test := range tests {
		chatRepository.EXPECT().GetChatMembersByChatId(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, chatID uint64) ([]model.ChatMembers, error) {
			return test.members, test.membersError
		}).AnyTimes()

		messagesRepository.EXPECT().DeleteMessageById(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, messageID string) error {
			return test.dbError
		}).AnyTimes()

		producerUsecase.EXPECT().ProduceMessage(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, message model.ProducerMessage) error {
			return test.producerError
		}).AnyTimes()

		err := usecase.PutInProducer(context.TODO(), test.body)
		time.Sleep(100 * time.Millisecond)
		if test.result == nil {
			require.NoError(t, err)
			continue
		}

		require.Equal(t, test.result.Error(), err.Error())
	}
}

func Test_Messages_UndefinedAction(t *testing.T) {
	tests := []testCase{
		{
			name: `undefined message; 
                      OK; 
                      zero members`,
			dbError:       nil,
			members:       []model.ChatMembers{},
			membersError:  nil,
			producerError: nil,
			result:        errors.New("не выбран ни один из трех 0, 1, 2"),
		},
	}

	messages := []model.WebSocketMessage{
		{
			Id:       "1",
			Action:   1337,
			Type:     config.NotSticker,
			Body:     "Hello world!",
			AuthorID: 1,
			ChatID:   1,
		},
	}

	var err error
	jsonMessages := make([][]byte, len(messages))
	for idx, message := range messages {
		if idx == 2 {
			jsonMessages[idx] = []byte("fdafdafd")
			continue
		}
		jsonMessages[idx], err = json.Marshal(message)
		require.NoError(t, err)
	}

	for idx, _ := range messages {
		tests[idx].body = jsonMessages[idx]
	}

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	chatRepository := chatMock.NewMockRepository(ctl)
	consumerUsecase := consumerMock.NewMockUsecase(ctl)
	producerUsecase := producerMock.NewMockUsecase(ctl)
	messagesRepository := messagesMock.NewMockRepository(ctl)
	usecase := NewMessagesUsecase(chatRepository, consumerUsecase, producerUsecase, messagesRepository)

	for _, test := range tests {
		chatRepository.EXPECT().GetChatMembersByChatId(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, chatID uint64) ([]model.ChatMembers, error) {
			return test.members, test.membersError
		}).AnyTimes()

		producerUsecase.EXPECT().ProduceMessage(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, message model.ProducerMessage) error {
			return test.producerError
		}).AnyTimes()

		err := usecase.PutInProducer(context.TODO(), test.body)
		time.Sleep(100 * time.Millisecond)
		if test.result == nil {
			require.NoError(t, err)
			continue
		}

		require.Equal(t, test.result.Error(), err.Error())
	}
}
