import { atom } from "recoil";

export type CharacteristicStateType = {
  name: string
  conditions_parameter: 'hp'|'damage'
	conditions_value: number|string
	conditions_expression: '>'|'<'|'='
	to_whom: 'myself'|'all_allies'|'random_allies'|'all_enemies'|'ramdom_enemies'
	parameter: 'attack'|'defence'|'critical_rate'|'agility'|'hp'|'mp'
	happen: '+'|'-'|'='
	how_much: number|string
}

const CharacteristicState = atom<CharacteristicStateType[]>({
  key: 'CharacteristicState',
  default: [],
});

export default CharacteristicState