import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import "./index.css";
import App from "./App.tsx";
import { Context } from "./pages/Context.tsx";

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <Context>
      <App />
    </Context>
  </StrictMode>
);
