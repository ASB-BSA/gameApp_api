import { atom } from "recoil";

export type CharacteristicStateType = {
  name: string
	timing: 'start'|'attack'|'damage'|'end'
  conditionsParameter: 'hp'|'damage'
	conditionsValue: number|string
	conditionsExpression: '>'|'<'|'='
	toWhom: 'myself'|'all_allies'|'random_allies'|'all_enemies'|'ramdom_enemies'
	parameter: 'attack'|'defence'|'critical_rate'|'agility'|'hp'|'mp'
	happen: '+'|'-'|'='
	howMuch: number|string
}

const CharacteristicState = atom<CharacteristicStateType[]>({
  key: 'CharacteristicState',
  default: [],
});

export default CharacteristicState