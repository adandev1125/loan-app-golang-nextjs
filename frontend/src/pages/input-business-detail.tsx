import axios from "axios";
import { useCallback, useRef, useState } from "react";
import { accountingProviders, balanceSheetHeaders } from "../constants";

export default function InputBusinessDetails() {
  const [balance, setBalance] = useState([]);
  const [fetchingBalance, setFetchingBalance] = useState(false);
  const [submitting, setSubmitting] = useState(false);

  const businessName = useRef("");
  const accountingProvider = useRef(accountingProviders[0]);
  const businessYear = useRef(2022);
  const loanAmount = useRef(0);

  const onSubmit = useCallback(async () => {
    // Validate inputs.
    if (loanAmount.current === 0 || isNaN(loanAmount.current)) {
      alert("Input loan amount.");
      return;
    }
    if (businessName.current.length === 0) {
      alert("Input business name.");
      return;
    }
    if (businessYear.current === 0 || isNaN(businessYear.current)) {
      alert("Input established year.");
      return;
    }

    // Set disable flag of submit button
    setSubmitting(true);

    // Sends apply request to api server.
    try {
      const { data } = await axios.post(
        process.env.API_HOST + "loan/apply_loan",
        {
          amount: loanAmount.current,
          businessName: businessName.current,
          businessYear: businessYear.current,
          accountingProvider: accountingProvider.current,
        }
      );

      if (data.error) {
        throw data.errorMsg;
      }

      window.location.href = "final-outcome/" + data.amount;
    } catch (error) {
      alert(error);
    }

    // Release disable flag of submit button
    setSubmitting(false);
  }, []);

  const onReviewBalance = useCallback(async () => {
    // Validate inputs.
    if (businessName.current.length === 0) {
      alert("Input business name.");
      return;
    }

    // Set disable flag of review balance button.
    setFetchingBalance(true);

    // Sends request for balance sheet review.
    try {
      const { data } = await axios.get(
        process.env.API_HOST + "loan/get_balance",
        {
          params: {
            business_name: businessName,
            accounting_provider: accountingProvider,
          },
        }
      );

      if (data.error) {
        throw data.errorMsg;
      }

      setBalance(data.balance);
    } catch (error) {
      alert(error);
    }

    // Release disable flag of review balance button.
    setFetchingBalance(false);
  }, []);

  const onChangeBusinessName = useCallback((e: any) => {
    businessName.current = e.target.value;
  }, []);

  const onChangeAccountingProvider = useCallback((e: any) => {
    accountingProvider.current = e.target.value;
  }, []);

  const onChangeBusinessyear = useCallback((e: any) => {
    businessYear.current = parseInt(e.target.value);
  }, []);

  const onChangeLoanAmount = useCallback((e: any) => {
    loanAmount.current = parseInt(e.target.value);
  }, []);

  return (
    <main className="flex flex-col justify-center items-center">
      <h1 className="mt-8 mb-16">Enter Business Details and Loan Amount</h1>

      <div className="grid grid-cols-2 gap-8">
        <div>
          <label htmlFor="name">Business name:</label>
          <input
            id="name"
            className="input"
            type="text"
            onInput={onChangeBusinessName}
          />
        </div>

        <div>
          <label htmlFor="year">Year established</label>
          <input
            id="year"
            className="input"
            type="number"
            defaultValue={2022}
            onInput={onChangeBusinessyear}
          />
        </div>

        <div>
          <label htmlFor="provider">Accounting Provider</label>
          <select
            id="provider"
            className="input"
            onChange={onChangeAccountingProvider}
            defaultValue={accountingProviders[0]}
          >
            {accountingProviders.map((provider, index) => (
              <option key={provider + index}>{provider}</option>
            ))}
          </select>
        </div>

        <div>
          <label htmlFor="amount">Loan Amount:</label>
          <input
            id="amount"
            className="input"
            type="number"
            defaultValue={0}
            onInput={onChangeLoanAmount}
          />
        </div>

        <button
          type="button"
          className="btn-outline"
          disabled={fetchingBalance}
          onClick={onReviewBalance}
        >
          Review Balance
        </button>

        <button
          type="button"
          className="btn-primary"
          onClick={onSubmit}
          disabled={submitting}
        >
          Submit
        </button>
      </div>

      <div className="grid grid-cols-4 gap-4 mt-16 px-8 py-4 bg-white rounded-2xl shadow-xl">
        <span className="col-span-4 text-center font-bold text-lg mb-4">
          Balance Sheet
        </span>

        {balanceSheetHeaders.map((header) => (
          <span
            key={"sheet_header_" + header.field}
            className="balance-sheet-header"
          >
            {header.caption}
          </span>
        ))}

        {balance.length === 0 ? (
          <span className="col-span-4 text-gray-500 text-center">No data.</span>
        ) : (
          balance.map((b, index) =>
            balanceSheetHeaders.map((header) => (
              <span
                key={"sheet_content_" + index + header.field}
                className="text-center"
              >
                {b[header.field]}
              </span>
            ))
          )
        )}
      </div>
    </main>
  );
}
