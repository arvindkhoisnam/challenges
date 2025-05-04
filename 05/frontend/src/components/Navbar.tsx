import DarkModeToggle from "./DarkModeToggle";
import { SiReactos } from "react-icons/si";
function Navbar() {
  const NAVS = ["home", "categories", "profile"];

  return (
    <div className="text-primary-text flex w-full items-center justify-between text-lg font-extralight">
      <span className="flex items-center gap-1">
        <SiReactos className="text-sky-500/90" size={20} />
        aibot
      </span>
      <div className="flex items-center gap-5 text-sm">
        {NAVS.map((nav, index) => (
          <a
            key={index}
            className="hover:bg-hover-bg hover:text-hover-text cursor-pointer rounded-md px-1.5 py-0.5"
          >
            {nav}
          </a>
        ))}
        <DarkModeToggle />
      </div>
    </div>
  );
}

export default Navbar;
