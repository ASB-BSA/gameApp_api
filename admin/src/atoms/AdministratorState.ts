import { atom } from "recoil";

export type AdministratorStateType = {
  ID: number
  user: string
  CreatedAt: string
}

const AdministratorState = atom<AdministratorStateType[]>({
  key: 'AdministratorState',
  default: [],
});

export default AdministratorState