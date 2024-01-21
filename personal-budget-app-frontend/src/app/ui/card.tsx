import { twMerge } from "tailwind-merge"

export default function Card({children, className}: {children: React.ReactNode, className: string}) {
    return (
      <>
        <div className={twMerge("overflow-hidden bg-white shadow sm:rounded-lg", className)}>
          <div className="px-4 py-5 sm:p-6">{children}</div>
        </div>
      </>
    )
  }
  