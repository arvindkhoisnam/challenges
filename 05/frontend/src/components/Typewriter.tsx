import { useEffect, useState } from "react";
import { motion } from "framer-motion";
const LETTER_DELAY = 0.045;
const BOX_FADE_DUARION = 0.125;
const FADE_DELAY = 3;
const MAIN_FADE_DURAION = 0.25;
const SWAP_DELAY_IN_MS = 2500;

function Typewriter() {
  const EXAMPLES = [
    "Hi, I'm aibot.",
    "Your virtual personal assistant.",
    "How may I help you today?",
  ];

  const [index, setIndex] = useState(0);
  useEffect(() => {
    const intervalId = setInterval(() => {
      setIndex(pv => (pv + 1) % EXAMPLES.length);
    }, SWAP_DELAY_IN_MS);

    return () => {
      clearInterval(intervalId);
    };
  }, [EXAMPLES.length]);
  return (
    <p className="mb-2.5 text-center text-3xl font-extralight">
      {EXAMPLES[index].split("").map((l, i) => (
        <motion.span
          key={`${index}-${i}`}
          className="relative"
          initial={{ opacity: 0 }}
          animate={{ opacity: 1 }}
          transition={{ delay: i * LETTER_DELAY, duration: 0 }}
        >
          <motion.span
            initial={{ opacity: 1 }}
            animate={{ opacity: 0 }}
            transition={{
              delay: FADE_DELAY,
              duration: MAIN_FADE_DURAION,
              ease: "easeInOut",
            }}
            className="text-primary-text"
          >
            {l}
          </motion.span>
          <motion.span
            initial={{ opacity: 0 }}
            animate={{ opacity: [0, 1, 0] }}
            transition={{
              delay: i * LETTER_DELAY,
              times: [0, 0.1, 1],
              duration: BOX_FADE_DUARION,
              ease: "easeInOut",
            }}
            className="bg-primary-text absolute top-[3px] right-0 bottom-[3px] left-[1px]"
          />
        </motion.span>
      ))}
    </p>
  );
}

export default Typewriter;
