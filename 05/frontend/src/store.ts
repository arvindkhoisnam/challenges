import { create } from "zustand";

type State = {
  started: boolean;
};

type Action = {
  updateState: (started: boolean) => void;
};

type Message = {
  role: "user" | "assistant";
  content: string;
};

type Messages = {
  messages: Message[];
};

type MessageAction = {
  updateMessage: (message: Message) => void;
  refresh: () => void;
};
const useChatStatus = create<State & Action>(set => ({
  started: false,
  updateState: started => set(() => ({ started })),
}));

const useMessages = create<Messages & MessageAction>(set => ({
  messages: [],
  updateMessage: message => set(state => ({ messages: [...state.messages, message] })),
  refresh: () => set(() => ({ messages: [] })),
}));
export { useChatStatus, useMessages };
