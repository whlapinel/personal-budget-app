
export default function AddTransactionPage() {
  return (
    <>
      <form className="flex flex-col items-center justify-center self-center">
        <h1>Add Transaction</h1>
        <div>
          <label htmlFor="date">Date</label>
          <input type="date" id="date" name="date" />
        </div>
        <div>
          <label htmlFor="description">Description</label>
          <input type="text" id="description" name="description" />
        </div>
        <div>
          <label htmlFor="amount">Amount</label>
          <input type="number" id="amount" name="amount" />
        </div>
        <div>
          <button type="submit">Add</button>
        </div>
      </form>
    </>
  )
}
