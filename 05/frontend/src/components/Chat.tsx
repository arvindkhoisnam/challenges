import { useMutation } from "@tanstack/react-query";
import axios from "axios";
import { useEffect, useRef, useState } from "react";
import { useMessages } from "../store";
import TextArea from "./TextArea";
import Messages from "./Messages";
import { useErrorBoundary } from "react-error-boundary";

const apiUrl = import.meta.env.VITE_PRIMARY_BACKEND;
function Chat() {
  const [prompt, setPrompt] = useState("");
  const { messages, updateMessage, refresh } = useMessages(state => state);
  const bottomRef = useRef<HTMLDivElement>(null);
  const textArea = useRef<HTMLTextAreaElement>(null);
  const { showBoundary } = useErrorBoundary();
  const { mutate, isPending } = useMutation({
    mutationFn: async () => {
      updateMessage({ role: "user", content: prompt });
      setPrompt("");
      const data = await axios.post(
        `${apiUrl}/chat`,
        { prompt: prompt },
        { withCredentials: true }
      );
      return data.data.data;
    },
    onSuccess: data => {
      refresh();
      data.forEach((msg: { role: "user" | "assistant"; content: string }) => updateMessage(msg));
    },
    onError: err => {
      showBoundary(err);
    },
  });
  useEffect(() => {
    bottomRef.current?.scrollIntoView({ behavior: "smooth" });
    textArea.current?.focus();
  }, [messages]);
  return (
    <div className="my-4 flex h-full min-h-0 flex-1 flex-col items-center gap-2">
      <Messages messages={messages} bottomRef={bottomRef} isPending={isPending} />
      <TextArea
        extendWidth={false}
        prompt={prompt}
        mutate={mutate}
        setPrompt={setPrompt}
        textArea={textArea}
        isPending={isPending}
      />
    </div>
  );
}

export default Chat;
