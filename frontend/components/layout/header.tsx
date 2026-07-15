import { Button } from "@/components/ui/button";

export function Header() {
  return (
    <header className="sticky top-0 z-40 border-b border-white/5 bg-slate-950/40 backdrop-blur-xl">
      <div className="mx-auto flex max-w-7xl items-center justify-between px-6 py-4 lg:px-10">
        <a href="/" className="text-sm font-semibold uppercase tracking-[0.28em] text-white">
          SaaS Starter
        </a>
        <nav className="hidden items-center gap-6 text-sm text-slate-300 md:flex">
          <a href="/dashboard">Dashboard</a>
          <a href="/login">Login</a>
          <a href="#architecture">Architecture</a>
        </nav>
        <Button href="/login" size="sm" variant="secondary">
          Sign in
        </Button>
      </div>
    </header>
  );
}