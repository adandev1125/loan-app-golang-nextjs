import axios from "axios";
import { useCallback, useState } from "react";

export default function Home() {
  const [initializing, setInitializing] = useState(false);

  const onStart = useCallback(async () => {
    setInitializing(true);

    try {
      const {data} = await axios.get(process.env.API_HOST + "loan/init");

      if (data.error) {
        throw data.msg;
      }

      window.location.href = "input-business-detail";
    } catch (error) {
      alert(error);
    }

    setInitializing(false);
  }, []);

  return (
    <main className="min-h-screen flex flex-col items-center justify-center">
      <h1 className="mb-4">Simple Loan App</h1>

      <span className="mb-16">For Demyst</span>

      <button
        type="button"
        className="btn-primary"
        onClick={onStart}
        disabled={initializing}
      >
        Get Started
      </button>
    </main>
  );
}
