import { CharacteristicStateType } from "../atoms/CharacteristicState";

const timingFunc = (e: CharacteristicStateType['timing']) => {
    switch(e) {
        case 'start':
            return 'ターン開始時'
        case 'damage':
            return 'ダメージを受けたとき'
        case 'attack':
            return '攻撃した時'
        case 'end':
            return 'ターン終了時'
    }
}

const paramFunc = (e: CharacteristicStateType['conditionsParameter']) => {
    switch(e) {
        case 'hp':
            return '残りHPが'
        case 'damage':
            return '受けるダメージが最大HPの'
    }
}

const expressionFunc = (e: CharacteristicStateType['conditionsExpression']) => {
    switch(e) {
        case '>':
            return '以下のとき'
        case '<':
            return '以上のとき'
        case '=':
            return 'のとき'
    }
}

const toWhomFunc = (e: CharacteristicStateType['toWhom']) => {
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
    const timing = timingFunc(data.timing)
    const param = paramFunc(data.conditionsParameter)
    const value = String(data.conditionsValue)
    const expression = expressionFunc(data.conditionsExpression)
    const toWhom = toWhomFunc(data.toWhom)
    const parameter = ParameterFunc(data.parameter)
    const howMuch = String(data.howMuch)
    const happen = HappenFunc(data.happen)

    const result = `${timing}、${param}${value}%${expression}、${toWhom}の${parameter}を${howMuch}${happen}`

    return result
}

export default ConvertToText