import { useState } from "react";
import "./App.css";

function App() {
  return (
    <div className="h-screen bg-neutral-900 flex flex-col justify-center items-center">
      <div>
        <h3 className="text-fuchsia-500 font-extralight text-2xl mb-5 text-center">
          Playwright says hello 👋
        </h3>
        <Bands />
      </div>
      <a
        id="link"
        href="https://google.com"
        target="new"
        className="text-fuchsia-500 text-center"
      >
        Go to Google
      </a>
    </div>
  );
}

function Bands() {
  const BANDS = ["Tool", "Animals as Leaders", "Karnivool"];
  const [bands, setBands] = useState(BANDS);
  const [newBand, setNewBand] = useState("");
  return (
    <div className="flex flex-col items-center">
      <h3 className="text-neutral-100 font-extralight text-2xl mb-5 text-center">
        🎵 FAVOURITE BANDS 🎶
      </h3>
      <ul
        className="p-4 bg-neutral-100 rounded-xl max-h-96 overflow-auto"
        id="content"
      >
        {bands.map((band, index) => (
          <li key={index} className="text-neutral-900 font-extralight mb-1">
            {index + 1}.{band}
          </li>
        ))}
      </ul>
      <form className="flex gap-2 my-2">
        <input
          id="text-input"
          value={newBand}
          onChange={(e) => setNewBand(e.target.value)}
          placeholder="Add band..."
          className="bg-neutral-200 text-neutral-900 p-2 rounded-lg font-extralight"
        />
        <button
          id="btn"
          className="bg-fuchsia-800 text-neutral-100 rounded-lg text-sm flex justify-center items-center px-4 cursor-pointer"
          onClick={(e) => {
            if (newBand.length === 0) {
              return;
            }
            e.preventDefault();
            setBands((bands) => [...bands, newBand]);
            setNewBand("");
          }}
        >
          Add
        </button>
      </form>
    </div>
  );
}
export default App;
