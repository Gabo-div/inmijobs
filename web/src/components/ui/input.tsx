import * as React from "react"

import { cn } from "@/lib/utils"

function Input({ className, type, ...props }: React.ComponentProps<"input">) {
  return (
    <input
      type={type}
      data-slot="input"
      className={cn(
        "h-10 w-full px-3 py-2 text-base border border-[#e0e0e0] rounded bg-white text-[#191919] placeholder:text-[#666666] outline-none transition-colors",
        "focus:border-[#0A66C2] focus:ring-1 focus:ring-[#0A66C2]",
        "disabled:cursor-not-allowed disabled:opacity-50",
        className
      )}
      {...props}
    />
  )
}

export { Input }
