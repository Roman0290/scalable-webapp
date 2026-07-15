import type { ReactNode } from "react";
import { cn } from "@/lib/utils";

type CardProps = {
  title: string;
  description: string;
  children?: ReactNode;
  className?: string;
};

export function Card({ title, description, children, className }: CardProps) {
  return (
    <article className={cn("rounded-3xl border border-white/10 bg-white/5 p-6", className)}>
      <h2 className="text-lg font-semibold text-white">{title}</h2>
      <p className="mt-3 text-sm leading-6 text-slate-300">{description}</p>
      {children}
    </article>
  );
}