import { useEffect, useState } from "react";
import { IoSunnyOutline } from "react-icons/io5";
import { PiMoonStarsLight } from "react-icons/pi";

function DarkModeToggle() {
  const [currTheme, setCurrTheme] = useState("");
  useEffect(() => {
    const rootElement = document.getElementById("root");
    if (rootElement?.classList.value === "dark") {
      setCurrTheme("dark");
    } else {
      setCurrTheme("light");
    }
  }, []);
  function DarkMode() {
    const rootElement = document.getElementById("root");
    if (rootElement?.classList.value === "dark") {
      rootElement?.classList.remove("dark");
      rootElement?.classList.add("light");
      setCurrTheme("light");
      return;
    }
    rootElement?.classList.remove("light");
    rootElement?.classList.add("dark");
    setCurrTheme("dark");
  }
  function LightMode() {
    const rootElement = document.getElementById("root");
    console.log(rootElement?.classList.value);
    if (rootElement?.classList.value === "light") {
      rootElement?.classList.remove("light");
      rootElement?.classList.add("dark");
      setCurrTheme("dark");
      return;
    }
    rootElement?.classList.remove("dark");
    rootElement?.classList.add("light");
    setCurrTheme("light");
  }
  return (
    <div className="text-primary-text flex items-center justify-center gap-2 rounded-xl border border-zinc-500 px-0.5 py-0.5">
      <div
        className={`${currTheme === "light" ? "bg-hover-bg text-hover-text" : ""} cursor-pointer rounded-full p-0.5 duration-300 ease-in-out`}
      >
        <IoSunnyOutline onClick={LightMode} size={15} />
      </div>
      <div
        className={`${currTheme === "dark" ? "bg-hover-bg text-hover-text" : ""} cursor-pointer rounded-full p-0.5 duration-300 ease-in-out`}
      >
        <PiMoonStarsLight onClick={DarkMode} size={15} />
      </div>
    </div>
  );
}

export default DarkModeToggle;
