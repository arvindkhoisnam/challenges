import { SiStackblitz } from "react-icons/si";
import Typewriter from "./Typewriter";
import axios from "axios";
import { useMutation } from "@tanstack/react-query";
import { useEffect, useRef, useState } from "react";
import { useChatStatus, useMessages } from "../store";
import TextArea from "./TextArea";
import { SiReactos } from "react-icons/si";
import { useErrorBoundary } from "react-error-boundary";
const apiUrl = import.meta.env.VITE_PRIMARY_BACKEND;

function Chatbox() {
  const [prompt, setPrompt] = useState("");
  const textArea = useRef<HTMLTextAreaElement>(null);
  const setChatStatus = useChatStatus(state => state.updateState);
  const addMessage = useMessages(state => state.updateMessage);
  const { showBoundary } = useErrorBoundary();
  const { mutate, isPending } = useMutation({
    mutationFn: async () => {
      const data = await axios.post(
        `${apiUrl}/chat`,
        { prompt: prompt },
        { withCredentials: true }
      );
      return data.data.data;
    },
    onSuccess: data => {
      data.forEach((msg: { role: "user" | "assistant"; content: string }) => addMessage(msg));
      setChatStatus(true);
    },
    onError: err => {
      showBoundary(err);
    },
  });

  useEffect(() => {
    textArea.current?.focus();
  }, []);
  return (
    <div className="absolute top-[50%] left-[50%] flex h-1/2 w-1/2 -translate-x-1/2 -translate-y-1/2 flex-col items-center justify-center">
      <Typewriter />
      <SiReactos className="absolute text-sky-500/10" size={400} />
      <div className="flex min-w-3/4 flex-col gap-2 rounded-lg border border-zinc-700/50 bg-zinc-700/30 p-2 backdrop-blur-md">
        <div className="flex items-center gap-0.5">
          <SiStackblitz color="#9f9fa9" size={10} />
          <p className="text-primary-text text-sm">Unlock more features with Pro plan.</p>
        </div>
        <TextArea
          extendWidth={true}
          prompt={prompt}
          mutate={mutate}
          setPrompt={setPrompt}
          isPending={isPending}
          textArea={textArea}
        />
      </div>
    </div>
  );
}
export default Chatbox;
