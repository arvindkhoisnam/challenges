import ReactConfetti from "react-confetti";
import { useUserInfo } from "./Context";

function Successfull() {
  const { userInfo } = useUserInfo();
  return (
    <div className="h-screen bg-zinc-800 flex flex-col justify-center items-center text-zinc-200">
      <ReactConfetti tweenDuration={10000} />
      <div className="p-6 bg-zinc-900 rounded-lg text-xl font-extralight min-w-1/2 flex flex-col gap-4 items-center">
        <div
          className="size-50 rounded-full bg-cover bg-center"
          style={{
            backgroundImage: `url(${userInfo.url})`,
          }}
          aria-label="profile-image"
        />
        <h3 id="result">Thank you for your submission {userInfo.email}.</h3>
        <p>
          Your message reads:{" "}
          <span className="text-violet-500">{userInfo.message}</span>
        </p>
      </div>
    </div>
  );
}

export default Successfull;
