package model

import (
	"context"
	"project/internal/chat"
	"project/internal/messages"
	protoMessage "project/internal/microservices/messages/protobuf"
	protoUser "project/internal/microservices/user/protobuf"
	myErrors "project/internal/pkg/errors"
	"project/internal/pkg/model_conversion"
	"project/internal/user"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	//"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
	//"github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/post/protobuf"
	"google.golang.org/protobuf/types/known/emptypb"

	"project/internal/microservices/chat/protobuf"
)

//func ConvertToProto(p models.Post) *protobuf.Post {
//	return &protobuf.Post{
//		ID:          p.ID,
//		UserID:      p.UserID,
//		Content:     p.Content,
//		Tier:        p.Tier,
//		IsAllowed:   p.IsAllowed,
//		DateCreated: timestamppb.New(p.DateCreated),
//		Tags:        p.Tags,
//		Author: &usersProto.LessUser{
//			Id:       p.Author.UserID,
//			Username: p.Author.Username,
//			Avatar:   p.Author.ImgPath,
//		},
//		LikesNum: p.LikesNum,
//		IsLiked:  p.IsLiked,
//	}
//}
//
//func ConvertToModel(p *protobuf.Post) models.Post {
//	return models.Post{
//		ID:          p.ID,
//		UserID:      p.UserID,
//		Content:     p.Content,
//		Tier:        p.Tier,
//		IsAllowed:   p.IsAllowed,
//		DateCreated: p.DateCreated.AsTime(),
//		Tags:        p.Tags,
//		Author: models.ResponseImageUsers{
//			UserID:   p.Author.GetId(),
//			Username: p.Author.GetUsername(),
//			ImgPath:  p.Author.GetAvatar(),
//		},
//		LikesNum: p.LikesNum,
//		IsLiked:  p.IsLiked,
//	}
//}

type ChatsService struct {
	chatRepo     chat.Repository
	userRepo     user.Repository
	messagesRepo messages.Repository
	protobuf.UnimplementedChatsServer
}

func NewChatsService(chatRepo chat.Repository, userRepo user.Repository, messagesRepo messages.Repository) protobuf.ChatsServer {
	return &ChatsService{chatRepo: chatRepo, userRepo: userRepo, messagesRepo: messagesRepo}
}

func (c ChatsService) CheckExistUserInChat(ctx context.Context, checkExistUserArguments *protobuf.ExistChatArguments) (*emptypb.Empty, error) {
	members := checkExistUserArguments.Chat.Members
	userID := checkExistUserArguments.UserID.UserID
	for _, member := range members {
		if member.Id == userID {
			return nil, status.Error(codes.Internal, myErrors.ErrUserIsAlreadyInChat.Error())
		}
	}

	return &emptypb.Empty{}, nil
}

func (c ChatsService) GetChatById(ctx context.Context, chatID *protobuf.ChatID) (*protobuf.Chat, error) {
	chat, err := c.chatRepo.GetChatById(context.Background(), chatID.ChatID)
	if err != nil {
		return nil, err
	}

	chatMembers, err := c.chatRepo.GetChatMembersByChatId(context.Background(), chatID.ChatID)
	if err != nil {
		return nil, err
	}

	var members []User
	for _, chatMember := range chatMembers {
		user, err := c.userRepo.GetUserById(context.Background(), chatMember.MemberId)
		if err != nil {
			return nil, err
		}

		members = append(members, model_conversion.FromAuthorizedUserToUser(user))
	}

	chatMessages, err := c.messagesRepo.GetChatMessages(context.Background(), chatID.ChatID)
	if err != nil {
		return nil, err
	}

	var messages []Message
	for _, chatMessage := range chatMessages {
		message, err := c.messagesRepo.GetMessageById(context.Background(), chatMessage.MessageId)
		if err != nil {
			return nil, err
		}

		messages = append(messages, message)
	}

	var protoUsers []*protoUser.User
	for _, member := range members {
		protoUsers = append(protoUsers, &protoUser.User{
			Id:       member.Id,
			Avatar:   member.Avatar,
			Username: member.Username,
			Nickname: member.Nickname,
			Email:    member.Email,
			Status:   member.Status,
		})
	}

	var protoMessages []*protoMessage.Message
	for _, message := range messages {
		protoMessages = append(protoMessages, &protoMessage.Message{
			Id:       message.Id,
			Body:     message.Body,
			AuthorId: message.AuthorId,
			ChatId:   message.ChatId,
		})
	}

	return &protobuf.Chat{
		Id:       chat.Id,
		Type:     chat.Type,
		Title:    chat.Title,
		Avatar:   chat.Avatar,
		Members:  protoUsers,
		Messages: protoMessages,
	}, nil

	//return model.Chat{
	//	Id:       chat.Id,
	//	Type:     chat.Type,
	//	Title:    chat.Title,
	//	Avatar:   chat.Avatar,
	//	Members:  members,
	//	Messages: messages,
	//}, nil
}

