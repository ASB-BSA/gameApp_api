import { CharacteristicStateType } from "../atoms/CharacteristicState";

const paramFunc = (e: CharacteristicStateType['conditions_parameter']) => {
    switch(e) {
        case 'hp':
            return '残りHPが'
        case 'damage':
            return '受けるダメージが最大HPの'
    }
}

const expressionFunc = (e: CharacteristicStateType['conditions_expression']) => {
    switch(e) {
        case '>':
            return '以下のとき'
        case '<':
            return '以上のとき'
        case '=':
            return 'のとき'
    }
}

const toWhomFunc = (e: CharacteristicStateType['to_whom']) => {
    switch(e) {
        case 'myself':
            return '自分'
        case 'all_allies':
            return '味方全体'
        case 'random_allies':
            return '味方単体'
        case 'all_enemies':
            return '敵全体'
        case 'ramdom_enemies':
            return '敵単体'
    }
}

const ParameterFunc = (e: CharacteristicStateType['parameter']) => {
    switch(e) {
        case 'attack':
            return '攻撃力'
        case 'defence':
            return '防御力'
        case 'critical_rate':
            return 'クリティカル率'
        case 'agility':
            return '素早さ'
        case 'hp':
            return 'HP'
        case 'mp':
            return 'MP'
    }
}

const HappenFunc = (e: CharacteristicStateType['happen']) => {
    switch(e) {
        case '+':
            return '増加させる'
        case '-':
            return '減少させる'
        case '=':
            return 'にする'
    }
}

const ConvertToText = (data:CharacteristicStateType):string => {
    console.log(data)
    const param = paramFunc(data.conditions_parameter)
    const value = String(data.conditions_value)
    const expression = expressionFunc(data.conditions_expression)
    const to_whom = toWhomFunc(data.to_whom)
    const parameter = ParameterFunc(data.parameter)
    const how_much = String(data.how_much)
    const happen = HappenFunc(data.happen)

    const result = `${param}${value}%${expression}、${to_whom}の${parameter}を${how_much}${happen}`

    return result
}

export default ConvertToText