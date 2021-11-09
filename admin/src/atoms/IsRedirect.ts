import { atom } from "recoil";

const IsRedirect = atom<boolean>({
  key: 'IsRedirect',
  default: false,
});

export default IsRedirect