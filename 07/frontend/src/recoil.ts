import { atom } from "recoil";

const userInfo = atom({
  key: "UserInfo",
  default: {
    email: "",
    message: "",
    url: "",
  },
});

export { userInfo };
