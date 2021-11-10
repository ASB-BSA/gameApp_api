import { atom } from "recoil";

export type SettingGroup = {
  ID: number
  group_name: string
  group_category: string
  settings: SettingType[]
}

export type SettingType = {
  ID: number
  group_id: string|number|undefined
  setting_name: string
  setting_label: string
  setting_value: string
  setting_type: 'string'|'int'|'number'
}

const SettingState = atom<SettingGroup[]>({
  key: 'SettingState',
  default: [],
});

export default SettingState