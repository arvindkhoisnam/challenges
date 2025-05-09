import React, { useState, type SetStateAction } from "react";
import { useNavigate } from "react-router";
import { useUserInfo } from "./Context";

function Home() {
  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  const [contact, setContact] = useState("");
  const [hobbies, setHobbies] = useState("");
  const [message, setMessage] = useState("");
  const [url, setUrl] = useState("");
  function onChangeHandler(
    e: string,
    setterFunction: React.Dispatch<SetStateAction<string>>
  ) {
    setterFunction(e);
  }
  const navigate = useNavigate();
  const { setUserInfo } = useUserInfo();
  function handleSubmit() {
    const info = {
      email,
      message,
      url,
    };
    setUserInfo(info);
    navigate("/successful");
  }
  return (
    <div className="h-screen bg-zinc-800 flex flex-col justify-center items-center text-zinc-200">
      <form
        className="p-4 bg-zinc-700 rounded-lg flex flex-col gap-6 min-w-1/2"
        onSubmit={(e) => {
          e.preventDefault();
          handleSubmit();
        }}
      >
        <h3 className="text-center text-2xl font-extralight" id="heading">
          Personal Info Collector
        </h3>
        <input
          id="name"
          value={name}
          onChange={(e) => onChangeHandler(e.target.value, setName)}
          placeholder="name"
          className="px-3 py-1 border border-zinc-500 rounded-md text-zinc-50 focus:outline-none"
        />
        <input
          id="email"
          value={email}
          onChange={(e) => onChangeHandler(e.target.value, setEmail)}
          placeholder="email"
          className="px-3 py-1 border border-zinc-500 rounded-md text-zinc-50 focus:outline-none"
        />{" "}
        <input
          id="contact"
          value={contact}
          onChange={(e) => onChangeHandler(e.target.value, setContact)}
          placeholder="contact"
          className="px-3 py-1 border border-zinc-500 rounded-md text-zinc-50 focus:outline-none"
        />
        <input
          id="hobbies"
          value={hobbies}
          onChange={(e) => onChangeHandler(e.target.value, setHobbies)}
          placeholder="hobbies"
          className="px-3 py-1 border border-zinc-500 rounded-md text-zinc-50 focus:outline-none"
        />
        <input
          id="message"
          value={message}
          onChange={(e) => onChangeHandler(e.target.value, setMessage)}
          placeholder="message"
          className="px-3 py-1 border border-zinc-500 rounded-md text-zinc-50 focus:outline-none"
        />
        <input
          id="image"
          value={url}
          onChange={(e) => onChangeHandler(e.target.value, setUrl)}
          placeholder="image"
          className="px-3 py-1 border border-zinc-500 rounded-md text-zinc-50 focus:outline-none"
        />
        <button
          id="button"
          className="px-3 py-1 bg-violet-600 text-violet-50 rounded-md cursor-pointer"
        >
          Submit
        </button>
      </form>
    </div>
  );
}

export default Home;
