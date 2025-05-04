import React from "react";
import { IoIosSettings } from "react-icons/io";
import { LiaTemperatureHighSolid } from "react-icons/lia";
import { SlArrowUpCircle } from "react-icons/sl";
import Spinner from "./Spinner";

function TextArea({
  extendWidth,
  isPending,
  prompt,
  mutate,
  setPrompt,
  textArea,
}: {
  extendWidth: boolean;
  isPending?: boolean;
  prompt: string;
  mutate: () => void;
  setPrompt: React.Dispatch<React.SetStateAction<string>>;
  textArea: React.Ref<HTMLTextAreaElement>;
}) {
  return (
    <form
      onSubmit={e => {
        e.preventDefault(); // 👈 Critical fix
        if (prompt.length < 1) {
          return;
        }
        mutate();
      }}
      className={`${extendWidth ? "w-full" : "w-1/2"} rounded-md border border-zinc-700 bg-zinc-900/30`}
    >
      <textarea
        ref={textArea}
        onKeyDown={e => {
          if (e.key === "Enter" && !e.shiftKey) {
            e.preventDefault();
            e.currentTarget.form?.requestSubmit();
          }
        }}
        value={prompt}
        onChange={e => setPrompt(e.target.value)}
        spellCheck="false"
        className="text-primary-text h-16 w-full resize-none p-3 text-sm focus:outline-none"
        placeholder="Ask anything..."
      />
      <div className="flex items-center justify-between p-1.5">
        <div className="flex gap-2">
          <IoIosSettings
            className="hover:text-primary-text cursor-pointer text-zinc-500"
            size={20}
          />
          <LiaTemperatureHighSolid
            className="hover:text-primary-text cursor-pointer text-zinc-500"
            size={20}
          />
        </div>
        {isPending ? (
          <Spinner />
        ) : (
          <SlArrowUpCircle
            onClick={() => {
              if (prompt.length < 1) {
                return;
              }
              mutate();
            }}
            size={20}
            className={`text-zinc-500 ${!prompt ? "cursor-not-allowed" : "hover:text-primary-text cursor-pointer"}`}
          />
        )}
      </div>
    </form>
  );
}

export default TextArea;
