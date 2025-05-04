import React, { useEffect, useRef, useState } from "react";
import { SiReactos } from "react-icons/si";
import Spinner from "./Spinner";

function Messages({
  isPending,
  messages,
  bottomRef,
}: {
  isPending: boolean;
  messages: { role: string; content: string }[];
  bottomRef: React.Ref<HTMLSpanElement>;
}) {
  const containerRef = useRef<HTMLDivElement | null>(null);
  const [isOverflowing, setIsOverflowing] = useState(false);

  useEffect(() => {
    const checkOverflow = () => {
      if (containerRef.current) {
        const hasOverflow = containerRef.current.scrollHeight > containerRef.current.clientHeight;
        setIsOverflowing(hasOverflow);
      }
    };

    checkOverflow();

    const resizeObserver = new ResizeObserver(checkOverflow);
    if (containerRef.current) {
      resizeObserver.observe(containerRef.current);
    }

    return () => resizeObserver.disconnect();
  });
  return (
    // <div className="scrollbar-thin scrollbar-thumb-zinc-700 scrollbar-track-scrollbar-track scrollbar-thumb-rounded flex min-h-0 w-3/4 flex-1 flex-col gap-5 overflow-y-auto mask-t-from-80% p-6 font-extralight">
    <div
      ref={containerRef}
      className={`scrollbar-thin scrollbar-thumb-zinc-700 scrollbar-track-scrollbar-track scrollbar-thumb-rounded flex min-h-0 w-3/4 flex-1 flex-col gap-5 overflow-y-auto p-6 font-extralight ${
        isOverflowing ? "mask-t-from-80%" : ""
      }`}
    >
      {messages.map((msg: { role: string; content: string }, index: number) => (
        <div
          key={index}
          className={`flex ${msg.role === "user" ? "justify-end" : "items-center justify-start gap-4"}`}
        >
          {msg.role === "assistant" ? (
            <span className="text-primary-text">
              <SiReactos size={18} />
            </span>
          ) : (
            ""
          )}
          <span
            ref={bottomRef}
            className={`text-sm ${msg.role === "user" ? "bg-hover-bg text-hover-text max-w-1/2 rounded-lg px-2 py-1" : "text-primary-text"}`}
          >
            {msg.content}
          </span>
        </div>
      ))}
      {isPending && (
        <div className="flex w-full items-center justify-start gap-2 py-2">
          <span className="text-primary-text">
            <SiReactos size={12} />
          </span>
          <Spinner />
        </div>
      )}
    </div>
  );
}

export default Messages;
