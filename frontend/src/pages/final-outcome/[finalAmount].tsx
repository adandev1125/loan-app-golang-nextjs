import { useRouter } from "next/router";
import { useCallback } from "react";

export default function FinalOutcome() {
  const router = useRouter();
  const { finalAmount } = router.query;

  const onGotoFirstPage = useCallback(() => {
    window.location.href = "/";
  }, []);

  return (
    <main className="min-h-screen flex flex-col items-center justify-center">
      <h1 className="mb-16">Successfuly loaned.</h1>

      <div className="flex items-center mb-32">
        <h2 className="mr-8">Loan Amount:</h2>
        <h1 className="text-4xl font-extrabold">{finalAmount}</h1>
      </div>

      <button type="button" className="btn-primary" onClick={onGotoFirstPage}>
        Go to First Page
      </button>
    </main>
  );
}
