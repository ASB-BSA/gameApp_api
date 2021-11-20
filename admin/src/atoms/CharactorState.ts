import { atom } from "recoil";

export type CharactorType = {
  ID: number|string
  name: string
  english: string
  img: string
  icon: string
}

const CharactorState = atom<CharactorType[]>({
  key: 'CharactorState',
  default: [],
});

export default CharactorState