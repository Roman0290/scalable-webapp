export default function LoginPage() {
  return (
    <section className="mx-auto max-w-2xl px-6 py-16 lg:px-10">
      <div className="glass-panel rounded-[2rem] p-8 lg:p-12">
        <p className="text-sm uppercase tracking-[0.25em] text-slate-400">Authentication</p>
        <h1 className="mt-3 text-3xl font-semibold text-white">Sign in</h1>
        <p className="mt-4 text-slate-300">
          Connect this screen to the auth service, Zod validation, and React Query mutations.
        </p>
      </div>
    </section>
  );
}