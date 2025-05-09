import { createContext, useContext, useState } from "react";

type UserInfo = {
  email: string;
  message: string;
  url: string;
};

type UserContextType = {
  userInfo: UserInfo;
  setUserInfo: React.Dispatch<React.SetStateAction<UserInfo>>;
};
const UserProvider = createContext<UserContextType | null>(null);

function Context({ children }: { children: React.ReactNode }) {
  const initialVal: UserInfo = {
    email: "",
    message: "",
    url: "",
  };
  const [userInfo, setUserInfo] = useState<UserInfo>(initialVal);
  return (
    <UserProvider.Provider value={{ userInfo, setUserInfo }}>
      {children}
    </UserProvider.Provider>
  );
}

function useUserInfo() {
  const context = useContext(UserProvider);

  if (!context) {
    throw new Error("useUserInfo must be used within a UserProvider");
  }
  return context;
}
export { Context, useUserInfo };
