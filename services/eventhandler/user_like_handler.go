package eventhandler

import (
	"reflect"
	"vbbs/model/constants"
	"vbbs/pkg/event"
	"vbbs/pkg/msg"
	"vbbs/services"
)

func init() {
	event.RegHandler(reflect.TypeOf(event.UserLikeEvent{}), handleUserLike)
	event.RegHandler(reflect.TypeOf(event.UserUnLikeEvent{}), handleUserUnLike)
}

func handleUserLike(i interface{}) {
	e := i.(event.UserLikeEvent)

	if e.EntityType == constants.EntityTopic {
		sendTopicLikeMsg(e.EntityId, e.UserId)
	} else if e.EntityType == constants.EntityComment {
		sendCommentLikeMsg(e.EntityId, e.UserId)
	}
}

func handleUserUnLike(i interface{}) {
	e := i.(event.UserUnLikeEvent)
	if e.EntityType == constants.EntityTopic {
		sendTopicUnLikeMsg(e.EntityId, e.UserId)
	}
}

// 话题收到点赞
func sendTopicLikeMsg(topicId, likeUserId int64) {
	topic := services.TopicService.Get(topicId)
	if topic == nil || topic.Status != constants.StatusOk {
		return
	}
	if topic.UserId == likeUserId {
		return
	}
	var (
		from         = likeUserId
		to           = topic.UserId
		title        = "点赞了你的话题"
		quoteContent = "《" + topic.GetTitle() + "》"
	)
	services.MessageService.SendMsg(from, to, msg.TypeTopicLike, title, "", quoteContent,
		&msg.TopicLikeExtraData{
			TopicId:    topicId,
			LikeUserId: likeUserId,
		})
}

// 话题收到点赞
func sendCommentLikeMsg(commentId, likeUserId int64) {
	comment := services.CommentService.Get(commentId)
	if comment == nil || comment.Status != constants.StatusOk {
		return
	}
	if comment.UserId == likeUserId {
		return
	}

	topic := services.TopicService.Get(comment.EntityId)
	if topic == nil || topic.Status != constants.StatusOk {
		return
	}

	var (
		from         = likeUserId
		to           = comment.UserId
		title        = "点赞了你的评论"
		quoteContent = "《" + topic.GetTitle() + "》"
	)
	services.MessageService.SendMsg(from, to, msg.TypeCommentLike, title, "", quoteContent,
		&msg.CommentLikeExtraData{
			CommentId:  commentId,
			LikeUserId: likeUserId,
		})
}

// 话题收到不喜欢
func sendTopicUnLikeMsg(topicId, likeUserId int64) {
	topic := services.TopicService.Get(topicId)
	if topic == nil || topic.Status != constants.StatusOk {
		return
	}
	if topic.UserId == likeUserId {
		return
	}
	var (
		from         = likeUserId
		to           = topic.UserId
		title        = "觉得你的话题很烂"
		quoteContent = "《" + topic.GetTitle() + "》"
	)
	services.MessageService.SendMsg(from, to, msg.TypeTopicLike, title, "", quoteContent,
		&msg.TopicLikeExtraData{
			TopicId:    topicId,
			LikeUserId: likeUserId,
		})
}
