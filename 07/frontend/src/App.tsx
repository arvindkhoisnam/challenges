import { BrowserRouter, Route, Routes } from "react-router";
import Home from "./pages/Home";
import Successfull from "./pages/Successfull";

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/successful" element={<Successfull />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
