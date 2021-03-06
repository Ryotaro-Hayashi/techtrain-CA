package controllers

import (
  "github.com/gin-gonic/gin"
  "techtrain-CA/database"
  "techtrain-CA/models"
)

type GachaController struct {
  GachaRepository      database.GachaRepository
  CollectionRepository database.CollectionRepository
  UserRepository       database.UserRepository
}

func NewGachaController(sqlHandler *database.SqlHandler) *GachaController {
  return &GachaController{
    GachaRepository: database.GachaRepository{
      SqlHandler: sqlHandler,
    },
    CollectionRepository: database.CollectionRepository{
      SqlHandler: sqlHandler,
    },
    UserRepository: database.UserRepository{
      SqlHandler: sqlHandler,
    },
  }
}

// トークンからuser_idを取得してランダムにcharacter_idを生成して保存して返す
func (controller *GachaController) Draw(c *gin.Context) {
  // リクエストに合う構造体を定義
  type GachaTimes struct {
    Times int
  }
  gachaTimes := GachaTimes{}

  err := c.Bind(&gachaTimes)
  if err != nil {
    c.JSON(500, err.Error())
    return
  }

  // ヘッダーのtokenを取得
  tokenString := c.Request.Header.Get("x-token")
  if tokenString == "" {
    c.JSON(500, "token must be needed.")
    return
  }

  // トークンでユーザーを検索
  user, err := controller.UserRepository.FindByToken(tokenString)
  if err != nil {
		c.JSON(500, err.Error())
		return
  }

  // 保存するキャラクターidを選択
  characterIds, err := controller.GachaRepository.Choose(gachaTimes.Times)
  if err != nil {
    c.JSON(500, err.Error())
    return
  }

  // キャラクターを保存して、保存されたcollectionのidを返す
  storedCharacterIds, err := controller.CollectionRepository.Store(user.Id, characterIds)
  if err != nil {
		c.JSON(500, err.Error())
		return
	}

  // 保存したcollectionをフォーマットを調整して返す
  gachaDrawResponses, err := controller.CollectionRepository.FindByIds(storedCharacterIds)

  // マップに保存したガチャ内容を格納
  result := map[string]models.GachaDrawResponses{"result": gachaDrawResponses}

  c.JSON(200, result)
}
