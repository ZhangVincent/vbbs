package eventhandler

import (
	"reflect"
	"vbbs/model/constants"
	"vbbs/pkg/event"
	"vbbs/pkg/msg"
	"vbbs/services"
)

func init() {
	event.RegHandler(reflect.TypeOf(event.UserFavoriteEvent{}), handleUserFavorite)
}

func handleUserFavorite(i interface{}) {
	e := i.(event.UserFavoriteEvent)

	if e.EntityType == constants.EntityTopic {
		sendTopicFavoriteMsg(e.EntityId, e.UserId)
	} else if e.EntityType == constants.EntityArticle {
		sendArticleFavoriteMsg(e.EntityId, e.UserId)
	}
}

// sendTopicFavoriteMsg 话题被收藏
func sendTopicFavoriteMsg(topicId, favoriteUserId int64) {
	topic := services.TopicService.Get(topicId)
	if topic == nil || topic.Status != constants.StatusOk {
		return
	}
	if topic.UserId == favoriteUserId {
		return
	}
	var (
		from         = favoriteUserId
		to           = topic.UserId
		title        = "收藏了你的话题"
		quoteContent = "《" + topic.GetTitle() + "》"
	)
	services.MessageService.SendMsg(from, to, msg.TypeTopicFavorite, title, "", quoteContent,
		&msg.TopicFavoriteExtraData{
			TopicId:        topicId,
			FavoriteUserId: favoriteUserId,
		})
}

// sendArticleFavoriteMsg 文章被收藏
func sendArticleFavoriteMsg(articleId, favoriteUserId int64) {
	article := services.ArticleService.Get(articleId)
	if article == nil || article.Status != constants.StatusOk {
		return
	}
	if article.UserId == favoriteUserId {
		return
	}
	var (
		from         = favoriteUserId
		to           = article.UserId
		title        = "收藏了你的话题"
		quoteContent = "《" + article.Title + "》"
	)
	services.MessageService.SendMsg(from, to, msg.TypeArticleFavorite, title, "", quoteContent,
		&msg.ArticleFavoriteExtraData{
			ArticleId:      articleId,
			FavoriteUserId: favoriteUserId,
		})
}
