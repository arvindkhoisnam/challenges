import { ErrorBoundary } from "react-error-boundary";
import Chat from "./components/Chat";
import Chatbox from "./components/Chatbox";
import Navbar from "./components/Navbar";
import { useChatStatus } from "./store";

function App() {
  const chatStatus = useChatStatus(state => state.started);
  return (
    <div className="bg-primary-bg h-screen">
      <div className="font relative mx-auto flex h-full max-w-7xl flex-col p-4 font-sans">
        <Navbar />
        <ErrorBoundary
          FallbackComponent={({ error, resetErrorBoundary }) => (
            <div className="absolute top-1/2 left-1/2 flex -translate-x-1/2 -translate-y-1/2 flex-col items-center justify-center gap-2">
              <p className="text-primary-text text-xl font-extralight">{error.message}</p>
              <button
                className="bg-primary-bg text-primary-text cursor-pointer rounded-lg border px-4 py-1"
                onClick={() => resetErrorBoundary()}
              >
                Try again
              </button>
            </div>
          )}
        >
          {chatStatus ? <Chat /> : <Chatbox />}
        </ErrorBoundary>
      </div>
    </div>
  );
}

export default App;