//func (s PostsService) GetAllByUserID(_ context.Context, userID *usersProto.UserID) (*protobuf.PostArray, error) {
//	posts, err := s.postsRepo.GetAllByUserID(userID.GetUserId())
//	if err != nil && err != sql.ErrNoRows {
//		return nil, status.Error(codes.Internal, err.Error())
//	}
//
//	result := make([]*protobuf.Post, 0)
//	for _, post := range posts {
//		result = append(result, ConvertToProto(post))
//	}
//
//	return &protobuf.PostArray{Posts: result}, nil
//}
//
//func (s PostsService) GetPostByID(_ context.Context, postID *protobuf.PostID) (*protobuf.Post, error) {
//	post, err := s.postsRepo.GetPostByID(postID.GetPostID())
//	if err != nil && err != sql.ErrNoRows {
//		return nil, status.Error(codes.Internal, err.Error())
//	}
//
//	return ConvertToProto(post), nil
//}
//
//func (s PostsService) Create(_ context.Context, post *protobuf.Post) (*protobuf.Post, error) {
//	newPost, err := s.postsRepo.Create(ConvertToModel(post))
//	if err != nil {
//		return nil, status.Error(codes.Internal, err.Error())
//	}
//
//	return ConvertToProto(newPost), nil
//}
//
//func (s PostsService) Update(_ context.Context, post *protobuf.Post) (*emptypb.Empty, error) {
//	err := s.postsRepo.Update(ConvertToModel(post))
//	if err != nil {
//		return nil, status.Error(codes.Internal, err.Error())
//	}
//
//	return &emptypb.Empty{}, nil
//}
//
//func (s PostsService) DeleteByID(_ context.Context, postID *protobuf.PostID) (*emptypb.Empty, error) {
//	err := s.postsRepo.DeleteByID(postID.GetPostID())
//	if err != nil {
//		return nil, status.Error(codes.Internal, err.Error())
//	}
//
//	return &emptypb.Empty{}, nil
//}
//
//func (s PostsService) GetPostsBySubscriptions(_ context.Context, userID *usersProto.UserID) (*protobuf.PostArray, error) {
//	posts, err := s.postsRepo.GetPostsBySubscriptions(userID.GetUserId())
//	if err != nil && err != sql.ErrNoRows {
//		return nil, status.Error(codes.Internal, err.Error())
//	}
//	result := make([]*protobuf.Post, 0)
//	for _, post := range posts {
//		result = append(result, ConvertToProto(post))
//	}
//
//	return &protobuf.PostArray{Posts: result}, nil
//}
//
//func (s PostsService) GetLikeByUserAndPostID(_ context.Context, postUserIDs *protobuf.PostUserIDs) (*protobuf.Like, error) {
//	like, err := s.postsRepo.GetLikeByUserAndPostID(postUserIDs.UserID, postUserIDs.PostID)
//	if err != nil && err != sql.ErrNoRows {
//		return nil, status.Error(codes.Internal, err.Error())
//	}
//
//	return &protobuf.Like{
//		UserID: like.UserID,
//		PostID: like.PostID,
//	}, nil
//}
//
//func (s PostsService) GetAllLikesByPostID(_ context.Context, postID *protobuf.PostID) (*protobuf.Likes, error) {
//	likes, err := s.postsRepo.GetAllLikesByPostID(postID.GetPostID())
//	if err != nil && err != sql.ErrNoRows {
//		return nil, status.Error(codes.Internal, err.Error())
//	}
//
//	result := make([]*protobuf.Like, 0)
//	for _, like := range likes {
//		result = append(result, &protobuf.Like{
//			UserID: like.UserID,
//			PostID: like.PostID,
//		})
//	}
//
//	return &protobuf.Likes{Likes: result}, nil
//}
//
//func (s PostsService) CreateLike(_ context.Context, postUserIDs *protobuf.PostUserIDs) (*emptypb.Empty, error) {
//	err := s.postsRepo.CreateLike(postUserIDs.UserID, postUserIDs.PostID)
//	if err != nil {
//		return nil, status.Error(codes.Internal, err.Error())
//	}
//
//	return &emptypb.Empty{}, nil
//}
//
//func (s PostsService) DeleteLikeByID(_ context.Context, postUserIDs *protobuf.PostUserIDs) (*emptypb.Empty, error) {
//	err := s.postsRepo.DeleteLikeByID(postUserIDs.UserID, postUserIDs.PostID)
//	if err != nil {
//		return nil, status.Error(codes.Internal, err.Error())
//	}
//
//	return &emptypb.Empty{}, nil
//}
//
//func (s PostsService) CreateTag(_ context.Context, tagName *protobuf.TagName) (*protobuf.TagID, error) {
//	id, err := s.postsRepo.CreateTag(tagName.GetTagName())
//	if err != nil {
//		return nil, status.Error(codes.Internal, err.Error())
//	}
//
//	return &protobuf.TagID{TagID: id}, nil
//}
//
//func (s PostsService) GetTagById(_ context.Context, tagID *protobuf.TagID) (*protobuf.Tag, error) {
//	tag, err := s.postsRepo.GetTagById(tagID.GetTagID())
//	if err != nil && err != sql.ErrNoRows {
//		return nil, status.Error(codes.Internal, err.Error())
//	}
//
//	return &protobuf.Tag{
//		Id:      tag.ID,
//		TagName: tag.TagName,
//	}, nil
//}
//
//func (s PostsService) GetTagByName(_ context.Context, tagName *protobuf.TagName) (*protobuf.Tag, error) {
//	tag, err := s.postsRepo.GetTagByName(tagName.GetTagName())
//	if err != nil && err != sql.ErrNoRows {
//		return nil, status.Error(codes.Internal, err.Error())
//	}
//
//	return &protobuf.Tag{
//		Id:      tag.ID,
//		TagName: tag.TagName,
//	}, nil
//}
//
//func (s PostsService) CreateDepTag(_ context.Context, tagDep *protobuf.TagDep) (*emptypb.Empty, error) {
//	err := s.postsRepo.CreateDepTag(tagDep.PostID, tagDep.TagID)
//	if err != nil {
//		return nil, status.Error(codes.Internal, err.Error())
//	}
//
//	return &emptypb.Empty{}, nil
//}
//
//func (s PostsService) GetTagDepsByPostId(_ context.Context, postID *protobuf.PostID) (*protobuf.TagDeps, error) {
//	tagDeps, err := s.postsRepo.GetTagDepsByPostId(postID.GetPostID())
//	if err != nil && err != sql.ErrNoRows {
//		return nil, status.Error(codes.Internal, err.Error())
//	}
//
//	result := make([]*protobuf.TagDep, 0)
//
//	for _, tagDep := range tagDeps {
//		result = append(result, &protobuf.TagDep{
//			PostID: tagDep.PostID,
//			TagID:  tagDep.TagID,
//		})
//	}
//
//	return &protobuf.TagDeps{TagDeps: result}, nil
//}
//
//func (s PostsService) DeleteDepTag(_ context.Context, tagDep *protobuf.TagDep) (*emptypb.Empty, error) {
//	err := s.postsRepo.DeleteDepTag(models.TagDep{
//		PostID: tagDep.GetPostID(),
//		TagID:  tagDep.GetTagID(),
//	})
//	if err != nil {
//		return nil, status.Error(codes.Internal, err.Error())
//	}
//
//	return &emptypb.Empty{}, nil
//}
//
//func (s PostsService) CreateComment(_ context.Context, comment *protobuf.Comment) (*protobuf.Comment, error) {
//	com, err := s.postsRepo.CreateComment(models.Comment{
//		PostID:  comment.GetPostID(),
//		UserID:  comment.GetUserID(),
//		Content: comment.GetContent(),
//	})
//	if err != nil {
//		return nil, status.Error(codes.Internal, err.Error())
//	}
//
//	return &protobuf.Comment{
//		ID:          com.ID,
//		PostID:      com.PostID,
//		UserID:      com.UserID,
//		Content:     com.Content,
//		DateCreated: timestamppb.New(com.DateCreated),
//	}, nil
//}
//
//func (s PostsService) GetCommentByID(_ context.Context, commentID *protobuf.CommentID) (*protobuf.Comment, error) {
//	comment, err := s.postsRepo.GetCommentByID(commentID.GetCommentID())
//	if err != nil {
//		return nil, status.Error(codes.Internal, err.Error())
//	}
//
//	return &protobuf.Comment{
//		ID:          comment.ID,
//		PostID:      comment.PostID,
//		UserID:      comment.UserID,
//		Content:     comment.Content,
//		DateCreated: timestamppb.New(comment.DateCreated),
//	}, nil
//}
//
//func (s PostsService) GetCommentsByPostID(_ context.Context, postID *protobuf.PostID) (*protobuf.CommentArray, error) {
//	comments, err := s.postsRepo.GetCommentsByPostId(postID.GetPostID())
//	if err != nil {
//		return nil, status.Error(codes.Internal, err.Error())
//	}
//
//	result := make([]*protobuf.Comment, 0)
//
//	for _, comment := range comments {
//		result = append(result, &protobuf.Comment{
//			ID:          comment.ID,
//			PostID:      comment.PostID,
//			UserID:      comment.UserID,
//			Content:     comment.Content,
//			DateCreated: timestamppb.New(comment.DateCreated),
//		})
//	}
//
//	return &protobuf.CommentArray{Comments: result}, nil
//}
//
//func (s PostsService) UpdateComment(_ context.Context, comment *protobuf.Comment) (*emptypb.Empty, error) {
//	err := s.postsRepo.UpdateComment(models.Comment{
//		ID:      comment.GetID(),
//		PostID:  comment.GetPostID(),
//		UserID:  comment.GetUserID(),
//		Content: comment.GetContent(),
//	})
//	if err != nil {
//		return nil, status.Error(codes.Internal, err.Error())
//	}
//
//	return &emptypb.Empty{}, nil
//}
//
//func (s PostsService) DeleteCommentByID(_ context.Context, commentID *protobuf.CommentID) (*emptypb.Empty, error) {
//	err := s.postsRepo.DeleteCommentByID(commentID.GetCommentID())
//	if err != nil {
//		return nil, status.Error(codes.Internal, err.Error())
//	}
//
//	return &emptypb.Empty{}, nil
//}
