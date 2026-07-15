import type { AnchorHTMLAttributes, ButtonHTMLAttributes, ReactNode } from "react";
import { cn } from "@/lib/utils";

type ButtonVariant = "primary" | "secondary";
type ButtonSize = "sm" | "md";

type SharedProps = {
  children: ReactNode;
  variant?: ButtonVariant;
  size?: ButtonSize;
  className?: string;
};

type ButtonAsButton = SharedProps & ButtonHTMLAttributes<HTMLButtonElement> & { href?: never };
type ButtonAsLink = SharedProps & AnchorHTMLAttributes<HTMLAnchorElement> & { href: string };

const baseStyles =
  "inline-flex items-center justify-center rounded-full font-medium transition focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-offset-2 focus-visible:ring-offset-slate-950";

const variantStyles: Record<ButtonVariant, string> = {
  primary: "bg-white text-slate-950 hover:bg-slate-200",
  secondary: "border border-white/10 bg-white/5 text-white hover:bg-white/10",
};

const sizeStyles: Record<ButtonSize, string> = {
  sm: "h-10 px-4 text-sm",
  md: "h-12 px-6 text-sm",
};

export function Button(props: ButtonAsButton | ButtonAsLink) {
  const { children, className, variant = "primary", size = "md", ...rest } = props;
  const classes = cn(baseStyles, variantStyles[variant], sizeStyles[size], className);

  if ("href" in props) {
    const linkProps = rest as AnchorHTMLAttributes<HTMLAnchorElement>;

    return (
      <a className={classes} {...linkProps} href={props.href}>
        {children}
      </a>
    );
  }

  const buttonProps = rest as ButtonHTMLAttributes<HTMLButtonElement>;

  return (
    <button className={classes} {...buttonProps}>
      {children}
    </button>
  );
}