package controllers

import (
	"boomin_game_api/src/database"
	"boomin_game_api/src/middlewares"
	"boomin_game_api/src/models"
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"

	"github.com/gofiber/fiber/v2"
)

type BattlePost struct {
	BattleId int `json:"battleId"`
}

func PostBattle(c *fiber.Ctx) error {
	p := new(BattlePost)

	userId, _ := middlewares.GetUserId(c)

	var battle models.Battle

	battle.ID = uint(p.BattleId)

	if result := database.DB.Where("is_active = ?", 1).First(&battle); result.Error != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "対戦情報が見つかりませんでした",
		})
	}

	if battle.UsersID != userId && battle.OpponentID != userId {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "対戦情報が見つかりませんでした",
		})
	}

	token, err := middlewares.GenerateBattleToken(uint(p.BattleId))

	if err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "対戦情報が見つかりませんでした",
		})
	}

	cookie := fiber.Cookie{
		Name:     "battle-jwt",
		Value:    token,
		Expires:  time.Now().AddDate(10, 0, 0),
		SameSite: "None",
		Secure:   true,
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	database.DB.Preload("User").Preload("UserTeams").Preload("UserTeams.Teams").Preload("OpponentUser").Preload("OpponentTeams").Preload("OpponentTeams.Teams").First(&battle)

	return c.JSON(battle)
}

func GetBattle(c *fiber.Ctx) error {
	id, _ := middlewares.GetBattleId(c)

	var battle models.Battle
	battle.ID = id

	if result := database.DB.Preload("User").Preload("UserTeams").Preload("UserTeams.Teams").Preload("OpponentUser").Preload("OpponentTeams").Preload("OpponentTeams.Teams").First(&battle); result.Error != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "対戦情報が見つかりませんでした",
		})
	}

	return c.JSON(battle)
}

func CreateBattleLog(c *fiber.Ctx) error {
	id, _ := middlewares.GetBattleId(c)
	var battle models.Battle
	battle.ID = id

	if result := database.DB.Preload("User").Preload("UserTeams").Preload("UserTeams.Teams").Preload("UserTeams.Teams.Characteristics").Preload("OpponentUser").Preload("OpponentTeams").Preload("OpponentTeams.Teams").Preload("OpponentTeams.Teams.Characteristics").Where("is_active = ?", "1").First(&battle); result.Error != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "対戦情報が見つかりませんでした",
		})
	}

	// キャラクターのみを抽出
	var characters []models.BattleCharacter
	var characters2 []models.BattleCharacter

	// どちらか全滅しているか確認
	for _, v := range battle.OpponentTeams.Teams {
		if v.Hp > 0 {
			characters = append(characters, v)
		}
	}
	for _, v := range battle.UserTeams.Teams {
		if v.Hp > 0 {
			characters = append(characters, v)
		}
	}

	if len(characters) == 0 || len(characters2) == 0 {
		fmt.Println("owari")
	}

	characters = append(characters, characters2...)

	sort.Slice(characters, func(i, j int) bool {
		return characters[i].Agility > characters[j].Agility
	})

	tx := database.DB.Begin()

	// ターン開始時のスキル発動
	// for _, v := range characters {
	// 	if v.Characteristics.Timing == "start" {
	// 		CharacteristicConditions(characters, v)
	// 	}
	// }

	// 攻撃処理
	for _, v := range characters {
		var targets []models.BattleCharacter

		for _, chara := range characters {
			if chara.BattleTeamsID != v.BattleTeamsID {
				targets = append(targets, chara)
			}
		}

		rand.Seed(time.Now().UnixNano())
		num := rand.Intn(len(targets))

		// 攻撃されるターゲット
		target := targets[num]

		// ダメージ計算
		damage, err := AttackProcess(target, v)

		if err != nil {
			println(err)
			continue
		}

		log := models.BattleLogs{
			BattleID:          battle.ID,
			BattleCharacterID: target.ID,
			AttackerID:        v.ID,
			Parameter:         "hp",
			LogType:           "attack",
			NumericalValue:    damage,
		}

		target.Hp = target.Hp - damage

		if res := tx.Updates(&target); res.Error != nil {
			tx.Rollback()
			return res.Error
		}

		if res := tx.Create(&log); res.Error != nil {
			tx.Rollback()
			return res.Error
		}

		fmt.Println("==========================")
		fmt.Println(v.ID)
		fmt.Println(target.ID)
		fmt.Println(damage)
	}

	return c.JSON(characters)
}

func AttackProcess(target models.BattleCharacter, attacker models.BattleCharacter) (int, error) {
	// 回避率の計算
	avoidance := (50 / target.Avoidance) * 100

	// 確率
	hit := rand.Intn(100) + 1

	// 攻撃が当たるかどうかのチェック
	if hit > avoidance {
		err := fmt.Errorf("%d: %s", 1, "攻撃が回避されました")
		return 0, err
	}

	// クリティカルかどうか
	isCritical := rand.Intn(100) <= (attacker.CriticalRate / 2)

	// ダメージ計算
	randomNumber := Round((rand.Float64()*(1.2-0.9) + 0.9), 1)
	damage := Round((float64((attacker.Attack/2)-(target.Defence/4)) * randomNumber), 0)

	if isCritical {
		damage = Round((damage * 1.8), 0)
	}

	return int(damage), nil
}

func CharacteristicConditions(data []models.BattleCharacter, attacker models.BattleCharacter) error {
	// 発動条件
	if attacker.Characteristics.ConditionsParameter == "hp" {
		nowHp := (attacker.Hp / attacker.DefaultsHp * 100)
		switch attacker.Characteristics.ConditionsExpression {
		case ">":
			if nowHp >= attacker.Characteristics.ConditionsValue {
				fmt.Println("発動")
				CharacteristicProcess(data, attacker)
			}

		case "<":
			if nowHp <= attacker.Characteristics.ConditionsValue {
				fmt.Println("発動")
				CharacteristicProcess(data, attacker)
			}

		case "=":
			if nowHp == attacker.Characteristics.ConditionsValue {
				fmt.Println("発動")
				CharacteristicProcess(data, attacker)
			}
		}
	}

	return nil
}

func CharacteristicProcess(data []models.BattleCharacter, attacker models.BattleCharacter) error {

	return nil
}

func Round(f float64, places int) float64 {
	shift := math.Pow(10, float64(places))
	return math.Floor(f*shift+.5) / shift
}
